package cmd

import (
	"context"

	"github.com/spf13/cobra"
	"github.com/meibel-ai/meibel-cli/internal/output"
)

var dataElementMetadataGetDataElementMetadataCmd = &cobra.Command{
	Use:   "get <datasource-id> <data-element-id>",
	Short: "Get Data Element Metadata",
	Long:  `Get Data Element Metadata

Arguments:
  datasource-id: required
  data-element-id: required`,
	Args:  cobra.ExactArgs(2),
	Example: "meibel data-element-metadata get <datasource-id> <data-element-id>",
	RunE: func(cmd *cobra.Command, args []string) error {
		ctx := context.Background()

		datasourceId := args[0]
		dataElementId := args[1]

		result, err := client.DataElementMetadata.GetDataElementMetadata(ctx, datasourceId, dataElementId)
		if err != nil {
			return err
		}

		return output.Print(result)
	},
}

func init() {
	dataElementMetadataCmd.AddCommand(dataElementMetadataGetDataElementMetadataCmd)

}
