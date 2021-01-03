package bridge

import (
	"fmt"

	"github.com/bloxapp/pools-network/x/bridge/keeper"
	"github.com/bloxapp/pools-network/x/bridge/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

// NewHandler ...
func NewHandler(k keeper.Keeper) sdk.Handler {
	return func(ctx sdk.Context, msg sdk.Msg) (*sdk.Result, error) {
		ctx = ctx.WithEventManager(sdk.NewEventManager())

		switch msg := msg.(type) {
		// this line is used by starport scaffolding # 1
		case *types.MsgEthereumClaim:
			return HandleMsgEthereumClaim(ctx, k, msg)
		default:
			errMsg := fmt.Sprintf("unrecognized %s message type: %T", types.ModuleName, msg)
			return nil, sdkerrors.Wrap(sdkerrors.ErrUnknownRequest, errMsg)
		}
	}
}

func HandleMsgEthereumClaim(ctx sdk.Context, keeper keeper.Keeper, msg *types.MsgEthereumClaim) (*sdk.Result, error) {
	// validate operator
	operator, found, err := keeper.PoolsKeeper.GetOperator(ctx, msg.ConsensusAddress)
	if !found {
		return nil, types.ErrOperatorNotFound
	}
	if err != nil {
		return nil, err
	}

	// validate bridge contract
	contract, found, err := keeper.GetEthereumBridgeContract(ctx, msg.ContractAddress)
	if !found {
		return nil, types.ErrBridgeContractNotFound
	}
	if err != nil {
		return nil, err
	}
	if contract.ChainId != msg.EthereumChainId {
		return nil, types.ErrWrongEthereumChainId
	}

	// validate nonce
	lastNonce := keeper.GetLastEthereumClaimNonce(ctx, operator.ConsensusAddress)
	if msg.Nonce != lastNonce.Uint64()+1 {
		return nil, sdkerrors.Wrap(types.ErrNonceInvalid, "non contiguous claim nonce")
	}
	keeper.SetLastEthereumClaimNonce(ctx, operator.ConsensusAddress, types.UInt64Nonce(msg.Nonce))

	// add claims
	for _, c := range msg.Data {
		// TODO - check slashing condition: same claim in different nonce
		if err := keeper.ProcessClaim(ctx, operator, contract, c); err != nil {
			return nil, err
		}
	}
	return &sdk.Result{
		Data:   nil,
		Log:    "",
		Events: nil,
	}, nil
}
