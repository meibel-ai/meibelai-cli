package cmd

import (
	"context"
	"fmt"

	"github.com/spf13/cobra"
	sdk "github.com/meibel-ai/meibel-go"
)

var (
	blueprintsInstancesUpdateActivityStatusUpdatedStatusValue string
)

var blueprintsInstancesUpdateActivityStatusCmd = &cobra.Command{
	Use:   "update-activity-status <blueprint-instance-id> <activity-id>",
	Short: "Update Activity Status",
	Long:  `Update Activity Status

Arguments:
  blueprint-instance-id: required
  activity-id: required`,
	Args:  cobra.ExactArgs(2),
	Example: "meibel blueprints instances update-activity-status <blueprint-instance-id> <activity-id>",
	RunE: func(cmd *cobra.Command, args []string) error {
		ctx := context.Background()

		blueprintInstanceId := args[0]
		activityId := args[1]

		err := client.BlueprintsInstances.UpdateActivityStatus(ctx, blueprintInstanceId, activityId, sdk.ActivityStatus(blueprintsInstancesUpdateActivityStatusUpdatedStatusValue))
		if err != nil {
			return err
		}

		fmt.Println("Success")
		return nil
	},
}

func init() {
	blueprintsInstancesCmd.AddCommand(blueprintsInstancesUpdateActivityStatusCmd)

	blueprintsInstancesUpdateActivityStatusCmd.Flags().StringVarP(&blueprintsInstancesUpdateActivityStatusUpdatedStatusValue, "updated-status-value", "", "", "The updated-status-value parameter")
	blueprintsInstancesUpdateActivityStatusCmd.MarkFlagRequired("updated-status-value")
}
