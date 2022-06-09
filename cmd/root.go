package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

func init() {
	cobra.OnInitialize(initConfig)

	rootCmd.CompletionOptions.HiddenDefaultCmd = true
	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", `config file (default "$HOME/.puppet-environment")`)
	rootCmd.PersistentFlags().StringVar(&envsFile, "environments-file", "/etc/puppetlabs/puppet/environments.yaml", "environments file")

	viper.BindPFlag("environments-file", rootCmd.Flags().Lookup("environments-file"))
}

var (
	EnvironmentFlag string
	TypeFlag        string
	SourceFlag      string
	VersionFlag     string
	cfgFile         string
	envsFile        string
)

var rootCmd = &cobra.Command{
	Use:   "puppet-environment",
	Short: "Manage Puppet environments and environment modules",
	Long: `The puppet-environment tool is used to manage and deploy Puppet code environments and modules.
		Environment definitions are stored in /etc/puppetlabs/puppet/environments.yaml.
		The r10k utility is used to instantiate environments defined there.`,
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func initConfig() {
	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	} else {
		// Find home directory.
		home, err := os.UserHomeDir()
		cobra.CheckErr(err)

		// Search config in home directory with name ".cobra" (without extension).
		viper.AddConfigPath(home)
		viper.SetConfigType("yaml")
		viper.SetConfigName(".puppet-environment")
	}

	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err == nil {
		fmt.Println("Using config file:", viper.ConfigFileUsed())
	}
}
