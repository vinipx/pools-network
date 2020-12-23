package ante

import (
	"testing"

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

func TestNewMsgEthereumClaimAnteHandler(t *testing.T) {
	tests := []struct {
		name             string
		msgEthereumClaim *types.MsgEthereumClaim
		errStr           string
	}{
		{
			name: "valid",
			msgEthereumClaim: &types.MsgEthereumClaim{
				Data: make([]*types.ClaimData, 24),
			},
			errStr: "",
		},
		{
			name: "max claims and votes, valid",
			msgEthereumClaim: &types.MsgEthereumClaim{
				Data: make([]*types.ClaimData, 25),
			},
			errStr: "",
		},
		{
			name: "too may claims",
			msgEthereumClaim: &types.MsgEthereumClaim{
				Data: make([]*types.ClaimData, 26),
			},
			errStr: "maximum number of claims is 25 but received 26: Ivalid MsgEthereumClaimAnteHandler",
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
