package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(updateCmd)
	updateCmd.Flags().StringVarP(&TypeFlag, "type", "t", "", "Module type")
	updateCmd.Flags().StringVarP(&SourceFlag, "source", "s", "", "Module source")
	updateCmd.Flags().StringVarP(&VersionFlag, "version", "v", "", "Module version")
	updateCmd.MarkFlagRequired("type")
	updateCmd.MarkFlagRequired("source")
}

var updateCmd = &cobra.Command{
	Use:   "update [name]",
	Short: "Update a Puppet code environment",
	Long:  `Update a deployed Puppet code environment defined in environments.yaml`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Not Implemented")
	},
}
