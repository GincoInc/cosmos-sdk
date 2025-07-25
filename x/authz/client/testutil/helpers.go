package authz

import (
	"github.com/GincoInc/cosmos-sdk/client"
	addresscodec "github.com/GincoInc/cosmos-sdk/codec/address"
	"github.com/GincoInc/cosmos-sdk/testutil"
	clitestutil "github.com/GincoInc/cosmos-sdk/testutil/cli"
	"github.com/GincoInc/cosmos-sdk/x/authz/client/cli"
)

func CreateGrant(clientCtx client.Context, args []string) (testutil.BufferWriter, error) {
	cmd := cli.NewCmdGrantAuthorization(addresscodec.NewBech32Codec("cosmos"))
	return clitestutil.ExecTestCLICmd(clientCtx, cmd, args)
}
