package keeper

import (
	"context"

	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/query"
	"github.com/bloxapp/pools-network/x/poolsnetwork/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) AllPoolTest(c context.Context, req *types.QueryAllPoolTestRequest) (*types.QueryAllPoolTestResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	var poolTests []*types.MsgPoolTest
	ctx := sdk.UnwrapSDKContext(c)

	store := ctx.KVStore(k.storeKey)
	poolTestStore := prefix.NewStore(store, types.KeyPrefix(types.PoolTestKey))

	pageRes, err := query.Paginate(poolTestStore, req.Pagination, func(key []byte, value []byte) error {
		var poolTest types.MsgPoolTest
		if err := k.cdc.UnmarshalBinaryBare(value, &poolTest); err != nil {
			return err
		}

		poolTests = append(poolTests, &poolTest)
		return nil
	})

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryAllPoolTestResponse{PoolTest: poolTests, Pagination: pageRes}, nil
}
