package keeper

import (
	"testing"
	"time"

	distrtypes "github.com/cosmos/cosmos-sdk/x/distribution/types"
	govtypes "github.com/cosmos/cosmos-sdk/x/gov/types"
	ibctransfertypes "github.com/cosmos/cosmos-sdk/x/ibc/applications/transfer/types"
	minttypes "github.com/cosmos/cosmos-sdk/x/mint/types"

	authkeeper "github.com/cosmos/cosmos-sdk/x/auth/keeper"
	authTypes "github.com/cosmos/cosmos-sdk/x/auth/types"
	bankkeeper "github.com/cosmos/cosmos-sdk/x/bank/keeper"
	bankTypes "github.com/cosmos/cosmos-sdk/x/bank/types"
	paramskeeper "github.com/cosmos/cosmos-sdk/x/params/keeper"
	paramsTypes "github.com/cosmos/cosmos-sdk/x/params/types"
	stakingKeeper "github.com/cosmos/cosmos-sdk/x/staking/keeper"
	stakingTypes "github.com/cosmos/cosmos-sdk/x/staking/types"

	poolsTypes "github.com/bloxapp/pools-network/x/poolsnetwork/types"
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/store"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/bank"
	"github.com/cosmos/cosmos-sdk/x/staking"
	"github.com/stretchr/testify/require"
	tmproto "github.com/tendermint/tendermint/proto/tendermint/types"
	db2 "github.com/tendermint/tm-db"
)

var TestingStakeParams = stakingTypes.Params{
	UnbondingTime:     100,
	MaxValidators:     10,
	MaxEntries:        10,
	HistoricalEntries: 10,
	BondDenom:         "stake",
}

var maccPerms = map[string][]string{
	authTypes.FeeCollectorName:     nil,
	distrtypes.ModuleName:          nil,
	minttypes.ModuleName:           {authTypes.Minter},
	stakingTypes.BondedPoolName:    {authTypes.Burner, authTypes.Staking},
	stakingTypes.NotBondedPoolName: {authTypes.Burner, authTypes.Staking},
	govtypes.ModuleName:            {authTypes.Burner},
	ibctransfertypes.ModuleName:    {authTypes.Minter, authTypes.Burner},
}

func Codecs() *codec.AminoCodec {
	legacy := codec.NewLegacyAmino()
	ret := codec.NewAminoCodec(legacy)
	staking.AppModuleBasic{}.RegisterLegacyAminoCodec(legacy)
	bank.AppModuleBasic{}.RegisterLegacyAminoCodec(legacy)
	poolsTypes.RegisterCodec(legacy)

	return ret
}

func CreateTestEnv(t *testing.T) (Keeper, sdk.Context) {
	t.Helper()

	poolsKey := sdk.NewTransientStoreKey(poolsTypes.StoreKey)
	authKey := sdk.NewKVStoreKey(authTypes.StoreKey)
	stakingKey := sdk.NewKVStoreKey(stakingTypes.StoreKey)
	bankKey := sdk.NewKVStoreKey(bankTypes.StoreKey)
	paramsKey := sdk.NewKVStoreKey(paramsTypes.StoreKey)
	tKeyParams := sdk.NewTransientStoreKey(paramsTypes.TStoreKey)

	db := db2.NewMemDB()
	ms := store.NewCommitMultiStore(db)
	ms.MountStoreWithDB(poolsKey, sdk.StoreTypeIAVL, db)
	ms.MountStoreWithDB(stakingKey, sdk.StoreTypeIAVL, db)
	ms.MountStoreWithDB(bankKey, sdk.StoreTypeIAVL, db)
	ms.MountStoreWithDB(paramsKey, sdk.StoreTypeIAVL, db)
	require.NoError(t, ms.LoadLatestVersion())

	// context
	const isCheckTx = false
	ctx := sdk.NewContext(ms, tmproto.Header{
		Height: 1234567,
		Time:   time.Date(2020, time.April, 22, 12, 0, 0, 0, time.UTC),
	}, isCheckTx, nil)

	// cdc
	cdc := Codecs()

	// keepers
	paramsKeeper := paramskeeper.NewKeeper(cdc, cdc.LegacyAmino, paramsKey, tKeyParams)
	accountKeeper := authkeeper.NewAccountKeeper(
		cdc,
		authKey,
		paramsKeeper.Subspace(authTypes.ModuleName),
		authTypes.ProtoBaseAccount,
		maccPerms,
	)
	bankKeeper := bankkeeper.NewBaseKeeper(
		cdc,
		bankKey,
		accountKeeper,
		paramsKeeper.Subspace(bankTypes.ModuleName),
		map[string]bool{},
	)

	stakingKeeper := stakingKeeper.NewKeeper(
		cdc,
		stakingKey,
		accountKeeper,
		bankKeeper,
		paramsKeeper.Subspace(stakingTypes.ModuleName),
	)

	poolsKeeper := NewKeeper(
		cdc,
		poolsKey,
		stakingKeeper,
	)

	return poolsKeeper, ctx
}
