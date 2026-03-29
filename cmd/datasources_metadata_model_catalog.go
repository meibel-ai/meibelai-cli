package cmd

import (
	"github.com/spf13/cobra"
)

var datasourcesMetadataModelCatalogCmd = &cobra.Command{
	Use:   "metadata-model-catalog",
	Short: "Manage metadata_model_catalog",
	Long:  `Commands for managing metadata_model_catalog resources.`,
}

func init() {
	datasourcesCmd.AddCommand(datasourcesMetadataModelCatalogCmd)
}
