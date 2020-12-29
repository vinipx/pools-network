package testing

import (
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/codec/types"
	"github.com/cosmos/cosmos-sdk/std"
	"github.com/cosmos/cosmos-sdk/types/module"
	"github.com/cosmos/cosmos-sdk/x/auth/tx"
)

// EncodingConfig specifies the concrete encoding types to use for a given app.
// This is provided for compatibility between protobuf and amino implementations.
type EncodingConfig struct {
	InterfaceRegistry types.InterfaceRegistry
	Marshaler         codec.Marshaler
	TxConfig          client.TxConfig
	Amino             *codec.LegacyAmino
}

// MakeEncodingConfig creates an EncodingConfig for an amino based test configuration.
func MakeEncodingConfig() EncodingConfig {
	amino := codec.NewLegacyAmino()
	interfaceRegistry := types.NewInterfaceRegistry()
	marshaler := codec.NewProtoCodec(interfaceRegistry)
	txCfg := tx.NewTxConfig(marshaler, tx.DefaultSignModes)

	return EncodingConfig{
		InterfaceRegistry: interfaceRegistry,
		Marshaler:         marshaler,
		TxConfig:          txCfg,
		Amino:             amino,
	}
}

func Codecs(manager module.BasicManager) EncodingConfig {
	encodingConfig := MakeEncodingConfig()
	std.RegisterLegacyAminoCodec(encodingConfig.Amino)
	std.RegisterInterfaces(encodingConfig.InterfaceRegistry)
	manager.RegisterLegacyAminoCodec(encodingConfig.Amino)
	manager.RegisterInterfaces(encodingConfig.InterfaceRegistry)

	//legacy := codec.NewLegacyAmino()
	//ret := codec.NewAminoCodec(legacy)
	//staking.AppModuleBasic{}.RegisterLegacyAminoCodec(legacy)
	//bank.AppModuleBasic{}.RegisterLegacyAminoCodec(legacy)
	//auth.AppModuleBasic{}.RegisterLegacyAminoCodec(legacy)
	//poolsTypes.RegisterCodec(legacy)
	//codec2.RegisterCrypto(legacy)
	return encodingConfig
}
