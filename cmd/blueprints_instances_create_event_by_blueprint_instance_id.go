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
	blueprintsInstancesCreateEventByBlueprintInstanceIdData string
	blueprintsInstancesCreateEventByBlueprintInstanceIdInteractive bool
)

var blueprintsInstancesCreateEventByBlueprintInstanceIdCmd = &cobra.Command{
	Use:   "create-event-by-id <blueprint-instance-id>",
	Short: "Create Event By Blueprint Instance Id",
	Long:  `Create Event By Blueprint Instance Id

Arguments:
  blueprint-instance-id: required`,
	Args:  cobra.ExactArgs(1),
	Example: "meibel blueprints instances create-event-by-id <blueprint-instance-id>",
	RunE: func(cmd *cobra.Command, args []string) error {
		ctx := context.Background()

		blueprintInstanceId := args[0]

		var body sdk.CustomEventRequest

		if blueprintsInstancesCreateEventByBlueprintInstanceIdData != "" {
			if err := json.Unmarshal([]byte(blueprintsInstancesCreateEventByBlueprintInstanceIdData), &body); err != nil {
				return fmt.Errorf("invalid JSON data: %w", err)
			}
		} else if blueprintsInstancesCreateEventByBlueprintInstanceIdInteractive || term.IsTerminal(int(os.Stdin.Fd())) {
			// Interactive form
			form := huh.NewForm(
				huh.NewGroup(
					huh.NewInput().Title("EventName").Description("Name of the custom event being logged.").Value(&body.EventName),
				),
			)

			if err := form.Run(); err != nil {
				return err
			}
		} else {
			return fmt.Errorf("--data flag required in non-interactive mode")
		}

		result, err := client.BlueprintsInstances.CreateEventByBlueprintInstanceId(ctx, blueprintInstanceId, body)
		if err != nil {
			return err
		}

		return output.Print(result)
	},
}

func init() {
	blueprintsInstancesCmd.AddCommand(blueprintsInstancesCreateEventByBlueprintInstanceIdCmd)

	blueprintsInstancesCreateEventByBlueprintInstanceIdCmd.Flags().StringVarP(&blueprintsInstancesCreateEventByBlueprintInstanceIdData, "data", "d", "", "JSON data for the request body")
	blueprintsInstancesCreateEventByBlueprintInstanceIdCmd.Flags().BoolVarP(&blueprintsInstancesCreateEventByBlueprintInstanceIdInteractive, "interactive", "i", false, "use interactive form input")
}
