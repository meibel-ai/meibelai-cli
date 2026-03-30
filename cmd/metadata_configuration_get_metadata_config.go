package cmd

import (
	"context"

	"github.com/spf13/cobra"
	"github.com/meibel-ai/meibel-cli/internal/output"
)

var metadataConfigurationGetMetadataConfigCmd = &cobra.Command{
	Use:   "get-config <datasource-id>",
	Short: "Get Metadata Config",
	Long:  `Get Metadata Config

Arguments:
  datasource-id: required`,
	Args:  cobra.ExactArgs(1),
	Example: "meibel metadata-configuration get-config <datasource-id>",
	RunE: func(cmd *cobra.Command, args []string) error {
		ctx := context.Background()

		datasourceId := args[0]

		result, err := client.MetadataConfiguration.GetMetadataConfig(ctx, datasourceId)
		if err != nil {
			return err
		}

		return output.Print(result)
	},
}

func init() {
	metadataConfigurationCmd.AddCommand(metadataConfigurationGetMetadataConfigCmd)

}
