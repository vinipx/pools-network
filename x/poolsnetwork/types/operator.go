package types

// Returns the operator's voting power
func (operator Operator) GetPower() uint64 {
	return operator.EthStake
}
