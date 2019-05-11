package cli

import (
	"log"

	"github.com/dgparker/pinger/internal/rest"
	"github.com/dgparker/pinger/internal/rpc"

	"github.com/spf13/cobra"
)

var dualCmd = &cobra.Command{
	Use: "dual",
	Short: "dualShortDesc",
	Long: "dualLongDesc",
	Run: startDual,
}

func init() {
	RootCmd.AddCommand(dualCmd)
}

func startDual(cmd *cobra.Command, args []string) {
	go func() {
		err := rest.Run()
		if err != nil {
			log.Fatal(err)
		}
	}()

	err := rpc.Run()
	if err != nil {
		log.Fatal(err)
	}
}
