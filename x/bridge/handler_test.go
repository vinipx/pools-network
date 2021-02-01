package bridge_test

import (
	"testing"

	"github.com/cosmos/cosmos-sdk/crypto/keys/ed25519"

	"github.com/bloxapp/pools-network/x/bridge"

	testing2 "github.com/bloxapp/pools-network/shared/testing"

	poolTypes "github.com/bloxapp/pools-network/x/poolsnetwork/types"

	sdkTypes "github.com/cosmos/cosmos-sdk/types"

	"github.com/stretchr/testify/require"

	sharedTypes "github.com/bloxapp/pools-network/shared/types"

	bridgeTypes "github.com/bloxapp/pools-network/x/bridge/types"

	"github.com/bloxapp/pools-network/x/bridge/keeper"
)

func setupEnv(t *testing.T) (keeper.Keeper, sdkTypes.Context, []sdkTypes.AccAddress) {
	t.Helper()
	app, ctx, accounts := testing2.SetupAppForTesting(false)

	// generate pk for operator
	pk := ed25519.GenPrivKey().PubKey()
	encoded, err := sdkTypes.Bech32ifyPubKey(sdkTypes.Bech32PubKeyTypeConsPub, pk)
	require.NoError(t, err)

	err = app.BridgeKeeper.PoolsKeeper.CreateOperator(ctx, poolTypes.Operator{
		EthereumAddress:  sharedTypes.EthereumAddress{1, 2, 3, 4},
		ConsensusAddress: sharedTypes.ConsensusAddress(accounts[0]),
		ConsensusPk:      encoded,
		EthStake:         10,
		CdtBalance:       10,
	})
	require.NoError(t, err)

	err = app.BridgeKeeper.SetEthereumBridgeContract(ctx, bridgeTypes.EthereumBridgeContact{
		ContractAddress: sharedTypes.EthereumAddress{1, 2, 3, 4},
		ChainId:         1,
	})
	require.NoError(t, err)

	return app.BridgeKeeper, ctx, accounts
}

func TestHandleMsgEthereumClaim(t *testing.T) {
	tests := []struct {
		name          string
		msg           *bridgeTypes.MsgEthereumClaim
		accountIndex  uint64
		expectedError string
	}{
		{
			name:          "valid",
			msg:           bridgeTypes.NewMsgEthereumClaim(1, 1, sharedTypes.EthereumAddress{1, 2, 3, 4}, nil),
			accountIndex:  0,
			expectedError: "",
		},
		{
			name:          "invalid operator",
			msg:           bridgeTypes.NewMsgEthereumClaim(1, 1, sharedTypes.EthereumAddress{1, 2, 3, 4}, nil),
			accountIndex:  1,
			expectedError: "Operator not found",
		},
		{
			name:          "invalid contract address",
			msg:           bridgeTypes.NewMsgEthereumClaim(1, 1, sharedTypes.EthereumAddress{1, 2, 3, 5}, nil),
			expectedError: "Ethereum bridge contract not found",
		},
		{
			name:          "invalid contract chain id",
			msg:           bridgeTypes.NewMsgEthereumClaim(1, 0, sharedTypes.EthereumAddress{1, 2, 3, 4}, nil),
			expectedError: "Ethereum chain id is wrong",
		},
		{
			name:          "invalid nonce",
			msg:           bridgeTypes.NewMsgEthereumClaim(0, 1, sharedTypes.EthereumAddress{1, 2, 3, 4}, nil),
			expectedError: "non contiguous claim nonce: Nonce invalid",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			keeper, ctx, accounts := setupEnv(t)

			// replace the nil consensus address
			test.msg.ConsensusAddress = sharedTypes.ConsensusAddress(accounts[test.accountIndex])

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
