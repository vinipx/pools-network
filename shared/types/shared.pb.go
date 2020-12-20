// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: shared/shared.proto

package types

import (
	fmt "fmt"
	_ "github.com/cosmos/cosmos-sdk/codec/types"
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

// Operator represents a pools network operator, not to be confused with validator which is an entity local to the Tendermint protocol.
// An Operator has the responsibility of executing various tasks within the pools network, post collateral and so on.
type Operator struct {
	Id              uint64            `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	EthereumAddress EthereumAddress   `protobuf:"bytes,2,opt,name=ethereum_address,json=ethereumAddress,proto3,customtype=EthereumAddress" json:"ethereum_address"`
	ConsensusPubkey *ConsensusAddress `protobuf:"bytes,3,opt,name=consensus_pubkey,json=consensusPubkey,proto3,customtype=ConsensusAddress" json:"consensus_pubkey,omitempty"`
	EthStake        uint64            `protobuf:"varint,4,opt,name=eth_stake,json=ethStake,proto3" json:"eth_stake,omitempty"`
	CdtBalance      uint64            `protobuf:"varint,5,opt,name=cdt_balance,json=cdtBalance,proto3" json:"cdt_balance,omitempty"`
}

func (m *Operator) Reset()         { *m = Operator{} }
func (m *Operator) String() string { return proto.CompactTextString(m) }
func (*Operator) ProtoMessage()    {}
func (*Operator) Descriptor() ([]byte, []int) {
	return fileDescriptor_9301c8e954a55bc6, []int{0}
}
func (m *Operator) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Operator) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Operator.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Operator) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Operator.Merge(m, src)
}
func (m *Operator) XXX_Size() int {
	return m.Size()
}
func (m *Operator) XXX_DiscardUnknown() {
	xxx_messageInfo_Operator.DiscardUnknown(m)
}

var xxx_messageInfo_Operator proto.InternalMessageInfo

func (m *Operator) GetId() uint64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *Operator) GetEthStake() uint64 {
	if m != nil {
		return m.EthStake
	}
	return 0
}

func (m *Operator) GetCdtBalance() uint64 {
	if m != nil {
		return m.CdtBalance
	}
	return 0
}

type Pool struct {
	Id         uint64   `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Eth2Pubkey []byte   `protobuf:"bytes,2,opt,name=eth2_pubkey,json=eth2Pubkey,proto3" json:"eth2_pubkey,omitempty"`
	Balance    uint64   `protobuf:"varint,3,opt,name=balance,proto3" json:"balance,omitempty"`
	Exited     bool     `protobuf:"varint,4,opt,name=exited,proto3" json:"exited,omitempty"`
	Slashed    bool     `protobuf:"varint,5,opt,name=slashed,proto3" json:"slashed,omitempty"`
	Committee  []uint64 `protobuf:"varint,6,rep,packed,name=committee,proto3" json:"committee,omitempty"`
}

func (m *Pool) Reset()         { *m = Pool{} }
func (m *Pool) String() string { return proto.CompactTextString(m) }
func (*Pool) ProtoMessage()    {}
func (*Pool) Descriptor() ([]byte, []int) {
	return fileDescriptor_9301c8e954a55bc6, []int{1}
}
func (m *Pool) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Pool) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Pool.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Pool) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Pool.Merge(m, src)
}
func (m *Pool) XXX_Size() int {
	return m.Size()
}
func (m *Pool) XXX_DiscardUnknown() {
	xxx_messageInfo_Pool.DiscardUnknown(m)
}

var xxx_messageInfo_Pool proto.InternalMessageInfo

func (m *Pool) GetId() uint64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *Pool) GetEth2Pubkey() []byte {
	if m != nil {
		return m.Eth2Pubkey
	}
	return nil
}

func (m *Pool) GetBalance() uint64 {
	if m != nil {
		return m.Balance
	}
	return 0
}

func (m *Pool) GetExited() bool {
	if m != nil {
		return m.Exited
	}
	return false
}

func (m *Pool) GetSlashed() bool {
	if m != nil {
		return m.Slashed
	}
	return false
}

func (m *Pool) GetCommittee() []uint64 {
	if m != nil {
		return m.Committee
	}
	return nil
}

