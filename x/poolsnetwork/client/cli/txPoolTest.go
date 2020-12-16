package cli

import (
  "strconv"
	"github.com/spf13/cobra"

    "github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	"github.com/bloxapp/pools-network/x/poolsnetwork/types"
)

func CmdCreatePoolTest() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "create-poolTest [pool_id] [pubKey] [slashed] [exited] [ssvCommittee]",
		Short: "Creates a new poolTest",
		Args:  cobra.ExactArgs(5),
		RunE: func(cmd *cobra.Command, args []string) error {
      argsPool_id := string(args[0])
      argsPubKey := string(args[1])
      argsSlashed, _ := strconv.ParseBool(args[2])
      argsExited, _ := strconv.ParseBool(args[3])
      argsSsvCommittee := string(args[4])
      
        	clientCtx := client.GetClientContextFromCmd(cmd)
			clientCtx, err := client.ReadTxCommandFlags(clientCtx, cmd.Flags())
			if err != nil {
				return err
			}

			msg := types.NewMsgPoolTest(clientCtx.GetFromAddress(), string(argsPool_id), string(argsPubKey), bool(argsSlashed), bool(argsExited), string(argsSsvCommittee))
			if err := msg.ValidateBasic(); err != nil {
				return err
			}
			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)

    return cmd
}
