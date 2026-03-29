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
	datasourcesRagUpdateRagConfigData string
	datasourcesRagUpdateRagConfigInteractive bool
)

var datasourcesRagUpdateRagConfigCmd = &cobra.Command{
	Use:   "update-config <datasource-id>",
	Short: "Update Rag Config",
	Long:  `Update Rag Config

Arguments:
  datasource-id: required`,
	Args:  cobra.ExactArgs(1),
	Example: "meibel datasources rag update-config <datasource-id>",
	RunE: func(cmd *cobra.Command, args []string) error {
		ctx := context.Background()

		datasourceId := args[0]

		var body sdk.UpdateRagConfigRequest

		if datasourcesRagUpdateRagConfigData != "" {
			if err := json.Unmarshal([]byte(datasourcesRagUpdateRagConfigData), &body); err != nil {
				return fmt.Errorf("invalid JSON data: %w", err)
			}
		} else if datasourcesRagUpdateRagConfigInteractive || term.IsTerminal(int(os.Stdin.Fd())) {
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

		result, err := client.DatasourcesRag.UpdateRagConfig(ctx, datasourceId, body)
		if err != nil {
			return err
		}

		return output.Print(result)
	},
}

func init() {
	datasourcesRagCmd.AddCommand(datasourcesRagUpdateRagConfigCmd)

	datasourcesRagUpdateRagConfigCmd.Flags().StringVarP(&datasourcesRagUpdateRagConfigData, "data", "d", "", "JSON data for the request body")
	datasourcesRagUpdateRagConfigCmd.Flags().BoolVarP(&datasourcesRagUpdateRagConfigInteractive, "interactive", "i", false, "use interactive form input")
}
