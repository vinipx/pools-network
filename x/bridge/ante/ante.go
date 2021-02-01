package ante

import (
	"fmt"

	bridgeTypes "github.com/bloxapp/pools-network/x/bridge/types"
	sdkTypes "github.com/cosmos/cosmos-sdk/types"
	sdkErrors "github.com/cosmos/cosmos-sdk/types/errors"
)

type paramsGetter interface {
	// GetParams should return the Bridge module's params set
	GetParams(ctx sdkTypes.Context) bridgeTypes.Params
}

func NewMsgEthereumClaimAnteHandler(keeper paramsGetter) MsgEthereumClaimAnteHandler {
	return MsgEthereumClaimAnteHandler{keeper: keeper}
}

type MsgEthereumClaimAnteHandler struct {
	keeper paramsGetter
}

var ErrInvalidMsg = fmt.Errorf("Invalid MsgEthereumClaimAnteHandler")

func (sud MsgEthereumClaimAnteHandler) AnteHandle(ctx sdkTypes.Context, tx sdkTypes.Tx, simulate bool, next sdkTypes.AnteHandler) (newCtx sdkTypes.Context, err error) {
	params := sud.keeper.GetParams(ctx)

	for _, msg := range tx.GetMsgs() {
		claims, ok := msg.(*bridgeTypes.MsgEthereumClaim)
		if !ok {
			continue
		}

		// max delegate number
		if uint64(len(claims.Data)) > params.MaxClaims {
			return ctx, sdkErrors.Wrapf(ErrInvalidMsg,
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
