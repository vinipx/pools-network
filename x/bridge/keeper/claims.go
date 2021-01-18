package keeper

import (
	"github.com/bloxapp/pools-network/shared/types"
	types2 "github.com/bloxapp/pools-network/x/bridge/types"
	types3 "github.com/bloxapp/pools-network/x/poolsnetwork/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

// GetEthereumBridgeContract checks and returns the contract
func (k Keeper) GetEthereumBridgeContract(ctx sdk.Context, address types.EthereumAddress) (contract types2.EthereumBridgeContact, found bool, err error) {
	store := ctx.KVStore(k.storeKey)
	byts := store.Get(address[:])

	if byts == nil || len(byts) == 0 {
		return types2.EthereumBridgeContact{}, false, nil
	}

	// unmarshal
	ret := types2.EthereumBridgeContact{}
	err = ret.Unmarshal(byts)
	if err != nil {
		return types2.EthereumBridgeContact{}, false, sdkerrors.Wrap(err, "Could not unmarshal EthereumBridgeContact")
	}

	return ret, true, nil
}

// SetEthereumBridgeContract persists the Ethereum contract address to the KVStore
func (k Keeper) SetEthereumBridgeContract(ctx sdk.Context, contract types2.EthereumBridgeContact) error {
	store := ctx.KVStore(k.storeKey)
	byts, err := contract.Marshal()
	if err != nil {
		return sdkerrors.Wrap(err, "Could not marshal the EthereumBridgeContact.")
	}
	store.Set(contract.ContractAddress[:], byts)
	return nil
}

// ProcessClaim process attestation after Keeper's validity checks on claim afirmation
func (k Keeper) ProcessClaim(ctx sdk.Context, operator types3.Operator, contract types2.EthereumBridgeContact, claim types2.ClaimData) error {
	// add attestation and mark finalized if enough votes
	att, err := k.AttestClaim(ctx, operator, contract, claim)
	if err != nil {
		return sdkerrors.Wrap(err, "could not attest claim")
	}

	// if finalized, process
	return k.ProcessAttestation(ctx, att)
}

// GetLastEthereumClaimNonce returns 0 if it's the operators first claim
func (k Keeper) GetLastEthereumClaimNonce(ctx sdk.Context, operatorAddress types.ConsensusAddress) types2.UInt64Nonce {
	store := ctx.KVStore(k.storeKey)
	bytes := store.Get(types2.GetOperatorLastClaimNonceKey(operatorAddress))
	if len(bytes) == 0 {
		return types2.NewUInt64Nonce(0)
	}
	return types2.UInt64NonceFromBytes(bytes)
}

// SetLastEthereumClaimNonce persists the nonce value into operator's address to the KVStore
func (k Keeper) SetLastEthereumClaimNonce(ctx sdk.Context, operatorAddress types.ConsensusAddress, nonce types2.UInt64Nonce) {
	store := ctx.KVStore(k.storeKey)
	store.Set(types2.GetOperatorLastClaimNonceKey(operatorAddress), nonce.Bytes())
}
