package types

import (
	"github.com/bloxapp/pools-network/shared/types"
)

const (
	// ModuleName defines the module name
	ModuleName = "bridge"

	// StoreKey defines the primary module store key
	StoreKey = ModuleName

	// RouterKey is the message route for slashing
	RouterKey = ModuleName

	// QuerierRoute defines the module's query routing key
	QuerierRoute = ModuleName

	// MemStoreKey defines the in-memory store key
	MemStoreKey = "mem_capability"
)

var (
	OperatorLastClaimNonce = []byte{0x1}
)

func GetOperatorLastClaimNonceKey(operator types.ConsensusAddress) []byte {
	return append(OperatorLastClaimNonce, operator...)
}

// Important - each tx hash can only have one claim
func GetClaimStoreKey(contract EthereumBridgeContact, address types.ConsensusAddress, claim ClaimData) []byte {
	ret := contract.ContractAddress[:]
	ret = append(ret, address...)
	ret = append(ret, claim.TxHash...)
	ret = append(ret, []byte("_claim")...)
	return ret
}

// Important - each tx hash can only have one claim
func GetClaimAttestationStoreKey(contract EthereumBridgeContact, claim ClaimData) []byte {
	ret := contract.ContractAddress[:]
	ret = append(ret, claim.TxHash...)
	ret = append(ret, []byte("_claim_attestation")...)
	return ret
}

func KeyPrefix(p string) []byte {
	return []byte(p)
}
