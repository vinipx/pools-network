package poolsnetwork

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/bloxapp/pools-network/x/poolsnetwork/types"
	"github.com/bloxapp/pools-network/x/poolsnetwork/keeper"
)

func handleMsgCreatePoolTest(ctx sdk.Context, k keeper.Keeper, poolTest *types.MsgPoolTest) (*sdk.Result, error) {
	k.CreatePoolTest(ctx, *poolTest)

	return &sdk.Result{Events: ctx.EventManager().ABCIEvents()}, nil
}
