package types

// DONTCOVER

import (
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

// x/bridge module sentinel errors
var (
	ErrOperatorNotFound       = sdkerrors.Register(ModuleName, 1100, "Operator not found")
	ErrBridgeContractNotFound = sdkerrors.Register(ModuleName, 1101, "Ethereum bridge contract not found")
	ErrWrongEthereumChainId   = sdkerrors.Register(ModuleName, 1102, "Ethereum chain id is wrong")
	ErrClaimDataInvalid       = sdkerrors.Register(ModuleName, 1103, "Delegate/ un-delegate claim data invalid")
	ErrNonceInvalid           = sdkerrors.Register(ModuleName, 1104, "Nonce invalid")
	ErrClaimExists            = sdkerrors.Register(ModuleName, 1105, "Claim already exists")
)
