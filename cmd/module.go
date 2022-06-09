package cmd

import (
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(moduleCmd)
}

var (
	ModuleTypeFlag    string
	ModuleSourceFlag  string
	ModuleVersionFlag string
)

var moduleCmd = &cobra.Command{
	Use:   "module [subcommand] [options]",
	Short: "Manage modules deployed to Puppet code environments",
	Long:  `Manage modules deployed to Puppet code environments defined in environments.yaml`,
}
