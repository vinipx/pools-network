package keeper

import (
	"testing"

	"github.com/stretchr/testify/require"

	shared "github.com/bloxapp/pools-network/shared/types"
	"github.com/bloxapp/pools-network/x/poolsnetwork/types"
)

func TestSetAndGetOperator(t *testing.T) {
	keeper, ctx := CreateTestEnv(t)

	err := keeper.SetOperator(ctx, types.Operator{
		Id:               12,
		EthereumAddress:  shared.EthereumAddress{1, 2, 3, 4},
		ConsensusAddress: shared.ConsensusAddress{1, 2, 3, 4, 5},
		EthStake:         191,
		CdtBalance:       2,
	})
	require.NoError(t, err)

	// find valid
	operator, found, err := keeper.GetOperator(ctx, shared.ConsensusAddress{1, 2, 3, 4, 5})
	require.NoError(t, err)
	require.True(t, found)
	require.EqualValues(t, 12, operator.Id)
	require.EqualValues(t, shared.EthereumAddress{1, 2, 3, 4}, operator.EthereumAddress)
	require.EqualValues(t, shared.ConsensusAddress{1, 2, 3, 4, 5}, operator.ConsensusAddress)
	require.EqualValues(t, 191, operator.EthStake)
	require.EqualValues(t, 2, operator.CdtBalance)

	// find invalid
	_, found, err = keeper.GetOperator(ctx, shared.ConsensusAddress{1, 2, 3, 4, 6})
	require.NoError(t, err)
	require.False(t, found)
}
