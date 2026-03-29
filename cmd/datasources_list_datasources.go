package cmd

import (
	"context"

	"github.com/meibel-ai/meibel-cli/internal/output"
	"github.com/spf13/cobra"
)

var datasourcesListDatasourcesCmd = &cobra.Command{
	Use:     "list",
	Short:   "List Datasources",
	Long:    `List Datasources`,
	Example: "meibel datasources list",
	RunE: func(cmd *cobra.Command, args []string) error {
		ctx := context.Background()

		iter := client.Datasources.ListDatasources(ctx)

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
	datasourcesCmd.AddCommand(datasourcesListDatasourcesCmd)

}
