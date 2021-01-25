package keeper_test

import (
	"testing"

	app2 "github.com/bloxapp/pools-network/app"

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
				TxHash:     []byte{1, 1, 1, 1},
				ClaimNonce: 1,
			},
			finalPower: 10,
			finalized:  false,
		},
		{
			name:                          "valid, not finalized attestation",
			createOperatorsFromAccountIds: []uint64{0, 1, 2, 3},
			attestClaimAccountIdx:         []uint64{0, 1},
			claim: types2.ClaimData{
				TxHash:     []byte{1, 1, 1, 1},
				ClaimNonce: 1,
			},
			finalPower: 20,
			finalized:  false,
		},
		{
			name:                          "valid, finalized attestation",
			createOperatorsFromAccountIds: []uint64{0, 1, 2, 3},
			attestClaimAccountIdx:         []uint64{0, 1, 2},
			claim: types2.ClaimData{
				TxHash:     []byte{1, 1, 1, 1},
				ClaimNonce: 1,
			},
			finalPower: 30,
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
				require.NoError(t, err)
				require.True(t, found)

				_, err = app.BridgeKeeper.AttestClaim(ctx, operator, contract, test.claim)
				require.NoError(t, err)

				// verify claim contains operator vote
				attestation, found, err := app.BridgeKeeper.GetAttestation(ctx, contract, test.claim)
				require.NoError(t, err)
				require.True(t, found)
				require.EqualValues(t,
					[]byte{0x1, 0x1, 0x1, 0x1},
					attestation.Claim.TxHash,
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
		Claim:            claim,
		Contract:         contract,
		AccumulatedPower: 10,
	})
	require.NoError(t, err)

	att, found, err := keeper.GetAttestation(ctx, contract, claim)
	require.NoError(t, err)
	require.True(t, found)
	require.EqualValues(t,
		[]byte{0x1, 0x1, 0x1, 0x1},
		att.Claim.TxHash,
	)
}

func TestProcessAttestation(t *testing.T) {
	// create operator
	app, ctx, accounts := testing2.SetupAppForTesting(false)

	account := accounts[0]
	consensusAddress := sharedTypes.ConsensusAddress(account)
	pk := randConsensusKey(t)

	operator := types.Operator{
		ConsensusAddress: consensusAddress,
		EthereumAddress:  sharedTypes.EthereumAddress{2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2},
		ConsensusPk:      pk,
		EthStake:         0,
	}
	err := app.PoolsKeeper.CreateOperator(ctx, operator)
	require.NoError(t, err)
	app.StakingKeeper.ApplyAndReturnValidatorSetUpdates(ctx)

	// make sure it has no power
	require.EqualValues(t, 0, app.StakingKeeper.GetLastValidatorPower(ctx, github_com_cosmos_cosmos_sdk_types.ValAddress(operator.ConsensusAddress)))

	// test
	tests := []struct {
		name          string
		attestation   *types2.ClaimAttestation
		verifyHandler func(t *testing.T, ctx github_com_cosmos_cosmos_sdk_types.Context, app *app2.App)
		expectedErr   string
	}{
		{
			name: "unsupported claim",
			attestation: &types2.ClaimAttestation{
				Claim: types2.ClaimData{
					ClaimType: types2.ClaimType(10),
				},
			},
			expectedErr: "Unsupported claim",
		},
		{
			name: "delegate",
			attestation: &types2.ClaimAttestation{
				Claim: types2.ClaimData{
					ClaimType:         types2.ClaimType_Delegate,
					EthereumAddresses: []sharedTypes.EthereumAddress{{1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1}, operator.EthereumAddress},
					Values:            []uint64{github_com_cosmos_cosmos_sdk_types.TokensFromConsensusPower(10).Uint64()},
				},
			},
			verifyHandler: func(t *testing.T, ctx github_com_cosmos_cosmos_sdk_types.Context, app *app2.App) {
				power := app.StakingKeeper.GetLastValidatorPower(ctx, github_com_cosmos_cosmos_sdk_types.ValAddress(operator.ConsensusAddress))
				require.EqualValues(t, 10, power)
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			err := app.BridgeKeeper.ProcessAttestation(ctx, test.attestation)
			if len(test.expectedErr) == 0 {
				require.NoError(t, err)
				app.StakingKeeper.ApplyAndReturnValidatorSetUpdates(ctx)
				test.verifyHandler(t, ctx, app)
			} else {
				require.EqualError(t, err, test.expectedErr)
			}
		})
	}
}
