package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	"github.com/bloxapp/pools-network/shared/types"
	types2 "github.com/bloxapp/pools-network/x/poolsnetwork/types"
)

func (k Keeper) GetOperator(ctx sdk.Context, address types.ConsensusAddress) (operator types2.Operator, found bool, err error) {
	store := ctx.KVStore(k.storeKey)
	byts := store.Get(address)

	if byts == nil || len(byts) == 0 {
		return types2.Operator{}, false, nil
	}

	// unmarshal
	ret := types2.Operator{}
	err = ret.Unmarshal(byts)
	if err != nil {
		return types2.Operator{}, false, sdkerrors.Wrap(err, "Could not unmarshal operator")
	}

	return ret, true, nil
}

func (k Keeper) SetOperator(ctx sdk.Context, operator types2.Operator) error {
	store := ctx.KVStore(k.storeKey)
	byts, err := operator.Marshal()
	if err != nil {
		return sdkerrors.Wrap(err, "Could not set operator")
	}
	store.Set(operator.ConsensusAddress, byts)
	return nil
}
