package keeper_test

import (
	"testing"

	testing2 "github.com/bloxapp/pools-network/shared/testing"

	"github.com/cosmos/cosmos-sdk/crypto/keys/ed25519"
	github_com_cosmos_cosmos_sdk_types "github.com/cosmos/cosmos-sdk/types"

	"github.com/stretchr/testify/require"

	sharedTypes "github.com/bloxapp/pools-network/shared/types"
	types2 "github.com/bloxapp/pools-network/x/bridge/types"
	"github.com/bloxapp/pools-network/x/poolsnetwork/types"
)

func randConsensusKey(t *testing.T) string {
	pk := ed25519.GenPrivKey().PubKey()
	encoded, err := github_com_cosmos_cosmos_sdk_types.Bech32ifyPubKey(github_com_cosmos_cosmos_sdk_types.Bech32PubKeyTypeConsPub, pk)
	require.NoError(t, err)

	return encoded
}

func TestAttestClaim(t *testing.T) {
	tests := []struct {
		name                          string
		createOperatorsFromAccountIds []uint64
		attestClaimAccountIdx         []uint64
		claim                         types2.ClaimData
		finalPower                    uint64 // the final power the attestation attested to
		finalized                     bool
	}{
		{
			name:                          "valid, not finalized attestation",
			createOperatorsFromAccountIds: []uint64{0, 1, 2, 3},
			attestClaimAccountIdx:         []uint64{0},
			claim: types2.ClaimData{
				TxHash: []byte{1, 1, 1, 1},
			},
			finalPower: github_com_cosmos_cosmos_sdk_types.TokensFromConsensusPower(10).Uint64(),
			finalized:  false,
		},
		{
			name:                          "valid, not finalized attestation",
			createOperatorsFromAccountIds: []uint64{0, 1, 2, 3},
			attestClaimAccountIdx:         []uint64{0, 1},
			claim: types2.ClaimData{
				TxHash: []byte{1, 1, 1, 1},
			},
			finalPower: github_com_cosmos_cosmos_sdk_types.TokensFromConsensusPower(20).Uint64(),
			finalized:  false,
		},
		{
			name:                          "valid, finalized attestation",
			createOperatorsFromAccountIds: []uint64{0, 1, 2, 3},
			attestClaimAccountIdx:         []uint64{0, 1, 2},
			claim: types2.ClaimData{
				TxHash: []byte{1, 1, 1, 1},
			},
			finalPower: github_com_cosmos_cosmos_sdk_types.TokensFromConsensusPower(30).Uint64(),
			finalized:  true,
		},
	}

	contract := types2.EthereumBridgeContact{
		ContractAddress: sharedTypes.EthereumAddress{1, 2, 3, 4},
		ChainId:         2,
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			app, ctx, accounts := testing2.SetupAppForTesting(false)

			// bridge
			err := app.BridgeKeeper.SetEthereumBridgeContract(ctx, contract)
			require.NoError(t, err)

			// operators
			for _, indx := range test.createOperatorsFromAccountIds {
				account := accounts[indx]
				consensusAddress := sharedTypes.ConsensusAddress(account)
				pk := randConsensusKey(t)

				operator := types.Operator{
					ConsensusAddress: consensusAddress,
					ConsensusPk:      pk,
					EthStake:         github_com_cosmos_cosmos_sdk_types.TokensFromConsensusPower(10).Uint64(),
				}
				err = app.PoolsKeeper.CreateOperator(ctx, operator)
				require.NoError(t, err)
			}
			app.StakingKeeper.ApplyAndReturnValidatorSetUpdates(ctx)

			// attest claim
			for _, indx := range test.attestClaimAccountIdx {
				account := accounts[indx]
				consensusAddress := sharedTypes.ConsensusAddress(account)
				operator, found, err := app.PoolsKeeper.GetOperator(ctx, consensusAddress)
				require.True(t, found)
				require.NoError(t, err)

				_, err = app.BridgeKeeper.AttestClaim(ctx, operator, contract, test.claim)
				require.NoError(t, err)

				// verify claim contains operator vote
				attestation, found, err := app.BridgeKeeper.GetAttestation(ctx, contract, test.claim)
				require.NoError(t, err)
				require.True(t, found)
				require.EqualValues(t,
					[]byte{0x1, 0x2, 0x3, 0x4, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x1, 0x1, 0x1, 0x1, 0x5f, 0x63, 0x6c, 0x61, 0x69, 0x6d, 0x5f, 0x61, 0x74, 0x74, 0x65, 0x73, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e},
					attestation.ClaimId,
				)
				require.Contains(t, attestation.Votes, consensusAddress.Hex())
				require.True(t, attestation.Votes[consensusAddress.Hex()])
			}

			// verify final power
			attestation, found, err := app.BridgeKeeper.GetAttestation(ctx, contract, test.claim)
			require.NoError(t, err)
			require.True(t, found)
			require.EqualValues(t, test.finalPower, attestation.AccumulatedPower)
			require.EqualValues(t, test.finalized, attestation.Finalized)
		})
	}
}

func TestGetAndSetClaimAttestation(t *testing.T) {
	keeper, ctx, _ := CreateTestEnv(t)

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
	err = keeper.SaveAttestation(ctx, &types2.ClaimAttestation{
		ClaimId:          types2.GetClaimAttestationStoreKey(contract, claim),
		AccumulatedPower: 10,
	})
	require.NoError(t, err)

	att, found, err := keeper.GetAttestation(ctx, contract, claim)
	require.NoError(t, err)
	require.True(t, found)
	require.EqualValues(t,
		[]byte{0x1, 0x2, 0x3, 0x4, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x1, 0x1, 0x1, 0x1, 0x5f, 0x63, 0x6c, 0x61, 0x69, 0x6d, 0x5f, 0x61, 0x74, 0x74, 0x65, 0x73, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e},
		att.ClaimId,
	)
}
