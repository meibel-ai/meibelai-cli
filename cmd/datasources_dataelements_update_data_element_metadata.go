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
	datasourcesDataelementsUpdateDataElementMetadataData string
	datasourcesDataelementsUpdateDataElementMetadataInteractive bool
)

var datasourcesDataelementsUpdateDataElementMetadataCmd = &cobra.Command{
	Use:   "update-data-element-metadata <datasource-id> <data-element-id>",
	Short: "Update Data Element Metadata",
	Long:  `Update Data Element Metadata

Arguments:
  datasource-id: required
  data-element-id: required`,
	Args:  cobra.ExactArgs(2),
	Example: "meibel datasources dataelements update-data-element-metadata <datasource-id> <data-element-id>",
	RunE: func(cmd *cobra.Command, args []string) error {
		ctx := context.Background()

		datasourceId := args[0]
		dataElementId := args[1]

		var body sdk.PutDataElementMetadataRequest

		if datasourcesDataelementsUpdateDataElementMetadataData != "" {
			if err := json.Unmarshal([]byte(datasourcesDataelementsUpdateDataElementMetadataData), &body); err != nil {
				return fmt.Errorf("invalid JSON data: %w", err)
			}
		} else if datasourcesDataelementsUpdateDataElementMetadataInteractive || term.IsTerminal(int(os.Stdin.Fd())) {
			// Interactive form
			form := huh.NewForm(
				huh.NewGroup(
					huh.NewInput().Title("Metadata").Description("").Value(&body.Metadata),
				),
			)

			if err := form.Run(); err != nil {
				return err
			}
		} else {
			return fmt.Errorf("--data flag required in non-interactive mode")
		}

		result, err := client.DatasourcesDataelements.UpdateDataElementMetadata(ctx, datasourceId, dataElementId, body)
		if err != nil {
			return err
		}

		return output.Print(result)
	},
}

func init() {
	datasourcesDataelementsCmd.AddCommand(datasourcesDataelementsUpdateDataElementMetadataCmd)

	datasourcesDataelementsUpdateDataElementMetadataCmd.Flags().StringVarP(&datasourcesDataelementsUpdateDataElementMetadataData, "data", "d", "", "JSON data for the request body")
	datasourcesDataelementsUpdateDataElementMetadataCmd.Flags().BoolVarP(&datasourcesDataelementsUpdateDataElementMetadataInteractive, "interactive", "i", false, "use interactive form input")
}
