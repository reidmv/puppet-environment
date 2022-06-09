package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

func init() {
	moduleCmd.AddCommand(moduleRemoveCmd)
	moduleRemoveCmd.Flags().StringVarP(&EnvironmentFlag, "environment", "e", "", "Puppet code environment")
	moduleRemoveCmd.MarkFlagRequired("environment")
}

var moduleRemoveCmd = &cobra.Command{
	Use:   "remove [name]",
	Short: "Remove module from a Puppet code environment",
	Long:  `Remove module from a Puppet code environment defined in environments.yaml`,
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Not Implemented")
	},
}
