package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(updateCmd)
	updateCmd.Flags().StringVarP(&typeFlag, "type", "t", "", "Module type")
	updateCmd.Flags().StringVarP(&sourceFlag, "source", "s", "", "Module source")
	updateCmd.Flags().StringVarP(&versionFlag, "version", "v", "", "Module version")
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
