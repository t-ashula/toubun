package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	k "github.com/t-ashula/toubun/core"
)

var versionCommand = &cobra.Command{
	Use:   "version",
	Short: "print version string",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("%s %s\n", k.Name, k.Version)
	},
}

func init() {
	RootCommand.AddCommand(versionCommand)
}
