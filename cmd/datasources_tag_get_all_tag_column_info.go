package cmd

import (
	"context"

	"github.com/spf13/cobra"
	"github.com/meibel-ai/meibel-cli/internal/output"
)

var datasourcesTagGetAllTagColumnInfoCmd = &cobra.Command{
	Use:   "get-all-column-info <datasource-id> <table-name>",
	Short: "Get All Tag Column Info",
	Long:  `Get All Tag Column Info

Arguments:
  datasource-id: required
  table-name: required`,
	Args:  cobra.ExactArgs(2),
	Example: "meibel datasources tag get-all-column-info <datasource-id> <table-name>",
	RunE: func(cmd *cobra.Command, args []string) error {
		ctx := context.Background()

		datasourceId := args[0]
		tableName := args[1]

		result, err := client.DatasourcesTag.GetAllTagColumnInfo(ctx, datasourceId, tableName)
		if err != nil {
			return err
		}

		return output.Print(result)
	},
}

func init() {
	datasourcesTagCmd.AddCommand(datasourcesTagGetAllTagColumnInfoCmd)

}
