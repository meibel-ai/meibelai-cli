package cmd

import (
	"context"

	"github.com/meibel-ai/meibel-cli/internal/output"
	"github.com/spf13/cobra"
)

var blueprintsInstancesGetEventsByBlueprintInstanceCmd = &cobra.Command{
	Use:   "get-events-by <blueprint-instance-id>",
	Short: "Get Events By Blueprint Instance",
	Long: `Get Events By Blueprint Instance

Arguments:
  blueprint-instance-id: required`,
	Args:    cobra.ExactArgs(1),
	Example: "meibel blueprints instances get-events-by <blueprint-instance-id>",
	RunE: func(cmd *cobra.Command, args []string) error {
		ctx := context.Background()

		blueprintInstanceId := args[0]

		result, err := client.BlueprintsInstances.GetEventsByBlueprintInstance(ctx, blueprintInstanceId)
		if err != nil {
			return err
		}

		return output.Print(result)
	},
}

func init() {
	blueprintsInstancesCmd.AddCommand(blueprintsInstancesGetEventsByBlueprintInstanceCmd)

}
