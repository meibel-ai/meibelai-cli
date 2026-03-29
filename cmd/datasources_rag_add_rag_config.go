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
	datasourcesRagAddRagConfigData        string
	datasourcesRagAddRagConfigInteractive bool
)

var datasourcesRagAddRagConfigCmd = &cobra.Command{
	Use:   "add-config <datasource-id>",
	Short: "Add Rag Config",
	Long: `Add Rag Config

Arguments:
  datasource-id: required`,
	Args:    cobra.ExactArgs(1),
	Example: "meibel datasources rag add-config <datasource-id>",
	RunE: func(cmd *cobra.Command, args []string) error {
		ctx := context.Background()

		datasourceId := args[0]

		var body sdk.AddRagConfigRequest

		if datasourcesRagAddRagConfigData != "" {
			if err := json.Unmarshal([]byte(datasourcesRagAddRagConfigData), &body); err != nil {
				return fmt.Errorf("invalid JSON data: %w", err)
			}
		} else if datasourcesRagAddRagConfigInteractive || term.IsTerminal(int(os.Stdin.Fd())) {
			// Interactive form
			form := huh.NewForm(
				huh.NewGroup(
					huh.NewInput().Title("CollectionId").Description("").Value(&body.CollectionId),
				),
			)

			if err := form.Run(); err != nil {
				return err
			}
		} else {
			return fmt.Errorf("--data flag required in non-interactive mode")
		}

		result, err := client.DatasourcesRag.AddRagConfig(ctx, datasourceId, body)
		if err != nil {
			return err
		}

		return output.Print(result)
	},
}

func init() {
	datasourcesRagCmd.AddCommand(datasourcesRagAddRagConfigCmd)

	datasourcesRagAddRagConfigCmd.Flags().StringVarP(&datasourcesRagAddRagConfigData, "data", "d", "", "JSON data for the request body")
	datasourcesRagAddRagConfigCmd.Flags().BoolVarP(&datasourcesRagAddRagConfigInteractive, "interactive", "i", false, "use interactive form input")
}
