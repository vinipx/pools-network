package keeper

import (
	"github.com/bloxapp/pools-network/x/poolsnetwork/types"
)

var _ types.QueryServer = Keeper{}
