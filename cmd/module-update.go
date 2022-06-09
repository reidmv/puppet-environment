package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

func init() {
	moduleCmd.AddCommand(moduleUpdateCmd)
	moduleUpdateCmd.Flags().StringVarP(&EnvironmentFlag, "environment", "e", "", "Puppet code environment")
	moduleUpdateCmd.Flags().StringVarP(&TypeFlag, "type", "t", "", "Module type")
	moduleUpdateCmd.Flags().StringVarP(&SourceFlag, "source", "s", "", "Module source")
	moduleUpdateCmd.Flags().StringVarP(&VersionFlag, "version", "v", "", "Module version")
	moduleUpdateCmd.MarkFlagRequired("environment")
	moduleUpdateCmd.MarkFlagRequired("type")
	moduleUpdateCmd.MarkFlagRequired("source")
}

var moduleUpdateCmd = &cobra.Command{
	Use:   "update [name]",
	Short: "Update module for a Puppet code environment",
	Long:  `Update module for a Puppet code environment defined in environments.yaml`,
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Not Implemented")
	},
}
