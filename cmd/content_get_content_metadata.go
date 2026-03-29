package cmd

import (
	"context"

	"github.com/spf13/cobra"
	"github.com/meibel-ai/meibel-cli/internal/output"
)

var contentGetContentMetadataCmd = &cobra.Command{
	Use:   "get-metadata <datasource-id> <path>",
	Short: "Get Content Metadata",
	Long:  `Get Content Metadata

Arguments:
  datasource-id: required
  path: required`,
	Args:  cobra.ExactArgs(2),
	Example: "meibel content get-metadata <datasource-id> <path>",
	RunE: func(cmd *cobra.Command, args []string) error {
		ctx := context.Background()

		datasourceId := args[0]
		path := args[1]

		result, err := client.Content.GetContentMetadata(ctx, datasourceId, path)
		if err != nil {
			return err
		}

		return output.Print(result)
	},
}

func init() {
	contentCmd.AddCommand(contentGetContentMetadataCmd)

}
