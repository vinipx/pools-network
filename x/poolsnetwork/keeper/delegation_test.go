package keeper_test

import (
	"testing"

	"github.com/stretchr/testify/require"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

func TestCreateDelegator(t *testing.T) {
	keeper, ctx, _ := CreateTestEnv(t)

	tests := []struct {
		name    string
		address sdk.AccAddress
		balance uint64
	}{
		{
			name:    "valid",
			address: sdk.AccAddress{1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1},
			balance: 100000,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			// before supply
			beforeSupply := keeper.BankKeeper.GetSupply(ctx).GetTotal()

			keeper.CreateDelegator(ctx, test.address, test.balance)
			acc, balance := keeper.GetDelegator(ctx, test.address)
			require.EqualValues(t, test.address, acc.GetAddress())
			require.EqualValues(t, test.balance, balance)

			// after supply
			afterSupply := keeper.BankKeeper.GetSupply(ctx).GetTotal()
			require.True(t, afterSupply.IsEqual(beforeSupply.Add(sdk.NewCoin("stake", sdk.NewIntFromUint64(test.balance)))), "after supply is wrong")
		})
	}
}

func TestDelegate(t *testing.T) {
	// the delegate funtion is part of the create operator function, we test it's functionality there
	TestCreateDelegator(t)
}
