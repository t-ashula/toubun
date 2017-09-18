package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

// RootCommand represents the base command when called without any subcommands
var RootCommand = &cobra.Command{
	Use:   "toubun",
	Short: "Update your code dependency libraries",
	Long:  "Update your code dependency libraries",
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := RootCommand.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize()
}
