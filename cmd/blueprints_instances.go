package cmd

import (
	"github.com/spf13/cobra"
)

var blueprintsInstancesCmd = &cobra.Command{
	Use:   "instances",
	Short: "Manage instances",
	Long:  `Commands for managing instances resources.`,
}

func init() {
	blueprintsCmd.AddCommand(blueprintsInstancesCmd)
}
