package cmd

import (
	"fmt"
	"log"

	"github.com/spf13/cobra"
)

func init() {
	moduleCmd.AddCommand(moduleListCmd)
	moduleListCmd.Flags().StringVarP(&environmentFlag, "environment", "e", "", "Puppet code environment")
	moduleListCmd.MarkFlagRequired("environment")
}

var moduleListCmd = &cobra.Command{
	Use:   "list",
	Short: "List modules for a Puppet code environment",
	Long:  `List modules for a Puppet code environment defined in environments.yaml`,
	Run: func(cmd *cobra.Command, args []string) {
		env, ok := environmentsFile.Environments[environmentFlag]
		if !ok {
			log.Fatal("Environment does not exist!")
		}
		for modname := range env.Modules {
			fmt.Println(modname)
		}
	},
}
