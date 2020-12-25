package keeper

import (
	"testing"
	"time"

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

	db := db2.NewMemDB()
	ms := store.NewCommitMultiStore(db)
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
	poolsKeeper := NewKeeper(
		cdc,
		poolsKey,
	)

	return poolsKeeper, ctx
}
