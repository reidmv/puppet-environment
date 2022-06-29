package cmd

import (
	"bytes"
	"errors"
	"fmt"
	"log"
	"os"
	"os/exec"
	"path/filepath"

	"github.com/reidmv/puppet-environment/internal/environment"
	"github.com/reidmv/puppet-environment/internal/filesync"
	"github.com/reidmv/puppet-environment/internal/r10k"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

func init() {
	cobra.OnInitialize(initConfig)

	rootCmd.CompletionOptions.HiddenDefaultCmd = true
	rootCmd.PersistentFlags().StringVar(&configFlag, "config", "", "path to .puppet-environment config")
	rootCmd.PersistentFlags().StringVar(&environmentsFileFlag, "environments-config", "", "path to environments.yaml config")
	rootCmd.PersistentFlags().StringVar(&environmentsFileFlag, "environments-root", "", "directory to deploy environments to")
	rootCmd.PersistentFlags().StringVar(&environmentsFileFlag, "r10k-config", "", "path to r10k.yaml config")
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
	PersistentPostRun: func(cmd *cobra.Command, args []string) {
		if codeManagerConfigured() {
			initFilesync()
			fmt.Println("Triggering file-sync commit...")
			if err := filesync.Commit(); err != nil {
				log.Fatal(err)
			}
		}
	},
}

// Used to disable PersistentPreRun and PersistentPostRun in other commands
func none(cmd *cobra.Command, args []string) {}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

// initConfig reads in the application config file and sets configuration default values.
// It should be executed before performing any other work (called by init()).
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

	// Bind config options to corresponding flags
	viper.BindPFlag("environments-config", rootCmd.PersistentFlags().Lookup("environments-config"))
	viper.SetDefault("environments-config", "/etc/puppetlabs/puppet/environments.yaml")

	viper.BindPFlag("environments-root", rootCmd.PersistentFlags().Lookup("environments-root"))
	viper.BindPFlag("r10k-config", rootCmd.PersistentFlags().Lookup("r10k-config"))
	if codeManagerConfigured() {
		// Default to using config locations assuming Code Manager is in use
		viper.SetDefault("r10k-config", "/opt/puppetlabs/server/data/code-manager/r10k.yaml")
		viper.SetDefault("environments-root", "/etc/puppetlabs/code-staging/environments")
	} else {
		// Default to using config locations assuming no Code Manager
		viper.SetDefault("r10k-config", "/etc/puppetlabs/r10k/r10k.yaml")
		viper.SetDefault("environments-root", "/etc/puppetlabs/code/environments")
	}

	// Determine the absolute path to r10k.yaml and set it in the r10k package.
	if path, err := filepath.Abs(viper.GetString("r10k-config")); err != nil {
		log.Fatal(err)
	} else {
		r10k.SetConfigPath(path)
	}
}

// initEnvironmentFile sets up the object used to read, write, and manipulate environment
// config data. It should be called before reading or writing to the environments.yaml file.
func initEnvironmentsFile() {
	abs, err := filepath.Abs(viper.GetString("environments-config"))
	if err != nil {
		log.Fatal(err)
	}
	environmentsFile = environment.EnvironmentsFile{
		Path: abs,
	}

	if _, err = os.Stat(environmentsFile.Path); errors.Is(err, os.ErrNotExist) {
		environmentsFile.Environments = environment.Environments{}
		return
	}

	if err = environmentsFile.Read(); err == nil {
		return
	} else {
		log.Fatal("Unable to read environments file")
	}
}

func initFilesync() {
	// Get the certname used by Puppet
	cmd := exec.Command("/opt/puppetlabs/bin/puppet", "config", "print", "certname")
	var out bytes.Buffer
	cmd.Stdout = &out
	err := cmd.Run()
	if err != nil {
		log.Fatalf("Error running `puppet config print certname`: %v", err)
	}
	certname := out.String()
	certname = certname[:len(certname)-1] // removing EOL

	cert := "/etc/puppetlabs/puppet/ssl/certs/" + certname + ".pem"
	privkey := "/etc/puppetlabs/puppet/ssl/private_keys/" + certname + ".pem"

	filesync.InitializeHttpClient(cert, privkey)
}

func codeManagerConfigured() bool {
	if _, err := os.Stat("/opt/puppetlabs/server/data/code-manager/r10k.yaml"); errors.Is(err, os.ErrNotExist) {
		// File does not exist, Code Manager is probably not in use
		return false
	} else {
		// File exists, assume Code Manager is in use
		return true
	}
}
