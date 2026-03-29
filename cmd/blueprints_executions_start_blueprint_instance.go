package cmd

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/spf13/cobra"
	"github.com/meibel-ai/meibel-cli/internal/output"
)

var (
	blueprintsExecutionsStartBlueprintInstanceData string
	blueprintsExecutionsStartBlueprintInstanceInteractive bool
)

var blueprintsExecutionsStartBlueprintInstanceCmd = &cobra.Command{
	Use:   "start-instance <blueprint-instance-id>",
	Short: "Start Blueprint Instance",
	Long:  `Start Blueprint Instance

Arguments:
  blueprint-instance-id: Unique identifier for the workflow instance`,
	Args:  cobra.ExactArgs(1),
	Example: "meibel blueprints executions start-instance <blueprint-instance-id>",
	RunE: func(cmd *cobra.Command, args []string) error {
		ctx := context.Background()

		blueprintInstanceId := args[0]

		var body interface{}

		if blueprintsExecutionsStartBlueprintInstanceData != "" {
			if err := json.Unmarshal([]byte(blueprintsExecutionsStartBlueprintInstanceData), &body); err != nil {
				return fmt.Errorf("invalid JSON data: %w", err)
			}
		} else {
			return fmt.Errorf("--data flag required (interactive form not available for this type)")
		}

		result, err := client.BlueprintsExecutions.StartBlueprintInstance(ctx, blueprintInstanceId, &body)
		if err != nil {
			return err
		}

		return output.Print(result)
	},
}

func init() {
	blueprintsExecutionsCmd.AddCommand(blueprintsExecutionsStartBlueprintInstanceCmd)

	blueprintsExecutionsStartBlueprintInstanceCmd.Flags().StringVarP(&blueprintsExecutionsStartBlueprintInstanceData, "data", "d", "", "JSON data for the request body")
	blueprintsExecutionsStartBlueprintInstanceCmd.Flags().BoolVarP(&blueprintsExecutionsStartBlueprintInstanceInteractive, "interactive", "i", false, "use interactive form input")
}
