package cmd

import (
	"context"

	"github.com/meibel-ai/meibel-cli/internal/output"
	"github.com/spf13/cobra"
)

var datasourcesDataelementsGetDataElementCmd = &cobra.Command{
	Use:   "get-data-element <datasource-id> <data-element-id>",
	Short: "Get Data Element",
	Long: `Get Data Element

Arguments:
  datasource-id: required
  data-element-id: required`,
	Args:    cobra.ExactArgs(2),
	Example: "meibel datasources dataelements get-data-element <datasource-id> <data-element-id>",
	RunE: func(cmd *cobra.Command, args []string) error {
		ctx := context.Background()

		datasourceId := args[0]
		dataElementId := args[1]

		result, err := client.DatasourcesDataelements.GetDataElement(ctx, datasourceId, dataElementId)
		if err != nil {
			return err
		}

		return output.Print(result)
	},
}

func init() {
	datasourcesDataelementsCmd.AddCommand(datasourcesDataelementsGetDataElementCmd)

}
