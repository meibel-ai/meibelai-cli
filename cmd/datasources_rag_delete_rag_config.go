package cmd

import (
	"context"
	"fmt"

	"github.com/meibel-ai/meibel-cli/internal/output"
	"github.com/spf13/cobra"
)

var (
	datasourcesRagDeleteRagConfigForce bool
)

var datasourcesRagDeleteRagConfigCmd = &cobra.Command{
	Use:   "delete-config <datasource-id>",
	Short: "Delete Rag Config",
	Long: `Delete Rag Config

Arguments:
  datasource-id: required`,
	Args:    cobra.ExactArgs(1),
	Example: "meibel datasources rag delete-config <datasource-id>",
	RunE: func(cmd *cobra.Command, args []string) error {
		ctx := context.Background()

		datasourceId := args[0]

		if !datasourcesRagDeleteRagConfigForce {
			fmt.Print("Are you sure? [y/N] ")
			var confirm string
			fmt.Scanln(&confirm)
			if confirm != "y" && confirm != "Y" {
				fmt.Println("Cancelled")
				return nil
			}
		}

		result, err := client.DatasourcesRag.DeleteRagConfig(ctx, datasourceId)
		if err != nil {
			return err
		}

		return output.Print(result)
	},
}

func init() {
	datasourcesRagCmd.AddCommand(datasourcesRagDeleteRagConfigCmd)

	datasourcesRagDeleteRagConfigCmd.Flags().BoolVarP(&datasourcesRagDeleteRagConfigForce, "force", "f", false, "skip confirmation prompt")
}
