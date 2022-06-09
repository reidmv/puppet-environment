package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(removeCmd)
}

var removeCmd = &cobra.Command{
	Use:   "remove [name]",
	Short: "Remove an existing Puppet code environment",
	Long:  `Remove a Puppet code environment defined in environments.yaml`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Not Implemented")
	},
}
