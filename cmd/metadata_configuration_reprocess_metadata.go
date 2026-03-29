package cmd

import (
	"context"

	"github.com/meibel-ai/meibel-cli/internal/output"
	"github.com/spf13/cobra"
)

var metadataConfigurationReprocessMetadataCmd = &cobra.Command{
	Use:   "reprocess <datasource-id>",
	Short: "Reprocess Metadata",
	Long: `Reprocess Metadata

Arguments:
  datasource-id: required`,
	Args:    cobra.ExactArgs(1),
	Example: "meibel metadata-configuration reprocess <datasource-id>",
	RunE: func(cmd *cobra.Command, args []string) error {
		ctx := context.Background()

		datasourceId := args[0]

		result, err := client.MetadataConfiguration.ReprocessMetadata(ctx, datasourceId)
		if err != nil {
			return err
		}

		return output.Print(result)
	},
}

func init() {
	metadataConfigurationCmd.AddCommand(metadataConfigurationReprocessMetadataCmd)

}
