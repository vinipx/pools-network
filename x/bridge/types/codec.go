package types

import (
	"github.com/cosmos/cosmos-sdk/codec"
	cdctypes "github.com/cosmos/cosmos-sdk/codec/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	// this line is used by starport scaffolding # 1
)

func RegisterCodec(cdc *codec.LegacyAmino) {
	cdc.RegisterConcrete(&MsgEthereumClaim{}, "pools/MsgEthereumClaim", nil)
	cdc.RegisterConcrete(&MsgBridgedUpdateConfirm{}, "pools/MsgBridgedUpdateConfirm", nil)
	// this line is used by starport scaffolding # 2
}

func RegisterInterfaces(registry cdctypes.InterfaceRegistry) {
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgEthereumClaim{},
		&MsgBridgedUpdateConfirm{},
	)
	// this line is used by starport scaffolding # 3
}

var (
	amino     = codec.NewLegacyAmino()
	ModuleCdc = codec.NewAminoCodec(amino)
)
