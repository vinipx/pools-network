package keeper

import (
	"github.com/bloxapp/pools-network/x/reports/types"
)

var _ types.QueryServer = Keeper{}
