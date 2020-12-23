package keeper

import (
	"fmt"

	keeper2 "github.com/bloxapp/pools-network/x/poolsnetwork/keeper"

	"github.com/cosmos/cosmos-sdk/x/staking/keeper"

	types2 "github.com/cosmos/cosmos-sdk/x/params/types"

	"github.com/tendermint/tendermint/libs/log"

	"github.com/bloxapp/pools-network/x/bridge/types"
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

type (
	Keeper struct {
		cdc        codec.Marshaler
		storeKey   sdk.StoreKey
		memKey     sdk.StoreKey
		paramstore types2.Subspace

		StakingKeeper keeper.Keeper
		PoolsKeeper   keeper2.Keeper
	}
)

func NewKeeper(cdc codec.Marshaler, paramstore types2.Subspace, storeKey, memKey sdk.StoreKey, stakingKeeper keeper.Keeper, poolsKeeper keeper2.Keeper) *Keeper {
	// set KeyTable if it has not already been set
	if !paramstore.HasKeyTable() {
		paramstore = paramstore.WithKeyTable(types.ParamKeyTable())
	}

	return &Keeper{
		cdc:           cdc,
		storeKey:      storeKey,
		memKey:        memKey,
		paramstore:    paramstore,
		StakingKeeper: stakingKeeper,
		PoolsKeeper:   poolsKeeper,
	}
}

func (k Keeper) Logger(ctx sdk.Context) log.Logger {
	return ctx.Logger().With("module", fmt.Sprintf("x/%s", types.ModuleName))
}
