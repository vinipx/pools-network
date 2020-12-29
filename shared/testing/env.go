package testing

import (
	"time"

	"github.com/bloxapp/pools-network/app"
	"github.com/tendermint/tendermint/abci/types"
	"github.com/tendermint/tendermint/libs/json"
	log2 "github.com/tendermint/tendermint/libs/log"
	db2 "github.com/tendermint/tm-db"
)

func SetupAppForTesting(isCheckTx bool) *app.App {
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

	return newApp
}
