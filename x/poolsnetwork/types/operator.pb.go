// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: poolsnetwork/v1beta/operator.proto

package types

import (
	fmt "fmt"
	github_com_bloxapp_pools_network_shared_types "github.com/bloxapp/pools-network/shared/types"
	github_com_cosmos_cosmos_sdk_x_staking_types "github.com/cosmos/cosmos-sdk/x/staking/types"
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
	Id                 uint64                                                         `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	EthereumAddress    github_com_bloxapp_pools_network_shared_types.EthereumAddress  `protobuf:"bytes,2,opt,name=ethereum_address,json=ethereumAddress,proto3,customtype=github.com/bloxapp/pools-network/shared/types.EthereumAddress" json:"ethereum_address"`
	ConsensusAddress   github_com_bloxapp_pools_network_shared_types.ConsensusAddress `protobuf:"bytes,3,opt,name=consensus_address,json=consensusAddress,proto3,customtype=github.com/bloxapp/pools-network/shared/types.ConsensusAddress" json:"consensus_address"`
	ConsensusPk        string                                                         `protobuf:"bytes,4,opt,name=consensus_pk,json=consensusPk,proto3" json:"consensus_pk,omitempty"`
	EthStake           uint64                                                         `protobuf:"varint,5,opt,name=eth_stake,json=ethStake,proto3" json:"eth_stake,omitempty"`
	CdtBalance         uint64                                                         `protobuf:"varint,6,opt,name=cdt_balance,json=cdtBalance,proto3" json:"cdt_balance,omitempty"`
	CosmosValidatorRef *github_com_cosmos_cosmos_sdk_x_staking_types.Validator        `protobuf:"bytes,7,opt,name=cosmos_validator_ref,json=cosmosValidatorRef,proto3,customtype=github.com/cosmos/cosmos-sdk/x/staking/types.Validator" json:"cosmos_validator_ref,omitempty"`
}

func (m *Operator) Reset()         { *m = Operator{} }
func (m *Operator) String() string { return proto.CompactTextString(m) }
func (*Operator) ProtoMessage()    {}
func (*Operator) Descriptor() ([]byte, []int) {
	return fileDescriptor_a0e60acd720822a7, []int{0}
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

func (m *Operator) GetConsensusPk() string {
	if m != nil {
		return m.ConsensusPk
	}
	return ""
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

type UpdateOperator struct {
	Nonce              uint64                                                          `protobuf:"varint,1,opt,name=nonce,proto3" json:"nonce,omitempty"`
	ConsensusAddress   *github_com_bloxapp_pools_network_shared_types.ConsensusAddress `protobuf:"bytes,2,opt,name=consensus_address,json=consensusAddress,proto3,customtype=github.com/bloxapp/pools-network/shared/types.ConsensusAddress" json:"consensus_address,omitempty"`
	NewEthereumAddress *github_com_bloxapp_pools_network_shared_types.EthereumAddress  `protobuf:"bytes,3,opt,name=new_ethereum_address,json=newEthereumAddress,proto3,customtype=github.com/bloxapp/pools-network/shared/types.EthereumAddress" json:"new_ethereum_address,omitempty"`
	NewEthStake        uint64                                                          `protobuf:"varint,4,opt,name=new_eth_stake,json=newEthStake,proto3" json:"new_eth_stake,omitempty"`
	Exit               bool                                                            `protobuf:"varint,5,opt,name=exit,proto3" json:"exit,omitempty"`
}

func (m *UpdateOperator) Reset()         { *m = UpdateOperator{} }
func (m *UpdateOperator) String() string { return proto.CompactTextString(m) }
func (*UpdateOperator) ProtoMessage()    {}
func (*UpdateOperator) Descriptor() ([]byte, []int) {
	return fileDescriptor_a0e60acd720822a7, []int{1}
}
func (m *UpdateOperator) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *UpdateOperator) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_UpdateOperator.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *UpdateOperator) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UpdateOperator.Merge(m, src)
}
func (m *UpdateOperator) XXX_Size() int {
	return m.Size()
}
func (m *UpdateOperator) XXX_DiscardUnknown() {
	xxx_messageInfo_UpdateOperator.DiscardUnknown(m)
}

var xxx_messageInfo_UpdateOperator proto.InternalMessageInfo

func (m *UpdateOperator) GetNonce() uint64 {
	if m != nil {
		return m.Nonce
	}
	return 0
}

func (m *UpdateOperator) GetNewEthStake() uint64 {
	if m != nil {
		return m.NewEthStake
	}
	return 0
}

func (m *UpdateOperator) GetExit() bool {
	if m != nil {
		return m.Exit
	}
	return false
}

func init() {
	proto.RegisterType((*Operator)(nil), "poolsnetwork.v1beta1.Operator")
	proto.RegisterType((*UpdateOperator)(nil), "poolsnetwork.v1beta1.UpdateOperator")
}

func init() {
	proto.RegisterFile("poolsnetwork/v1beta/operator.proto", fileDescriptor_a0e60acd720822a7)
}

var fileDescriptor_a0e60acd720822a7 = []byte{
	// 468 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xb4, 0x53, 0xcf, 0x6f, 0xd3, 0x30,
	0x14, 0x6e, 0xba, 0x6e, 0x74, 0xee, 0x18, 0xc3, 0xca, 0x21, 0x02, 0x29, 0x2d, 0x3d, 0xf5, 0xb2,
	0x5a, 0x13, 0x82, 0x03, 0x12, 0x48, 0x2b, 0x1a, 0xd7, 0xa1, 0x20, 0x38, 0x70, 0x89, 0x9c, 0xf8,
	0xad, 0x89, 0x92, 0xe6, 0x59, 0xb1, 0xbb, 0x96, 0xff, 0x82, 0x3f, 0x8a, 0xc3, 0x8e, 0x3d, 0xa2,
	0x1d, 0x2a, 0xd4, 0xfe, 0x23, 0x28, 0x76, 0xda, 0xfd, 0x00, 0x09, 0x21, 0xc4, 0x29, 0xf6, 0x97,
	0xcf, 0xef, 0x7b, 0xef, 0xfb, 0x6c, 0xd2, 0x97, 0x88, 0xb9, 0x2a, 0x40, 0xcf, 0xb0, 0xcc, 0xd8,
	0xe5, 0x49, 0x04, 0x9a, 0x33, 0x94, 0x50, 0x72, 0x8d, 0xe5, 0x50, 0x96, 0xa8, 0x91, 0xba, 0xb7,
	0x39, 0x43, 0xcb, 0x39, 0x79, 0xe2, 0x8e, 0x71, 0x8c, 0x86, 0xc0, 0xaa, 0x95, 0xe5, 0xf6, 0x17,
	0x3b, 0xa4, 0x7d, 0x5e, 0x1f, 0xa7, 0x87, 0xa4, 0x99, 0x0a, 0xcf, 0xe9, 0x39, 0x83, 0x56, 0xd0,
	0x4c, 0x05, 0x95, 0xe4, 0x08, 0x74, 0x02, 0x25, 0x4c, 0x27, 0x21, 0x17, 0xa2, 0x04, 0xa5, 0xbc,
	0x66, 0xcf, 0x19, 0x1c, 0x8c, 0xce, 0xae, 0x96, 0xdd, 0xc6, 0xf5, 0xb2, 0xfb, 0x7a, 0x9c, 0xea,
	0x64, 0x1a, 0x0d, 0x63, 0x9c, 0xb0, 0x28, 0xc7, 0x39, 0x97, 0x92, 0x19, 0xf5, 0xe3, 0x4d, 0x8b,
	0x2a, 0xe1, 0x25, 0x08, 0xa6, 0xbf, 0x48, 0x50, 0xc3, 0xb3, 0xba, 0xda, 0xa9, 0x2d, 0x16, 0x3c,
	0x82, 0xbb, 0x00, 0x55, 0xe4, 0x71, 0x8c, 0x85, 0x82, 0x42, 0x4d, 0xd5, 0x56, 0x72, 0xc7, 0x48,
	0xbe, 0xab, 0x25, 0xdf, 0xfc, 0x9d, 0xe4, 0xdb, 0x4d, 0xb9, 0x8d, 0xe6, 0x51, 0x7c, 0x0f, 0xa1,
	0xcf, 0xc8, 0xc1, 0x8d, 0xa8, 0xcc, 0xbc, 0x56, 0xcf, 0x19, 0xec, 0x07, 0x9d, 0x2d, 0xf6, 0x3e,
	0xa3, 0x4f, 0xc9, 0x3e, 0xe8, 0x24, 0x54, 0x9a, 0x67, 0xe0, 0xed, 0x1a, 0x83, 0xda, 0xa0, 0x93,
	0x0f, 0xd5, 0x9e, 0x76, 0x49, 0x27, 0x16, 0x3a, 0x8c, 0x78, 0xce, 0x8b, 0x18, 0xbc, 0x3d, 0xf3,
	0x9b, 0xc4, 0x42, 0x8f, 0x2c, 0x42, 0x73, 0xe2, 0xc6, 0xa8, 0x26, 0xa8, 0xc2, 0x4b, 0x9e, 0xa7,
	0xa2, 0xf2, 0x3a, 0x2c, 0xe1, 0xc2, 0x7b, 0x60, 0x06, 0x7b, 0x75, 0xbd, 0xec, 0xbe, 0xbc, 0x35,
	0x94, 0xa5, 0xd6, 0x9f, 0x63, 0x25, 0x32, 0x36, 0x67, 0x95, 0x70, 0x5a, 0x8c, 0xeb, 0x99, 0x3e,
	0x6d, 0xaa, 0x04, 0xd4, 0xb2, 0x6e, 0x00, 0xb8, 0xe8, 0x7f, 0x6b, 0x92, 0xc3, 0x8f, 0x52, 0x70,
	0x0d, 0xdb, 0x60, 0x5d, 0xb2, 0x5b, 0x60, 0xd5, 0x9b, 0xcd, 0xd6, 0x6e, 0x28, 0xfe, 0xce, 0x6c,
	0x9b, 0xef, 0xe8, 0xbf, 0x18, 0xad, 0x88, 0x5b, 0xc0, 0x2c, 0xfc, 0xe5, 0x4e, 0xd9, 0x80, 0x4f,
	0xff, 0xfd, 0x3e, 0xd1, 0x02, 0x66, 0xf7, 0x30, 0xda, 0x27, 0x0f, 0x6b, 0xd1, 0x3a, 0xbe, 0x96,
	0xf1, 0xa0, 0x63, 0xa9, 0x36, 0x41, 0x4a, 0x5a, 0x30, 0x4f, 0xb5, 0x49, 0xb6, 0x1d, 0x98, 0xf5,
	0xe8, 0xfc, 0x6a, 0xe5, 0x3b, 0x8b, 0x95, 0xef, 0xfc, 0x58, 0xf9, 0xce, 0xd7, 0xb5, 0xdf, 0x58,
	0xac, 0xfd, 0xc6, 0xf7, 0xb5, 0xdf, 0xf8, 0xfc, 0xe2, 0x8f, 0x4d, 0xce, 0xd9, 0x9d, 0x67, 0x6a,
	0x9a, 0x8d, 0xf6, 0xcc, 0x8b, 0x7b, 0xfe, 0x33, 0x00, 0x00, 0xff, 0xff, 0x5e, 0x18, 0x52, 0x7f,
	0xc3, 0x03, 0x00, 0x00,
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
	if m.CosmosValidatorRef != nil {
		{
			size := m.CosmosValidatorRef.Size()
			i -= size
			if _, err := m.CosmosValidatorRef.MarshalTo(dAtA[i:]); err != nil {
				return 0, err
			}
			i = encodeVarintOperator(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x3a
	}
	if m.CdtBalance != 0 {
		i = encodeVarintOperator(dAtA, i, uint64(m.CdtBalance))
		i--
		dAtA[i] = 0x30
	}
	if m.EthStake != 0 {
		i = encodeVarintOperator(dAtA, i, uint64(m.EthStake))
		i--
		dAtA[i] = 0x28
	}
	if len(m.ConsensusPk) > 0 {
		i -= len(m.ConsensusPk)
		copy(dAtA[i:], m.ConsensusPk)
		i = encodeVarintOperator(dAtA, i, uint64(len(m.ConsensusPk)))
		i--
		dAtA[i] = 0x22
	}
	{
		size := m.ConsensusAddress.Size()
		i -= size
		if _, err := m.ConsensusAddress.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintOperator(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x1a
	{
		size := m.EthereumAddress.Size()
		i -= size
		if _, err := m.EthereumAddress.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintOperator(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x12
	if m.Id != 0 {
		i = encodeVarintOperator(dAtA, i, uint64(m.Id))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *UpdateOperator) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *UpdateOperator) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *UpdateOperator) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Exit {
		i--
		if m.Exit {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i--
		dAtA[i] = 0x28
	}
	if m.NewEthStake != 0 {
		i = encodeVarintOperator(dAtA, i, uint64(m.NewEthStake))
		i--
		dAtA[i] = 0x20
	}
	if m.NewEthereumAddress != nil {
		{
			size := m.NewEthereumAddress.Size()
			i -= size
			if _, err := m.NewEthereumAddress.MarshalTo(dAtA[i:]); err != nil {
				return 0, err
			}
			i = encodeVarintOperator(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x1a
	}
	if m.ConsensusAddress != nil {
		{
			size := m.ConsensusAddress.Size()
			i -= size
			if _, err := m.ConsensusAddress.MarshalTo(dAtA[i:]); err != nil {
				return 0, err
			}
			i = encodeVarintOperator(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x12
	}
	if m.Nonce != 0 {
		i = encodeVarintOperator(dAtA, i, uint64(m.Nonce))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func encodeVarintOperator(dAtA []byte, offset int, v uint64) int {
	offset -= sovOperator(v)
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
		n += 1 + sovOperator(uint64(m.Id))
	}
	l = m.EthereumAddress.Size()
	n += 1 + l + sovOperator(uint64(l))
	l = m.ConsensusAddress.Size()
	n += 1 + l + sovOperator(uint64(l))
	l = len(m.ConsensusPk)
	if l > 0 {
		n += 1 + l + sovOperator(uint64(l))
	}
	if m.EthStake != 0 {
		n += 1 + sovOperator(uint64(m.EthStake))
	}
	if m.CdtBalance != 0 {
		n += 1 + sovOperator(uint64(m.CdtBalance))
	}
	if m.CosmosValidatorRef != nil {
		l = m.CosmosValidatorRef.Size()
		n += 1 + l + sovOperator(uint64(l))
	}
	return n
}

func (m *UpdateOperator) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Nonce != 0 {
		n += 1 + sovOperator(uint64(m.Nonce))
	}
	if m.ConsensusAddress != nil {
		l = m.ConsensusAddress.Size()
		n += 1 + l + sovOperator(uint64(l))
	}
	if m.NewEthereumAddress != nil {
		l = m.NewEthereumAddress.Size()
		n += 1 + l + sovOperator(uint64(l))
	}
	if m.NewEthStake != 0 {
		n += 1 + sovOperator(uint64(m.NewEthStake))
	}
	if m.Exit {
		n += 2
	}
	return n
}

func sovOperator(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozOperator(x uint64) (n int) {
	return sovOperator(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Operator) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowOperator
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
					return ErrIntOverflowOperator
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
					return ErrIntOverflowOperator
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
				return ErrInvalidLengthOperator
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthOperator
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
				return fmt.Errorf("proto: wrong wireType = %d for field ConsensusAddress", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowOperator
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
				return ErrInvalidLengthOperator
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthOperator
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.ConsensusAddress.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ConsensusPk", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowOperator
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
				return ErrInvalidLengthOperator
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthOperator
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ConsensusPk = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 5:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field EthStake", wireType)
			}
			m.EthStake = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowOperator
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
		case 6:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field CdtBalance", wireType)
			}
			m.CdtBalance = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowOperator
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
		case 7:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field CosmosValidatorRef", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowOperator
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
				return ErrInvalidLengthOperator
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthOperator
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			var v github_com_cosmos_cosmos_sdk_x_staking_types.Validator
			m.CosmosValidatorRef = &v
			if err := m.CosmosValidatorRef.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipOperator(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthOperator
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthOperator
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
func (m *UpdateOperator) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowOperator
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
			return fmt.Errorf("proto: UpdateOperator: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: UpdateOperator: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Nonce", wireType)
			}
			m.Nonce = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowOperator
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Nonce |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ConsensusAddress", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowOperator
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
				return ErrInvalidLengthOperator
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthOperator
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			var v github_com_bloxapp_pools_network_shared_types.ConsensusAddress
			m.ConsensusAddress = &v
			if err := m.ConsensusAddress.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field NewEthereumAddress", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowOperator
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
				return ErrInvalidLengthOperator
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthOperator
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			var v github_com_bloxapp_pools_network_shared_types.EthereumAddress
			m.NewEthereumAddress = &v
			if err := m.NewEthereumAddress.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field NewEthStake", wireType)
			}
			m.NewEthStake = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowOperator
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.NewEthStake |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 5:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Exit", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowOperator
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
			m.Exit = bool(v != 0)
		default:
			iNdEx = preIndex
			skippy, err := skipOperator(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthOperator
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthOperator
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
func skipOperator(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowOperator
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
					return 0, ErrIntOverflowOperator
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
					return 0, ErrIntOverflowOperator
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
				return 0, ErrInvalidLengthOperator
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupOperator
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthOperator
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthOperator        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowOperator          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupOperator = fmt.Errorf("proto: unexpected end of group")
)
