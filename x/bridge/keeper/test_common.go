package keeper

import (
	"testing"
	"time"

	bankTypes "github.com/cosmos/cosmos-sdk/x/bank/types"

	poolsnetworkkeeper "github.com/bloxapp/pools-network/x/poolsnetwork/keeper"
	poolsTypes "github.com/bloxapp/pools-network/x/poolsnetwork/types"

	bankkeeper "github.com/cosmos/cosmos-sdk/x/bank/keeper"

	distrtypes "github.com/cosmos/cosmos-sdk/x/distribution/types"
	govtypes "github.com/cosmos/cosmos-sdk/x/gov/types"
	ibctransfertypes "github.com/cosmos/cosmos-sdk/x/ibc/applications/transfer/types"
	minttypes "github.com/cosmos/cosmos-sdk/x/mint/types"

	authkeeper "github.com/cosmos/cosmos-sdk/x/auth/keeper"
	authTypes "github.com/cosmos/cosmos-sdk/x/auth/types"

	paramsTypes "github.com/cosmos/cosmos-sdk/x/params/types"

	paramskeeper "github.com/cosmos/cosmos-sdk/x/params/keeper"

	tmproto "github.com/tendermint/tendermint/proto/tendermint/types"

	"github.com/stretchr/testify/require"

	"github.com/cosmos/cosmos-sdk/store"
	db2 "github.com/tendermint/tm-db"

	bridgeTypes "github.com/bloxapp/pools-network/x/bridge/types"

	"github.com/cosmos/cosmos-sdk/x/bank"

	"github.com/cosmos/cosmos-sdk/x/staking"

	"github.com/cosmos/cosmos-sdk/codec"

	stakingKeeper "github.com/cosmos/cosmos-sdk/x/staking/keeper"

	sdk "github.com/cosmos/cosmos-sdk/types"
	stakingTypes "github.com/cosmos/cosmos-sdk/x/staking/types"
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
	bridgeTypes.RegisterCodec(legacy)

	return ret
}

func CreateTestEnv(t *testing.T) (Keeper, sdk.Context) {
	t.Helper()

	bridgeKey := sdk.NewKVStoreKey(bridgeTypes.StoreKey)
	stakingKey := sdk.NewKVStoreKey(stakingTypes.StoreKey)
	authKey := sdk.NewKVStoreKey(authTypes.StoreKey)
	bankKey := sdk.NewKVStoreKey(bankTypes.StoreKey)
	paramsKey := sdk.NewKVStoreKey(paramsTypes.StoreKey)
	tKeyParams := sdk.NewTransientStoreKey(paramsTypes.TStoreKey)
	poolsKey := sdk.NewTransientStoreKey(poolsTypes.StoreKey)

	db := db2.NewMemDB()
	ms := store.NewCommitMultiStore(db)
	ms.MountStoreWithDB(bridgeKey, sdk.StoreTypeIAVL, db)
	ms.MountStoreWithDB(stakingKey, sdk.StoreTypeIAVL, db)
	ms.MountStoreWithDB(bankKey, sdk.StoreTypeIAVL, db)
	ms.MountStoreWithDB(paramsKey, sdk.StoreTypeIAVL, db)
	ms.MountStoreWithDB(poolsKey, sdk.StoreTypeIAVL, db)
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

	poolsKeeper := poolsnetworkkeeper.NewKeeper(
		cdc,
		poolsKey,
	)

	bridgeKeeper := NewKeeper(
		cdc,
		paramsKeeper.Subspace(bridgeTypes.ModuleName),
		bridgeKey,
		stakingKeeper,
		poolsKeeper,
	)

	return bridgeKeeper, ctx
}
