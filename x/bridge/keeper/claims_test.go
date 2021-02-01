package keeper_test

import (
	sharedTypes "github.com/bloxapp/pools-network/shared/types"
	bridgeTypes "github.com/bloxapp/pools-network/x/bridge/types"
	poolTypes "github.com/bloxapp/pools-network/x/poolsnetwork/types"
	sdkTypes "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestGetAndSetLastEthereumClaimNonce(t *testing.T) {
	keeper, ctx, _ := CreateTestEnv(t)
	address := sharedTypes.ConsensusAddress([]byte{1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1})
	keeper.SetLastEthereumClaimNonce(ctx, address, 1)
	res := keeper.GetLastEthereumClaimNonce(ctx, address)
	require.EqualValues(t, uint64(1), res.Uint64())
}

func TestGetAndSetEthereumBridgeContract(t *testing.T) {
	keeper, ctx, _ := CreateTestEnv(t)
	err := keeper.SetEthereumBridgeContract(ctx, bridgeTypes.EthereumBridgeContact{
		ContractAddress: sharedTypes.EthereumAddress{1, 2, 3, 4},
		ChainId:         2,
	})
	require.NoError(t, err)

	// find valid
	c, found, err := keeper.GetEthereumBridgeContract(ctx, sharedTypes.EthereumAddress{1, 2, 3, 4})
	require.NoError(t, err)
	require.True(t, found)
	require.EqualValues(t, sharedTypes.EthereumAddress{1, 2, 3, 4}, c.ContractAddress)
	require.EqualValues(t, 2, c.ChainId)

	// find invalid
	_, found, err = keeper.GetEthereumBridgeContract(ctx, sharedTypes.EthereumAddress{1, 2, 3, 5})
	require.NoError(t, err)
	require.False(t, found)
}

func TestProcessClaim(t *testing.T) {
	keeper, ctx, accounts := CreateTestEnv(t)

	encoded1 := randConsensusKey(t)
	operator1 := poolTypes.Operator{
		EthereumAddress:  sharedTypes.EthereumAddress{0, 0, 0, 0},
		ConsensusAddress: sharedTypes.ConsensusAddress(accounts[0]),
		ConsensusPk:      encoded1,
		EthStake:         sdkTypes.TokensFromConsensusPower(10).Uint64(),
	}

	contract := bridgeTypes.EthereumBridgeContact{
		ContractAddress: sharedTypes.EthereumAddress{1, 2, 3, 4},
		ChainId:         2,
	}

	tests := []struct {
		name        string
		claim       bridgeTypes.ClaimData
		operator    poolTypes.Operator
		contract    bridgeTypes.EthereumBridgeContact
		expectedErr string
	}{
		{
			name: "valid",
			claim: bridgeTypes.ClaimData{
				TxHash:             []byte{1, 2, 3, 4},
				ClaimType:          bridgeTypes.ClaimType_Delegate,
				EthereumAddresses:  []sharedTypes.EthereumAddress{{1, 2, 3, 4}},
				ConsensusAddresses: []sharedTypes.ConsensusAddress{sharedTypes.ConsensusAddress(accounts[0])},
			},
			operator:    operator1,
			contract:    contract,
			expectedErr: "",
		},
	}

	// setup env
	err := keeper.SetEthereumBridgeContract(ctx, contract)
	require.NoError(t, err)

	err = keeper.PoolsKeeper.CreateOperator(ctx, operator1)
	require.NoError(t, err)

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			err := keeper.ProcessClaim(
				ctx,
				test.operator,
				test.contract,
				test.claim,
			)
			if len(test.expectedErr) > 0 {
				require.NotNil(t, err)
				require.EqualError(t, err, test.expectedErr)
			} else {
				require.NoError(t, err)
			}
		})
	}
}
