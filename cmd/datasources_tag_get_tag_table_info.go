package cmd

import (
	"context"

	"github.com/spf13/cobra"
	"github.com/meibel-ai/meibel-cli/internal/output"
)

var datasourcesTagGetTagTableInfoCmd = &cobra.Command{
	Use:   "get-table-info <datasource-id> <table-name>",
	Short: "Get Tag Table Info",
	Long:  `Get Tag Table Info

Arguments:
  datasource-id: required
  table-name: required`,
	Args:  cobra.ExactArgs(2),
	Example: "meibel datasources tag get-table-info <datasource-id> <table-name>",
	RunE: func(cmd *cobra.Command, args []string) error {
		ctx := context.Background()

		datasourceId := args[0]
		tableName := args[1]

		result, err := client.DatasourcesTag.GetTagTableInfo(ctx, datasourceId, tableName)
		if err != nil {
			return err
		}

		return output.Print(result)
	},
}

func init() {
	datasourcesTagCmd.AddCommand(datasourcesTagGetTagTableInfoCmd)

}
