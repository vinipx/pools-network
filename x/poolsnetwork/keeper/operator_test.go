package keeper_test

import (
	"testing"

	"github.com/tendermint/tendermint/crypto/ed25519"

	github_com_cosmos_cosmos_sdk_types "github.com/cosmos/cosmos-sdk/types"

	"github.com/stretchr/testify/require"

	shared "github.com/bloxapp/pools-network/shared/types"
	"github.com/bloxapp/pools-network/x/poolsnetwork/types"
)

func TestDeleteOperator(t *testing.T) {
	keeper, ctx, accounts := CreateTestEnv(t)

	sk := ed25519.GenPrivKey()
	pk := sk.PubKey()
	encoded, err := github_com_cosmos_cosmos_sdk_types.Bech32ifyPubKey(github_com_cosmos_cosmos_sdk_types.Bech32PubKeyTypeConsPub, pk)
	require.NoError(t, err)

	operator := types.Operator{
		EthereumAddress:  shared.EthereumAddress{1, 2, 3, 4},
		ConsensusAddress: shared.ConsensusAddress(accounts[0]),
		ConsensusPk:      encoded,
		EthStake:         191,
		CdtBalance:       2,
	}
	err = keeper.CreateOperator(ctx, operator)
	require.NoError(t, err)

	// delete
	keeper.DeleteOperator(ctx, shared.ConsensusAddress(accounts[0]))

	// verify
	_, found, err := keeper.GetOperator(ctx, operator.ConsensusAddress)
	require.NoError(t, err)
	require.False(t, found)

	_, found = keeper.StakingKeeper.GetValidator(ctx, github_com_cosmos_cosmos_sdk_types.ValAddress(shared.ConsensusAddress(accounts[0])))
	require.False(t, found)
}

func TestCreateOperator(t *testing.T) {
	keeper, ctx, accounts := CreateTestEnv(t)

	sk := ed25519.GenPrivKey()
	pk := sk.PubKey()
	encoded, err := github_com_cosmos_cosmos_sdk_types.Bech32ifyPubKey(github_com_cosmos_cosmos_sdk_types.Bech32PubKeyTypeConsPub, pk)
	require.NoError(t, err)

	err = keeper.CreateOperator(ctx, types.Operator{
		EthereumAddress:  shared.EthereumAddress{1, 2, 3, 4},
		ConsensusAddress: shared.ConsensusAddress(accounts[0]),
		ConsensusPk:      encoded,
		EthStake:         github_com_cosmos_cosmos_sdk_types.TokensFromConsensusPower(10).Uint64(),
		CdtBalance:       2,
	})
	require.NoError(t, err)

	// find valid
	operator, found, err := keeper.GetOperator(ctx, shared.ConsensusAddress(accounts[0]))
	require.NoError(t, err)
	require.True(t, found)
	require.NotNil(t, operator.CosmosValidatorRef)
	require.EqualValues(t, encoded, operator.CosmosValidatorRef.ConsensusPubkey)
	require.EqualValues(t, shared.EthereumAddress{1, 2, 3, 4}, operator.EthereumAddress)
	require.EqualValues(t, 10000000, operator.EthStake)
	require.EqualValues(t, 2, operator.CdtBalance)

	keeper.StakingKeeper.ApplyAndReturnValidatorSetUpdates(ctx)

	// verify initial delegation
	delegations := keeper.StakingKeeper.GetValidatorDelegations(ctx, github_com_cosmos_cosmos_sdk_types.ValAddress(shared.ConsensusAddress(accounts[0])))
	require.Len(t, delegations, 1)
	power := keeper.StakingKeeper.GetLastValidatorPower(ctx, github_com_cosmos_cosmos_sdk_types.ValAddress(shared.ConsensusAddress(accounts[0])))
	require.EqualValues(t, int64(10), power)

	// find invalid
	_, found, err = keeper.GetOperator(ctx, shared.ConsensusAddress{1, 2, 3, 4, 6})
	require.NoError(t, err)
	require.False(t, found)
}
