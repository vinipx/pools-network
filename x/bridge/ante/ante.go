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
		if uint64(len(claims.Data)) > params.MaxClaims {
			return ctx, sdkerrors.Wrapf(ErrInvalidMsg,
				"maximum number of claims is %d but received %d",
				params.MaxClaims, len(claims.Data),
			)
		}

		// validate claims
		for _, c := range claims.Data {
			if err := c.ValidateBasic(); err != nil {
				return ctx, err
			}
		}
	}

	return next(ctx, tx, simulate)
}
