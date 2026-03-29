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
	datasourcesRagReprocessDatasourceMetadataData        string
	datasourcesRagReprocessDatasourceMetadataInteractive bool
)

var datasourcesRagReprocessDatasourceMetadataCmd = &cobra.Command{
	Use:   "reprocess-metadata <datasource-id>",
	Short: "Reprocess Datasource Metadata",
	Long: `Reprocess Datasource Metadata

Arguments:
  datasource-id: required`,
	Args:    cobra.ExactArgs(1),
	Example: "meibel datasources rag reprocess-metadata <datasource-id>",
	RunE: func(cmd *cobra.Command, args []string) error {
		ctx := context.Background()

		datasourceId := args[0]

		var body sdk.ReprocessDatasourceRequest

		if datasourcesRagReprocessDatasourceMetadataData != "" {
			if err := json.Unmarshal([]byte(datasourcesRagReprocessDatasourceMetadataData), &body); err != nil {
				return fmt.Errorf("invalid JSON data: %w", err)
			}
		} else if datasourcesRagReprocessDatasourceMetadataInteractive || term.IsTerminal(int(os.Stdin.Fd())) {
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

		result, err := client.DatasourcesRag.ReprocessDatasourceMetadata(ctx, datasourceId, body)
		if err != nil {
			return err
		}

		return output.Print(result)
	},
}

func init() {
	datasourcesRagCmd.AddCommand(datasourcesRagReprocessDatasourceMetadataCmd)

	datasourcesRagReprocessDatasourceMetadataCmd.Flags().StringVarP(&datasourcesRagReprocessDatasourceMetadataData, "data", "d", "", "JSON data for the request body")
	datasourcesRagReprocessDatasourceMetadataCmd.Flags().BoolVarP(&datasourcesRagReprocessDatasourceMetadataInteractive, "interactive", "i", false, "use interactive form input")
}
