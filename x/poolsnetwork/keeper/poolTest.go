package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/bloxapp/pools-network/x/poolsnetwork/types"
	"github.com/cosmos/cosmos-sdk/store/prefix"
)

func (k Keeper) CreatePoolTest(ctx sdk.Context, poolTest types.MsgPoolTest) {
	store :=  prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.PoolTestKey))
	b := k.cdc.MustMarshalBinaryBare(&poolTest)
	store.Set(types.KeyPrefix(types.PoolTestKey), b)
}

func (k Keeper) GetAllPoolTest(ctx sdk.Context) (msgs []types.MsgPoolTest) {
    store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.PoolTestKey))
	iterator := sdk.KVStorePrefixIterator(store, types.KeyPrefix(types.PoolTestKey))

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var msg types.MsgPoolTest
		k.cdc.MustUnmarshalBinaryBare(iterator.Value(), &msg)
        msgs = append(msgs, msg)
	}

    return
}
