package cmd

import (
	"context"
	"fmt"

	"github.com/spf13/cobra"
	"github.com/meibel-ai/meibel-cli/internal/output"
)

var (
	dataElementsDeleteDataElementForce bool
)

var dataElementsDeleteDataElementCmd = &cobra.Command{
	Use:   "delete <datasource-id> <data-element-id>",
	Short: "Delete Data Element",
	Long:  `Delete Data Element

Arguments:
  datasource-id: required
  data-element-id: required`,
	Args:  cobra.ExactArgs(2),
	Example: "meibel data-elements delete <datasource-id> <data-element-id>",
	RunE: func(cmd *cobra.Command, args []string) error {
		ctx := context.Background()

		datasourceId := args[0]
		dataElementId := args[1]

		if !dataElementsDeleteDataElementForce {
			fmt.Print("Are you sure? [y/N] ")
			var confirm string
			fmt.Scanln(&confirm)
			if confirm != "y" && confirm != "Y" {
				fmt.Println("Cancelled")
				return nil
			}
		}

		result, err := client.DataElements.DeleteDataElement(ctx, datasourceId, dataElementId)
		if err != nil {
			return err
		}

		return output.Print(result)
	},
}

func init() {
	dataElementsCmd.AddCommand(dataElementsDeleteDataElementCmd)

	dataElementsDeleteDataElementCmd.Flags().BoolVarP(&dataElementsDeleteDataElementForce, "force", "f", false, "skip confirmation prompt")
}
