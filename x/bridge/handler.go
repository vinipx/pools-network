package bridge

import (
	"fmt"
	poolTypes "github.com/bloxapp/pools-network/x/poolsnetwork/types"

	"github.com/bloxapp/pools-network/x/bridge/keeper"
	bridgeTypes "github.com/bloxapp/pools-network/x/bridge/types"
	sdkTypes "github.com/cosmos/cosmos-sdk/types"
	sdkErrors "github.com/cosmos/cosmos-sdk/types/errors"
)

// NewHandler ...
func NewHandler(k keeper.Keeper) sdkTypes.Handler {
	return func(ctx sdkTypes.Context, msg sdkTypes.Msg) (*sdkTypes.Result, error) {
		ctx = ctx.WithEventManager(sdkTypes.NewEventManager())

		switch msg := msg.(type) {
		// this line is used by starport scaffolding # 1
		case *bridgeTypes.MsgEthereumClaim:
			return HandleMsgEthereumClaim(ctx, k, msg)
		default:
			errMsg := fmt.Sprintf("unrecognized %s message type: %T", poolTypes.ModuleName, msg)
			return nil, sdkErrors.Wrap(sdkErrors.ErrUnknownRequest, errMsg)
		}
	}
}

func HandleMsgEthereumClaim(ctx sdkTypes.Context, keeper keeper.Keeper, msg *bridgeTypes.MsgEthereumClaim) (*sdkTypes.Result, error) {
	// validate operator
	operator, found, err := keeper.PoolsKeeper.GetOperator(ctx, msg.ConsensusAddress)
	if !found {
		return nil, poolTypes.ErrOperatorNotFound
	}
	if err != nil {
		return nil, err
	}

	// validate bridge contract
	contract, found, err := keeper.GetEthereumBridgeContract(ctx, msg.ContractAddress)
	if !found {
		return nil, bridgeTypes.ErrBridgeContractNotFound
	}
	if err != nil {
		return nil, err
	}
	if contract.ChainId != msg.EthereumChainId {
		return nil, bridgeTypes.ErrWrongEthereumChainId
	}

	// validate nonce
	lastNonce := keeper.GetLastEthereumClaimNonce(ctx, operator.ConsensusAddress)
	if msg.Nonce != lastNonce.Uint64()+1 {
		return nil, sdkErrors.Wrap(bridgeTypes.ErrNonceInvalid, "non contiguous claim nonce")
	}
	keeper.SetLastEthereumClaimNonce(ctx, operator.ConsensusAddress, bridgeTypes.UInt64Nonce(msg.Nonce))

	// add claims
	for _, c := range msg.Data {
		// TODO - check slashing condition: same claim in different nonce
		if err := keeper.ProcessClaim(ctx, operator, contract, c); err != nil {
			return nil, err
		}
	}
	return &sdkTypes.Result{
		Data:   nil,
		Log:    "",
		Events: nil,
	}, nil
}
