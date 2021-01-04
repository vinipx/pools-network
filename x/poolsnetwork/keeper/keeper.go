package keeper

import (
	"fmt"

	bankKeeper "github.com/cosmos/cosmos-sdk/x/bank/keeper"

	authKeeper "github.com/cosmos/cosmos-sdk/x/auth/keeper"

	stakingKeeper "github.com/cosmos/cosmos-sdk/x/staking/keeper"

	"github.com/tendermint/tendermint/libs/log"

	"github.com/bloxapp/pools-network/x/poolsnetwork/types"
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

type (
	Keeper struct {
		cdc      codec.Marshaler
		storeKey sdk.StoreKey

		StakingKeeper stakingKeeper.Keeper
		AccountKeeper authKeeper.AccountKeeper
		BankKeeper    bankKeeper.Keeper
	}
)

func NewKeeper(
	cdc codec.Marshaler,
	storeKey sdk.StoreKey,
	stakingKeeper stakingKeeper.Keeper,
	accountKeeper authKeeper.AccountKeeper,
	bankingKeeper bankKeeper.Keeper,
) Keeper {
	return Keeper{
		cdc:           cdc,
		storeKey:      storeKey,
		StakingKeeper: stakingKeeper,
		AccountKeeper: accountKeeper,
		BankKeeper:    bankingKeeper,
	}
}

func (k Keeper) Logger(ctx sdk.Context) log.Logger {
	return ctx.Logger().With("module", fmt.Sprintf("x/%s", types.ModuleName))
}

// GetLastTotalPower returns the operator set total stakign power
func (k Keeper) GetLastTotalPower(ctx sdk.Context) uint64 {
	return k.StakingKeeper.GetLastTotalPower(ctx).Uint64()
}
