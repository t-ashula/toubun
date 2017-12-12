package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var initCommand = &cobra.Command{
	Use:   "init",
	Short: "initialize toubun config",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Print("not yet.")
	},
}

func init() {
	RootCommand.AddCommand(initCommand)
}
