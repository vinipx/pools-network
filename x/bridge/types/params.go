package types

import paramstypes "github.com/cosmos/cosmos-sdk/x/params/types"

// Default parameter values
const (
	DefaultMaxDelegateVotes   uint64 = 10
	DefaultMaxUndelegateVotes uint64 = 10
	DefaultMaxCreatePools     uint64 = 10
	DefaultMaxCreateOperator  uint64 = 10
)

// Parameter keys
var (
	KeyMaxDelegateVotes   = []byte("MaxDelegateVotes")
	KeyMaxUndelegateVotes = []byte("MaxUndelegateVotes")
	KeyMaxCreatePools     = []byte("MaxCreatePools")
	KeyMaxCreateOperator  = []byte("MaxCreateOperator")
)

var _ paramstypes.ParamSet = &Params{}

func DefaultParams() Params {
	return Params{
		MaxDelegateVotesClaims:   DefaultMaxDelegateVotes,
		MaxUndelegateVotesClaims: DefaultMaxUndelegateVotes,
		MaxCreatePoolsClaims:     DefaultMaxCreatePools,
		MaxCreateOperatorClaims:  DefaultMaxCreateOperator,
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
		paramstypes.NewParamSetPair(KeyMaxDelegateVotes, &p.MaxDelegateVotesClaims, validateMaxDelegateVotesClaims),
		paramstypes.NewParamSetPair(KeyMaxUndelegateVotes, &p.MaxUndelegateVotesClaims, validateMaxUndelegateVotesClaims),
		paramstypes.NewParamSetPair(KeyMaxCreatePools, &p.MaxCreatePoolsClaims, validateMaxCreatePoolsClaims),
		paramstypes.NewParamSetPair(KeyMaxCreateOperator, &p.MaxCreateOperatorClaims, validateMaxCreateOperatorClaims),
	}
}

// Validate checks that the parameters have valid values.
func (p *Params) Validate() error {
	if err := validateMaxCreateOperatorClaims(p.MaxCreateOperatorClaims); err != nil {
		return err
	}
	if err := validateMaxCreatePoolsClaims(p.MaxCreatePoolsClaims); err != nil {
		return err
	}
	if err := validateMaxDelegateVotesClaims(p.MaxDelegateVotesClaims); err != nil {
		return err
	}
	if err := validateMaxUndelegateVotesClaims(p.MaxUndelegateVotesClaims); err != nil {
		return err
	}

	return nil
}

func validateMaxDelegateVotesClaims(i interface{}) error {
	return nil // TODO - validate bridge params
}

func validateMaxUndelegateVotesClaims(i interface{}) error {
	return nil // TODO - validate bridge params
}

func validateMaxCreatePoolsClaims(i interface{}) error {
	return nil // TODO - validate bridge params
}

func validateMaxCreateOperatorClaims(i interface{}) error {
	return nil // TODO - validate bridge params
}
