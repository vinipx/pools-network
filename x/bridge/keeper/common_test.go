package keeper_test

import (
	"testing"

	"github.com/bloxapp/pools-network/x/bridge/keeper"

	testing2 "github.com/bloxapp/pools-network/shared/testing"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func CreateTestEnv(t *testing.T) (keeper.Keeper, sdk.Context) {
	t.Helper()

	app, ctx := testing2.SetupAppForTesting(false)

	return app.BridgeKeeper, ctx
}
