package testing

import (
	"time"

	types2 "github.com/cosmos/cosmos-sdk/x/bank/types"

	tmproto "github.com/tendermint/tendermint/proto/tendermint/types"

	"github.com/cosmos/cosmos-sdk/crypto/keys/ed25519"
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/bloxapp/pools-network/app"
	"github.com/tendermint/tendermint/abci/types"
	"github.com/tendermint/tendermint/libs/json"
	log2 "github.com/tendermint/tendermint/libs/log"
	db2 "github.com/tendermint/tm-db"
)

func SetupAppForTesting(isCheckTx bool) (*app.App, sdk.Context, []sdk.AccAddress) {
	db := db2.NewMemDB()
	appName := "test"
	newApp := app.New(
		appName,
		log2.NewNopLogger(),
		db,
		nil,
		true,
		map[int64]bool{},
		app.DefaultNodeHome(appName),
		5,
		app.MakeEncodingConfig(),
	)

	if !isCheckTx {
		genesisState := app.NewDefaultGenesisState()
		stateBytes, err := json.MarshalIndent(genesisState, "", " ")
		if err != nil {
			panic(err)
		}

		newApp.InitChain(types.RequestInitChain{
			Time:          time.Time{},
			Validators:    []types.ValidatorUpdate{},
			AppStateBytes: stateBytes,
		})
	}

	ctx := newApp.BaseApp.NewContext(false, tmproto.Header{})

	return newApp, ctx, CreateTestAccounts(ctx, newApp, 100)
}

func CreateTestAccounts(ctx sdk.Context, app *app.App, number int) []sdk.AccAddress {
	addresses := createRandomAccounts(number)
	amount := sdk.TokensFromConsensusPower(10).MulRaw(100)
	initCoins := sdk.NewCoins(sdk.NewCoin(app.StakingKeeper.BondDenom(ctx), amount))

	for _, address := range addresses {
		acc := app.AccountKeeper.NewAccountWithAddress(ctx, address)
		app.AccountKeeper.SetAccount(ctx, acc)
		if err := app.BankKeeper.AddCoins(ctx, address, initCoins); err != nil {
			panic(err)
		}
	}

	// update total supply
	totalSupply := sdk.NewCoins(sdk.NewCoin(app.StakingKeeper.BondDenom(ctx), amount.Mul(sdk.NewIntFromUint64(uint64(number)))))
	prevSupply := app.BankKeeper.GetSupply(ctx)
	app.BankKeeper.SetSupply(ctx, types2.NewSupply(prevSupply.GetTotal().Add(totalSupply...)))

	return addresses
}

// createRandomAccounts is a strategy used by addTestAddrs() in order to generated addresses in random order.
func createRandomAccounts(accNum int) []sdk.AccAddress {
	testAddrs := make([]sdk.AccAddress, accNum)
	for i := 0; i < accNum; i++ {
		pk := ed25519.GenPrivKey().PubKey()
		testAddrs[i] = sdk.AccAddress(pk.Address())
	}

	return testAddrs
}
