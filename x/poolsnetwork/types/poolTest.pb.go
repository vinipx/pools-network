// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: poolsnetwork/v1beta/poolTest.proto

package types

import (
	fmt "fmt"
	github_com_cosmos_cosmos_sdk_types "github.com/cosmos/cosmos-sdk/types"
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

type MsgPoolTest struct {
	Id           string                                        `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Creator      github_com_cosmos_cosmos_sdk_types.AccAddress `protobuf:"bytes,2,opt,name=creator,proto3,casttype=github.com/cosmos/cosmos-sdk/types.AccAddress" json:"creator,omitempty"`
	PoolId       string                                        `protobuf:"bytes,3,opt,name=pool_id,json=poolId,proto3" json:"pool_id,omitempty"`
	PubKey       string                                        `protobuf:"bytes,4,opt,name=pubKey,proto3" json:"pubKey,omitempty"`
	Slashed      bool                                          `protobuf:"varint,5,opt,name=slashed,proto3" json:"slashed,omitempty"`
	Exited       bool                                          `protobuf:"varint,6,opt,name=exited,proto3" json:"exited,omitempty"`
	SsvCommittee string                                        `protobuf:"bytes,7,opt,name=ssvCommittee,proto3" json:"ssvCommittee,omitempty"`
}

func (m *MsgPoolTest) Reset()         { *m = MsgPoolTest{} }
func (m *MsgPoolTest) String() string { return proto.CompactTextString(m) }
func (*MsgPoolTest) ProtoMessage()    {}
func (*MsgPoolTest) Descriptor() ([]byte, []int) {
	return fileDescriptor_4239f42a3a24866a, []int{0}
}
func (m *MsgPoolTest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgPoolTest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgPoolTest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgPoolTest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgPoolTest.Merge(m, src)
}
func (m *MsgPoolTest) XXX_Size() int {
	return m.Size()
}
func (m *MsgPoolTest) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgPoolTest.DiscardUnknown(m)
}

var xxx_messageInfo_MsgPoolTest proto.InternalMessageInfo

func (m *MsgPoolTest) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *MsgPoolTest) GetCreator() github_com_cosmos_cosmos_sdk_types.AccAddress {
	if m != nil {
		return m.Creator
	}
	return nil
}

func (m *MsgPoolTest) GetPoolId() string {
	if m != nil {
		return m.PoolId
	}
	return ""
}

func (m *MsgPoolTest) GetPubKey() string {
	if m != nil {
		return m.PubKey
	}
	return ""
}

func (m *MsgPoolTest) GetSlashed() bool {
	if m != nil {
		return m.Slashed
	}
	return false
}

func (m *MsgPoolTest) GetExited() bool {
	if m != nil {
		return m.Exited
	}
	return false
}

func (m *MsgPoolTest) GetSsvCommittee() string {
	if m != nil {
		return m.SsvCommittee
	}
	return ""
}

func init() {
	proto.RegisterType((*MsgPoolTest)(nil), "poolsnetwork.poolsnetwork.v1beta1.MsgPoolTest")
}

func init() {
	proto.RegisterFile("poolsnetwork/v1beta/poolTest.proto", fileDescriptor_4239f42a3a24866a)
}

var fileDescriptor_4239f42a3a24866a = []byte{
	// 316 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x54, 0x91, 0xc1, 0x4e, 0x3a, 0x31,
	0x10, 0xc6, 0x29, 0xff, 0xbf, 0xbb, 0x5a, 0x89, 0x87, 0xc6, 0x68, 0xe3, 0x61, 0x45, 0x4e, 0x5c,
	0xd8, 0x0d, 0x31, 0x3e, 0x00, 0x78, 0x32, 0xc4, 0x68, 0x88, 0x27, 0x2f, 0x66, 0x77, 0x3b, 0x59,
	0x1a, 0x58, 0xa7, 0xd9, 0x29, 0x08, 0x6f, 0xe1, 0x63, 0x79, 0xe4, 0xe8, 0xc9, 0x18, 0x78, 0x02,
	0xaf, 0x9e, 0xcc, 0x96, 0x25, 0x81, 0x53, 0xfb, 0x9b, 0xce, 0xf7, 0xb5, 0xf3, 0x95, 0xb7, 0x0c,
	0xe2, 0x84, 0x5e, 0xc1, 0xbe, 0x61, 0x31, 0x8e, 0x66, 0xdd, 0x04, 0x6c, 0x1c, 0x95, 0xb5, 0x27,
	0x20, 0x1b, 0x9a, 0x02, 0x2d, 0x8a, 0xab, 0xdd, 0x9e, 0x70, 0x0f, 0x36, 0x82, 0xee, 0xc5, 0x69,
	0x86, 0x19, 0xba, 0xee, 0xa8, 0xdc, 0x6d, 0x84, 0xad, 0x1f, 0xc6, 0x8f, 0xef, 0x29, 0x7b, 0xac,
	0xec, 0xc4, 0x09, 0xaf, 0x6b, 0x25, 0x59, 0x93, 0xb5, 0x8f, 0x86, 0x75, 0xad, 0xc4, 0x80, 0xfb,
	0x69, 0x01, 0xb1, 0xc5, 0x42, 0xd6, 0x9b, 0xac, 0xdd, 0xe8, 0x77, 0x7f, 0xbf, 0x2e, 0x3b, 0x99,
	0xb6, 0xa3, 0x69, 0x12, 0xa6, 0x98, 0x47, 0x29, 0x52, 0x8e, 0x54, 0x2d, 0x1d, 0x52, 0xe3, 0xc8,
	0x2e, 0x0c, 0x50, 0xd8, 0x4b, 0xd3, 0x9e, 0x52, 0x05, 0x10, 0x0d, 0xb7, 0x0e, 0xe2, 0x9c, 0xfb,
	0xe5, 0xd3, 0x5e, 0xb4, 0x92, 0xff, 0xdc, 0x0d, 0x5e, 0x89, 0x77, 0x4a, 0x9c, 0x71, 0xcf, 0x4c,
	0x93, 0x01, 0x2c, 0xe4, 0xff, 0xaa, 0xee, 0x48, 0x48, 0xee, 0xd3, 0x24, 0xa6, 0x11, 0x28, 0x79,
	0xd0, 0x64, 0xed, 0xc3, 0xe1, 0x16, 0x4b, 0x05, 0xcc, 0xb5, 0x05, 0x25, 0x3d, 0x77, 0x50, 0x91,
	0x68, 0xf1, 0x06, 0xd1, 0xec, 0x16, 0xf3, 0x5c, 0x5b, 0x0b, 0x20, 0x7d, 0xe7, 0xb7, 0x57, 0xeb,
	0x3f, 0x7c, 0xac, 0x02, 0xb6, 0x5c, 0x05, 0xec, 0x7b, 0x15, 0xb0, 0xf7, 0x75, 0x50, 0x5b, 0xae,
	0x83, 0xda, 0xe7, 0x3a, 0xa8, 0x3d, 0xdf, 0xec, 0x0c, 0x96, 0x4c, 0x70, 0x1e, 0x1b, 0xe3, 0x92,
	0xa6, 0xce, 0x36, 0xfe, 0x79, 0xb4, 0xf7, 0x1b, 0x6e, 0xd6, 0xc4, 0x73, 0x59, 0x5e, 0xff, 0x05,
	0x00, 0x00, 0xff, 0xff, 0x69, 0xe7, 0x0c, 0xd3, 0xaa, 0x01, 0x00, 0x00,
}

func (m *MsgPoolTest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgPoolTest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgPoolTest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.SsvCommittee) > 0 {
		i -= len(m.SsvCommittee)
		copy(dAtA[i:], m.SsvCommittee)
		i = encodeVarintPoolTest(dAtA, i, uint64(len(m.SsvCommittee)))
		i--
		dAtA[i] = 0x3a
	}
	if m.Exited {
		i--
		if m.Exited {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i--
		dAtA[i] = 0x30
	}
	if m.Slashed {
		i--
		if m.Slashed {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i--
		dAtA[i] = 0x28
	}
	if len(m.PubKey) > 0 {
		i -= len(m.PubKey)
		copy(dAtA[i:], m.PubKey)
		i = encodeVarintPoolTest(dAtA, i, uint64(len(m.PubKey)))
		i--
		dAtA[i] = 0x22
	}
	if len(m.PoolId) > 0 {
		i -= len(m.PoolId)
		copy(dAtA[i:], m.PoolId)
		i = encodeVarintPoolTest(dAtA, i, uint64(len(m.PoolId)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.Creator) > 0 {
		i -= len(m.Creator)
		copy(dAtA[i:], m.Creator)
		i = encodeVarintPoolTest(dAtA, i, uint64(len(m.Creator)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Id) > 0 {
		i -= len(m.Id)
		copy(dAtA[i:], m.Id)
		i = encodeVarintPoolTest(dAtA, i, uint64(len(m.Id)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintPoolTest(dAtA []byte, offset int, v uint64) int {
	offset -= sovPoolTest(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *MsgPoolTest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Id)
	if l > 0 {
		n += 1 + l + sovPoolTest(uint64(l))
	}
	l = len(m.Creator)
	if l > 0 {
		n += 1 + l + sovPoolTest(uint64(l))
	}
	l = len(m.PoolId)
	if l > 0 {
		n += 1 + l + sovPoolTest(uint64(l))
	}
	l = len(m.PubKey)
	if l > 0 {
		n += 1 + l + sovPoolTest(uint64(l))
	}
	if m.Slashed {
		n += 2
	}
	if m.Exited {
		n += 2
	}
	l = len(m.SsvCommittee)
	if l > 0 {
		n += 1 + l + sovPoolTest(uint64(l))
	}
	return n
}

func sovPoolTest(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozPoolTest(x uint64) (n int) {
	return sovPoolTest(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *MsgPoolTest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowPoolTest
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
			return fmt.Errorf("proto: MsgPoolTest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgPoolTest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Id", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPoolTest
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
				return ErrInvalidLengthPoolTest
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthPoolTest
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Id = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Creator", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPoolTest
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthPoolTest
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthPoolTest
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Creator = append(m.Creator[:0], dAtA[iNdEx:postIndex]...)
			if m.Creator == nil {
				m.Creator = []byte{}
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field PoolId", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPoolTest
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
				return ErrInvalidLengthPoolTest
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthPoolTest
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.PoolId = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field PubKey", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPoolTest
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
				return ErrInvalidLengthPoolTest
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthPoolTest
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.PubKey = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 5:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Slashed", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPoolTest
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.Slashed = bool(v != 0)
		case 6:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Exited", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPoolTest
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.Exited = bool(v != 0)
		case 7:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field SsvCommittee", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPoolTest
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
				return ErrInvalidLengthPoolTest
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthPoolTest
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.SsvCommittee = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipPoolTest(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthPoolTest
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthPoolTest
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipPoolTest(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowPoolTest
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
					return 0, ErrIntOverflowPoolTest
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
					return 0, ErrIntOverflowPoolTest
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
				return 0, ErrInvalidLengthPoolTest
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupPoolTest
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthPoolTest
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthPoolTest        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowPoolTest          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupPoolTest = fmt.Errorf("proto: unexpected end of group")
)
