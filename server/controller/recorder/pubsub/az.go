/**
 * Copyright (c) 2023 Yunshan Networks
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

package pubsub

import (
	"github.com/deepflowio/deepflow/server/controller/common"
	"github.com/deepflowio/deepflow/server/controller/recorder/pubsub/message"
)

type AZ struct {
	SingleResourcePubSubComponent[*message.AZAdd, message.AZAdd, *message.AZUpdate, message.AZUpdate, *message.AZFieldsUpdate, message.AZFieldsUpdate, *message.AZDelete, message.AZDelete]
}

func NewAZ() *AZ {
	return &AZ{
		SingleResourcePubSubComponent[*message.AZAdd, message.AZAdd, *message.AZUpdate, message.AZUpdate, *message.AZFieldsUpdate, message.AZFieldsUpdate, *message.AZDelete, message.AZDelete]{
			PubSubComponent: newPubSubComponent(PubSubTypeAZ),
			resourceType:    common.RESOURCE_TYPE_AZ_EN,
		},
	}
}
