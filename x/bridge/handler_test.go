package bridge

import (
	"testing"

	types3 "github.com/bloxapp/pools-network/x/poolsnetwork/types"

	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/stretchr/testify/require"

	types2 "github.com/bloxapp/pools-network/shared/types"

	"github.com/bloxapp/pools-network/x/bridge/types"

	"github.com/bloxapp/pools-network/x/bridge/keeper"
)

func setupEnv(t *testing.T) (keeper.Keeper, sdk.Context) {
	keeper, ctx := keeper.CreateTestEnv(t)

	err := keeper.PoolsKeeper.CreateOperator(ctx, types3.Operator{
		EthereumAddress:  types2.EthereumAddress{1, 2, 3, 4},
		ConsensusAddress: types2.ConsensusAddress{5, 6, 7, 8},
		EthStake:         10,
		CdtBalance:       10,
	})
	require.NoError(t, err)

	err = keeper.SetEthereumBridgeContract(ctx, types.EthereumBridgeContact{
		ContractAddress: types2.EthereumAddress{1, 2, 3, 4},
		ChainId:         1,
	})
	require.NoError(t, err)

	return keeper, ctx
}

func TestHandleMsgEthereumClaim(t *testing.T) {
	tests := []struct {
		name          string
		msg           *types.MsgEthereumClaim
		expectedError string
	}{
		{
			name:          "valid",
			msg:           types.NewMsgEthereumClaim(1, 1, types2.EthereumAddress{1, 2, 3, 4}, types2.ConsensusAddress{5, 6, 7, 8}),
			expectedError: "",
		},
		{
			name:          "invalid operator",
			msg:           types.NewMsgEthereumClaim(1, 1, types2.EthereumAddress{1, 2, 3, 4}, types2.ConsensusAddress{5, 6, 7, 9}),
			expectedError: "Operator not found",
		},
		{
			name:          "invalid contract address",
			msg:           types.NewMsgEthereumClaim(1, 1, types2.EthereumAddress{1, 2, 3, 5}, types2.ConsensusAddress{5, 6, 7, 8}),
			expectedError: "Ethereum bridge contract not found",
		},
		{
			name:          "invalid contract chain id",
			msg:           types.NewMsgEthereumClaim(1, 0, types2.EthereumAddress{1, 2, 3, 4}, types2.ConsensusAddress{5, 6, 7, 8}),
			expectedError: "Ethereum chain id is wrong",
		},
		{
			name:          "invalid nonce",
			msg:           types.NewMsgEthereumClaim(0, 1, types2.EthereumAddress{1, 2, 3, 4}, types2.ConsensusAddress{5, 6, 7, 8}),
			expectedError: "non contiguous claim nonce: Nonce invalid",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			keeper, ctx := setupEnv(t)
			res, err := handleMsgEthereumClaim(
				ctx,
				keeper,
				test.msg,
			)

			if len(test.expectedError) == 0 {
				require.NoError(t, err)
				require.EqualValues(t, "", res.Log)

				require.EqualValues(t, test.msg.Nonce, keeper.GetLastEthereumClaimNonce(ctx, test.msg.ConsensusAddress).Uint64())
				t.Fail() // TODO - check claim is stored correctly, processed and so on
			} else {
				require.NotNil(t, err)
				require.EqualError(t, err, test.expectedError)
			}
		})
	}
}
