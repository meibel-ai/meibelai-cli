package cmd

import (
	"github.com/spf13/cobra"
)

var datasourcesTagCmd = &cobra.Command{
	Use:   "tag",
	Short: "Manage tag",
	Long:  `Commands for managing tag resources.`,
}

func init() {
	datasourcesCmd.AddCommand(datasourcesTagCmd)
}
