package keeper_test

import (
	"testing"

	"github.com/bloxapp/pools-network/x/poolsnetwork/keeper"

	tmproto "github.com/tendermint/tendermint/proto/tendermint/types"

	testing2 "github.com/bloxapp/pools-network/shared/testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

func CreateTestEnv(t *testing.T) (keeper.Keeper, sdk.Context) {
	t.Helper()

	app := testing2.SetupAppForTesting(false)
	ctx := app.BaseApp.NewContext(false, tmproto.Header{})

	return app.PoolsKeeper, ctx
}