func init() {
	proto.RegisterType((*Operator)(nil), "poolsnetwork.v1beta1.Operator")
	proto.RegisterType((*Pool)(nil), "poolsnetwork.v1beta1.Pool")
}

func init() { proto.RegisterFile("shared/shared.proto", fileDescriptor_9301c8e954a55bc6) }

var fileDescriptor_9301c8e954a55bc6 = []byte{
	// 396 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x64, 0x52, 0xcd, 0x6e, 0xd3, 0x40,
	0x10, 0xce, 0x26, 0x21, 0xb8, 0x5b, 0x44, 0xa3, 0x25, 0x02, 0x53, 0x90, 0x1d, 0xf5, 0x94, 0x4b,
	0x63, 0x15, 0x1e, 0x00, 0x61, 0x84, 0x38, 0x52, 0x99, 0x1b, 0x17, 0x6b, 0xed, 0x1d, 0x6c, 0x2b,
	0xb6, 0xc7, 0xf2, 0x8e, 0xa1, 0x79, 0x0b, 0x5e, 0x81, 0xb7, 0xe9, 0xb1, 0x47, 0x54, 0x89, 0x08,
	0x25, 0x2f, 0x82, 0xbc, 0xb6, 0x83, 0x10, 0xa7, 0xdd, 0xef, 0x67, 0x66, 0xbf, 0xd1, 0x2c, 0x7f,
	0xa2, 0x53, 0x59, 0x83, 0xf2, 0xba, 0x63, 0x5d, 0xd5, 0x48, 0x28, 0x16, 0x15, 0x62, 0xae, 0x4b,
	0xa0, 0x6f, 0x58, 0x6f, 0xd6, 0x5f, 0xaf, 0x22, 0x20, 0x79, 0x75, 0xbe, 0x48, 0x30, 0x41, 0x63,
	0xf0, 0xda, 0x5b, 0xe7, 0x3d, 0x7f, 0x9e, 0x20, 0x26, 0x39, 0x78, 0x06, 0x45, 0xcd, 0x17, 0x4f,
	0x96, 0xdb, 0x4e, 0xba, 0xf8, 0xc5, 0xb8, 0xf5, 0xb1, 0x82, 0x5a, 0x12, 0xd6, 0xe2, 0x31, 0x1f,
	0x67, 0xca, 0x66, 0x4b, 0xb6, 0x9a, 0x06, 0xe3, 0x4c, 0x09, 0x9f, 0xcf, 0x81, 0x52, 0xa8, 0xa1,
	0x29, 0x42, 0xa9, 0x54, 0x0d, 0x5a, 0xdb, 0xe3, 0x25, 0x5b, 0x3d, 0xf2, 0x9f, 0xdd, 0xee, 0xdc,
	0xd1, 0xfd, 0xce, 0x3d, 0x7b, 0xdf, 0xeb, 0x6f, 0x3b, 0x39, 0x38, 0x83, 0x7f, 0x09, 0xf1, 0x86,
	0xcf, 0x63, 0x2c, 0x35, 0x94, 0xba, 0xd1, 0x61, 0xd5, 0x44, 0x1b, 0xd8, 0xda, 0x13, 0xd3, 0x63,
	0x71, 0xbf, 0x73, 0xe7, 0xef, 0x06, 0xed, 0xd8, 0xe0, 0xe8, 0xbe, 0x36, 0x66, 0xf1, 0x82, 0x9f,
	0x00, 0xa5, 0xa1, 0x26, 0xb9, 0x01, 0x7b, 0x6a, 0xb2, 0x59, 0x40, 0xe9, 0xa7, 0x16, 0x0b, 0x97,
	0x9f, 0xc6, 0x8a, 0xc2, 0x48, 0xe6, 0xb2, 0x8c, 0xc1, 0x7e, 0x60, 0x64, 0x1e, 0x2b, 0xf2, 0x3b,
	0xe6, 0xe2, 0x07, 0xe3, 0xd3, 0x6b, 0xc4, 0xfc, 0xbf, 0xd9, 0x5c, 0x7e, 0x0a, 0x94, 0xbe, 0x1a,
	0x22, 0x99, 0xb1, 0x02, 0xde, 0x52, 0xfd, 0xbb, 0x36, 0x7f, 0x38, 0xb4, 0x9d, 0x98, 0xaa, 0x01,
	0x8a, 0xa7, 0x7c, 0x06, 0x37, 0x19, 0x81, 0x32, 0x71, 0xac, 0xa0, 0x47, 0x6d, 0x85, 0xce, 0xa5,
	0x4e, 0x41, 0x99, 0x20, 0x56, 0x30, 0x40, 0xf1, 0x92, 0x9f, 0xc4, 0x58, 0x14, 0x19, 0x11, 0x80,
	0x3d, 0x5b, 0x4e, 0x56, 0xd3, 0xe0, 0x2f, 0xe1, 0x7f, 0xb8, 0xdd, 0x3b, 0xec, 0x6e, 0xef, 0xb0,
	0xdf, 0x7b, 0x87, 0x7d, 0x3f, 0x38, 0xa3, 0xbb, 0x83, 0x33, 0xfa, 0x79, 0x70, 0x46, 0x9f, 0x2f,
	0x93, 0x8c, 0xd2, 0x26, 0x5a, 0xc7, 0x58, 0x78, 0x51, 0x8e, 0x37, 0xb2, 0xaa, 0x3c, 0xb3, 0xf7,
	0xcb, 0x7e, 0xf1, 0xfd, 0x9f, 0xf0, 0x68, 0x5b, 0x81, 0x8e, 0x66, 0x66, 0xa7, 0xaf, 0xff, 0x04,
	0x00, 0x00, 0xff, 0xff, 0x4c, 0x9a, 0x00, 0x69, 0x31, 0x02, 0x00, 0x00,
}

