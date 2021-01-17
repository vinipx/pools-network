package keeper

import (
	"github.com/bloxapp/pools-network/x/bridge/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// GetMaxClaims returns the max possible delegate claims in a MsgEthereumClaim
func (k Keeper) GetMaxClaims(ctx sdk.Context) (res uint64) {
	k.paramstore.Get(ctx, types.KeyMaxClaims, &res)
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
