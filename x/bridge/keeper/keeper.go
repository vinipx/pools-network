package keeper

import (
	"fmt"

	types3 "github.com/bloxapp/pools-network/shared/types"
	poolTypes "github.com/bloxapp/pools-network/x/poolsnetwork/types"

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
		paramstore types2.Subspace

		PoolsKeeper PoolKeeper
	}
)

// PoolKeeper contains all necessary interfaces for the pool keeper, created to prevenet cyclic import
type PoolKeeper interface {
	GetOperator(ctx sdk.Context, address types3.ConsensusAddress) (operator poolTypes.Operator, found bool, err error)
	CreateOperator(ctx sdk.Context, operator poolTypes.Operator) error
	GetLastTotalPower(ctx sdk.Context) uint64
}

func NewKeeper(cdc codec.Marshaler, paramstore types2.Subspace, storeKey sdk.StoreKey, poolsKeeper PoolKeeper) Keeper {
	// set KeyTable if it has not already been set
	if !paramstore.HasKeyTable() {
		paramstore = paramstore.WithKeyTable(types.ParamKeyTable())
	}

	return Keeper{
		cdc:         cdc,
		storeKey:    storeKey,
		paramstore:  paramstore,
		PoolsKeeper: poolsKeeper,
	}
}

func (k Keeper) Logger(ctx sdk.Context) log.Logger {
	return ctx.Logger().With("module", fmt.Sprintf("x/%s", types.ModuleName))
}
