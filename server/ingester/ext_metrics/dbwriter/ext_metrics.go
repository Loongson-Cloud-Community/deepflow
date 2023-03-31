/*
 * Copyright (c) 2022 Yunshan Networks
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package dbwriter

import (
	"github.com/deepflowio/deepflow/server/ingester/flow_tag"
	"github.com/deepflowio/deepflow/server/libs/ckdb"
	"github.com/deepflowio/deepflow/server/libs/datatype"
	"github.com/deepflowio/deepflow/server/libs/pool"
	"github.com/deepflowio/deepflow/server/libs/zerodoc"
)

const (
	DefaultPartition = ckdb.TimeFuncTwelveHour
)

type ExtMetrics struct {
	Timestamp uint32 // s
	MsgType   datatype.MessageType

	UniversalTag zerodoc.UniversalTag

	// in deepflow_system: table name
	// in ext_metrids: virtual_table_name
	VTableName string

	TagNames  []string
	TagValues []string

	MetricsFloatNames  []string
	MetricsFloatValues []float64
}

func (m *ExtMetrics) DatabaseName() string {
	if m.MsgType == datatype.MESSAGE_TYPE_DFSTATS {
		return DEEPFLOW_SYSTEM_DB
	} else {
		return EXT_METRICS_DB
	}
}

func (m *ExtMetrics) TableName() string {
	if m.MsgType == datatype.MESSAGE_TYPE_DFSTATS {
		return m.VTableName
	} else {
		return EXT_METRICS_TABLE
	}
}

func (m *ExtMetrics) VirtualTableName() string {
	if m.MsgType == datatype.MESSAGE_TYPE_DFSTATS {
		return ""
	} else {
		return m.VTableName
	}
}

// Note: The order of Write() must be consistent with the order of append() in Columns.
func (m *ExtMetrics) WriteBlock(block *ckdb.Block) {
	block.WriteDateTime(m.Timestamp)
	if m.MsgType != datatype.MESSAGE_TYPE_DFSTATS {
		m.UniversalTag.WriteBlock(block)
		block.Write(m.VTableName)
	}
	block.Write(
		m.TagNames,
		m.TagValues,
		m.MetricsFloatNames,
		m.MetricsFloatValues,
	)
}

// Note: The order of append() must be consistent with the order of Write() in WriteBlock.
func (m *ExtMetrics) Columns() []*ckdb.Column {
	columns := []*ckdb.Column{}

	columns = append(columns, ckdb.NewColumnWithGroupBy("time", ckdb.DateTime))
	if m.MsgType != datatype.MESSAGE_TYPE_DFSTATS {
		columns = zerodoc.GenUniversalTagColumns(columns)

		// FIXME: Currently there is no virtual_table_name column in the deepflow_system database,
		// but it will be unified in the future.
		columns = append(columns, ckdb.NewColumn("virtual_table_name", ckdb.LowCardinalityString).SetComment("虚拟表名k"))
	}
	columns = append(columns,
		ckdb.NewColumn("tag_names", ckdb.ArrayString).SetComment("额外的tag"),
		ckdb.NewColumn("tag_values", ckdb.ArrayString).SetComment("额外的tag对应的值"),
		ckdb.NewColumn("metrics_float_names", ckdb.ArrayString).SetComment("额外的float类型metrics"),
		ckdb.NewColumn("metrics_float_values", ckdb.ArrayFloat64).SetComment("额外的float metrics值"),
	)

	return columns
}

func (m *ExtMetrics) Release() {
	ReleaseExtMetrics(m)
}

func (m *ExtMetrics) GenCKTable(cluster, storagePolicy string, ttl int, coldStorage *ckdb.ColdStorage) *ckdb.Table {
	timeKey := "time"
	engine := ckdb.MergeTree

	// order key
	orderKeys := []string{}
	if m.MsgType != datatype.MESSAGE_TYPE_DFSTATS {
		// FIXME: Currently there is no virtual_table_name column in the deepflow_system database,
		// but it will be unified in the future.
		orderKeys = append(orderKeys, "virtual_table_name")

		// order key in universal tags
		orderKeys = append(orderKeys, "l3_epc_id")
		orderKeys = append(orderKeys, "ip4")
		orderKeys = append(orderKeys, "ip6")
	}
	orderKeys = append(orderKeys, timeKey)

	return &ckdb.Table{
		Database:        m.DatabaseName(),
		LocalName:       m.TableName() + ckdb.LOCAL_SUBFFIX,
		GlobalName:      m.TableName(),
		Columns:         m.Columns(),
		TimeKey:         timeKey,
		TTL:             ttl,
		PartitionFunc:   DefaultPartition,
		Engine:          engine,
		Cluster:         cluster,
		StoragePolicy:   storagePolicy,
		ColdStorage:     *coldStorage,
		OrderKeys:       orderKeys,
		PrimaryKeyCount: len(orderKeys),
	}
}

// Check if there is a TagName/TagValue/MetricsName not in fieldCache or fieldValueCache, and store the newly appeared item in cache.
func (m *ExtMetrics) GenerateNewFlowTags(cache *flow_tag.FlowTagCache) {
	tableName := m.TableName()
	if m.VirtualTableName() != "" {
		tableName = m.VirtualTableName()
	}

	// reset temporary buffers
	flowTagInfo := &cache.FlowTagInfoBuffer
	*flowTagInfo = flow_tag.FlowTagInfo{
		Table:   tableName,
		VpcId:   m.UniversalTag.L3EpcID,
		PodNsId: m.UniversalTag.PodNSID,
	}
	cache.Fields = cache.Fields[:0]
	cache.FieldValues = cache.FieldValues[:0]

	var tableExist, fieldNameExist, fieldValueExist bool
	var tableId, fieldNameId, fieldValueId uint32
	if tableId, tableExist = cache.TableIdMap[tableName]; !tableExist {
		tableId = uint32(len(cache.TableIdMap))
		cache.TableIdMap[tableName] = tableId
	}
	flowTagInfoKey := &cache.FlowTagInfoKeyBuffer
	flowTagInfoKey.Reset()
	flowTagInfoKey.SetTableId(tableId)
	flowTagInfoKey.SetVpcId(flowTagInfo.VpcId)
	flowTagInfoKey.SetPodNsId(flowTagInfo.PodNsId)

	// tags
	flowTagInfo.FieldType = flow_tag.FieldTag
	flowTagInfoKey.SetFieldType(flow_tag.FieldTag)
	for i, name := range m.TagNames {
		flowTagInfo.FieldName = name
		if fieldNameId, fieldNameExist = cache.FieldNameIdMap[name]; !fieldNameExist {
			fieldNameId = uint32(len(cache.FieldNameIdMap))
			cache.FieldNameIdMap[name] = fieldNameId
		}
		flowTagInfoKey.SetFieldNameId(fieldNameId)

		// tag + value
		flowTagInfo.FieldValue = m.TagValues[i]
		if fieldValueId, fieldValueExist = cache.FieldValueIdMap[flowTagInfo.FieldValue]; !fieldValueExist {
			fieldValueId = uint32(len(cache.FieldValueIdMap))
			cache.FieldValueIdMap[flowTagInfo.FieldValue] = fieldValueId
		}
		flowTagInfoKey.SetFieldValueId(fieldValueId)
		v1 := m.Timestamp
		if old := cache.FieldValueCache.AddOrGet(flowTagInfoKey.Low, flowTagInfoKey.High, &v1); old != nil {
			oldv, _ := old.(*uint32)
			if *oldv+cache.CacheFlushTimeout >= m.Timestamp {
				// If there is no new fieldValue, of course there will be no new field.
				// So we can just skip the rest of the process in the loop.
				continue
			} else {
				*oldv = m.Timestamp
			}
		}
		tagFieldValue := flow_tag.AcquireFlowTag()
		tagFieldValue.Timestamp = m.Timestamp
		tagFieldValue.FlowTagInfo = *flowTagInfo
		cache.FieldValues = append(cache.FieldValues, tagFieldValue)

		// only tag
		flowTagInfo.FieldValue = ""
		flowTagInfoKey.SetFieldValueId(0)
		v2 := m.Timestamp
		if old := cache.FieldCache.AddOrGet(flowTagInfoKey.Low, flowTagInfoKey.High, &v2); old != nil {
			oldv, _ := old.(*uint32)
			if *oldv+cache.CacheFlushTimeout >= m.Timestamp {
				continue
			} else {
				*oldv = m.Timestamp
			}
		}
		tagField := flow_tag.AcquireFlowTag()
		tagField.Timestamp = m.Timestamp
		tagField.FlowTagInfo = *flowTagInfo
		cache.Fields = append(cache.Fields, tagField)
	}

	// metrics
	flowTagInfo.FieldType = flow_tag.FieldMetrics
	flowTagInfoKey.SetFieldType(flow_tag.FieldMetrics)
	flowTagInfo.FieldValue = ""
	flowTagInfoKey.SetFieldValueId(0)
	for _, name := range m.MetricsFloatNames {
		flowTagInfo.FieldName = name
		if fieldNameId, fieldNameExist = cache.FieldNameIdMap[name]; !fieldNameExist {
			fieldNameId = uint32(len(cache.FieldNameIdMap))
			cache.FieldNameIdMap[name] = fieldNameId
		}
		flowTagInfoKey.SetFieldNameId(fieldNameId)

		v := m.Timestamp
		if old := cache.FieldCache.AddOrGet(flowTagInfoKey.Low, flowTagInfoKey.High, &v); old != nil {
			oldv, _ := old.(*uint32)
			if *oldv+cache.CacheFlushTimeout >= m.Timestamp {
				continue
			} else {
				*oldv = m.Timestamp
			}
		}
		tagField := flow_tag.AcquireFlowTag()
		tagField.Timestamp = m.Timestamp
		tagField.FlowTagInfo = *flowTagInfo
		cache.Fields = append(cache.Fields, tagField)
	}
}

var extMetricsPool = pool.NewLockFreePool(func() interface{} {
	return &ExtMetrics{}
})

func AcquireExtMetrics() *ExtMetrics {
	return extMetricsPool.Get().(*ExtMetrics)
}

var emptyUniversalTag = zerodoc.UniversalTag{}

func ReleaseExtMetrics(m *ExtMetrics) {
	m.UniversalTag = emptyUniversalTag
	m.TagNames = m.TagNames[:0]
	m.TagValues = m.TagValues[:0]
	m.MetricsFloatNames = m.MetricsFloatNames[:0]
	m.MetricsFloatValues = m.MetricsFloatValues[:0]
	extMetricsPool.Put(m)
}
