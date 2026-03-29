package cmd

import (
	"context"
	"fmt"

	"github.com/spf13/cobra"
	"github.com/meibel-ai/meibel-cli/internal/output"
)

var (
	datasourcesContentDeleteDatasourceContentForce bool
)

var datasourcesContentDeleteDatasourceContentCmd = &cobra.Command{
	Use:   "delete <datasource-id> <path>",
	Short: "Delete content",
	Long:  `Delete a file or directory from the datasource

Arguments:
  datasource-id: required
  path: required`,
	Args:  cobra.ExactArgs(2),
	Example: "meibel datasources content delete <datasource-id> <path>",
	RunE: func(cmd *cobra.Command, args []string) error {
		ctx := context.Background()

		datasourceId := args[0]
		path := args[1]

		if !datasourcesContentDeleteDatasourceContentForce {
			fmt.Print("Are you sure? [y/N] ")
			var confirm string
			fmt.Scanln(&confirm)
			if confirm != "y" && confirm != "Y" {
				fmt.Println("Cancelled")
				return nil
			}
		}

		result, err := client.DatasourcesContent.DeleteDatasourceContent(ctx, datasourceId, path)
		if err != nil {
			return err
		}

		return output.Print(result)
	},
}

func init() {
	datasourcesContentCmd.AddCommand(datasourcesContentDeleteDatasourceContentCmd)

	datasourcesContentDeleteDatasourceContentCmd.Flags().BoolVarP(&datasourcesContentDeleteDatasourceContentForce, "force", "f", false, "skip confirmation prompt")
}
