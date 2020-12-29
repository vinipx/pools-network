package keeper_test

import (
	"testing"

	"github.com/bloxapp/pools-network/x/bridge/keeper"

	testing2 "github.com/bloxapp/pools-network/shared/testing"
	tmproto "github.com/tendermint/tendermint/proto/tendermint/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

func CreateTestEnv(t *testing.T) (keeper.Keeper, sdk.Context) {
	t.Helper()

	app := testing2.SetupAppForTesting(false)
	ctx := app.BaseApp.NewContext(false, tmproto.Header{})

	return app.BridgeKeeper, ctx
}
