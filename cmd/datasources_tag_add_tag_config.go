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
	datasourcesTagAddTagConfigData        string
	datasourcesTagAddTagConfigInteractive bool
)

var datasourcesTagAddTagConfigCmd = &cobra.Command{
	Use:   "add-config <datasource-id>",
	Short: "Add Tag Config",
	Long: `Add Tag Config

Arguments:
  datasource-id: required`,
	Args:    cobra.ExactArgs(1),
	Example: "meibel datasources tag add-config <datasource-id>",
	RunE: func(cmd *cobra.Command, args []string) error {
		ctx := context.Background()

		datasourceId := args[0]

		var body sdk.AddTagConfigRequest

		if datasourcesTagAddTagConfigData != "" {
			if err := json.Unmarshal([]byte(datasourcesTagAddTagConfigData), &body); err != nil {
				return fmt.Errorf("invalid JSON data: %w", err)
			}
		} else if datasourcesTagAddTagConfigInteractive || term.IsTerminal(int(os.Stdin.Fd())) {
			// Interactive form
			form := huh.NewForm(
				huh.NewGroup(
					huh.NewInput().Title("WorkingBucket").Description("").Value(&body.WorkingBucket),
				),
			)

			if err := form.Run(); err != nil {
				return err
			}
		} else {
			return fmt.Errorf("--data flag required in non-interactive mode")
		}

		result, err := client.DatasourcesTag.AddTagConfig(ctx, datasourceId, body)
		if err != nil {
			return err
		}

		return output.Print(result)
	},
}

func init() {
	datasourcesTagCmd.AddCommand(datasourcesTagAddTagConfigCmd)

	datasourcesTagAddTagConfigCmd.Flags().StringVarP(&datasourcesTagAddTagConfigData, "data", "d", "", "JSON data for the request body")
	datasourcesTagAddTagConfigCmd.Flags().BoolVarP(&datasourcesTagAddTagConfigInteractive, "interactive", "i", false, "use interactive form input")
}
