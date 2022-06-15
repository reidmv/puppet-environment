package cmd

import (
	"errors"
	"fmt"
	"log"
	"os"

	"github.com/reidmv/puppet-environment/internal/environment"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

func init() {
	cobra.OnInitialize(initConfig)

	rootCmd.CompletionOptions.HiddenDefaultCmd = true
	rootCmd.PersistentFlags().StringVar(&configFlag, "config", "", "config file")
	rootCmd.PersistentFlags().StringVar(&environmentsFileFlag, "environments-file", "", "environments yaml file")

	viper.BindPFlag("environments-file", rootCmd.PersistentFlags().Lookup("environments-file"))
	viper.SetDefault("environments-file", "/etc/puppetlabs/puppet/environments.yaml")
}

var (
	environmentsFile     environment.EnvironmentsFile
	environmentFlag      string
	typeFlag             string
	sourceFlag           string
	versionFlag          string
	configFlag           string
	environmentsFileFlag string
)

var rootCmd = &cobra.Command{
	Use:   "puppet-environment",
	Short: "Manage Puppet environments and environment modules",
	Long: `The puppet-environment tool is used to manage and deploy Puppet code environments and modules.
		Environment definitions are stored in /etc/puppetlabs/puppet/environments.yaml.
		The r10k utility is used to instantiate environments defined there.`,
	PersistentPreRun: func(cmd *cobra.Command, args []string) {
		initEnvironmentsFile()
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func initConfig() {
	if configFlag != "" {
		// Use config file from the flag.
		viper.SetConfigFile(configFlag)
	} else {
		// Find home directory.
		home, err := os.UserHomeDir()
		cobra.CheckErr(err)

		// Search config in home directory with name ".puppet-environment" (without extension).
		viper.AddConfigPath(home)
		viper.SetConfigType("yaml")
		viper.SetConfigName(".puppet-environment")
	}

	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err == nil {
		fmt.Println("Using config file:", viper.ConfigFileUsed())
	}
}

func initEnvironmentsFile() {
	environmentsFile = environment.EnvironmentsFile{
		Path: viper.GetString("environments-file"),
	}

	_, err := os.Stat(environmentsFile.Path)

	if errors.Is(err, os.ErrNotExist) {
		environmentsFile.Environments = environment.Environments{}
		return
	}

	if err = environmentsFile.Read(); err == nil {
		return
	} else {
		log.Fatal("Unable to read environments file")
	}
}
