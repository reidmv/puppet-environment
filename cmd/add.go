package cmd

import (
	"log"

	"github.com/reidmv/puppet-environment/internal/environment"
	"github.com/reidmv/puppet-environment/internal/r10k"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
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
		name := args[0]
		if _, ok := environmentsFile.Environments[name]; ok {
			log.Fatal("Environment already exists!")
		}

		environmentsFile.Environments[name] = &environment.Environment{
			Type:    typeFlag,
			Source:  sourceFlag,
			Version: versionFlag,
		}

		environmentsFile.Write()
		err := r10k.DeployEnvironment(name, environmentsFile.Path, viper.GetString("environments-path"))
		if err != nil {
			log.Fatal(err)
		}
	},
}
