package cmd

import (
	"github.com/spf13/cobra"
)

var tagDescriptionsCmd = &cobra.Command{
	Use:   "tag-descriptions",
	Short: "Manage Tag Descriptions",
	Long:  `Commands for managing Tag Descriptions resources.`,
}

func init() {
	rootCmd.AddCommand(tagDescriptionsCmd)
}
