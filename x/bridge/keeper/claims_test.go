package keeper_test

import (
	"testing"

	github_com_cosmos_cosmos_sdk_types "github.com/cosmos/cosmos-sdk/types"

	types3 "github.com/bloxapp/pools-network/x/poolsnetwork/types"

	types2 "github.com/bloxapp/pools-network/x/bridge/types"

	"github.com/stretchr/testify/require"

	"github.com/bloxapp/pools-network/shared/types"
)

func TestGetAndSetLastEthereumClaimNonce(t *testing.T) {
	keeper, ctx, _ := CreateTestEnv(t)
	address := types.ConsensusAddress([]byte{1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1})
	keeper.SetLastEthereumClaimNonce(ctx, address, 1)
	res := keeper.GetLastEthereumClaimNonce(ctx, address)
	require.EqualValues(t, uint64(1), res.Uint64())
}

func TestGetAndSetEthereumBridgeAddress(t *testing.T) {
	keeper, ctx, _ := CreateTestEnv(t)
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

func TestAddClaim(t *testing.T) {
	keeper, ctx, accounts := CreateTestEnv(t)

	_, encoded1 := randConsensusKey(t)
	operator1 := types3.Operator{
		EthereumAddress:  types.EthereumAddress{0, 0, 0, 0},
		ConsensusAddress: types.ConsensusAddress(accounts[0]),
		ConsensusPk:      encoded1,
		EthStake:         github_com_cosmos_cosmos_sdk_types.TokensFromConsensusPower(10).Uint64(),
	}

	contract := types2.EthereumBridgeContact{
		ContractAddress: types.EthereumAddress{1, 2, 3, 4},
		ChainId:         2,
	}

	tests := []struct {
		name        string
		claim       *types2.ClaimData
		operator    types3.Operator
		contract    types2.EthereumBridgeContact
		expectedErr string
	}{
		{
			name: "valid",
			claim: &types2.ClaimData{
				TxHash:             []byte{1, 2, 3, 4},
				ClaimType:          types2.ClaimType_Delegate,
				EthereumAddresses:  []types.EthereumAddress{{1, 2, 3, 4}},
				ConsensusAddresses: []types.ConsensusAddress{types.ConsensusAddress(accounts[0])},
			},
			operator:    operator1,
			contract:    contract,
			expectedErr: "",
		},
		{
			name: "duplicate claim, should error",
			claim: &types2.ClaimData{
				TxHash:             []byte{1, 2, 3, 4},
				ClaimType:          types2.ClaimType_Delegate,
				EthereumAddresses:  []types.EthereumAddress{{1, 2, 3, 4}},
				ConsensusAddresses: []types.ConsensusAddress{types.ConsensusAddress(accounts[0])},
			},
			operator:    operator1,
			contract:    contract,
			expectedErr: "Claim already exists",
		},
	}

	// setup env
	err := keeper.SetEthereumBridgeContract(ctx, contract)
	require.NoError(t, err)

	err = keeper.PoolsKeeper.CreateOperator(ctx, operator1)
	require.NoError(t, err)

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			err := keeper.AddClaim(
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
