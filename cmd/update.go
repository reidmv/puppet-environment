package cmd

import (
	"log"

	"github.com/reidmv/puppet-environment/internal/r10k"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

func init() {
	rootCmd.AddCommand(updateCmd)
	updateCmd.Flags().StringVarP(&typeFlag, "type", "t", "", "Module type")
	updateCmd.Flags().StringVarP(&sourceFlag, "source", "s", "", "Module source")
	updateCmd.Flags().StringVarP(&versionFlag, "version", "v", "", "Module version")
}

var updateCmd = &cobra.Command{
	Use:   "update [name]",
	Short: "Update a Puppet code environment",
	Long:  `Update a deployed Puppet code environment defined in environments.yaml`,
	Run: func(cmd *cobra.Command, args []string) {
		name := args[0]
		env, ok := environmentsFile.Environments[name]
		if !ok {
			log.Fatal("Environment does not exist!")
		}
		if cmd.Flags().Changed("type") {
			env.Type = typeFlag
		}
		if cmd.Flags().Changed("source") {
			env.Source = sourceFlag
		}
		if cmd.Flags().Changed("version") {
			env.Version = versionFlag
		}

		environmentsFile.Write()
		err := r10k.DeployEnvironment(name, environmentsFile.Path, viper.GetString("environments-root"))
		if err != nil {
			log.Fatal(err)
		}
	},
}
