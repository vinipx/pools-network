package types

// DONTCOVER

import (
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

// x/poolsnetwork module sentinel errors
var (
	ErrNoStakingValidatorForOperator = sdkerrors.Register(ModuleName, 1100, "No Staking Validator For Operator")
)
