package bridge_test

import (
	"testing"

	"github.com/cosmos/cosmos-sdk/crypto/keys/ed25519"

	"github.com/bloxapp/pools-network/x/bridge"

	testing2 "github.com/bloxapp/pools-network/shared/testing"

	types3 "github.com/bloxapp/pools-network/x/poolsnetwork/types"

	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/stretchr/testify/require"

	types2 "github.com/bloxapp/pools-network/shared/types"

	"github.com/bloxapp/pools-network/x/bridge/types"

	"github.com/bloxapp/pools-network/x/bridge/keeper"
)

func setupEnv(t *testing.T) (keeper.Keeper, sdk.Context, []sdk.AccAddress) {
	t.Helper()
	app, ctx, accounts := testing2.SetupAppForTesting(false)

	// generate pk for operator
	pk := ed25519.GenPrivKey().PubKey()
	encoded, err := sdk.Bech32ifyPubKey(sdk.Bech32PubKeyTypeConsPub, pk)
	require.NoError(t, err)

	err = app.BridgeKeeper.PoolsKeeper.CreateOperator(ctx, types3.Operator{
		EthereumAddress:  types2.EthereumAddress{1, 2, 3, 4},
		ConsensusAddress: types2.ConsensusAddress(accounts[0]),
		ConsensusPk:      encoded,
		EthStake:         10,
		CdtBalance:       10,
	})
	require.NoError(t, err)

	err = app.BridgeKeeper.SetEthereumBridgeContract(ctx, types.EthereumBridgeContact{
		ContractAddress: types2.EthereumAddress{1, 2, 3, 4},
		ChainId:         1,
	})
	require.NoError(t, err)

	return app.BridgeKeeper, ctx, accounts
}

func TestHandleMsgEthereumClaim(t *testing.T) {
	tests := []struct {
		name          string
		msg           *types.MsgEthereumClaim
		accountIndex  uint64
		expectedError string
	}{
		{
			name:          "valid",
			msg:           types.NewMsgEthereumClaim(1, 1, types2.EthereumAddress{1, 2, 3, 4}, nil),
			accountIndex:  0,
			expectedError: "",
		},
		{
			name:          "invalid operator",
			msg:           types.NewMsgEthereumClaim(1, 1, types2.EthereumAddress{1, 2, 3, 4}, nil),
			accountIndex:  1,
			expectedError: "Operator not found",
		},
		{
			name:          "invalid contract address",
			msg:           types.NewMsgEthereumClaim(1, 1, types2.EthereumAddress{1, 2, 3, 5}, nil),
			expectedError: "Ethereum bridge contract not found",
		},
		{
			name:          "invalid contract chain id",
			msg:           types.NewMsgEthereumClaim(1, 0, types2.EthereumAddress{1, 2, 3, 4}, nil),
			expectedError: "Ethereum chain id is wrong",
		},
		{
			name:          "invalid nonce",
			msg:           types.NewMsgEthereumClaim(0, 1, types2.EthereumAddress{1, 2, 3, 4}, nil),
			expectedError: "non contiguous claim nonce: Nonce invalid",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			keeper, ctx, accounts := setupEnv(t)

			// replace the nil consensus address
			test.msg.ConsensusAddress = types2.ConsensusAddress(accounts[test.accountIndex])

			res, err := bridge.HandleMsgEthereumClaim(
				ctx,
				keeper,
				test.msg,
			)

			if len(test.expectedError) == 0 {
				require.NoError(t, err)
				require.EqualValues(t, "", res.Log)

				require.EqualValues(t, test.msg.Nonce, keeper.GetLastEthereumClaimNonce(ctx, test.msg.ConsensusAddress).Uint64())
				// TODO - check claim is stored correctly, processed and so on
			} else {
				require.NotNil(t, err)
				require.EqualError(t, err, test.expectedError)
			}
		})
	}
}
