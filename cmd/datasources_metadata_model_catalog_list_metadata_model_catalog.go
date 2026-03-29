package cmd

import (
	"context"

	"github.com/meibel-ai/meibel-cli/internal/output"
	sdk "github.com/meibel-ai/meibel-go"
	"github.com/spf13/cobra"
)

var (
	datasourcesMetadataModelCatalogListMetadataModelCatalogScope string
)

var datasourcesMetadataModelCatalogListMetadataModelCatalogCmd = &cobra.Command{
	Use:     "list",
	Short:   "List Metadata Model Catalog",
	Long:    `List Metadata Model Catalog`,
	Example: "meibel datasources metadata-model-catalog list --scope=<value>",
	RunE: func(cmd *cobra.Command, args []string) error {
		ctx := context.Background()

		opts := &sdk.ListMetadataModelCatalogOptions{}
		if datasourcesMetadataModelCatalogListMetadataModelCatalogScope != "" {
			opts.Scope = &datasourcesMetadataModelCatalogListMetadataModelCatalogScope
		}

		iter := client.DatasourcesMetadataModelCatalog.ListMetadataModelCatalog(ctx, opts)

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
	datasourcesMetadataModelCatalogCmd.AddCommand(datasourcesMetadataModelCatalogListMetadataModelCatalogCmd)

	datasourcesMetadataModelCatalogListMetadataModelCatalogCmd.Flags().StringVarP(&datasourcesMetadataModelCatalogListMetadataModelCatalogScope, "scope", "", "", "The scope parameter")
}
