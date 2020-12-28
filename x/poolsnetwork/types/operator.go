package types

// Returns the operator's voting power
func (operator Operator) GetPower() uint64 {
	return operator.EthStake
}

func (operator Operator) CopyWithoutValidatorRef() Operator {
	operator.CosmosValidatorRef = nil
	return operator
}
