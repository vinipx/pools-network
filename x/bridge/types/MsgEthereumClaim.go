package types

import (
	"github.com/bloxapp/pools-network/shared/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

var _ sdk.Msg = &MsgEthereumClaim{}

func NewMsgEthereumClaim(
	chainId uint64,
	contractAddress *types.EthereumAddress,
	pubKey *types.ConsensusAddress,
) *MsgEthereumClaim {
	return &MsgEthereumClaim{
		EthereumChainId: chainId,
		ContractAddress: contractAddress,
		ConsensusPubkey: pubKey,
		Delegates:       make([]*DelegateVote, 0),
		Undelegates:     make([]*UnDelegateVote, 0),
		CreatePools:     make([]*CreatePool, 0),
		CreateOperators: make([]*CreateOperator, 0),
	}
}

func (msg *MsgEthereumClaim) AddDelegate(d *DelegateVote) *MsgEthereumClaim {
	msg.Delegates = append(msg.Delegates, d)
	return msg
}

func (msg *MsgEthereumClaim) AddUnDelegate(d *UnDelegateVote) *MsgEthereumClaim {
	msg.Undelegates = append(msg.Undelegates, d)
	return msg
}

func (msg *MsgEthereumClaim) AddCreatePool(p *CreatePool) *MsgEthereumClaim {
	msg.CreatePools = append(msg.CreatePools, p)
	return msg
}

func (msg *MsgEthereumClaim) AddCreateOperator(o *CreateOperator) *MsgEthereumClaim {
	msg.CreateOperators = append(msg.CreateOperators, o)
	return msg
}

func (msg *MsgEthereumClaim) Route() string {
	return RouterKey
}

func (msg *MsgEthereumClaim) Type() string {
	return "EthereumClaim"
}

func (msg *MsgEthereumClaim) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{sdk.AccAddress(*msg.ConsensusPubkey)}
}

func (msg *MsgEthereumClaim) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgEthereumClaim) ValidateBasic() error {
	// TODO
	return nil
}
