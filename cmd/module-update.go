package cmd

import (
	"log"

	"github.com/reidmv/puppet-environment/internal/environment"
	"github.com/reidmv/puppet-environment/internal/r10k"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

func init() {
	moduleCmd.AddCommand(moduleUpdateCmd)
	moduleUpdateCmd.Flags().StringVarP(&environmentFlag, "environment", "e", "", "Puppet code environment")
	moduleUpdateCmd.Flags().StringVarP(&typeFlag, "type", "t", "", "Module type")
	moduleUpdateCmd.Flags().StringVarP(&sourceFlag, "source", "s", "", "Module source")
	moduleUpdateCmd.Flags().StringVarP(&versionFlag, "version", "v", "", "Module version")
	moduleUpdateCmd.MarkFlagRequired("environment")
}

var moduleUpdateCmd = &cobra.Command{
	Use:   "update [name]",
	Short: "Update module for a Puppet code environment",
	Long:  `Update module for a Puppet code environment defined in environments.yaml`,
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		name := args[0]
		env, ok := environmentsFile.Environments[environmentFlag]
		if !ok {
			log.Fatalf("Environment %s does not exist!", environmentFlag)
		}
		if _, ok = env.Modules[name]; !ok {
			log.Fatalf("Module %s does not exist!", name)
		}
		if cmd.Flags().Changed("type") {
			env.Modules[name].Type = typeFlag
		}
		if cmd.Flags().Changed("source") {
			env.Modules[name].Source = sourceFlag
		}
		if cmd.Flags().Changed("version") {
			env.Modules[name].Version = versionFlag
		}

		environmentsFile.Write()
		err := r10k.DeployModule(environmentFlag, environment.ModuleName(name), environmentsFile.Path, viper.GetString("environments-path"))
		if err != nil {
			log.Fatal(err)
		}
	},
}
