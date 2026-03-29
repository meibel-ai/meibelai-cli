package cmd

import (
	"context"

	"github.com/spf13/cobra"
	"github.com/meibel-ai/meibel-cli/internal/output"
	sdk "github.com/meibel-ai/meibel-go"
)

var (
	blueprintsInstancesGetAllBlueprintInstancesIncludeChildren bool
	blueprintsInstancesGetAllBlueprintInstancesIncludeActivities bool
	blueprintsInstancesGetAllBlueprintInstancesIncludeEvents bool
)

var blueprintsInstancesGetAllBlueprintInstancesCmd = &cobra.Command{
	Use:   "get-all",
	Short: "Get All Blueprint Instances",
	Long:  `Get All Blueprint Instances`,
	Example: "meibel blueprints instances get-all --include-children=<value> --include-activities=<value>",
	RunE: func(cmd *cobra.Command, args []string) error {
		ctx := context.Background()

		opts := &sdk.GetAllBlueprintInstancesOptions{}
		if blueprintsInstancesGetAllBlueprintInstancesIncludeChildren {
			opts.IncludeChildren = &blueprintsInstancesGetAllBlueprintInstancesIncludeChildren
		}
		if blueprintsInstancesGetAllBlueprintInstancesIncludeActivities {
			opts.IncludeActivities = &blueprintsInstancesGetAllBlueprintInstancesIncludeActivities
		}
		if blueprintsInstancesGetAllBlueprintInstancesIncludeEvents {
			opts.IncludeEvents = &blueprintsInstancesGetAllBlueprintInstancesIncludeEvents
		}

		result, err := client.BlueprintsInstances.GetAllBlueprintInstances(ctx, opts)
		if err != nil {
			return err
		}

		return output.Print(result)
	},
}

func init() {
	blueprintsInstancesCmd.AddCommand(blueprintsInstancesGetAllBlueprintInstancesCmd)

	blueprintsInstancesGetAllBlueprintInstancesCmd.Flags().BoolVarP(&blueprintsInstancesGetAllBlueprintInstancesIncludeChildren, "include-children", "", false, "The include-children parameter")
	blueprintsInstancesGetAllBlueprintInstancesCmd.Flags().BoolVarP(&blueprintsInstancesGetAllBlueprintInstancesIncludeActivities, "include-activities", "", false, "The include-activities parameter")
	blueprintsInstancesGetAllBlueprintInstancesCmd.Flags().BoolVarP(&blueprintsInstancesGetAllBlueprintInstancesIncludeEvents, "include-events", "", false, "The include-events parameter")
}
