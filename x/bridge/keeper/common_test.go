package keeper_test

import (
	"testing"

	"github.com/bloxapp/pools-network/x/bridge/keeper"

	testing2 "github.com/bloxapp/pools-network/shared/testing"
	sdkTypes "github.com/cosmos/cosmos-sdk/types"
)

func CreateTestEnv(t *testing.T) (keeper.Keeper, sdkTypes.Context, []sdkTypes.AccAddress) {
	t.Helper()
	app, ctx, accounts := testing2.SetupAppForTesting(false)

	return app.BridgeKeeper, ctx, accounts
}
