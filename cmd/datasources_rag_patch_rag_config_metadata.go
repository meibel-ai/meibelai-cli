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
	datasourcesRagPatchRagConfigMetadataData string
	datasourcesRagPatchRagConfigMetadataInteractive bool
)

var datasourcesRagPatchRagConfigMetadataCmd = &cobra.Command{
	Use:   "patch-config-metadata <datasource-id>",
	Short: "Patch Rag Config Metadata",
	Long:  `Patch Rag Config Metadata

Arguments:
  datasource-id: required`,
	Args:  cobra.ExactArgs(1),
	Example: "meibel datasources rag patch-config-metadata <datasource-id>",
	RunE: func(cmd *cobra.Command, args []string) error {
		ctx := context.Background()

		datasourceId := args[0]

		var body sdk.PatchRagConfigMetadataRequest

		if datasourcesRagPatchRagConfigMetadataData != "" {
			if err := json.Unmarshal([]byte(datasourcesRagPatchRagConfigMetadataData), &body); err != nil {
				return fmt.Errorf("invalid JSON data: %w", err)
			}
		} else if datasourcesRagPatchRagConfigMetadataInteractive || term.IsTerminal(int(os.Stdin.Fd())) {
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

		result, err := client.DatasourcesRag.PatchRagConfigMetadata(ctx, datasourceId, body)
		if err != nil {
			return err
		}

		return output.Print(result)
	},
}

func init() {
	datasourcesRagCmd.AddCommand(datasourcesRagPatchRagConfigMetadataCmd)

	datasourcesRagPatchRagConfigMetadataCmd.Flags().StringVarP(&datasourcesRagPatchRagConfigMetadataData, "data", "d", "", "JSON data for the request body")
	datasourcesRagPatchRagConfigMetadataCmd.Flags().BoolVarP(&datasourcesRagPatchRagConfigMetadataInteractive, "interactive", "i", false, "use interactive form input")
}
