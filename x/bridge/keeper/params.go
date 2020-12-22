package keeper

import (
	"github.com/bloxapp/pools-network/x/bridge/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// returns the max possible delegate claims in a MsgEthereumClaim
func (k Keeper) GetMaxDelegateVoteClaims(ctx sdk.Context) (res uint64) {
	k.paramstore.Get(ctx, types.KeyMaxDelegateVotes, &res)
	return
}

// returns the max possible un-delegate claims in a MsgEthereumClaim
func (k Keeper) GetMaxUnDelegateVoteClaims(ctx sdk.Context) (res uint64) {
	k.paramstore.Get(ctx, types.KeyMaxUndelegateVotes, &res)
	return
}

// returns the max possible new operator claims in a MsgEthereumClaim
func (k Keeper) GetMaxNewOperatorClaims(ctx sdk.Context) (res uint64) {
	k.paramstore.Get(ctx, types.KeyMaxCreateOperator, &res)
	return
}

// returns the max possible new pools claims in a MsgEthereumClaim
func (k Keeper) GetMaxNewPoolClaims(ctx sdk.Context) (res uint64) {
	k.paramstore.Get(ctx, types.KeyMaxCreatePools, &res)
	return
}

// SetParams sets the auth module's parameters.
func (k Keeper) SetParams(ctx sdk.Context, params types.Params) {
	k.paramstore.SetParamSet(ctx, &params)
}

// GetParams gets the auth module's parameters.
func (k Keeper) GetParams(ctx sdk.Context) (params types.Params) {
	k.paramstore.GetParamSet(ctx, &params)
	return
}
