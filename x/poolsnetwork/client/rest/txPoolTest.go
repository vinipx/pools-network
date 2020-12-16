package rest

import (
	"net/http"
	"strconv"

    "github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/tx"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/rest"
	"github.com/bloxapp/pools-network/x/poolsnetwork/types"
)

// Used to not have an error if strconv is unused
var _ = strconv.Itoa(42)

type createPoolTestRequest struct {
	BaseReq rest.BaseReq `json:"base_req"`
	Creator string `json:"creator"`
	Pool_id string `json:"pool_id"`
	PubKey string `json:"pubKey"`
	Slashed string `json:"slashed"`
	Exited string `json:"exited"`
	SsvCommittee string `json:"ssvCommittee"`
	
}

func createPoolTestHandler(clientCtx client.Context) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req createPoolTestRequest
		if !rest.ReadRESTReq(w, r, clientCtx.LegacyAmino, &req) {
			rest.WriteErrorResponse(w, http.StatusBadRequest, "failed to parse request")
			return
		}

		baseReq := req.BaseReq.Sanitize()
		if !baseReq.ValidateBasic(w) {
			return
		}

		creator, err := sdk.AccAddressFromBech32(req.Creator)
		if err != nil {
			rest.WriteErrorResponse(w, http.StatusBadRequest, err.Error())
			return
		}

		
		parsedPool_id := req.Pool_id
		
		parsedPubKey := req.PubKey
		
		parsedSlashed, err := strconv.ParseBool(req.Slashed)
		if err != nil {
			rest.WriteErrorResponse(w, http.StatusBadRequest, err.Error())
			return
		}	
			
		parsedExited, err := strconv.ParseBool(req.Exited)
		if err != nil {
			rest.WriteErrorResponse(w, http.StatusBadRequest, err.Error())
			return
		}	
			
		parsedSsvCommittee := req.SsvCommittee
		

		msg := types.NewMsgPoolTest(
			creator,
			parsedPool_id,
			parsedPubKey,
			parsedSlashed,
			parsedExited,
			parsedSsvCommittee,
			
		)

		tx.WriteGeneratedTxResponse(clientCtx, w, req.BaseReq, msg)
	}
}
