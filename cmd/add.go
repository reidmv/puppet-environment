package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(addCmd)
	addCmd.Flags().StringVarP(&typeFlag, "type", "t", "", "Module type")
	addCmd.Flags().StringVarP(&sourceFlag, "source", "s", "", "Module source")
	addCmd.Flags().StringVarP(&versionFlag, "version", "v", "", "Module version")
	addCmd.MarkFlagRequired("type")
	addCmd.MarkFlagRequired("source")
}

var addCmd = &cobra.Command{
	Use:   "add [name]",
	Short: "Add a new Puppet code environment",
	Long:  `Add a new Puppet code environment definition to environments.yaml, and deploy it`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Not Implemented")
	},
}