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
	datasourcesRagUpdateChunkingStrategyData        string
	datasourcesRagUpdateChunkingStrategyInteractive bool
)

var datasourcesRagUpdateChunkingStrategyCmd = &cobra.Command{
	Use:   "update-chunking-strategy <datasource-id>",
	Short: "Update Chunking Strategy",
	Long: `Update Chunking Strategy

Arguments:
  datasource-id: required`,
	Args:    cobra.ExactArgs(1),
	Example: "meibel datasources rag update-chunking-strategy <datasource-id>",
	RunE: func(cmd *cobra.Command, args []string) error {
		ctx := context.Background()

		datasourceId := args[0]

		var body sdk.UpdateChunkingStrategyRequest

		if datasourcesRagUpdateChunkingStrategyData != "" {
			if err := json.Unmarshal([]byte(datasourcesRagUpdateChunkingStrategyData), &body); err != nil {
				return fmt.Errorf("invalid JSON data: %w", err)
			}
		} else if datasourcesRagUpdateChunkingStrategyInteractive || term.IsTerminal(int(os.Stdin.Fd())) {
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

		result, err := client.DatasourcesRag.UpdateChunkingStrategy(ctx, datasourceId, body)
		if err != nil {
			return err
		}

		return output.Print(result)
	},
}

func init() {
	datasourcesRagCmd.AddCommand(datasourcesRagUpdateChunkingStrategyCmd)

	datasourcesRagUpdateChunkingStrategyCmd.Flags().StringVarP(&datasourcesRagUpdateChunkingStrategyData, "data", "d", "", "JSON data for the request body")
	datasourcesRagUpdateChunkingStrategyCmd.Flags().BoolVarP(&datasourcesRagUpdateChunkingStrategyInteractive, "interactive", "i", false, "use interactive form input")
}
