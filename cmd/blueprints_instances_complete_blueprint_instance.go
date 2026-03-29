package cmd

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/spf13/cobra"
	"github.com/meibel-ai/meibel-cli/internal/output"
)

var (
	blueprintsInstancesCompleteBlueprintInstanceData string
	blueprintsInstancesCompleteBlueprintInstanceInteractive bool
)

var blueprintsInstancesCompleteBlueprintInstanceCmd = &cobra.Command{
	Use:   "complete <blueprint-instance-id>",
	Short: "Complete a Blueprint Instance",
	Long:  `This endpoint is used to mark a Blueprint Instance as completed. It will update the status of the Blueprint Instance to 'COMPLETED' and log the completion event.

Arguments:
  blueprint-instance-id: Unique identifier for the workflow instance`,
	Args:  cobra.ExactArgs(1),
	Example: "meibel blueprints instances complete <blueprint-instance-id>",
	RunE: func(cmd *cobra.Command, args []string) error {
		ctx := context.Background()

		blueprintInstanceId := args[0]

		var body interface{}

		if blueprintsInstancesCompleteBlueprintInstanceData != "" {
			if err := json.Unmarshal([]byte(blueprintsInstancesCompleteBlueprintInstanceData), &body); err != nil {
				return fmt.Errorf("invalid JSON data: %w", err)
			}
		} else {
			return fmt.Errorf("--data flag required (interactive form not available for this type)")
		}

		result, err := client.BlueprintsInstances.CompleteBlueprintInstance(ctx, blueprintInstanceId, &body)
		if err != nil {
			return err
		}

		return output.Print(result)
	},
}

func init() {
	blueprintsInstancesCmd.AddCommand(blueprintsInstancesCompleteBlueprintInstanceCmd)

	blueprintsInstancesCompleteBlueprintInstanceCmd.Flags().StringVarP(&blueprintsInstancesCompleteBlueprintInstanceData, "data", "d", "", "JSON data for the request body")
	blueprintsInstancesCompleteBlueprintInstanceCmd.Flags().BoolVarP(&blueprintsInstancesCompleteBlueprintInstanceInteractive, "interactive", "i", false, "use interactive form input")
}
