package cmd

import (
	"context"

	"github.com/meibel-ai/meibel-cli/internal/output"
	"github.com/spf13/cobra"
)

var blueprintsInstancesGetActivitiesByBlueprintInstanceCmd = &cobra.Command{
	Use:   "get-activities-by <blueprint-instance-id>",
	Short: "Get Activities By Blueprint Instance",
	Long: `Get Activities By Blueprint Instance

Arguments:
  blueprint-instance-id: required`,
	Args:    cobra.ExactArgs(1),
	Example: "meibel blueprints instances get-activities-by <blueprint-instance-id>",
	RunE: func(cmd *cobra.Command, args []string) error {
		ctx := context.Background()

		blueprintInstanceId := args[0]

		result, err := client.BlueprintsInstances.GetActivitiesByBlueprintInstance(ctx, blueprintInstanceId)
		if err != nil {
			return err
		}

		return output.Print(result)
	},
}

func init() {
	blueprintsInstancesCmd.AddCommand(blueprintsInstancesGetActivitiesByBlueprintInstanceCmd)

}
