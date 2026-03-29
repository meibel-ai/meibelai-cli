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
	datasourcesRagAddChunkingStrategyData string
	datasourcesRagAddChunkingStrategyInteractive bool
)

var datasourcesRagAddChunkingStrategyCmd = &cobra.Command{
	Use:   "add-chunking-strategy <datasource-id>",
	Short: "Add Chunking Strategy",
	Long:  `Add Chunking Strategy

Arguments:
  datasource-id: required`,
	Args:  cobra.ExactArgs(1),
	Example: "meibel datasources rag add-chunking-strategy <datasource-id>",
	RunE: func(cmd *cobra.Command, args []string) error {
		ctx := context.Background()

		datasourceId := args[0]

		var body sdk.AddChunkingStrategyRequest

		if datasourcesRagAddChunkingStrategyData != "" {
			if err := json.Unmarshal([]byte(datasourcesRagAddChunkingStrategyData), &body); err != nil {
				return fmt.Errorf("invalid JSON data: %w", err)
			}
		} else if datasourcesRagAddChunkingStrategyInteractive || term.IsTerminal(int(os.Stdin.Fd())) {
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

		result, err := client.DatasourcesRag.AddChunkingStrategy(ctx, datasourceId, body)
		if err != nil {
			return err
		}

		return output.Print(result)
	},
}

func init() {
	datasourcesRagCmd.AddCommand(datasourcesRagAddChunkingStrategyCmd)

	datasourcesRagAddChunkingStrategyCmd.Flags().StringVarP(&datasourcesRagAddChunkingStrategyData, "data", "d", "", "JSON data for the request body")
	datasourcesRagAddChunkingStrategyCmd.Flags().BoolVarP(&datasourcesRagAddChunkingStrategyInteractive, "interactive", "i", false, "use interactive form input")
}
