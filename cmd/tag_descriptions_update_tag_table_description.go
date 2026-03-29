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
	tagDescriptionsUpdateTagTableDescriptionData string
	tagDescriptionsUpdateTagTableDescriptionInteractive bool
)

var tagDescriptionsUpdateTagTableDescriptionCmd = &cobra.Command{
	Use:   "update-table <datasource-id> <table-name>",
	Short: "Update Tag Table Description",
	Long:  `Update Tag Table Description

Arguments:
  datasource-id: required
  table-name: required`,
	Args:  cobra.ExactArgs(2),
	Example: "meibel tag-descriptions update-table <datasource-id> <table-name>",
	RunE: func(cmd *cobra.Command, args []string) error {
		ctx := context.Background()

		datasourceId := args[0]
		tableName := args[1]

		var body sdk.UpdateTagDescriptionRequest

		if tagDescriptionsUpdateTagTableDescriptionData != "" {
			if err := json.Unmarshal([]byte(tagDescriptionsUpdateTagTableDescriptionData), &body); err != nil {
				return fmt.Errorf("invalid JSON data: %w", err)
			}
		} else if tagDescriptionsUpdateTagTableDescriptionInteractive || term.IsTerminal(int(os.Stdin.Fd())) {
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

		result, err := client.TagDescriptions.UpdateTagTableDescription(ctx, datasourceId, tableName, body)
		if err != nil {
			return err
		}

		return output.Print(result)
	},
}

func init() {
	tagDescriptionsCmd.AddCommand(tagDescriptionsUpdateTagTableDescriptionCmd)

	tagDescriptionsUpdateTagTableDescriptionCmd.Flags().StringVarP(&tagDescriptionsUpdateTagTableDescriptionData, "data", "d", "", "JSON data for the request body")
	tagDescriptionsUpdateTagTableDescriptionCmd.Flags().BoolVarP(&tagDescriptionsUpdateTagTableDescriptionInteractive, "interactive", "i", false, "use interactive form input")
}
