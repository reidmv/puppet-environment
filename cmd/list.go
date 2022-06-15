package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(listCmd)
}

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List all Puppet code environments",
	Long:  `List each Puppet code environment defined in environments.yaml`,
	Run: func(cmd *cobra.Command, args []string) {
		for name, _ := range environmentsFile.Environments {
			fmt.Println(name)
		}
	},
}
