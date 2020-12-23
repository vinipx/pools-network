package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/bloxapp/pools-network/shared/types"
	types2 "github.com/bloxapp/pools-network/x/poolsnetwork/types"
)

func (k Keeper) GetOperator(ctx sdk.Context, address types.ConsensusAddress) (operator types2.Operator, found bool) {
	return types2.Operator{
		Id:               0,
		EthereumAddress:  types.EthereumAddress{},
		ConsensusAddress: nil,
		EthStake:         0,
		CdtBalance:       0,
	}, true
}
