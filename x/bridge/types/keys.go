package types

import (
	"encoding/binary"

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

func GetClaimAttestationStoreKey(contract EthereumBridgeContact, claim ClaimData) []byte {
	nonceByts := make([]byte, 8)
	binary.LittleEndian.PutUint64(nonceByts, claim.ClaimNonce)

	chainIdByts := make([]byte, 8)
	binary.LittleEndian.PutUint64(chainIdByts, contract.ChainId)

	ret := contract.ContractAddress[:]
	ret = append(ret, chainIdByts...)
	ret = append(ret, claim.TxHash...)
	ret = append(ret, nonceByts...)
	ret = append(ret, []byte("_claim_attestation")...)
	return ret
}

func KeyPrefix(p string) []byte {
	return []byte(p)
}
