package ante

import (
	"fmt"

	"github.com/bloxapp/pools-network/x/bridge/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

type paramsGetter interface {
	// GetParams should return the Bridge module's params set
	GetParams(ctx sdk.Context) types.Params
}

func NewMsgEthereumClaimAnteHandler(keeper paramsGetter) MsgEthereumClaimAnteHandler {
	return MsgEthereumClaimAnteHandler{keeper: keeper}
}

type MsgEthereumClaimAnteHandler struct {
	keeper paramsGetter
}

var ErrInvalidMsg = fmt.Errorf("Ivalid MsgEthereumClaimAnteHandler")

func (sud MsgEthereumClaimAnteHandler) AnteHandle(ctx sdk.Context, tx sdk.Tx, simulate bool, next sdk.AnteHandler) (newCtx sdk.Context, err error) {
	params := sud.keeper.GetParams(ctx)

	for _, msg := range tx.GetMsgs() {
		claims, ok := msg.(*types.MsgEthereumClaim)
		if !ok {
			continue
		}

		// max delegate number
		if uint64(len(claims.Delegates)) > params.MaxDelegateVotesClaims {
			return ctx, sdkerrors.Wrapf(ErrInvalidMsg,
				"maximum number of delegate votes is %d but received %d",
				params.MaxDelegateVotesClaims, len(claims.Delegates),
			)
		}

		// max un-delegate number
		if uint64(len(claims.Undelegates)) > params.MaxUndelegateVotesClaims {
			return ctx, sdkerrors.Wrapf(ErrInvalidMsg,
				"maximum number of un-delegate votes is %d but received %d",
				params.MaxUndelegateVotesClaims, len(claims.Undelegates),
			)
		}

		// max create pools
		if uint64(len(claims.CreatePools)) > params.MaxCreatePoolsClaims {
			return ctx, sdkerrors.Wrapf(ErrInvalidMsg,
				"maximum number of create pool claims is %d but received %d",
				params.MaxCreatePoolsClaims, len(claims.CreatePools),
			)
		}

		// max create operator
		if uint64(len(claims.CreateOperators)) > params.MaxCreateOperatorClaims {
			return ctx, sdkerrors.Wrapf(ErrInvalidMsg,
				"maximum number of create operator claimss is %d but received %d characters",
				params.MaxCreateOperatorClaims, len(claims.CreateOperators),
			)
		}
	}

	return next(ctx, tx, simulate)
}
