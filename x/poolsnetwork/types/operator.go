package types

import "github.com/cosmos/cosmos-sdk/types"

// Returns the operator's voting power
func (operator Operator) GetPower() uint64 {
	return uint64(types.TokensToConsensusPower(types.NewIntFromUint64(operator.EthStake)))
}

func (operator Operator) CopyWithoutValidatorRef() Operator {
	operator.CosmosValidatorRef = nil
	return operator
}
