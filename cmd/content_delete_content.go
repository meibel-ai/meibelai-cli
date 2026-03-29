package cmd

import (
	"context"
	"fmt"

	"github.com/spf13/cobra"
	"github.com/meibel-ai/meibel-cli/internal/output"
)

var (
	contentDeleteContentForce bool
)

var contentDeleteContentCmd = &cobra.Command{
	Use:   "delete <datasource-id> <path>",
	Short: "Delete Content",
	Long:  `Delete Content

Arguments:
  datasource-id: required
  path: required`,
	Args:  cobra.ExactArgs(2),
	Example: "meibel content delete <datasource-id> <path>",
	RunE: func(cmd *cobra.Command, args []string) error {
		ctx := context.Background()

		datasourceId := args[0]
		path := args[1]

		if !contentDeleteContentForce {
			fmt.Print("Are you sure? [y/N] ")
			var confirm string
			fmt.Scanln(&confirm)
			if confirm != "y" && confirm != "Y" {
				fmt.Println("Cancelled")
				return nil
			}
		}

		result, err := client.Content.DeleteContent(ctx, datasourceId, path)
		if err != nil {
			return err
		}

		return output.Print(result)
	},
}

func init() {
	contentCmd.AddCommand(contentDeleteContentCmd)

	contentDeleteContentCmd.Flags().BoolVarP(&contentDeleteContentForce, "force", "f", false, "skip confirmation prompt")
}
