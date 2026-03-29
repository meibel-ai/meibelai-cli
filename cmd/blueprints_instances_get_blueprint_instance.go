package cmd

import (
	"context"

	"github.com/spf13/cobra"
	"github.com/meibel-ai/meibel-cli/internal/output"
	sdk "github.com/meibel-ai/meibel-go"
)

var (
	blueprintsInstancesGetBlueprintInstanceIncludeChildren bool
	blueprintsInstancesGetBlueprintInstanceIncludeActivities bool
	blueprintsInstancesGetBlueprintInstanceIncludeEvents bool
)

var blueprintsInstancesGetBlueprintInstanceCmd = &cobra.Command{
	Use:   "get <blueprint-instance-id>",
	Short: "Get Blueprint Instance",
	Long:  `Get Blueprint Instance

Arguments:
  blueprint-instance-id: required`,
	Args:  cobra.ExactArgs(1),
	Example: "meibel blueprints instances get <blueprint-instance-id> --include-children=<value> --include-activities=<value>",
	RunE: func(cmd *cobra.Command, args []string) error {
		ctx := context.Background()

		blueprintInstanceId := args[0]

		opts := &sdk.GetBlueprintInstanceOptions{}
		if blueprintsInstancesGetBlueprintInstanceIncludeChildren {
			opts.IncludeChildren = &blueprintsInstancesGetBlueprintInstanceIncludeChildren
		}
		if blueprintsInstancesGetBlueprintInstanceIncludeActivities {
			opts.IncludeActivities = &blueprintsInstancesGetBlueprintInstanceIncludeActivities
		}
		if blueprintsInstancesGetBlueprintInstanceIncludeEvents {
			opts.IncludeEvents = &blueprintsInstancesGetBlueprintInstanceIncludeEvents
		}

		result, err := client.BlueprintsInstances.GetBlueprintInstance(ctx, blueprintInstanceId, opts)
		if err != nil {
			return err
		}

		return output.Print(result)
	},
}

func init() {
	blueprintsInstancesCmd.AddCommand(blueprintsInstancesGetBlueprintInstanceCmd)

	blueprintsInstancesGetBlueprintInstanceCmd.Flags().BoolVarP(&blueprintsInstancesGetBlueprintInstanceIncludeChildren, "include-children", "", false, "The include-children parameter")
	blueprintsInstancesGetBlueprintInstanceCmd.Flags().BoolVarP(&blueprintsInstancesGetBlueprintInstanceIncludeActivities, "include-activities", "", false, "The include-activities parameter")
	blueprintsInstancesGetBlueprintInstanceCmd.Flags().BoolVarP(&blueprintsInstancesGetBlueprintInstanceIncludeEvents, "include-events", "", false, "The include-events parameter")
}
