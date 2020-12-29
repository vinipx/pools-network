package keeper_test

import (
	"testing"

	"github.com/bloxapp/pools-network/x/poolsnetwork/keeper"

	testing2 "github.com/bloxapp/pools-network/shared/testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

func CreateTestEnv(t *testing.T) (keeper.Keeper, sdk.Context, []sdk.AccAddress) {
	t.Helper()

	app, ctx, accounts := testing2.SetupAppForTesting(false)

	return app.PoolsKeeper, ctx, accounts
}
