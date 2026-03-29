package cmd

import (
	"github.com/spf13/cobra"
)

var dataElementMetadataCmd = &cobra.Command{
	Use:   "data-element-metadata",
	Short: "Manage Data Element Metadata",
	Long:  `Commands for managing Data Element Metadata resources.`,
}

func init() {
	rootCmd.AddCommand(dataElementMetadataCmd)
}
