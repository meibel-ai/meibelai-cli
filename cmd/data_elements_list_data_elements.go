package cmd

import (
	"context"

	"github.com/spf13/cobra"
	"github.com/meibel-ai/meibel-cli/internal/output"
)

var dataElementsListDataElementsCmd = &cobra.Command{
	Use:   "list <datasource-id>",
	Short: "List Data Elements",
	Long:  `List Data Elements

Arguments:
  datasource-id: required`,
	Args:  cobra.ExactArgs(1),
	Example: "meibel data-elements list <datasource-id>",
	RunE: func(cmd *cobra.Command, args []string) error {
		ctx := context.Background()

		datasourceId := args[0]

		iter := client.DataElements.ListDataElements(ctx, datasourceId)

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
	dataElementsCmd.AddCommand(dataElementsListDataElementsCmd)

}
