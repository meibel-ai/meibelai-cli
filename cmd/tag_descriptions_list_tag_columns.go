package cmd

import (
	"context"

	"github.com/spf13/cobra"
	"github.com/meibel-ai/meibel-cli/internal/output"
)

var tagDescriptionsListTagColumnsCmd = &cobra.Command{
	Use:   "list-columns <datasource-id> <table-name>",
	Short: "List Tag Columns",
	Long:  `List Tag Columns

Arguments:
  datasource-id: required
  table-name: required`,
	Args:  cobra.ExactArgs(2),
	Example: "meibel tag-descriptions list-columns <datasource-id> <table-name>",
	RunE: func(cmd *cobra.Command, args []string) error {
		ctx := context.Background()

		datasourceId := args[0]
		tableName := args[1]

		iter := client.TagDescriptions.ListTagColumns(ctx, datasourceId, tableName)

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
	tagDescriptionsCmd.AddCommand(tagDescriptionsListTagColumnsCmd)

}
