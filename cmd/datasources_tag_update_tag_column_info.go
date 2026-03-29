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
	datasourcesTagUpdateTagColumnInfoData string
	datasourcesTagUpdateTagColumnInfoInteractive bool
)

var datasourcesTagUpdateTagColumnInfoCmd = &cobra.Command{
	Use:   "update-column-info <datasource-id> <table-name> <column-name>",
	Short: "Update Tag Column Info",
	Long:  `Update Tag Column Info

Arguments:
  datasource-id: required
  table-name: required
  column-name: required`,
	Args:  cobra.ExactArgs(3),
	Example: "meibel datasources tag update-column-info <datasource-id> <table-name> <column-name>",
	RunE: func(cmd *cobra.Command, args []string) error {
		ctx := context.Background()

		datasourceId := args[0]
		tableName := args[1]
		columnName := args[2]

		var body sdk.UpdateTagColumnRequest

		if datasourcesTagUpdateTagColumnInfoData != "" {
			if err := json.Unmarshal([]byte(datasourcesTagUpdateTagColumnInfoData), &body); err != nil {
				return fmt.Errorf("invalid JSON data: %w", err)
			}
		} else if datasourcesTagUpdateTagColumnInfoInteractive || term.IsTerminal(int(os.Stdin.Fd())) {
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

		result, err := client.DatasourcesTag.UpdateTagColumnInfo(ctx, datasourceId, tableName, columnName, body)
		if err != nil {
			return err
		}

		return output.Print(result)
	},
}

func init() {
	datasourcesTagCmd.AddCommand(datasourcesTagUpdateTagColumnInfoCmd)

	datasourcesTagUpdateTagColumnInfoCmd.Flags().StringVarP(&datasourcesTagUpdateTagColumnInfoData, "data", "d", "", "JSON data for the request body")
	datasourcesTagUpdateTagColumnInfoCmd.Flags().BoolVarP(&datasourcesTagUpdateTagColumnInfoInteractive, "interactive", "i", false, "use interactive form input")
}
