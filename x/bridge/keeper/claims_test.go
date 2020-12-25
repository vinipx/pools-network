package keeper

import (
	"testing"

	types2 "github.com/bloxapp/pools-network/x/bridge/types"

	"github.com/stretchr/testify/require"

	"github.com/bloxapp/pools-network/shared/types"
)

func TestGetAndSetLastEthereumClaimNonce(t *testing.T) {
	keeper, ctx := CreateTestEnv(t)
	address := types.ConsensusAddress([]byte{1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1})
	keeper.SetLastEthereumClaimNonce(ctx, address, 1)
	res := keeper.GetLastEthereumClaimNonce(ctx, address)
	require.EqualValues(t, uint64(1), res.Uint64())
}

func TestGetAndSetEthereumBridgeAddress(t *testing.T) {
	keeper, ctx := CreateTestEnv(t)
	err := keeper.SetEthereumBridgeContract(ctx, types2.EthereumBridgeContact{
		ContractAddress: types.EthereumAddress{1, 2, 3, 4},
		ChainId:         2,
	})
	require.NoError(t, err)

	// find valid
	c, found, err := keeper.GetEthereumBridgeContract(ctx, types.EthereumAddress{1, 2, 3, 4})
	require.True(t, found)
	require.NoError(t, err)
	require.EqualValues(t, types.EthereumAddress{1, 2, 3, 4}, c.ContractAddress)
	require.EqualValues(t, 2, c.ChainId)

	// find invalid
	_, found, err = keeper.GetEthereumBridgeContract(ctx, types.EthereumAddress{1, 2, 3, 5})
	require.False(t, found)
	require.NoError(t, err)
}
