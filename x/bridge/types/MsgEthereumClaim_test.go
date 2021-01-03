package types

import (
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/bloxapp/pools-network/shared/types"
)

func TestMsgEthereumClaimValidateBasic(t *testing.T) {
	tests := []struct {
		name          string
		msg           MsgEthereumClaim
		expectedError string
	}{
		{
			name: "valid",
			msg: MsgEthereumClaim{
				Nonce:            1,
				EthereumChainId:  1,
				ContractAddress:  types.EthereumAddress{1, 2, 3, 4},
				ConsensusAddress: types.ConsensusAddress{1, 2, 3, 4},
				Data: []ClaimData{
					{
						TxHash:             []byte{1, 2, 3, 4},
						ClaimNonce:         1,
						ClaimType:          ClaimType_Delegate,
						EthereumAddresses:  []types.EthereumAddress{{1, 2, 3, 4}, {1, 2, 3, 4}},
						ConsensusAddresses: []types.ConsensusAddress{{1, 2, 3, 4}},
						Values:             []uint64{1},
					},
				},
			},
		},
		{
			name: "invalid consensus address",
			msg: MsgEthereumClaim{
				Nonce:            1,
				EthereumChainId:  1,
				ContractAddress:  types.EthereumAddress{1, 2, 3, 4},
				ConsensusAddress: types.ConsensusAddress{},
				Data: []ClaimData{
					{
						TxHash:             []byte{1, 2, 3, 4},
						ClaimNonce:         1,
						ClaimType:          ClaimType_Delegate,
						EthereumAddresses:  []types.EthereumAddress{{1, 2, 3, 4}, {1, 2, 3, 4}},
						ConsensusAddresses: []types.ConsensusAddress{{1, 2, 3, 4}},
						Values:             []uint64{1},
					},
				},
			},
			expectedError: "Consensus address is invalid: Claim data invalid",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			if len(test.expectedError) == 0 {
				require.NoError(t, test.msg.ValidateBasic())
			} else {
				require.EqualError(t, test.msg.ValidateBasic(), test.expectedError)
			}
		})
	}
}

