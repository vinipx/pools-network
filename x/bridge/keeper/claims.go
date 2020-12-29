package keeper

import (
	"github.com/bloxapp/pools-network/shared/types"
	types2 "github.com/bloxapp/pools-network/x/bridge/types"
	types3 "github.com/bloxapp/pools-network/x/poolsnetwork/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

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

func (k Keeper) SetEthereumBridgeContract(ctx sdk.Context, contract types2.EthereumBridgeContact) error {
	store := ctx.KVStore(k.storeKey)
	byts, err := contract.Marshal()
	if err != nil {
		return sdkerrors.Wrap(err, "Could not set ethereum contract address")
	}
	store.Set(contract.ContractAddress[:], byts)
	return nil
}

func (k Keeper) AddClaim(ctx sdk.Context, operator types3.Operator, contract types2.EthereumBridgeContact, claim *types2.ClaimData) error {
	// add claim to db
	_, found, err := k.getClaim(ctx, operator, contract, claim.TxHash)
	if err != nil {
		return err
	}
	if found {
		return types2.ErrClaimExists
	}
	if err := k.storeClaim(ctx, operator, contract, claim); err != nil {
		return err
	}

	// add attestation and mark finalized if enough votes
	att, err := k.attestClaim(ctx, operator, contract, *claim)
	if err != nil {
		return sdkerrors.Wrap(err, "could not attest claim")
	}

	// if finalized, process
	return k.processAttestation(ctx, att)
}

func (k Keeper) storeClaim(ctx sdk.Context, operator types3.Operator, contract types2.EthereumBridgeContact, claim *types2.ClaimData) error {
	store := ctx.KVStore(k.storeKey)
	byts, err := claim.Marshal()
	if err != nil {
		return sdkerrors.Wrap(err, "Could not store claim")
	}
	store.Set(types2.GetClaimStoreKey(contract, operator.ConsensusAddress, *claim), byts)
	return nil
}

func (k Keeper) getClaim(ctx sdk.Context, operator types3.Operator, contract types2.EthereumBridgeContact, claimTxHash []byte) (claim *types2.ClaimData, found bool, err error) {
	store := ctx.KVStore(k.storeKey)
	byts := store.Get(types2.GetClaimStoreKey(contract, operator.ConsensusAddress, types2.ClaimData{TxHash: claimTxHash}))
	if byts == nil || len(byts) == 0 {
		return nil, false, nil
	}

	ret := &types2.ClaimData{}
	if err := ret.Unmarshal(byts); err != nil {
		return nil, false, err
	}
	return ret, true, nil
}

// Returns 0 if it's the operators first claim
func (k Keeper) GetLastEthereumClaimNonce(ctx sdk.Context, operatorAddress types.ConsensusAddress) types2.UInt64Nonce {
	store := ctx.KVStore(k.storeKey)
	bytes := store.Get(types2.GetOperatorLastClaimNonceKey(operatorAddress))
	if len(bytes) == 0 {
		return types2.NewUInt64Nonce(0)
	}
	return types2.UInt64NonceFromBytes(bytes)
}

func (k Keeper) SetLastEthereumClaimNonce(ctx sdk.Context, operatorAddress types.ConsensusAddress, nonce types2.UInt64Nonce) {
	store := ctx.KVStore(k.storeKey)
	store.Set(types2.GetOperatorLastClaimNonceKey(operatorAddress), nonce.Bytes())
}
