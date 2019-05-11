package cli

import (
	"log"

	"github.com/spf13/cobra"
)

var (
	// RootCmd ...
	RootCmd = &cobra.Command{
		Use: "Pinger",
		Short: "shortDesc",
		Long: "longDesc",
	}
)

func init() {
	// initialize flags here ...
}

// Execute ...
func Execute() {
	err := RootCmd.Execute()
	if err != nil {
		log.Fatal(err)
	}
}
