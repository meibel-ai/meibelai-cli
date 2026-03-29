package cmd

import (
	"context"
	"fmt"

	"github.com/spf13/cobra"
	"github.com/meibel-ai/meibel-cli/internal/output"
)

var (
	datasourcesTagDeleteTagConfigForce bool
)

var datasourcesTagDeleteTagConfigCmd = &cobra.Command{
	Use:   "delete-config <datasource-id>",
	Short: "Delete Tag Config",
	Long:  `Delete Tag Config

Arguments:
  datasource-id: required`,
	Args:  cobra.ExactArgs(1),
	Example: "meibel datasources tag delete-config <datasource-id>",
	RunE: func(cmd *cobra.Command, args []string) error {
		ctx := context.Background()

		datasourceId := args[0]

		if !datasourcesTagDeleteTagConfigForce {
			fmt.Print("Are you sure? [y/N] ")
			var confirm string
			fmt.Scanln(&confirm)
			if confirm != "y" && confirm != "Y" {
				fmt.Println("Cancelled")
				return nil
			}
		}

		result, err := client.DatasourcesTag.DeleteTagConfig(ctx, datasourceId)
		if err != nil {
			return err
		}

		return output.Print(result)
	},
}

func init() {
	datasourcesTagCmd.AddCommand(datasourcesTagDeleteTagConfigCmd)

	datasourcesTagDeleteTagConfigCmd.Flags().BoolVarP(&datasourcesTagDeleteTagConfigForce, "force", "f", false, "skip confirmation prompt")
}
