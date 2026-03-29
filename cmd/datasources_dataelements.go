package cmd

import (
	"github.com/spf13/cobra"
)

var datasourcesDataelementsCmd = &cobra.Command{
	Use:   "dataelements",
	Short: "Manage dataelements",
	Long:  `Commands for managing dataelements resources.`,
}

func init() {
	datasourcesCmd.AddCommand(datasourcesDataelementsCmd)
}
