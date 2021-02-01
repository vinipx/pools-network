package ante

import (
	"testing"

	sharedTypes "github.com/bloxapp/pools-network/shared/types"

	"github.com/stretchr/testify/require"

	bridgeTypes "github.com/bloxapp/pools-network/x/bridge/types"
	sdkTypes "github.com/cosmos/cosmos-sdk/types"
)

type TxTest struct {
	msgs []sdkTypes.Msg
}

func (tx TxTest) GetMsgs() []sdkTypes.Msg {
	return tx.msgs
}

func (tx TxTest) ValidateBasic() error {
	return nil
}

type paramsProvider struct {
	p bridgeTypes.Params
}

func (p paramsProvider) GetParams(ctx sdkTypes.Context) bridgeTypes.Params {
	return p.p
}

func validClaim() bridgeTypes.ClaimData {
	return bridgeTypes.ClaimData{
		TxHash:             []byte{1, 2, 3, 4},
		ClaimNonce:         1,
		ClaimType:          bridgeTypes.ClaimType_Delegate,
		EthereumAddresses:  []sharedTypes.EthereumAddress{{1, 2, 3, 4}, {1, 2, 3, 4}},
		ConsensusAddresses: []sharedTypes.ConsensusAddress{{1, 2, 3, 4}},
		Values:             []uint64{1},
	}
}

func populateWithClaims(claim bridgeTypes.ClaimData, length uint64) []bridgeTypes.ClaimData {
	ret := make([]bridgeTypes.ClaimData, length)
	for i := uint64(0); i < length; i++ {
		ret[i] = claim
	}
	return ret
}

func TestNewMsgEthereumClaimAnteHandler(t *testing.T) {
	tests := []struct {
		name             string
		msgEthereumClaim *bridgeTypes.MsgEthereumClaim
		errStr           string
	}{
		{
			name: "valid",
			msgEthereumClaim: &bridgeTypes.MsgEthereumClaim{
				Data: populateWithClaims(validClaim(), 24),
			},
			errStr: "",
		},
		{
			name: "max claims and votes, valid",
			msgEthereumClaim: &bridgeTypes.MsgEthereumClaim{
				Data: populateWithClaims(validClaim(), 25),
			},
			errStr: "",
		},
		{
			name: "too may claims",
			msgEthereumClaim: &bridgeTypes.MsgEthereumClaim{
				Data: populateWithClaims(validClaim(), 26),
			},
			errStr: "maximum number of claims is 25 but received 26: Invalid MsgEthereumClaimAnteHandler",
		},
		{
			name: "invalid claim data",
			msgEthereumClaim: &bridgeTypes.MsgEthereumClaim{
				Data: populateWithClaims(bridgeTypes.ClaimData{
					TxHash:             []byte{},
					ClaimNonce:         1,
					ClaimType:          bridgeTypes.ClaimType_Delegate,
					EthereumAddresses:  []sharedTypes.EthereumAddress{{1, 2, 3, 4}},
					ConsensusAddresses: []sharedTypes.ConsensusAddress{{1, 2, 3, 4}},
					Values:             []uint64{1, 2, 3, 4},
				}, 1),
			},
			errStr: "Tx hash is invalid: Claim data invalid",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			paramsProvider := paramsProvider{p: bridgeTypes.DefaultParams()}

			handler := sdkTypes.ChainAnteDecorators(NewMsgEthereumClaimAnteHandler(paramsProvider))
			ctx := sdkTypes.Context{}
			tx := TxTest{msgs: []sdkTypes.Msg{test.msgEthereumClaim}}

			_, err := handler(ctx, tx, true)
			if len(test.errStr) > 0 {
				require.EqualError(t, err, test.errStr)
			} else {
				require.NoError(t, err)
			}
		})
	}
}
