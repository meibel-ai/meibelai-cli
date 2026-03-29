package cmd

import (
	"github.com/spf13/cobra"
)

var metadataConfigurationCmd = &cobra.Command{
	Use:   "metadata-configuration",
	Short: "Manage Metadata Configuration",
	Long:  `Commands for managing Metadata Configuration resources.`,
}

func init() {
	rootCmd.AddCommand(metadataConfigurationCmd)
}
