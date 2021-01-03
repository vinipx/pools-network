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
						Values:             []uint64{1, 2, 3, 4},
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
						Values:             []uint64{1, 2, 3, 4},
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
