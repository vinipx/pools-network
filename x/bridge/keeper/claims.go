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
	// add attestation
	// check if attestation finalized

	return nil
}

func (k Keeper) storeClaim(ctx sdk.Context, operator types3.Operator, contract types2.EthereumBridgeContact, claim *types2.ClaimData) error {
	return nil
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
