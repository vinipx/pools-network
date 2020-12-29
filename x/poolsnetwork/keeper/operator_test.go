package keeper

import (
	"testing"

	"github.com/tendermint/tendermint/crypto/ed25519"

	github_com_cosmos_cosmos_sdk_types "github.com/cosmos/cosmos-sdk/types"

	"github.com/stretchr/testify/require"

	shared "github.com/bloxapp/pools-network/shared/types"
	"github.com/bloxapp/pools-network/x/poolsnetwork/types"
)

func TestDeleteOperator(t *testing.T) {
	keeper, ctx := CreateTestEnv(t)

	sk := ed25519.GenPrivKey()
	pk := sk.PubKey()
	encoded, err := github_com_cosmos_cosmos_sdk_types.Bech32ifyPubKey(github_com_cosmos_cosmos_sdk_types.Bech32PubKeyTypeConsPub, pk)
	require.NoError(t, err)

	operator := types.Operator{
		EthereumAddress:  shared.EthereumAddress{1, 2, 3, 4},
		ConsensusAddress: shared.ConsensusAddress{1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1},
		ConsensusPk:      encoded,
		EthStake:         191,
		CdtBalance:       2,
	}
	err = keeper.CreateOperator(ctx, operator)
	require.NoError(t, err)

	// delete
	keeper.DeleteOperator(ctx, shared.ConsensusAddress{1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1})

	// verify
	_, found, err := keeper.GetOperator(ctx, operator.ConsensusAddress)
	require.NoError(t, err)
	require.False(t, found)

	_, found = keeper.StakingKeeper.GetValidator(ctx, github_com_cosmos_cosmos_sdk_types.ValAddress(shared.ConsensusAddress{1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1}))
	require.False(t, found)
}

func TestCreateOperator(t *testing.T) {
	keeper, ctx := CreateTestEnv(t)

	sk := ed25519.GenPrivKey()
	pk := sk.PubKey()
	encoded, err := github_com_cosmos_cosmos_sdk_types.Bech32ifyPubKey(github_com_cosmos_cosmos_sdk_types.Bech32PubKeyTypeConsPub, pk)
	require.NoError(t, err)

	err = keeper.CreateOperator(ctx, types.Operator{
		EthereumAddress:  shared.EthereumAddress{1, 2, 3, 4},
		ConsensusAddress: shared.ConsensusAddress{1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1},
		ConsensusPk:      encoded,
		EthStake:         191,
		CdtBalance:       2,
	})
	require.NoError(t, err)

	// find valid
	operator, found, err := keeper.GetOperator(ctx, shared.ConsensusAddress{1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1})
	require.NoError(t, err)
	require.True(t, found)
	require.NotNil(t, operator.CosmosValidatorRef)
	require.EqualValues(t, encoded, operator.CosmosValidatorRef.ConsensusPubkey)
	require.EqualValues(t, shared.EthereumAddress{1, 2, 3, 4}, operator.EthereumAddress)
	require.EqualValues(t, 191, operator.EthStake)
	require.EqualValues(t, 2, operator.CdtBalance)

	// find invalid
	_, found, err = keeper.GetOperator(ctx, shared.ConsensusAddress{1, 2, 3, 4, 6})
	require.NoError(t, err)
	require.False(t, found)
}
