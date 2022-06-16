package cmd

import (
	"log"

	"github.com/spf13/cobra"
)

func init() {
	moduleCmd.AddCommand(moduleRemoveCmd)
	moduleRemoveCmd.Flags().StringVarP(&environmentFlag, "environment", "e", "", "Puppet code environment")
	moduleRemoveCmd.MarkFlagRequired("environment")
}

var moduleRemoveCmd = &cobra.Command{
	Use:   "remove [name]",
	Short: "Remove module from a Puppet code environment",
	Long:  `Remove module from a Puppet code environment defined in environments.yaml`,
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		name := args[0]
		env, ok := environmentsFile.Environments[environmentFlag]
		if !ok {
			log.Fatal("Environment does not exist!")
		}
		if _, ok = env.Modules[name]; !ok {
			log.Fatal("Module does not exist!")
		}
		delete(env.Modules, name)

		environmentsFile.Write()
	},
}
