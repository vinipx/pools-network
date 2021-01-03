package types

import (
	"github.com/bloxapp/pools-network/shared/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var _ sdk.Msg = &MsgEthereumClaim{}

func NewMsgEthereumClaim(
	nonce uint64,
	chainId uint64,
	contractAddress types.EthereumAddress,
	consensusAddress types.ConsensusAddress,
) *MsgEthereumClaim {
	return &MsgEthereumClaim{
		Nonce:            nonce,
		EthereumChainId:  chainId,
		ContractAddress:  contractAddress,
		ConsensusAddress: consensusAddress,
		Data:             make([]ClaimData, 0),
	}
}

func (msg *MsgEthereumClaim) AddClaim(d ClaimData) *MsgEthereumClaim {
	msg.Data = append(msg.Data, d)
	return msg
}

func (msg MsgEthereumClaim) Route() string {
	return RouterKey
}

func (msg MsgEthereumClaim) Type() string {
	return "MsgEthereumClaim"
}

func (msg MsgEthereumClaim) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{sdk.AccAddress(msg.ConsensusAddress)}
}

func (msg MsgEthereumClaim) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(&msg)
	return sdk.MustSortJSON(bz)
}

func (msg MsgEthereumClaim) ValidateBasic() error {
	if len(msg.ContractAddress) == 0 {
		return sdkerrors.Wrap(ErrClaimDataInvalid, "Contract address is invalid")
	}
	if len(msg.ConsensusAddress) == 0 {
		return sdkerrors.Wrap(ErrClaimDataInvalid, "Consensus address is invalid")
	}

	for _, c := range msg.Data {
		if err := c.ValidateBasic(); err != nil {
			return err
		}
	}
	return nil
}

func (c ClaimData) ValidateBasic() error {
	if c.TxHash == nil || len(c.TxHash) == 0 {
		return sdkerrors.Wrap(ErrClaimDataInvalid, "Tx hash is invalid")
	}

	switch c.ClaimType {
	case ClaimType_Delegate, ClaimType_Undelegate:
		if len(c.EthereumAddresses) != 2 {
			return sdkerrors.Wrap(ErrClaimDataInvalid, "Delegate/ Undelegate: Ethereum addresses length must be 2")
		}
		if err := c.EthereumAddresses[0].Validate(); err != nil {
			return sdkerrors.Wrap(ErrClaimDataInvalid, "Delegate/ Undelegate: Ethereum addresses invalid")
		}
		if err := c.EthereumAddresses[1].Validate(); err != nil {
			return sdkerrors.Wrap(ErrClaimDataInvalid, "Delegate/ Undelegate: Ethereum addresses invalid")
		}
		if len(c.Values) != 1 {
			return sdkerrors.Wrap(ErrClaimDataInvalid, "Delegate/ Undelegate: values length must be 1")
		}
	case ClaimType_CreatePool:
		return nil // TODO
	case ClaimType_CreateOperator:
		if len(c.EthereumAddresses) != 1 {
			return sdkerrors.Wrap(ErrClaimDataInvalid, "CreateOperator: Ethereum addresses length must be 1")
		}
		if err := c.EthereumAddresses[0].Validate(); err != nil {
			return sdkerrors.Wrap(ErrClaimDataInvalid, "CreateOperator: Ethereum addresses invalid")
		}
		if len(c.ConsensusAddresses) != 1 {
			return sdkerrors.Wrap(ErrClaimDataInvalid, "CreateOperator: Consensus addresses length must be 1")
		}
		if err := c.ConsensusAddresses[0].Validate(); err != nil {
			return sdkerrors.Wrap(ErrClaimDataInvalid, "CreateOperator: Consensus addresses invalid")
		}
		if len(c.Values) != 2 {
			return sdkerrors.Wrap(ErrClaimDataInvalid, "CreateOperator: values length must be 2")
		}
	default:
		return sdkerrors.Wrap(ErrClaimDataInvalid, "Unknown claim type")
	}

	return nil
}
