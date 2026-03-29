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
	datasourcesTagAddTagTableInfoData        string
	datasourcesTagAddTagTableInfoInteractive bool
)

var datasourcesTagAddTagTableInfoCmd = &cobra.Command{
	Use:   "add-table-info <datasource-id> <table-name>",
	Short: "Add Tag Table Info",
	Long: `Add Tag Table Info

Arguments:
  datasource-id: required
  table-name: required`,
	Args:    cobra.ExactArgs(2),
	Example: "meibel datasources tag add-table-info <datasource-id> <table-name>",
	RunE: func(cmd *cobra.Command, args []string) error {
		ctx := context.Background()

		datasourceId := args[0]
		tableName := args[1]

		var body sdk.AddTagTableRequest

		if datasourcesTagAddTagTableInfoData != "" {
			if err := json.Unmarshal([]byte(datasourcesTagAddTagTableInfoData), &body); err != nil {
				return fmt.Errorf("invalid JSON data: %w", err)
			}
		} else if datasourcesTagAddTagTableInfoInteractive || term.IsTerminal(int(os.Stdin.Fd())) {
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

		result, err := client.DatasourcesTag.AddTagTableInfo(ctx, datasourceId, tableName, body)
		if err != nil {
			return err
		}

		return output.Print(result)
	},
}

func init() {
	datasourcesTagCmd.AddCommand(datasourcesTagAddTagTableInfoCmd)

	datasourcesTagAddTagTableInfoCmd.Flags().StringVarP(&datasourcesTagAddTagTableInfoData, "data", "d", "", "JSON data for the request body")
	datasourcesTagAddTagTableInfoCmd.Flags().BoolVarP(&datasourcesTagAddTagTableInfoInteractive, "interactive", "i", false, "use interactive form input")
}
