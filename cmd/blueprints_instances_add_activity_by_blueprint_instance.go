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
	blueprintsInstancesAddActivityByBlueprintInstanceData string
	blueprintsInstancesAddActivityByBlueprintInstanceInteractive bool
)

var blueprintsInstancesAddActivityByBlueprintInstanceCmd = &cobra.Command{
	Use:   "add-activity-by <blueprint-instance-id>",
	Short: "Add Activity By Blueprint Instance",
	Long:  `Add Activity By Blueprint Instance

Arguments:
  blueprint-instance-id: required`,
	Args:  cobra.ExactArgs(1),
	Example: "meibel blueprints instances add-activity-by <blueprint-instance-id>",
	RunE: func(cmd *cobra.Command, args []string) error {
		ctx := context.Background()

		blueprintInstanceId := args[0]

		var body sdk.AddActivityRequest

		if blueprintsInstancesAddActivityByBlueprintInstanceData != "" {
			if err := json.Unmarshal([]byte(blueprintsInstancesAddActivityByBlueprintInstanceData), &body); err != nil {
				return fmt.Errorf("invalid JSON data: %w", err)
			}
		} else if blueprintsInstancesAddActivityByBlueprintInstanceInteractive || term.IsTerminal(int(os.Stdin.Fd())) {
			// Interactive form
			form := huh.NewForm(
				huh.NewGroup(
					huh.NewInput().Title("ActivityType").Description("").Value(&body.ActivityType),
				),
			)

			if err := form.Run(); err != nil {
				return err
			}
		} else {
			return fmt.Errorf("--data flag required in non-interactive mode")
		}

		result, err := client.BlueprintsInstances.AddActivityByBlueprintInstance(ctx, blueprintInstanceId, body)
		if err != nil {
			return err
		}

		return output.Print(result)
	},
}

func init() {
	blueprintsInstancesCmd.AddCommand(blueprintsInstancesAddActivityByBlueprintInstanceCmd)

	blueprintsInstancesAddActivityByBlueprintInstanceCmd.Flags().StringVarP(&blueprintsInstancesAddActivityByBlueprintInstanceData, "data", "d", "", "JSON data for the request body")
	blueprintsInstancesAddActivityByBlueprintInstanceCmd.Flags().BoolVarP(&blueprintsInstancesAddActivityByBlueprintInstanceInteractive, "interactive", "i", false, "use interactive form input")
}
