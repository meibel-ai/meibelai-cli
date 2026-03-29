package cmd

import (
	"context"
	"encoding/json"
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/charmbracelet/huh"
	"golang.org/x/term"
	"github.com/meibel-ai/meibel-cli/internal/output"
	sdk "github.com/meibel-ai/meibel-go"
)

var (
	metadataConfigurationUpdateMetadataConfigData string
	metadataConfigurationUpdateMetadataConfigInteractive bool
)

var metadataConfigurationUpdateMetadataConfigCmd = &cobra.Command{
	Use:   "update-config <datasource-id>",
	Short: "Update Metadata Config",
	Long:  `Update Metadata Config

Arguments:
  datasource-id: required`,
	Args:  cobra.ExactArgs(1),
	Example: "meibel metadata-configuration update-config <datasource-id>",
	RunE: func(cmd *cobra.Command, args []string) error {
		ctx := context.Background()

		datasourceId := args[0]

		var body sdk.MetadataConfigRequest

		if metadataConfigurationUpdateMetadataConfigData != "" {
			if err := json.Unmarshal([]byte(metadataConfigurationUpdateMetadataConfigData), &body); err != nil {
				return fmt.Errorf("invalid JSON data: %w", err)
			}
		} else if metadataConfigurationUpdateMetadataConfigInteractive || term.IsTerminal(int(os.Stdin.Fd())) {
			// Interactive form
			form := huh.NewForm(
				huh.NewGroup(
				),
			)

			if err := form.Run(); err != nil {
				return err
			}
		} else {
			return fmt.Errorf("--data flag required in non-interactive mode")
		}

		result, err := client.MetadataConfiguration.UpdateMetadataConfig(ctx, datasourceId, body)
		if err != nil {
			return err
		}

		return output.Print(result)
	},
}

func init() {
	metadataConfigurationCmd.AddCommand(metadataConfigurationUpdateMetadataConfigCmd)

	metadataConfigurationUpdateMetadataConfigCmd.Flags().StringVarP(&metadataConfigurationUpdateMetadataConfigData, "data", "d", "", "JSON data for the request body")
	metadataConfigurationUpdateMetadataConfigCmd.Flags().BoolVarP(&metadataConfigurationUpdateMetadataConfigInteractive, "interactive", "i", false, "use interactive form input")
}
