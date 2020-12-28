package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	stakingTypes "github.com/cosmos/cosmos-sdk/x/staking/types"

	"github.com/bloxapp/pools-network/shared/types"
	poolTypes "github.com/bloxapp/pools-network/x/poolsnetwork/types"
)

func (k Keeper) GetOperator(ctx sdk.Context, address types.ConsensusAddress) (operator poolTypes.Operator, found bool, err error) {
	store := ctx.KVStore(k.storeKey)
	byts := store.Get(address)

	if byts == nil || len(byts) == 0 {
		return poolTypes.Operator{}, false, nil
	}

	// unmarshal
	ret := poolTypes.Operator{}
	err = ret.Unmarshal(byts)
	if err != nil {
		return poolTypes.Operator{}, false, sdkerrors.Wrap(err, "Could not unmarshal operator")
	}

	// attach cosmos validator ref
	val, found := k.StakingKeeper.GetValidator(ctx, sdk.ValAddress(ret.ConsensusAddress))
	if !found {
		return poolTypes.Operator{}, false, sdkerrors.Wrap(poolTypes.ErrNoStakingValidatorForOperator, "")
	}
	ret.CosmosValidatorRef = &val

	return ret, true, nil
}

func (k Keeper) SetOperator(ctx sdk.Context, operator poolTypes.Operator) error {
	store := ctx.KVStore(k.storeKey)

	revert := func() {
		k.DeleteOperator(ctx, operator)
	}

	cpy := operator.CopyWithoutValidatorRef()
	byts, err := cpy.Marshal()
	if err != nil {
		revert()
		return sdkerrors.Wrap(err, "Could not set operator")
	}
	store.Set(cpy.ConsensusAddress, byts)

	// An operator is a wrapper around the native staking validator found in the staking module
	// https://github.com/cosmos/cosmos-sdk/tree/master/x/staking
	// When setting an operator we should also be setting a dedicated validator
	pk, err := sdk.GetPubKeyFromBech32(sdk.Bech32PubKeyTypeConsPub, operator.ConsensusPk)
	if err != nil {
		revert()
		return sdkerrors.Wrap(err, "Could not set validator for staking module")
	}
	val := stakingTypes.NewValidator(sdk.ValAddress(operator.ConsensusAddress), pk, stakingTypes.Description{})

	k.StakingKeeper.SetValidator(ctx, val)

	return nil
}

func (k Keeper) DeleteOperator(ctx sdk.Context, operator poolTypes.Operator) {
	store := ctx.KVStore(k.storeKey)
	store.Delete(operator.ConsensusAddress)
}
