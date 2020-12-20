package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/google/uuid"
)

var _ sdk.Msg = &MsgPoolTest{}

func NewMsgPoolTest(creator sdk.AccAddress, pool_id string, pubKey string, slashed bool, exited bool, ssvCommittee string) *MsgPoolTest {
	return &MsgPoolTest{
		Id:           uuid.New().String(),
		Creator:      creator,
		PubKey:       pubKey,
		Slashed:      slashed,
		Exited:       exited,
		SsvCommittee: ssvCommittee,
	}
}

func (msg *MsgPoolTest) Route() string {
	return RouterKey
}

func (msg *MsgPoolTest) Type() string {
	return "CreatePoolTest"
}

func (msg *MsgPoolTest) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{sdk.AccAddress(msg.Creator)}
}

func (msg *MsgPoolTest) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgPoolTest) ValidateBasic() error {
	if msg.Creator.Empty() {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, "creator can't be empty")
	}
	return nil
}