func (m *Operator) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Operator) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Operator) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.CdtBalance != 0 {
		i = encodeVarintShared(dAtA, i, uint64(m.CdtBalance))
		i--
		dAtA[i] = 0x28
	}
	if m.EthStake != 0 {
		i = encodeVarintShared(dAtA, i, uint64(m.EthStake))
		i--
		dAtA[i] = 0x20
	}
	if m.ConsensusPubkey != nil {
		{
			size := m.ConsensusPubkey.Size()
			i -= size
			if _, err := m.ConsensusPubkey.MarshalTo(dAtA[i:]); err != nil {
				return 0, err
			}
			i = encodeVarintShared(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x1a
	}
	{
		size := m.EthereumAddress.Size()
		i -= size
		if _, err := m.EthereumAddress.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintShared(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x12
	if m.Id != 0 {
		i = encodeVarintShared(dAtA, i, uint64(m.Id))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *Pool) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Pool) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Pool) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Committee) > 0 {
		dAtA2 := make([]byte, len(m.Committee)*10)
		var j1 int
		for _, num := range m.Committee {
			for num >= 1<<7 {
				dAtA2[j1] = uint8(uint64(num)&0x7f | 0x80)
				num >>= 7
				j1++
			}
			dAtA2[j1] = uint8(num)
			j1++
		}
		i -= j1
		copy(dAtA[i:], dAtA2[:j1])
		i = encodeVarintShared(dAtA, i, uint64(j1))
		i--
		dAtA[i] = 0x32
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
	if m.Exited {
		i--
		if m.Exited {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i--
		dAtA[i] = 0x20
	}
	if m.Balance != 0 {
		i = encodeVarintShared(dAtA, i, uint64(m.Balance))
		i--
		dAtA[i] = 0x18
	}
	if len(m.Eth2Pubkey) > 0 {
		i -= len(m.Eth2Pubkey)
		copy(dAtA[i:], m.Eth2Pubkey)
		i = encodeVarintShared(dAtA, i, uint64(len(m.Eth2Pubkey)))
		i--
		dAtA[i] = 0x12
	}
	if m.Id != 0 {
		i = encodeVarintShared(dAtA, i, uint64(m.Id))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func encodeVarintShared(dAtA []byte, offset int, v uint64) int {
	offset -= sovShared(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *Operator) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Id != 0 {
		n += 1 + sovShared(uint64(m.Id))
	}
	l = m.EthereumAddress.Size()
	n += 1 + l + sovShared(uint64(l))
	if m.ConsensusPubkey != nil {
		l = m.ConsensusPubkey.Size()
		n += 1 + l + sovShared(uint64(l))
	}
	if m.EthStake != 0 {
		n += 1 + sovShared(uint64(m.EthStake))
	}
	if m.CdtBalance != 0 {
		n += 1 + sovShared(uint64(m.CdtBalance))
	}
	return n
}

func (m *Pool) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Id != 0 {
		n += 1 + sovShared(uint64(m.Id))
	}
	l = len(m.Eth2Pubkey)
	if l > 0 {
		n += 1 + l + sovShared(uint64(l))
	}
	if m.Balance != 0 {
		n += 1 + sovShared(uint64(m.Balance))
	}
	if m.Exited {
		n += 2
	}
	if m.Slashed {
		n += 2
	}
	if len(m.Committee) > 0 {
		l = 0
		for _, e := range m.Committee {
			l += sovShared(uint64(e))
		}
		n += 1 + sovShared(uint64(l)) + l
	}
	return n
}

func sovShared(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozShared(x uint64) (n int) {
	return sovShared(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Operator) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowShared
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
			return fmt.Errorf("proto: Operator: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Operator: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Id", wireType)
			}
			m.Id = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowShared
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Id |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field EthereumAddress", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowShared
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
				return ErrInvalidLengthShared
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthShared
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.EthereumAddress.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ConsensusPubkey", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowShared
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
				return ErrInvalidLengthShared
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthShared
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			var v ConsensusAddress
			m.ConsensusPubkey = &v
			if err := m.ConsensusPubkey.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field EthStake", wireType)
			}
			m.EthStake = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowShared
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.EthStake |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 5:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field CdtBalance", wireType)
			}
			m.CdtBalance = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowShared
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.CdtBalance |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipShared(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthShared
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthShared
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
func (m *Pool) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowShared
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
			return fmt.Errorf("proto: Pool: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Pool: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Id", wireType)
			}
			m.Id = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowShared
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Id |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Eth2Pubkey", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowShared
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
				return ErrInvalidLengthShared
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthShared
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Eth2Pubkey = append(m.Eth2Pubkey[:0], dAtA[iNdEx:postIndex]...)
			if m.Eth2Pubkey == nil {
				m.Eth2Pubkey = []byte{}
			}
			iNdEx = postIndex
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Balance", wireType)
			}
			m.Balance = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowShared
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Balance |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Exited", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowShared
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
		case 5:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Slashed", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowShared
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
			if wireType == 0 {
				var v uint64
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return ErrIntOverflowShared
					}
					if iNdEx >= l {
						return io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					v |= uint64(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				m.Committee = append(m.Committee, v)
			} else if wireType == 2 {
				var packedLen int
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return ErrIntOverflowShared
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
					return ErrInvalidLengthShared
				}
				postIndex := iNdEx + packedLen
				if postIndex < 0 {
					return ErrInvalidLengthShared
				}
				if postIndex > l {
					return io.ErrUnexpectedEOF
				}
				var elementCount int
				var count int
				for _, integer := range dAtA[iNdEx:postIndex] {
					if integer < 128 {
						count++
					}
				}
				elementCount = count
				if elementCount != 0 && len(m.Committee) == 0 {
					m.Committee = make([]uint64, 0, elementCount)
				}
				for iNdEx < postIndex {
					var v uint64
					for shift := uint(0); ; shift += 7 {
						if shift >= 64 {
							return ErrIntOverflowShared
						}
						if iNdEx >= l {
							return io.ErrUnexpectedEOF
						}
						b := dAtA[iNdEx]
						iNdEx++
						v |= uint64(b&0x7F) << shift
						if b < 0x80 {
							break
						}
					}
					m.Committee = append(m.Committee, v)
				}
			} else {
				return fmt.Errorf("proto: wrong wireType = %d for field Committee", wireType)
			}
		default:
			iNdEx = preIndex
			skippy, err := skipShared(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthShared
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthShared
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
func skipShared(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowShared
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
					return 0, ErrIntOverflowShared
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
					return 0, ErrIntOverflowShared
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
				return 0, ErrInvalidLengthShared
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupShared
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthShared
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthShared        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowShared          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupShared = fmt.Errorf("proto: unexpected end of group")
)
