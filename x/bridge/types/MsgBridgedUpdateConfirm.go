package types

import sdk "github.com/cosmos/cosmos-sdk/types"

func (msg MsgBridgedUpdateConfirm) Route() string {
	return RouterKey
}

func (msg MsgBridgedUpdateConfirm) Type() string {
	return "MsgBridgedUpdateConfirm"
}

func (msg MsgBridgedUpdateConfirm) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{} // TODO
}

func (msg MsgBridgedUpdateConfirm) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(&msg)
	return sdk.MustSortJSON(bz)
}

func (msg MsgBridgedUpdateConfirm) ValidateBasic() error {
	return nil
}
