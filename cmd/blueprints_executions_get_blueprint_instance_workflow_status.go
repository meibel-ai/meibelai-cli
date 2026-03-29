package cmd

import (
	"context"

	"github.com/meibel-ai/meibel-cli/internal/output"
	"github.com/spf13/cobra"
)

var blueprintsExecutionsGetBlueprintInstanceWorkflowStatusCmd = &cobra.Command{
	Use:   "get-instance-workflow-status <blueprint-instance-id>",
	Short: "Get Blueprint Instance Workflow Status",
	Long: `Get Blueprint Instance Workflow Status

Arguments:
  blueprint-instance-id: Unique identifier for the workflow instance`,
	Args:    cobra.ExactArgs(1),
	Example: "meibel blueprints executions get-instance-workflow-status <blueprint-instance-id>",
	RunE: func(cmd *cobra.Command, args []string) error {
		ctx := context.Background()

		blueprintInstanceId := args[0]

		result, err := client.BlueprintsExecutions.GetBlueprintInstanceWorkflowStatus(ctx, blueprintInstanceId)
		if err != nil {
			return err
		}

		return output.Print(result)
	},
}

func init() {
	blueprintsExecutionsCmd.AddCommand(blueprintsExecutionsGetBlueprintInstanceWorkflowStatusCmd)

}
