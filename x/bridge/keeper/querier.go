package keeper

import (
	// this line is used by starport scaffolding # 1
	bridgeTypes "github.com/bloxapp/pools-network/x/bridge/types"
	"github.com/cosmos/cosmos-sdk/codec"
	sdkTypes "github.com/cosmos/cosmos-sdk/types"
	sdkErrors "github.com/cosmos/cosmos-sdk/types/errors"

	abci "github.com/tendermint/tendermint/abci/types"
)

// NewQuerier implements the sdkTypes.Querier function to serve as module level router for state queries
func NewQuerier(k Keeper, legacyQuerierCdc *codec.LegacyAmino) sdkTypes.Querier {
	return func(ctx sdkTypes.Context, path []string, req abci.RequestQuery) ([]byte, error) {
		var (
			res []byte
			err error
		)

		switch path[0] {
		// this line is used by starport scaffolding # 1
		default:
			err = sdkErrors.Wrapf(sdkErrors.ErrUnknownRequest, "unknown %s query endpoint: %s", bridgeTypes.ModuleName, path[0])
		}

		return res, err
	}
}
