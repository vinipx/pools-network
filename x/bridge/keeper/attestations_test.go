package keeper

import (
	"testing"

	"github.com/cosmos/cosmos-sdk/crypto/keys/ed25519"
	github_com_cosmos_cosmos_sdk_types "github.com/cosmos/cosmos-sdk/types"

	"github.com/stretchr/testify/require"

	sharedTypes "github.com/bloxapp/pools-network/shared/types"
	types2 "github.com/bloxapp/pools-network/x/bridge/types"
	"github.com/bloxapp/pools-network/x/poolsnetwork/types"
)

func randConsensusKey(t *testing.T) (*ed25519.PrivKey, string) {
	sk := ed25519.GenPrivKey()
	pk := sk.PubKey()
	encoded, err := github_com_cosmos_cosmos_sdk_types.Bech32ifyPubKey(github_com_cosmos_cosmos_sdk_types.Bech32PubKeyTypeConsPub, pk)
	require.NoError(t, err)

	return sk, encoded
}

func TestAttestClaim(t *testing.T) {
	keeper, ctx := CreateTestEnv(t)

	// setup
	contract := types2.EthereumBridgeContact{
		ContractAddress: sharedTypes.EthereumAddress{1, 2, 3, 4},
		ChainId:         2,
	}
	err := keeper.SetEthereumBridgeContract(ctx, contract)
	require.NoError(t, err)

	//
	_, encoded1 := randConsensusKey(t)
	operator1 := types.Operator{
		ConsensusAddress: sharedTypes.ConsensusAddress{1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1},
		ConsensusPk:      encoded1,
		EthStake:         100,
	}
	err = keeper.PoolsKeeper.CreateOperator(ctx, operator1)
	require.NoError(t, err)

	_, encoded2 := randConsensusKey(t)
	operator2 := types.Operator{
		ConsensusAddress: sharedTypes.ConsensusAddress{2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2},
		ConsensusPk:      encoded2,
		EthStake:         200,
	}
	err = keeper.PoolsKeeper.CreateOperator(ctx, operator2)
	require.NoError(t, err)

	// attest
	claim := types2.ClaimData{
		TxHash: []byte{1, 1, 1, 1},
	}

	_, err = keeper.attestClaim(ctx, operator1, contract, claim)
	require.NoError(t, err)
	_, err = keeper.attestClaim(ctx, operator2, contract, claim)
	require.NoError(t, err)

	att, err := keeper.getAttestation(ctx, contract, claim)
	require.NoError(t, err)
	require.EqualValues(t,
		[]byte{0x1, 0x2, 0x3, 0x4, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x1, 0x1, 0x1, 0x1, 0x5f, 0x63, 0x6c, 0x61, 0x69, 0x6d, 0x5f, 0x61, 0x74, 0x74, 0x65, 0x73, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e},
		att.ClaimId,
	)
	require.EqualValues(t, uint64(300), att.AccumulatedPower)
	require.Contains(t, att.Votes, sharedTypes.ConsensusAddress{1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1}.Hex())
	require.True(t, att.Votes[sharedTypes.ConsensusAddress{1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1}.Hex()])
	require.Contains(t, att.Votes, sharedTypes.ConsensusAddress{2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2}.Hex())
	require.True(t, att.Votes[sharedTypes.ConsensusAddress{2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2}.Hex()])
}

func TestGetAndSetClaimAttestation(t *testing.T) {
	keeper, ctx := CreateTestEnv(t)

	// setup
	contract := types2.EthereumBridgeContact{
		ContractAddress: sharedTypes.EthereumAddress{1, 2, 3, 4},
		ChainId:         2,
	}
	err := keeper.SetEthereumBridgeContract(ctx, contract)
	require.NoError(t, err)

	claim := types2.ClaimData{
		TxHash: []byte{1, 1, 1, 1},
	}

	// save and get
	err = keeper.saveAttestation(ctx, &types2.ClaimAttestation{
		ClaimId:          types2.GetClaimAttestationStoreKey(contract, claim),
		AccumulatedPower: 10,
	})
	require.NoError(t, err)

	att, err := keeper.getAttestation(ctx, contract, claim)
	require.NoError(t, err)
	require.EqualValues(t,
		[]byte{0x1, 0x2, 0x3, 0x4, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x1, 0x1, 0x1, 0x1, 0x5f, 0x63, 0x6c, 0x61, 0x69, 0x6d, 0x5f, 0x61, 0x74, 0x74, 0x65, 0x73, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e},
		att.ClaimId,
	)
}
