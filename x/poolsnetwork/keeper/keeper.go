package keeper

import (
	"fmt"

	"github.com/cosmos/cosmos-sdk/x/staking/keeper"

	"github.com/tendermint/tendermint/libs/log"

	"github.com/bloxapp/pools-network/x/poolsnetwork/types"
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

type (
	Keeper struct {
		cdc      codec.Marshaler
		storeKey sdk.StoreKey

		StakingKeeper keeper.Keeper
	}
)

func NewKeeper(cdc codec.Marshaler, storeKey sdk.StoreKey, stakingKeeper keeper.Keeper) Keeper {
	return Keeper{
		cdc:           cdc,
		storeKey:      storeKey,
		StakingKeeper: stakingKeeper,
	}
}

func (k Keeper) Logger(ctx sdk.Context) log.Logger {
	return ctx.Logger().With("module", fmt.Sprintf("x/%s", types.ModuleName))
}

// GetLastTotalPower returns the operator set total stakign power
func (k Keeper) GetLastTotalPower(ctx sdk.Context) uint64 {
	return k.StakingKeeper.GetLastTotalPower(ctx).Uint64()
}
