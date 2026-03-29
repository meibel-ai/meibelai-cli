package cmd

import (
	"context"
	"fmt"

	"github.com/meibel-ai/meibel-cli/internal/output"
	"github.com/spf13/cobra"
)

var (
	datasourcesTagDeleteTagTableInfoForce bool
)

var datasourcesTagDeleteTagTableInfoCmd = &cobra.Command{
	Use:   "delete-table-info <datasource-id> <table-name>",
	Short: "Delete Tag Table Info",
	Long: `Delete Tag Table Info

Arguments:
  datasource-id: required
  table-name: required`,
	Args:    cobra.ExactArgs(2),
	Example: "meibel datasources tag delete-table-info <datasource-id> <table-name>",
	RunE: func(cmd *cobra.Command, args []string) error {
		ctx := context.Background()

		datasourceId := args[0]
		tableName := args[1]

		if !datasourcesTagDeleteTagTableInfoForce {
			fmt.Print("Are you sure? [y/N] ")
			var confirm string
			fmt.Scanln(&confirm)
			if confirm != "y" && confirm != "Y" {
				fmt.Println("Cancelled")
				return nil
			}
		}

		result, err := client.DatasourcesTag.DeleteTagTableInfo(ctx, datasourceId, tableName)
		if err != nil {
			return err
		}

		return output.Print(result)
	},
}

func init() {
	datasourcesTagCmd.AddCommand(datasourcesTagDeleteTagTableInfoCmd)

	datasourcesTagDeleteTagTableInfoCmd.Flags().BoolVarP(&datasourcesTagDeleteTagTableInfoForce, "force", "f", false, "skip confirmation prompt")
}
