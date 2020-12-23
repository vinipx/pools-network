package types

import paramstypes "github.com/cosmos/cosmos-sdk/x/params/types"

// Default parameter values
const (
	DefaultMaxClaims uint64 = 25
)

// Parameter keys
var (
	KeyMaxClaims = []byte("KeyMaxClaims")
)

var _ paramstypes.ParamSet = &Params{}

func DefaultParams() Params {
	return Params{
		MaxClaims: DefaultMaxClaims,
	}
}

// ParamKeyTable for auth module
func ParamKeyTable() paramstypes.KeyTable {
	return paramstypes.NewKeyTable().RegisterParamSet(&Params{})
}

// ParamSetPairs implements the ParamSet interface and returns all the key/value pairs
// pairs of auth module's parameters.
// nolint
func (p *Params) ParamSetPairs() paramstypes.ParamSetPairs {
	return paramstypes.ParamSetPairs{
		paramstypes.NewParamSetPair(KeyMaxClaims, &p.MaxClaims, validateMaxClaims),
	}
}

// Validate checks that the parameters have valid values.
func (p *Params) Validate() error {
	return nil
}

func validateMaxClaims(i interface{}) error {
	return nil // TODO - validate bridge params
}
