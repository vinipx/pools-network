package keeper

import (
	"github.com/bloxapp/pools-network/x/bridge/types"
)

var _ types.QueryServer = Keeper{}
