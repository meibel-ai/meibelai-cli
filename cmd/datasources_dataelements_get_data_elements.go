package cmd

import (
	"context"

	"github.com/meibel-ai/meibel-cli/internal/output"
	"github.com/spf13/cobra"
)

var datasourcesDataelementsGetDataElementsCmd = &cobra.Command{
	Use:   "get-data-elements <datasource-id>",
	Short: "Get Data Elements",
	Long: `Get Data Elements

Arguments:
  datasource-id: required`,
	Args:    cobra.ExactArgs(1),
	Example: "meibel datasources dataelements get-data-elements <datasource-id>",
	RunE: func(cmd *cobra.Command, args []string) error {
		ctx := context.Background()

		datasourceId := args[0]

		result, err := client.DatasourcesDataelements.GetDataElements(ctx, datasourceId)
		if err != nil {
			return err
		}

		return output.Print(result)
	},
}

func init() {
	datasourcesDataelementsCmd.AddCommand(datasourcesDataelementsGetDataElementsCmd)

}
