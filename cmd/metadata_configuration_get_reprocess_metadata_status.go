package cmd

import (
	"context"

	"github.com/spf13/cobra"
	"github.com/meibel-ai/meibel-cli/internal/output"
)

var metadataConfigurationGetReprocessMetadataStatusCmd = &cobra.Command{
	Use:   "get-reprocess-status <datasource-id>",
	Short: "Get Reprocess Metadata Status",
	Long:  `Get Reprocess Metadata Status

Arguments:
  datasource-id: required`,
	Args:  cobra.ExactArgs(1),
	Example: "meibel metadata-configuration get-reprocess-status <datasource-id>",
	RunE: func(cmd *cobra.Command, args []string) error {
		ctx := context.Background()

		datasourceId := args[0]

		result, err := client.MetadataConfiguration.GetReprocessMetadataStatus(ctx, datasourceId)
		if err != nil {
			return err
		}

		return output.Print(result)
	},
}

func init() {
	metadataConfigurationCmd.AddCommand(metadataConfigurationGetReprocessMetadataStatusCmd)

}
