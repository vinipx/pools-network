package keeper

import (
	"testing"

	"github.com/stretchr/testify/require"

	sharedTypes "github.com/bloxapp/pools-network/shared/types"
	types2 "github.com/bloxapp/pools-network/x/bridge/types"
	"github.com/bloxapp/pools-network/x/poolsnetwork/types"
)

func TestAttestClaim(t *testing.T) {
	keeper, ctx := CreateTestEnv(t)

	// setup
	contract := types2.EthereumBridgeContact{
		ContractAddress: sharedTypes.EthereumAddress{1, 2, 3, 4},
		ChainId:         2,
	}
	err := keeper.SetEthereumBridgeContract(ctx, contract)
	require.NoError(t, err)

	operator1 := types.Operator{
		ConsensusAddress: sharedTypes.ConsensusAddress{1, 2, 3, 4},
		EthStake:         100,
	}
	err = keeper.PoolsKeeper.SetOperator(ctx, operator1)
	require.NoError(t, err)

	operator2 := types.Operator{
		ConsensusAddress: sharedTypes.ConsensusAddress{5, 6, 7, 8},
		EthStake:         200,
	}
	err = keeper.PoolsKeeper.SetOperator(ctx, operator2)
	require.NoError(t, err)

	// attest
	claim := types2.ClaimData{
		TxHash: []byte{1, 1, 1, 1},
	}

	require.NoError(t, keeper.attestClaim(ctx, operator1, contract, claim))
	require.NoError(t, keeper.attestClaim(ctx, operator2, contract, claim))

	att, err := keeper.getAttestation(ctx, contract, claim)
	require.NoError(t, err)
	require.EqualValues(t,
		[]byte{0x1, 0x2, 0x3, 0x4, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x1, 0x1, 0x1, 0x1, 0x5f, 0x63, 0x6c, 0x61, 0x69, 0x6d, 0x5f, 0x61, 0x74, 0x74, 0x65, 0x73, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e},
		att.ClaimId,
	)
	require.EqualValues(t, uint64(300), att.AccumulatedPower)
	require.Contains(t, att.Votes, sharedTypes.ConsensusAddress{1, 2, 3, 4}.Hex())
	require.True(t, att.Votes[sharedTypes.ConsensusAddress{1, 2, 3, 4}.Hex()])
	require.Contains(t, att.Votes, sharedTypes.ConsensusAddress{5, 6, 7, 8}.Hex())
	require.True(t, att.Votes[sharedTypes.ConsensusAddress{5, 6, 7, 8}.Hex()])
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
