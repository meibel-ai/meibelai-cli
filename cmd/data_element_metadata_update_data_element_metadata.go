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
	dataElementMetadataUpdateDataElementMetadataData string
	dataElementMetadataUpdateDataElementMetadataInteractive bool
)

var dataElementMetadataUpdateDataElementMetadataCmd = &cobra.Command{
	Use:   "update <datasource-id> <data-element-id>",
	Short: "Update Data Element Metadata",
	Long:  `Update Data Element Metadata

Arguments:
  datasource-id: required
  data-element-id: required`,
	Args:  cobra.ExactArgs(2),
	Example: "meibel data-element-metadata update <datasource-id> <data-element-id>",
	RunE: func(cmd *cobra.Command, args []string) error {
		ctx := context.Background()

		datasourceId := args[0]
		dataElementId := args[1]

		var body sdk.DataElementMetadata

		if dataElementMetadataUpdateDataElementMetadataData != "" {
			if err := json.Unmarshal([]byte(dataElementMetadataUpdateDataElementMetadataData), &body); err != nil {
				return fmt.Errorf("invalid JSON data: %w", err)
			}
		} else if dataElementMetadataUpdateDataElementMetadataInteractive || term.IsTerminal(int(os.Stdin.Fd())) {
			// Interactive form
			form := huh.NewForm(
				huh.NewGroup(
					huh.NewInput().Title("Metadata").Description("Arbitrary key-value metadata").Value(&body.Metadata),
				),
			)

			if err := form.Run(); err != nil {
				return err
			}
		} else {
			return fmt.Errorf("--data flag required in non-interactive mode")
		}

		result, err := client.DataElementMetadata.UpdateDataElementMetadata(ctx, datasourceId, dataElementId, body)
		if err != nil {
			return err
		}

		return output.Print(result)
	},
}

func init() {
	dataElementMetadataCmd.AddCommand(dataElementMetadataUpdateDataElementMetadataCmd)

	dataElementMetadataUpdateDataElementMetadataCmd.Flags().StringVarP(&dataElementMetadataUpdateDataElementMetadataData, "data", "d", "", "JSON data for the request body")
	dataElementMetadataUpdateDataElementMetadataCmd.Flags().BoolVarP(&dataElementMetadataUpdateDataElementMetadataInteractive, "interactive", "i", false, "use interactive form input")
}
