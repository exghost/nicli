package nicli

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "nicli",
	Short: "nicli - CLI tool to help facilitate NI development",
	Run: func(cmd *cobra.Command, args []string) {

	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintf(os.Stderr, "Uh oh, looks like you dun messed up! '%s'", err)
		os.Exit(1)
	}
}
