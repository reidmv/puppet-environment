package cmd

import (
	"log"

	"github.com/reidmv/puppet-environment/internal/environment"
	"github.com/reidmv/puppet-environment/internal/r10k"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
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
			log.Fatalf("Environment %s does not exist!", environmentFlag)
		}
		if incl, key := env.Modules.Include(name); incl {
			if name == key {
				log.Fatalf("Module %s already exists!", name)
			} else {
				log.Fatalf("Cannot add %s; name conflicts with %s", name, key)
			}
		}
		env.Modules.Set(name, &environment.Module{
			Type:    typeFlag,
			Source:  sourceFlag,
			Version: versionFlag,
		})

		environmentsFile.Write()
		err := r10k.DeployModule(environmentFlag, name, environmentsFile.Path, viper.GetString("environments-root"))
		if err != nil {
			log.Fatal(err)
		}
	},
}
