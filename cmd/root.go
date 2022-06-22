package cmd

import (
	"errors"
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/reidmv/puppet-environment/internal/environment"
	"github.com/reidmv/puppet-environment/internal/r10k"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

func init() {
	cobra.OnInitialize(initConfig)

	rootCmd.CompletionOptions.HiddenDefaultCmd = true
	rootCmd.PersistentFlags().StringVar(&configFlag, "config", "", "path to .puppet-environment config")

	rootCmd.PersistentFlags().StringVar(&environmentsFileFlag, "environments-config", "", "path to environments.yaml config")
	viper.BindPFlag("environments-config", rootCmd.PersistentFlags().Lookup("environments-config"))
	viper.SetDefault("environments-config", "/etc/puppetlabs/puppet/environments.yaml")

	rootCmd.PersistentFlags().StringVar(&environmentsFileFlag, "environments-root", "", "directory to deploy environments to")
	viper.BindPFlag("environments-root", rootCmd.PersistentFlags().Lookup("environments-root"))
	viper.SetDefault("environments-root", "/etc/puppetlabs/code/environments")

	rootCmd.PersistentFlags().StringVar(&environmentsFileFlag, "r10k-config", "", "path to r10k.yaml config")
	viper.BindPFlag("r10k-config", rootCmd.PersistentFlags().Lookup("r10k-config"))
	viper.SetDefault("r10k-config", "/etc/puppetlabs/r10k/r10k.yaml")
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

	if path, err := filepath.Abs(viper.GetString("r10k-config")); err != nil {
		log.Fatal(err)
	} else {
		r10k.SetConfigPath(path)
	}
}

func initEnvironmentsFile() {
	abs, err := filepath.Abs(viper.GetString("environments-config"))
	if err != nil {
		log.Fatal(err)
	}
	environmentsFile = environment.EnvironmentsFile{
		Path: abs,
	}

	_, err = os.Stat(environmentsFile.Path)

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
