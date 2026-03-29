package cmd

import (
	"context"

	"github.com/meibel-ai/meibel-cli/internal/output"
	"github.com/spf13/cobra"
)

var tagDescriptionsListTagTablesCmd = &cobra.Command{
	Use:   "list-tables <datasource-id>",
	Short: "List Tag Tables",
	Long: `List Tag Tables

Arguments:
  datasource-id: required`,
	Args:    cobra.ExactArgs(1),
	Example: "meibel tag-descriptions list-tables <datasource-id>",
	RunE: func(cmd *cobra.Command, args []string) error {
		ctx := context.Background()

		datasourceId := args[0]

		iter := client.TagDescriptions.ListTagTables(ctx, datasourceId)

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
	tagDescriptionsCmd.AddCommand(tagDescriptionsListTagTablesCmd)

}
