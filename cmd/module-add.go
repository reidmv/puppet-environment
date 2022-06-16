package cmd

import (
	"log"

	"github.com/reidmv/puppet-environment/internal/environment"
	"github.com/spf13/cobra"
)

func init() {
	moduleCmd.AddCommand(moduleAddCmd)
	moduleAddCmd.Flags().StringVarP(&environmentFlag, "environment", "e", "", "Puppet code environment")
	moduleAddCmd.Flags().StringVarP(&typeFlag, "type", "t", "", "Module type")
	moduleAddCmd.Flags().StringVarP(&sourceFlag, "source", "s", "", "Module source")
	moduleAddCmd.Flags().StringVarP(&versionFlag, "version", "v", "", "Module version")
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
		name := args[0]
		env, ok := environmentsFile.Environments[environmentFlag]
		if !ok {
			log.Fatal("Environment does not exist!")
		}
		if _, ok = env.Modules[name]; ok {
			log.Fatal("Module already exists!")
		}
		env.Modules[name] = &environment.Module{
			Type:    typeFlag,
			Source:  sourceFlag,
			Version: versionFlag,
		}

		environmentsFile.Write()
	},
}
