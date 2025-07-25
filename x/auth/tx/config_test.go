package tx_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"

	"github.com/GincoInc/cosmos-sdk/codec"
	"github.com/GincoInc/cosmos-sdk/codec/testutil"
	"github.com/GincoInc/cosmos-sdk/codec/types"
	"github.com/GincoInc/cosmos-sdk/std"
	"github.com/GincoInc/cosmos-sdk/testutil/testdata"
	sdk "github.com/GincoInc/cosmos-sdk/types"
	"github.com/GincoInc/cosmos-sdk/x/auth/tx"
	txtestutil "github.com/GincoInc/cosmos-sdk/x/auth/tx/testutil"
)

func TestGenerator(t *testing.T) {
	interfaceRegistry := testutil.CodecOptions{}.NewInterfaceRegistry()
	std.RegisterInterfaces(interfaceRegistry)
	interfaceRegistry.RegisterImplementations((*sdk.Msg)(nil), &testdata.TestMsg{})
	protoCodec := codec.NewProtoCodec(interfaceRegistry)
	suite.Run(t, txtestutil.NewTxConfigTestSuite(tx.NewTxConfig(protoCodec, tx.DefaultSignModes)))
}

func TestConfigOptions(t *testing.T) {
	interfaceRegistry := types.NewInterfaceRegistry()
	protoCodec := codec.NewProtoCodec(interfaceRegistry)
	configOptions := tx.ConfigOptions{}
	txConfig, err := tx.NewTxConfigWithOptions(protoCodec, configOptions)
	require.NoError(t, err)
	require.NotNil(t, txConfig)
	handler := txConfig.SignModeHandler()
	require.NotNil(t, handler)
}
