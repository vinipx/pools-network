package keeper

import (
	"fmt"

	sharedTypes "github.com/bloxapp/pools-network/shared/types"
	poolTypes "github.com/bloxapp/pools-network/x/poolsnetwork/types"

	sdkParamTypes "github.com/cosmos/cosmos-sdk/x/params/types"

	"github.com/tendermint/tendermint/libs/log"

	bridgeTypes "github.com/bloxapp/pools-network/x/bridge/types"
	"github.com/cosmos/cosmos-sdk/codec"
	sdkTypes "github.com/cosmos/cosmos-sdk/types"
)

// Keeper maintains the link to storage and exposes getter/setter methods for the various parts of the state machine
type (
	Keeper struct {
		cdc        codec.Marshaler
		storeKey   sdkTypes.StoreKey
		paramstore sdkParamTypes.Subspace

		PoolsKeeper PoolKeeper
	}
)

// PoolKeeper contains all necessary interfaces for the pool keeper, created to prevenet cyclic import
type PoolKeeper interface {
	GetOperator(ctx sdkTypes.Context, address sharedTypes.ConsensusAddress) (operator poolTypes.Operator, found bool, err error)
	GetOperatorByEthereumAddress(ctx sdkTypes.Context, address sharedTypes.EthereumAddress) (operator poolTypes.Operator, found bool, err error)
	CreateOperator(ctx sdkTypes.Context, operator poolTypes.Operator) error
	GetLastTotalPower(ctx sdkTypes.Context) uint64
	CreateDelegator(ctx sdkTypes.Context, address sdkTypes.AccAddress, balance uint64)
	Delegate(ctx sdkTypes.Context, from sdkTypes.AccAddress, to poolTypes.Operator, amount sdkTypes.Int) error
}

// NewKeeper returns a Keeper structure based on a marshaller, paramStore, storeKey, and poolsKeeper
func NewKeeper(cdc codec.Marshaler, paramstore sdkParamTypes.Subspace, storeKey sdkTypes.StoreKey, poolsKeeper PoolKeeper) Keeper {
	// set KeyTable if it has not already been set
	if !paramstore.HasKeyTable() {
		paramstore = paramstore.WithKeyTable(bridgeTypes.ParamKeyTable())
	}

	return Keeper{
		cdc:         cdc,
		storeKey:    storeKey,
		paramstore:  paramstore,
		PoolsKeeper: poolsKeeper,
	}
}

// Logger returns a logger to represent its modules name
func (k Keeper) Logger(ctx sdkTypes.Context) log.Logger {
	return ctx.Logger().With("module", fmt.Sprintf("x/%s", bridgeTypes.ModuleName))
}
