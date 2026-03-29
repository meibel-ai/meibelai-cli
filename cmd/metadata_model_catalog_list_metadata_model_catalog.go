package cmd

import (
	"context"

	"github.com/meibel-ai/meibel-cli/internal/output"
	sdk "github.com/meibel-ai/meibel-go"
	"github.com/spf13/cobra"
)

var (
	metadataModelCatalogListMetadataModelCatalogScope string
)

var metadataModelCatalogListMetadataModelCatalogCmd = &cobra.Command{
	Use:     "list",
	Short:   "List Metadata Model Catalog",
	Long:    `List Metadata Model Catalog`,
	Example: "meibel metadata-model-catalog list --scope=<value>",
	RunE: func(cmd *cobra.Command, args []string) error {
		ctx := context.Background()

		opts := &sdk.ListMetadataModelCatalogOptions{}
		if metadataModelCatalogListMetadataModelCatalogScope != "" {
			opts.Scope = &metadataModelCatalogListMetadataModelCatalogScope
		}

		iter := client.MetadataModelCatalog.ListMetadataModelCatalog(ctx, opts)

		var items []interface{}
		for iter.Next(ctx) {
			items = append(items, iter.Item())
		}
		if err := iter.Err(); err != nil {
			return err
		}

		return output.Print(items)
	},
}

func init() {
	metadataModelCatalogCmd.AddCommand(metadataModelCatalogListMetadataModelCatalogCmd)

	metadataModelCatalogListMetadataModelCatalogCmd.Flags().StringVarP(&metadataModelCatalogListMetadataModelCatalogScope, "scope", "", "", "The scope parameter")
}
