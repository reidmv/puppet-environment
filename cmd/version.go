package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(versionCmd)
}

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the version number of puppet-environment",
	Long:  `All software has versions. This is puppet-environment's`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("puppet-environment v0.1 -- HEAD")
	},
	PersistentPreRun:  none,
	PersistentPostRun: none,
}
