package cmd

import (
	"context"
	"encoding/json"
	"fmt"
	"os"

	"github.com/charmbracelet/huh"
	"github.com/meibel-ai/meibel-cli/internal/output"
	sdk "github.com/meibel-ai/meibel-go"
	"github.com/spf13/cobra"
	"golang.org/x/term"
)

var (
	datasourcesDataelementsUpdateDataElementData        string
	datasourcesDataelementsUpdateDataElementInteractive bool
)

var datasourcesDataelementsUpdateDataElementCmd = &cobra.Command{
	Use:   "update-data-element <datasource-id> <data-element-id>",
	Short: "Update Data Element",
	Long: `Update Data Element

Arguments:
  datasource-id: required
  data-element-id: required`,
	Args:    cobra.ExactArgs(2),
	Example: "meibel datasources dataelements update-data-element <datasource-id> <data-element-id>",
	RunE: func(cmd *cobra.Command, args []string) error {
		ctx := context.Background()

		datasourceId := args[0]
		dataElementId := args[1]

		var body sdk.DatasourceServiceClientModelsUpdateDataElementRequestUpdateDataElementRequest

		if datasourcesDataelementsUpdateDataElementData != "" {
			if err := json.Unmarshal([]byte(datasourcesDataelementsUpdateDataElementData), &body); err != nil {
				return fmt.Errorf("invalid JSON data: %w", err)
			}
		} else if datasourcesDataelementsUpdateDataElementInteractive || term.IsTerminal(int(os.Stdin.Fd())) {
			// Interactive form
			form := huh.NewForm(
				huh.NewGroup(),
			)

			if err := form.Run(); err != nil {
				return err
			}
		} else {
			return fmt.Errorf("--data flag required in non-interactive mode")
		}

		result, err := client.DatasourcesDataelements.UpdateDataElement(ctx, datasourceId, dataElementId, body)
		if err != nil {
			return err
		}

		return output.Print(result)
	},
}

func init() {
	datasourcesDataelementsCmd.AddCommand(datasourcesDataelementsUpdateDataElementCmd)

	datasourcesDataelementsUpdateDataElementCmd.Flags().StringVarP(&datasourcesDataelementsUpdateDataElementData, "data", "d", "", "JSON data for the request body")
	datasourcesDataelementsUpdateDataElementCmd.Flags().BoolVarP(&datasourcesDataelementsUpdateDataElementInteractive, "interactive", "i", false, "use interactive form input")
}
