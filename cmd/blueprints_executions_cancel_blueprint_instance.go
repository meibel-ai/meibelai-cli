package cmd

import (
	"context"

	"github.com/spf13/cobra"
	"github.com/meibel-ai/meibel-cli/internal/output"
)

var blueprintsExecutionsCancelBlueprintInstanceCmd = &cobra.Command{
	Use:   "cancel-instance <blueprint-instance-id>",
	Short: "Cancel Blueprint Instance",
	Long:  `Cancel Blueprint Instance

Arguments:
  blueprint-instance-id: Unique identifier for the workflow instance`,
	Args:  cobra.ExactArgs(1),
	Example: "meibel blueprints executions cancel-instance <blueprint-instance-id>",
	RunE: func(cmd *cobra.Command, args []string) error {
		ctx := context.Background()

		blueprintInstanceId := args[0]

		result, err := client.BlueprintsExecutions.CancelBlueprintInstance(ctx, blueprintInstanceId)
		if err != nil {
			return err
		}

		return output.Print(result)
	},
}

func init() {
	blueprintsExecutionsCmd.AddCommand(blueprintsExecutionsCancelBlueprintInstanceCmd)

}
