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
				Delegates:       []*types.DelegateVote{},
				Undelegates:     []*types.UnDelegateVote{},
				CreatePools:     []*types.CreatePool{},
				CreateOperators: []*types.CreateOperator{},
			},
			errStr: "",
		},
		{
			name: "max claims and votes, valid",
			msgEthereumClaim: &types.MsgEthereumClaim{
				Delegates:       make([]*types.DelegateVote, 10),
				Undelegates:     make([]*types.UnDelegateVote, 10),
				CreatePools:     make([]*types.CreatePool, 10),
				CreateOperators: make([]*types.CreateOperator, 10),
			},
			errStr: "",
		},
		{
			name: "too may delegate votes",
			msgEthereumClaim: &types.MsgEthereumClaim{
				Delegates:       make([]*types.DelegateVote, 11),
				Undelegates:     []*types.UnDelegateVote{},
				CreatePools:     []*types.CreatePool{},
				CreateOperators: []*types.CreateOperator{},
			},
			errStr: "maximum number of delegate votes is 10 but received 11: Ivalid MsgEthereumClaimAnteHandler",
		},
		{
			name: "too may un-delegate votes",
			msgEthereumClaim: &types.MsgEthereumClaim{
				Delegates:       []*types.DelegateVote{},
				Undelegates:     make([]*types.UnDelegateVote, 11),
				CreatePools:     []*types.CreatePool{},
				CreateOperators: []*types.CreateOperator{},
			},
			errStr: "maximum number of un-delegate votes is 10 but received 11: Ivalid MsgEthereumClaimAnteHandler",
		},
		{
			name: "too create pool claims",
			msgEthereumClaim: &types.MsgEthereumClaim{
				Delegates:       []*types.DelegateVote{},
				Undelegates:     []*types.UnDelegateVote{},
				CreatePools:     make([]*types.CreatePool, 11),
				CreateOperators: []*types.CreateOperator{},
			},
			errStr: "maximum number of create pool claims is 10 but received 11: Ivalid MsgEthereumClaimAnteHandler",
		},
		{
			name: "too create operator claims",
			msgEthereumClaim: &types.MsgEthereumClaim{
				Delegates:       []*types.DelegateVote{},
				Undelegates:     []*types.UnDelegateVote{},
				CreatePools:     []*types.CreatePool{},
				CreateOperators: make([]*types.CreateOperator, 11),
			},
			errStr: "maximum number of create operator claimss is 10 but received 11 characters: Ivalid MsgEthereumClaimAnteHandler",
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
