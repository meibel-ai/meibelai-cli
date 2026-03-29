package cmd

import (
	"context"

	"github.com/spf13/cobra"
	"github.com/meibel-ai/meibel-cli/internal/output"
)

var dataElementMetadataGetDataElementMetadataResultCmd = &cobra.Command{
	Use:   "get-result <datasource-id> <data-element-id> <request-id>",
	Short: "Get Data Element Metadata Result",
	Long:  `Get Data Element Metadata Result

Arguments:
  datasource-id: required
  data-element-id: required
  request-id: required`,
	Args:  cobra.ExactArgs(3),
	Example: "meibel data-element-metadata get-result <datasource-id> <data-element-id> <request-id>",
	RunE: func(cmd *cobra.Command, args []string) error {
		ctx := context.Background()

		datasourceId := args[0]
		dataElementId := args[1]
		requestId := args[2]

		result, err := client.DataElementMetadata.GetDataElementMetadataResult(ctx, datasourceId, dataElementId, requestId)
		if err != nil {
			return err
		}

		return output.Print(result)
	},
}

func init() {
	dataElementMetadataCmd.AddCommand(dataElementMetadataGetDataElementMetadataResultCmd)

}
