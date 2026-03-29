package cmd

import (
	"context"
	"fmt"

	sdk "github.com/meibel-ai/meibel-go"
	"github.com/spf13/cobra"
)

var (
	blueprintsInstancesUpdateBlueprintInstanceStatusUpdatedStatusValue string
	blueprintsInstancesUpdateBlueprintInstanceStatusWorkflowRunId      string
)

var blueprintsInstancesUpdateBlueprintInstanceStatusCmd = &cobra.Command{
	Use:   "update-status <blueprint-instance-id>",
	Short: "Update Blueprint Instance Status",
	Long: `Update Blueprint Instance Status

Arguments:
  blueprint-instance-id: required`,
	Args:    cobra.ExactArgs(1),
	Example: "meibel blueprints instances update-status <blueprint-instance-id> --workflow-run-id=<value>",
	RunE: func(cmd *cobra.Command, args []string) error {
		ctx := context.Background()

		blueprintInstanceId := args[0]

		opts := &sdk.UpdateBlueprintInstanceStatusOptions{}
		if blueprintsInstancesUpdateBlueprintInstanceStatusWorkflowRunId != "" {
			opts.WorkflowRunId = &blueprintsInstancesUpdateBlueprintInstanceStatusWorkflowRunId
		}

		err := client.BlueprintsInstances.UpdateBlueprintInstanceStatus(ctx, blueprintInstanceId, sdk.BlueprintInstanceStatus(blueprintsInstancesUpdateBlueprintInstanceStatusUpdatedStatusValue), opts)
		if err != nil {
			return err
		}

		fmt.Println("Success")
		return nil
	},
}

func init() {
	blueprintsInstancesCmd.AddCommand(blueprintsInstancesUpdateBlueprintInstanceStatusCmd)

	blueprintsInstancesUpdateBlueprintInstanceStatusCmd.Flags().StringVarP(&blueprintsInstancesUpdateBlueprintInstanceStatusUpdatedStatusValue, "updated-status-value", "", "", "The updated-status-value parameter")
	blueprintsInstancesUpdateBlueprintInstanceStatusCmd.MarkFlagRequired("updated-status-value")
	blueprintsInstancesUpdateBlueprintInstanceStatusCmd.Flags().StringVarP(&blueprintsInstancesUpdateBlueprintInstanceStatusWorkflowRunId, "workflow-run-id", "", "", "The workflow-run-id parameter")
}
