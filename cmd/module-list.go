package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

func init() {
	moduleCmd.AddCommand(moduleListCmd)
	moduleListCmd.Flags().StringVarP(&EnvironmentFlag, "environment", "e", "", "Puppet code environment")
	moduleListCmd.MarkFlagRequired("environment")
}

var moduleListCmd = &cobra.Command{
	Use:   "list",
	Short: "List modules for a Puppet code environment",
	Long:  `List modules for a Puppet code environment defined in environments.yaml`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Not Implemented")
	},
}
