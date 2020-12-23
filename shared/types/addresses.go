package types

import (
	"github.com/cosmos/cosmos-sdk/types"
	"github.com/ethereum/go-ethereum/common"
)

type EthereumAddress common.Address

func (address EthereumAddress) Size() int {
	return 20
}

func (address EthereumAddress) MarshalTo(dAtA []byte) (int, error) {
	return copy(dAtA, address[:]), nil
}

func (address EthereumAddress) Validate() error {
	return nil // TODO - should pass any validation?
}

func (address *EthereumAddress) Unmarshal(data []byte) error {
	var arr [20]byte
	copy(arr[:], data[:])

	*address = arr
	return nil
}

type ConsensusAddress types.ValAddress

func (address ConsensusAddress) Size() int {
	return len(address)
}

func (address ConsensusAddress) MarshalTo(dAtA []byte) (int, error) {
	return copy(dAtA, address[:]), nil
}

func (address ConsensusAddress) Validate() error {
	return nil // TODO - should pass any validation?
}

func (address *ConsensusAddress) Unmarshal(data []byte) error {
	*address = data
	return nil
}