func TestClaimDataValidateBasic(t *testing.T) {
	tests := []struct {
		name          string
		claim         ClaimData
		expectedError string
	}{
		{
			name: "valid delegate",
			claim: ClaimData{
				TxHash:             []byte{1, 2, 3, 4},
				ClaimNonce:         1,
				ClaimType:          ClaimType_Delegate,
				EthereumAddresses:  []types.EthereumAddress{{1, 2, 3, 4}, {1, 2, 3, 4}},
				ConsensusAddresses: []types.ConsensusAddress{{1, 2, 3, 4}},
				Values:             []uint64{1},
			},
		},
		{
			name: "Invalid tx hash",
			claim: ClaimData{
				TxHash:             []byte{},
				ClaimNonce:         1,
				ClaimType:          ClaimType_Delegate,
				EthereumAddresses:  []types.EthereumAddress{{1, 2, 3, 4}, {1, 2, 3, 4}},
				ConsensusAddresses: []types.ConsensusAddress{{1, 2, 3, 4}},
				Values:             []uint64{1},
			},
			expectedError: "Tx hash is invalid: Claim data invalid",
		},
		{
			name: "Invalid tx hash",
			claim: ClaimData{
				TxHash:             nil,
				ClaimNonce:         1,
				ClaimType:          ClaimType_Delegate,
				EthereumAddresses:  []types.EthereumAddress{{1, 2, 3, 4}, {1, 2, 3, 4}},
				ConsensusAddresses: []types.ConsensusAddress{{1, 2, 3, 4}},
				Values:             []uint64{1},
			},
			expectedError: "Tx hash is invalid: Claim data invalid",
		},
		{
			name: "Invalid delegate - ethereum address",
			claim: ClaimData{
				TxHash:             []byte{1, 2, 3, 4},
				ClaimNonce:         1,
				ClaimType:          ClaimType_Delegate,
				EthereumAddresses:  []types.EthereumAddress{{1, 2, 3, 4}},
				ConsensusAddresses: []types.ConsensusAddress{{1, 2, 3, 4}},
				Values:             []uint64{1},
			},
			expectedError: "Delegate/ Undelegate: Ethereum addresses length must be 2: Claim data invalid",
		},
		{
			name: "Invalid delegate - ethereum address",
			claim: ClaimData{
				TxHash:             []byte{1, 2, 3, 4},
				ClaimNonce:         1,
				ClaimType:          ClaimType_Delegate,
				EthereumAddresses:  []types.EthereumAddress{{1, 2, 3, 4}, {1, 2, 3, 4}, {1, 2, 3, 4}},
				ConsensusAddresses: []types.ConsensusAddress{{1, 2, 3, 4}},
				Values:             []uint64{1},
			},
			expectedError: "Delegate/ Undelegate: Ethereum addresses length must be 2: Claim data invalid",
		},
		{
			name: "Invalid delegate - values",
			claim: ClaimData{
				TxHash:             []byte{1, 2, 3, 4},
				ClaimNonce:         1,
				ClaimType:          ClaimType_Delegate,
				EthereumAddresses:  []types.EthereumAddress{{1, 2, 3, 4}, {1, 2, 3, 4}},
				ConsensusAddresses: []types.ConsensusAddress{{1, 2, 3, 4}},
				Values:             []uint64{},
			},
			expectedError: "Delegate/ Undelegate: values length must be 1: Claim data invalid",
		},
		{
			name: "Invalid delegate - values",
			claim: ClaimData{
				TxHash:             []byte{1, 2, 3, 4},
				ClaimNonce:         1,
				ClaimType:          ClaimType_Delegate,
				EthereumAddresses:  []types.EthereumAddress{{1, 2, 3, 4}, {1, 2, 3, 4}},
				ConsensusAddresses: []types.ConsensusAddress{{1, 2, 3, 4}},
				Values:             []uint64{1, 1},
			},
			expectedError: "Delegate/ Undelegate: values length must be 1: Claim data invalid",
		},
		{
			name: "valid un-delegate",
			claim: ClaimData{
				TxHash:             []byte{1, 2, 3, 4},
				ClaimNonce:         1,
				ClaimType:          ClaimType_Undelegate,
				EthereumAddresses:  []types.EthereumAddress{{1, 2, 3, 4}, {1, 2, 3, 4}},
				ConsensusAddresses: []types.ConsensusAddress{{1, 2, 3, 4}},
				Values:             []uint64{1},
			},
		},
		{
			name: "Invalid un-delegate - ethereum address",
			claim: ClaimData{
				TxHash:             []byte{1, 2, 3, 4},
				ClaimNonce:         1,
				ClaimType:          ClaimType_Undelegate,
				EthereumAddresses:  []types.EthereumAddress{{1, 2, 3, 4}},
				ConsensusAddresses: []types.ConsensusAddress{{1, 2, 3, 4}},
				Values:             []uint64{1},
			},
			expectedError: "Delegate/ Undelegate: Ethereum addresses length must be 2: Claim data invalid",
		},
		{
			name: "Invalid un-delegate - ethereum address",
			claim: ClaimData{
				TxHash:             []byte{1, 2, 3, 4},
				ClaimNonce:         1,
				ClaimType:          ClaimType_Undelegate,
				EthereumAddresses:  []types.EthereumAddress{{1, 2, 3, 4}, {1, 2, 3, 5}, {1, 2, 3, 6}},
				ConsensusAddresses: []types.ConsensusAddress{{1, 2, 3, 4}},
				Values:             []uint64{1},
			},
			expectedError: "Delegate/ Undelegate: Ethereum addresses length must be 2: Claim data invalid",
		},
		{
			name: "Invalid un-delegate - values",
			claim: ClaimData{
				TxHash:             []byte{1, 2, 3, 4},
				ClaimNonce:         1,
				ClaimType:          ClaimType_Undelegate,
				EthereumAddresses:  []types.EthereumAddress{{1, 2, 3, 4}, {1, 2, 3, 5}},
				ConsensusAddresses: []types.ConsensusAddress{{1, 2, 3, 4}},
				Values:             []uint64{},
			},
			expectedError: "Delegate/ Undelegate: values length must be 1: Claim data invalid",
		},
		{
			name: "Invalid un-delegate - values",
			claim: ClaimData{
				TxHash:             []byte{1, 2, 3, 4},
				ClaimNonce:         1,
				ClaimType:          ClaimType_Undelegate,
				EthereumAddresses:  []types.EthereumAddress{{1, 2, 3, 4}, {1, 2, 3, 5}},
				ConsensusAddresses: []types.ConsensusAddress{{1, 2, 3, 4}},
				Values:             []uint64{1, 1},
			},
			expectedError: "Delegate/ Undelegate: values length must be 1: Claim data invalid",
		},
		{
			name: "valid create operator",
			claim: ClaimData{
				TxHash:             []byte{1, 2, 3, 4},
				ClaimNonce:         1,
				ClaimType:          ClaimType_CreateOperator,
				EthereumAddresses:  []types.EthereumAddress{{1, 2, 3, 4}},
				ConsensusAddresses: []types.ConsensusAddress{{1, 2, 3, 4}},
				Values:             []uint64{1, 2},
			},
		},
		{
			name: "Invalid create operator - ethereum address",
			claim: ClaimData{
				TxHash:             []byte{1, 2, 3, 4},
				ClaimNonce:         1,
				ClaimType:          ClaimType_CreateOperator,
				EthereumAddresses:  []types.EthereumAddress{},
				ConsensusAddresses: []types.ConsensusAddress{{1, 2, 3, 4}},
				Values:             []uint64{1, 2},
			},
			expectedError: "CreateOperator: Ethereum addresses length must be 1: Claim data invalid",
		},
		{
			name: "Invalid create operator - ethereum address",
			claim: ClaimData{
				TxHash:             []byte{1, 2, 3, 4},
				ClaimNonce:         1,
				ClaimType:          ClaimType_CreateOperator,
				EthereumAddresses:  []types.EthereumAddress{{1, 2, 3, 4}, {1, 2, 3, 4}},
				ConsensusAddresses: []types.ConsensusAddress{{1, 2, 3, 4}},
				Values:             []uint64{1, 2},
			},
			expectedError: "CreateOperator: Ethereum addresses length must be 1: Claim data invalid",
		},
		{
			name: "Invalid create operator - consensus address",
			claim: ClaimData{
				TxHash:             []byte{1, 2, 3, 4},
				ClaimNonce:         1,
				ClaimType:          ClaimType_CreateOperator,
				EthereumAddresses:  []types.EthereumAddress{{1, 2, 3, 4}},
				ConsensusAddresses: []types.ConsensusAddress{},
				Values:             []uint64{1, 2},
			},
			expectedError: "CreateOperator: Consensus addresses length must be 1: Claim data invalid",
		},
		{
			name: "Invalid create operator - consensus address",
			claim: ClaimData{
				TxHash:             []byte{1, 2, 3, 4},
				ClaimNonce:         1,
				ClaimType:          ClaimType_CreateOperator,
				EthereumAddresses:  []types.EthereumAddress{{1, 2, 3, 4}},
				ConsensusAddresses: []types.ConsensusAddress{{1, 2, 3, 4}, {1, 2, 3, 4}},
				Values:             []uint64{1, 2},
			},
			expectedError: "CreateOperator: Consensus addresses length must be 1: Claim data invalid",
		},
		{
			name: "Invalid create operator - value",
			claim: ClaimData{
				TxHash:             []byte{1, 2, 3, 4},
				ClaimNonce:         1,
				ClaimType:          ClaimType_CreateOperator,
				EthereumAddresses:  []types.EthereumAddress{{1, 2, 3, 4}},
				ConsensusAddresses: []types.ConsensusAddress{{1, 2, 3, 4}},
				Values:             []uint64{1},
			},
			expectedError: "CreateOperator: values length must be 2: Claim data invalid",
		},
		{
			name: "Invalid create operator - value",
			claim: ClaimData{
				TxHash:             []byte{1, 2, 3, 4},
				ClaimNonce:         1,
				ClaimType:          ClaimType_CreateOperator,
				EthereumAddresses:  []types.EthereumAddress{{1, 2, 3, 4}},
				ConsensusAddresses: []types.ConsensusAddress{{1, 2, 3, 4}},
				Values:             []uint64{1, 2, 3},
			},
			expectedError: "CreateOperator: values length must be 2: Claim data invalid",
		},
		{
			name: "valid create pool",
			claim: ClaimData{
				TxHash:             []byte{1, 2, 3, 4},
				ClaimNonce:         1,
				ClaimType:          ClaimType_CreatePool,
				EthereumAddresses:  []types.EthereumAddress{{1, 2, 3, 4}, {1, 2, 3, 4}},
				ConsensusAddresses: []types.ConsensusAddress{{1, 2, 3, 4}},
				Values:             []uint64{1, 2, 3, 4},
			},
		},
		{
			name: "Uknown claim type",
			claim: ClaimData{
				TxHash:             []byte{1, 2, 3, 4},
				ClaimNonce:         1,
				ClaimType:          ClaimType(10),
				EthereumAddresses:  []types.EthereumAddress{{1, 2, 3, 4}, {1, 2, 3, 4}},
				ConsensusAddresses: []types.ConsensusAddress{{1, 2, 3, 4}},
				Values:             []uint64{1, 2, 3, 4},
			},
			expectedError: "Unknown claim type: Claim data invalid",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			if len(test.expectedError) == 0 {
				require.NoError(t, test.claim.ValidateBasic())
			} else {
				require.EqualError(t, test.claim.ValidateBasic(), test.expectedError)
			}
		})
	}
}
