package ante

import (
	"testing"

	types2 "github.com/bloxapp/pools-network/shared/types"

	"github.com/stretchr/testify/require"

	"github.com/bloxapp/pools-network/x/bridge/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

type TxTest struct {
	msgs []sdk.Msg
}

func (tx TxTest) GetMsgs() []sdk.Msg {
	return tx.msgs
}

func (tx TxTest) ValidateBasic() error {
	return nil
}

type paramsProvider struct {
	p types.Params
}

func (p paramsProvider) GetParams(ctx sdk.Context) types.Params {
	return p.p
}

func validClaim() types.ClaimData {
	return types.ClaimData{
		TxHash:             []byte{1, 2, 3, 4},
		ClaimNonce:         1,
		ClaimType:          types.ClaimType_Delegate,
		EthereumAddresses:  []types2.EthereumAddress{{1, 2, 3, 4}, {1, 2, 3, 4}},
		ConsensusAddresses: []types2.ConsensusAddress{{1, 2, 3, 4}},
		Values:             []uint64{1},
	}
}

func populateWithClaims(claim types.ClaimData, length uint64) []types.ClaimData {
	ret := make([]types.ClaimData, length)
	for i := uint64(0); i < length; i++ {
		ret[i] = claim
	}
	return ret
}

func TestNewMsgEthereumClaimAnteHandler(t *testing.T) {
	tests := []struct {
		name             string
		msgEthereumClaim *types.MsgEthereumClaim
		errStr           string
	}{
		{
			name: "valid",
			msgEthereumClaim: &types.MsgEthereumClaim{
				Data: populateWithClaims(validClaim(), 24),
			},
			errStr: "",
		},
		{
			name: "max claims and votes, valid",
			msgEthereumClaim: &types.MsgEthereumClaim{
				Data: populateWithClaims(validClaim(), 25),
			},
			errStr: "",
		},
		{
			name: "too may claims",
			msgEthereumClaim: &types.MsgEthereumClaim{
				Data: populateWithClaims(validClaim(), 26),
			},
			errStr: "maximum number of claims is 25 but received 26: Invalid MsgEthereumClaimAnteHandler",
		},
		{
			name: "invalid claim data",
			msgEthereumClaim: &types.MsgEthereumClaim{
				Data: populateWithClaims(types.ClaimData{
					TxHash:             []byte{},
					ClaimNonce:         1,
					ClaimType:          types.ClaimType_Delegate,
					EthereumAddresses:  []types2.EthereumAddress{{1, 2, 3, 4}},
					ConsensusAddresses: []types2.ConsensusAddress{{1, 2, 3, 4}},
					Values:             []uint64{1, 2, 3, 4},
				}, 1),
			},
			errStr: "Tx hash is invalid: Claim data invalid",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			paramsProvider := paramsProvider{p: types.DefaultParams()}

			handler := sdk.ChainAnteDecorators(NewMsgEthereumClaimAnteHandler(paramsProvider))
			ctx := sdk.Context{}
			tx := TxTest{msgs: []sdk.Msg{test.msgEthereumClaim}}

			_, err := handler(ctx, tx, true)
			if len(test.errStr) > 0 {
				require.EqualError(t, err, test.errStr)
			} else {
				require.NoError(t, err)
			}
		})
	}
}
