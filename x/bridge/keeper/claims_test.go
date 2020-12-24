package keeper

import (
	"testing"

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
