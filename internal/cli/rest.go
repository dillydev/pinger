package cli

import (
	"log"

	"github.com/dgparker/pinger/internal/rest"

	"github.com/spf13/cobra"
)

var restCmd = &cobra.Command{
	Use: "rest",
	Short: "rpcShortDesc",
	Long: "rpcLongDesc",
	Run: startRest,
}

func init() {
	RootCmd.AddCommand(restCmd)
}

func startRest(cmd *cobra.Command, args []string) {
	err := rest.Run()
	if err != nil {
		log.Fatal(err)
	}
}
