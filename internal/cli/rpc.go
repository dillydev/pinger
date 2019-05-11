package cli

import (
	"log"

	"github.com/dgparker/pinger/internal/rpc"

	"github.com/spf13/cobra"
)

var rpcCmd = &cobra.Command{
	Use: "rpc",
	Short: "rpcShortDesc",
	Long: "rpcLongDesc",
	Run: startRPC,
}

func init() {
	RootCmd.AddCommand(rpcCmd)
}

func startRPC(cmd *cobra.Command, args []string) {
	err := rpc.Run()
	if err != nil {
		log.Fatal(err)
	}
}
