package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

func init() {
	moduleCmd.AddCommand(moduleAddCmd)
	moduleAddCmd.Flags().StringVarP(&EnvironmentFlag, "environment", "e", "", "Puppet code environment")
	moduleAddCmd.Flags().StringVarP(&TypeFlag, "type", "t", "", "Module type")
	moduleAddCmd.Flags().StringVarP(&SourceFlag, "source", "s", "", "Module source")
	moduleAddCmd.Flags().StringVarP(&VersionFlag, "version", "v", "", "Module version")
	moduleAddCmd.MarkFlagRequired("environment")
	moduleAddCmd.MarkFlagRequired("type")
	moduleAddCmd.MarkFlagRequired("source")
}

var moduleAddCmd = &cobra.Command{
	Use:   "add [name]",
	Short: "Add module to a Puppet code environment",
	Long:  `Add module to a Puppet code environment defined in environments.yaml`,
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Not Implemented")
	},
}
