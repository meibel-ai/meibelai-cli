package cmd

import (
	"context"

	"github.com/meibel-ai/meibel-cli/internal/output"
	"github.com/spf13/cobra"
)

var datasourcesTagGetAllTagTableInfoCmd = &cobra.Command{
	Use:   "get-all-table-info <datasource-id>",
	Short: "Get All Tag Table Info",
	Long: `Get All Tag Table Info

Arguments:
  datasource-id: required`,
	Args:    cobra.ExactArgs(1),
	Example: "meibel datasources tag get-all-table-info <datasource-id>",
	RunE: func(cmd *cobra.Command, args []string) error {
		ctx := context.Background()

		datasourceId := args[0]

		result, err := client.DatasourcesTag.GetAllTagTableInfo(ctx, datasourceId)
		if err != nil {
			return err
		}

		return output.Print(result)
	},
}

func init() {
	datasourcesTagCmd.AddCommand(datasourcesTagGetAllTagTableInfoCmd)

}
