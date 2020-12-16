package rest

import (
	"github.com/gorilla/mux"

	"github.com/cosmos/cosmos-sdk/client"
    "github.com/bloxapp/pools-network/x/poolsnetwork/types"
)

const (
    MethodGet = "GET"
)

// RegisterRoutes registers poolsnetwork-related REST handlers to a router
func RegisterRoutes(clientCtx client.Context, r *mux.Router) {
    // this line is used by starport scaffolding # 2
	registerQueryRoutes(clientCtx, r)
	registerTxHandlers(clientCtx, r)

}

func registerQueryRoutes(clientCtx client.Context, r *mux.Router) {
    // this line is used by starport scaffolding # 3
    r.HandleFunc("custom/poolsnetwork/" + types.QueryListPoolTest, listPoolTestHandler(clientCtx)).Methods("GET")

}

func registerTxHandlers(clientCtx client.Context, r *mux.Router) {
    // this line is used by starport scaffolding # 4
    r.HandleFunc("/poolsnetwork/poolTest", createPoolTestHandler(clientCtx)).Methods("POST")

}

