// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: stats.proto

package pb

import (
	encoding_binary "encoding/binary"
	fmt "fmt"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	io "io"
	math "math"
	math_bits "math/bits"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

type Stats struct {
	Timestamp            uint64    `protobuf:"varint,1,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
	Name                 string    `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	TagNames             []string  `protobuf:"bytes,3,rep,name=tag_names,json=tagNames,proto3" json:"tag_names,omitempty"`
	TagValues            []string  `protobuf:"bytes,4,rep,name=tag_values,json=tagValues,proto3" json:"tag_values,omitempty"`
	MetricsFloatNames    []string  `protobuf:"bytes,7,rep,name=metrics_float_names,json=metricsFloatNames,proto3" json:"metrics_float_names,omitempty"`
	MetricsFloatValues   []float64 `protobuf:"fixed64,8,rep,packed,name=metrics_float_values,json=metricsFloatValues,proto3" json:"metrics_float_values,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *Stats) Reset()         { *m = Stats{} }
func (m *Stats) String() string { return proto.CompactTextString(m) }
func (*Stats) ProtoMessage()    {}
func (*Stats) Descriptor() ([]byte, []int) {
	return fileDescriptor_b4756a0aec8b9d44, []int{0}
}
func (m *Stats) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Stats) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Stats.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Stats) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Stats.Merge(m, src)
}
func (m *Stats) XXX_Size() int {
	return m.Size()
}
func (m *Stats) XXX_DiscardUnknown() {
	xxx_messageInfo_Stats.DiscardUnknown(m)
}

var xxx_messageInfo_Stats proto.InternalMessageInfo

func (m *Stats) GetTimestamp() uint64 {
	if m != nil {
		return m.Timestamp
	}
	return 0
}

func (m *Stats) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Stats) GetTagNames() []string {
	if m != nil {
		return m.TagNames
	}
	return nil
}

func (m *Stats) GetTagValues() []string {
	if m != nil {
		return m.TagValues
	}
	return nil
}

func (m *Stats) GetMetricsFloatNames() []string {
	if m != nil {
		return m.MetricsFloatNames
	}
	return nil
}

func (m *Stats) GetMetricsFloatValues() []float64 {
	if m != nil {
		return m.MetricsFloatValues
	}
	return nil
}

func init() {
	proto.RegisterType((*Stats)(nil), "stats.Stats")
}

func init() { proto.RegisterFile("stats.proto", fileDescriptor_b4756a0aec8b9d44) }

var fileDescriptor_b4756a0aec8b9d44 = []byte{
	// 235 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x54, 0x8f, 0xbd, 0x4e, 0xc4, 0x30,
	0x10, 0x84, 0xb5, 0x97, 0x1c, 0x5c, 0x96, 0x06, 0x16, 0x8a, 0x88, 0x9f, 0x28, 0xa2, 0x4a, 0x43,
	0x0e, 0x89, 0x37, 0xa0, 0xa0, 0xa4, 0x08, 0x12, 0x05, 0xcd, 0xc9, 0x3e, 0xf9, 0x4c, 0xa4, 0x33,
	0x8e, 0xe2, 0x0d, 0xcf, 0x48, 0x49, 0x47, 0x8b, 0xf2, 0x24, 0xc8, 0x9b, 0x48, 0x40, 0x37, 0xb3,
	0xdf, 0xcc, 0x58, 0xc6, 0xa3, 0xc0, 0x8a, 0x43, 0xdd, 0xf5, 0x9e, 0x3d, 0x2d, 0xc5, 0x9c, 0xdf,
	0xd8, 0x96, 0x5f, 0x07, 0x5d, 0x6f, 0xbd, 0x5b, 0x5b, 0x6f, 0xfd, 0x5a, 0xa8, 0x1e, 0x76, 0xe2,
	0xc4, 0x88, 0x9a, 0x5a, 0xd7, 0x5f, 0x80, 0xcb, 0xa7, 0x58, 0xa4, 0x4b, 0xcc, 0xb8, 0x75, 0x26,
	0xb0, 0x72, 0x5d, 0x0e, 0x25, 0x54, 0x69, 0xf3, 0x7b, 0x20, 0xc2, 0xf4, 0x4d, 0x39, 0x93, 0x2f,
	0x4a, 0xa8, 0xb2, 0x46, 0x34, 0x5d, 0x60, 0xc6, 0xca, 0x6e, 0xa2, 0x0e, 0x79, 0x52, 0x26, 0x55,
	0xd6, 0xac, 0x58, 0xd9, 0xc7, 0xe8, 0xe9, 0x0a, 0x31, 0xc2, 0x77, 0xb5, 0x1f, 0x4c, 0xc8, 0x53,
	0xa1, 0x31, 0xfe, 0x2c, 0x07, 0xaa, 0xf1, 0xd4, 0x19, 0xee, 0xdb, 0x6d, 0xd8, 0xec, 0xf6, 0x5e,
	0xf1, 0xbc, 0x72, 0x28, 0xb9, 0x93, 0x19, 0x3d, 0x44, 0x32, 0xcd, 0xdd, 0xe2, 0xd9, 0xff, 0xfc,
	0x3c, 0xbc, 0x2a, 0x93, 0x0a, 0x1a, 0xfa, 0x5b, 0x98, 0x5e, 0xb8, 0x3f, 0xfe, 0x18, 0x0b, 0xf8,
	0x1c, 0x0b, 0xf8, 0x1e, 0x0b, 0x78, 0x59, 0x74, 0x5a, 0x1f, 0xc8, 0x97, 0xef, 0x7e, 0x02, 0x00,
	0x00, 0xff, 0xff, 0x5d, 0x89, 0x05, 0xf2, 0x37, 0x01, 0x00, 0x00,
}

func (m *Stats) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Stats) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Stats) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	if len(m.MetricsFloatValues) > 0 {
		for iNdEx := len(m.MetricsFloatValues) - 1; iNdEx >= 0; iNdEx-- {
			f1 := math.Float64bits(float64(m.MetricsFloatValues[iNdEx]))
			i -= 8
			encoding_binary.LittleEndian.PutUint64(dAtA[i:], uint64(f1))
		}
		i = encodeVarintStats(dAtA, i, uint64(len(m.MetricsFloatValues)*8))
		i--
		dAtA[i] = 0x42
	}
	if len(m.MetricsFloatNames) > 0 {
		for iNdEx := len(m.MetricsFloatNames) - 1; iNdEx >= 0; iNdEx-- {
			i -= len(m.MetricsFloatNames[iNdEx])
			copy(dAtA[i:], m.MetricsFloatNames[iNdEx])
			i = encodeVarintStats(dAtA, i, uint64(len(m.MetricsFloatNames[iNdEx])))
			i--
			dAtA[i] = 0x3a
		}
	}
	if len(m.TagValues) > 0 {
		for iNdEx := len(m.TagValues) - 1; iNdEx >= 0; iNdEx-- {
			i -= len(m.TagValues[iNdEx])
			copy(dAtA[i:], m.TagValues[iNdEx])
			i = encodeVarintStats(dAtA, i, uint64(len(m.TagValues[iNdEx])))
			i--
			dAtA[i] = 0x22
		}
	}
	if len(m.TagNames) > 0 {
		for iNdEx := len(m.TagNames) - 1; iNdEx >= 0; iNdEx-- {
			i -= len(m.TagNames[iNdEx])
			copy(dAtA[i:], m.TagNames[iNdEx])
			i = encodeVarintStats(dAtA, i, uint64(len(m.TagNames[iNdEx])))
			i--
			dAtA[i] = 0x1a
		}
	}
	if len(m.Name) > 0 {
		i -= len(m.Name)
		copy(dAtA[i:], m.Name)
		i = encodeVarintStats(dAtA, i, uint64(len(m.Name)))
		i--
		dAtA[i] = 0x12
	}
	if m.Timestamp != 0 {
		i = encodeVarintStats(dAtA, i, uint64(m.Timestamp))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func encodeVarintStats(dAtA []byte, offset int, v uint64) int {
	offset -= sovStats(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *Stats) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Timestamp != 0 {
		n += 1 + sovStats(uint64(m.Timestamp))
	}
	l = len(m.Name)
	if l > 0 {
		n += 1 + l + sovStats(uint64(l))
	}
	if len(m.TagNames) > 0 {
		for _, s := range m.TagNames {
			l = len(s)
			n += 1 + l + sovStats(uint64(l))
		}
	}
	if len(m.TagValues) > 0 {
		for _, s := range m.TagValues {
			l = len(s)
			n += 1 + l + sovStats(uint64(l))
		}
	}
	if len(m.MetricsFloatNames) > 0 {
		for _, s := range m.MetricsFloatNames {
			l = len(s)
			n += 1 + l + sovStats(uint64(l))
		}
	}
	if len(m.MetricsFloatValues) > 0 {
		n += 1 + sovStats(uint64(len(m.MetricsFloatValues)*8)) + len(m.MetricsFloatValues)*8
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func sovStats(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozStats(x uint64) (n int) {
	return sovStats(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Stats) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowStats
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: Stats: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Stats: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Timestamp", wireType)
			}
			m.Timestamp = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowStats
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Timestamp |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Name", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowStats
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthStats
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthStats
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Name = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field TagNames", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowStats
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthStats
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthStats
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.TagNames = append(m.TagNames, string(dAtA[iNdEx:postIndex]))
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field TagValues", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowStats
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthStats
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthStats
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.TagValues = append(m.TagValues, string(dAtA[iNdEx:postIndex]))
			iNdEx = postIndex
		case 7:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field MetricsFloatNames", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowStats
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthStats
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthStats
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.MetricsFloatNames = append(m.MetricsFloatNames, string(dAtA[iNdEx:postIndex]))
			iNdEx = postIndex
		case 8:
			if wireType == 1 {
				var v uint64
				if (iNdEx + 8) > l {
					return io.ErrUnexpectedEOF
				}
				v = uint64(encoding_binary.LittleEndian.Uint64(dAtA[iNdEx:]))
				iNdEx += 8
				v2 := float64(math.Float64frombits(v))
				m.MetricsFloatValues = append(m.MetricsFloatValues, v2)
			} else if wireType == 2 {
				var packedLen int
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return ErrIntOverflowStats
					}
					if iNdEx >= l {
						return io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					packedLen |= int(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				if packedLen < 0 {
					return ErrInvalidLengthStats
				}
				postIndex := iNdEx + packedLen
				if postIndex < 0 {
					return ErrInvalidLengthStats
				}
				if postIndex > l {
					return io.ErrUnexpectedEOF
				}
				var elementCount int
				elementCount = packedLen / 8
				if elementCount != 0 && len(m.MetricsFloatValues) == 0 {
					m.MetricsFloatValues = make([]float64, 0, elementCount)
				}
				for iNdEx < postIndex {
					var v uint64
					if (iNdEx + 8) > l {
						return io.ErrUnexpectedEOF
					}
					v = uint64(encoding_binary.LittleEndian.Uint64(dAtA[iNdEx:]))
					iNdEx += 8
					v2 := float64(math.Float64frombits(v))
					m.MetricsFloatValues = append(m.MetricsFloatValues, v2)
				}
			} else {
				return fmt.Errorf("proto: wrong wireType = %d for field MetricsFloatValues", wireType)
			}
		default:
			iNdEx = preIndex
			skippy, err := skipStats(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthStats
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipStats(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowStats
			}
			if iNdEx >= l {
				return 0, io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		wireType := int(wire & 0x7)
		switch wireType {
		case 0:
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowStats
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
		case 1:
			iNdEx += 8
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowStats
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				length |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if length < 0 {
				return 0, ErrInvalidLengthStats
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupStats
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthStats
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthStats        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowStats          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupStats = fmt.Errorf("proto: unexpected end of group")
)