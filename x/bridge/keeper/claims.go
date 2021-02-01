package keeper

import (
	sharedTypes "github.com/bloxapp/pools-network/shared/types"
	bridgeTypes "github.com/bloxapp/pools-network/x/bridge/types"
	poolTypes "github.com/bloxapp/pools-network/x/poolsnetwork/types"
	sdkTypes "github.com/cosmos/cosmos-sdk/types"
	sdkErrors "github.com/cosmos/cosmos-sdk/types/errors"
)

// GetEthereumBridgeContract checks and returns the contract
func (k Keeper) GetEthereumBridgeContract(ctx sdkTypes.Context, address sharedTypes.EthereumAddress) (contract bridgeTypes.EthereumBridgeContact, found bool, err error) {
	store := ctx.KVStore(k.storeKey)
	byts := store.Get(address[:])

	if byts == nil || len(byts) == 0 {
		return bridgeTypes.EthereumBridgeContact{}, false, nil
	}

	// unmarshal
	ret := bridgeTypes.EthereumBridgeContact{}
	err = ret.Unmarshal(byts)
	if err != nil {
		return bridgeTypes.EthereumBridgeContact{}, false, sdkErrors.Wrap(err, "Could not unmarshal EthereumBridgeContact")
	}

	return ret, true, nil
}

// SetEthereumBridgeContract persists the Ethereum contract address to the KVStore
func (k Keeper) SetEthereumBridgeContract(ctx sdkTypes.Context, contract bridgeTypes.EthereumBridgeContact) error {
	store := ctx.KVStore(k.storeKey)
	byts, err := contract.Marshal()
	if err != nil {
		return sdkErrors.Wrap(err, "Could not marshal the EthereumBridgeContact.")
	}
	store.Set(contract.ContractAddress[:], byts)
	return nil
}

// ProcessClaim process attestation after Keeper's validity checks on claim affirmation
func (k Keeper) ProcessClaim(ctx sdkTypes.Context, operator poolTypes.Operator, contract bridgeTypes.EthereumBridgeContact, claim bridgeTypes.ClaimData) error {
	// add attestation and mark finalized if enough votes
	att, err := k.AttestClaim(ctx, operator, contract, claim)
	if err != nil {
		return sdkErrors.Wrap(err, "could not attest claim")
	}

	// if finalized, process
	return k.ProcessAttestation(ctx, att)
}

// GetLastEthereumClaimNonce returns 0 if it's the operators first claim
func (k Keeper) GetLastEthereumClaimNonce(ctx sdkTypes.Context, operatorAddress sharedTypes.ConsensusAddress) bridgeTypes.UInt64Nonce {
	store := ctx.KVStore(k.storeKey)
	bytes := store.Get(bridgeTypes.GetOperatorLastClaimNonceKey(operatorAddress))
	if len(bytes) == 0 {
		return bridgeTypes.NewUInt64Nonce(0)
	}
	return bridgeTypes.UInt64NonceFromBytes(bytes)
}

// SetLastEthereumClaimNonce persists the nonce value into operator's address to the KVStore
func (k Keeper) SetLastEthereumClaimNonce(ctx sdkTypes.Context, operatorAddress sharedTypes.ConsensusAddress, nonce bridgeTypes.UInt64Nonce) {
	store := ctx.KVStore(k.storeKey)
	store.Set(bridgeTypes.GetOperatorLastClaimNonceKey(operatorAddress), nonce.Bytes())
}
