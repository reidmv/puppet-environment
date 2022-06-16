package cmd

import (
	"log"

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
		name := args[0]
		if _, ok := environmentsFile.Environments[name]; !ok {
			log.Fatal("Environment does not exist!")
		} else {
			delete(environmentsFile.Environments, name)
		}

		environmentsFile.Write()
	},
}
