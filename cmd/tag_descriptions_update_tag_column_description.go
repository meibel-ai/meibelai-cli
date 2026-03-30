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
	tagDescriptionsUpdateTagColumnDescriptionData string
	tagDescriptionsUpdateTagColumnDescriptionInteractive bool
)

var tagDescriptionsUpdateTagColumnDescriptionCmd = &cobra.Command{
	Use:   "update-column <datasource-id> <table-name> <column-name>",
	Short: "Update Tag Column Description",
	Long:  `Update Tag Column Description

Arguments:
  datasource-id: required
  table-name: required
  column-name: required`,
	Args:  cobra.ExactArgs(3),
	Example: "meibel tag-descriptions update-column <datasource-id> <table-name> <column-name>",
	RunE: func(cmd *cobra.Command, args []string) error {
		ctx := context.Background()

		datasourceId := args[0]
		tableName := args[1]
		columnName := args[2]

		var body sdk.UpdateTagDescriptionRequest

		if tagDescriptionsUpdateTagColumnDescriptionData != "" {
			if err := json.Unmarshal([]byte(tagDescriptionsUpdateTagColumnDescriptionData), &body); err != nil {
				return fmt.Errorf("invalid JSON data: %w", err)
			}
		} else if tagDescriptionsUpdateTagColumnDescriptionInteractive || term.IsTerminal(int(os.Stdin.Fd())) {
			// Interactive form
			form := huh.NewForm(
				huh.NewGroup(
					huh.NewInput().Title("Description").Description("Description for AI context").Value(&body.Description),
				),
			)

			if err := form.Run(); err != nil {
				return err
			}
		} else {
			return fmt.Errorf("--data flag required in non-interactive mode")
		}

		result, err := client.TagDescriptions.UpdateTagColumnDescription(ctx, datasourceId, tableName, columnName, body)
		if err != nil {
			return err
		}

		return output.Print(result)
	},
}

func init() {
	tagDescriptionsCmd.AddCommand(tagDescriptionsUpdateTagColumnDescriptionCmd)

	tagDescriptionsUpdateTagColumnDescriptionCmd.Flags().StringVarP(&tagDescriptionsUpdateTagColumnDescriptionData, "data", "d", "", "JSON data for the request body")
	tagDescriptionsUpdateTagColumnDescriptionCmd.Flags().BoolVarP(&tagDescriptionsUpdateTagColumnDescriptionInteractive, "interactive", "i", false, "use interactive form input")
}
