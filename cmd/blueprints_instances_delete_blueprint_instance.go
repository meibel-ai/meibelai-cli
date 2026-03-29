package cmd

import (
	"context"
	"fmt"

	"github.com/spf13/cobra"
)

var (
	blueprintsInstancesDeleteBlueprintInstanceForce bool
)

var blueprintsInstancesDeleteBlueprintInstanceCmd = &cobra.Command{
	Use:   "delete <blueprint-instance-id>",
	Short: "Delete Blueprint Instance",
	Long:  `Delete Blueprint Instance

Arguments:
  blueprint-instance-id: required`,
	Args:  cobra.ExactArgs(1),
	Example: "meibel blueprints instances delete <blueprint-instance-id>",
	RunE: func(cmd *cobra.Command, args []string) error {
		ctx := context.Background()

		blueprintInstanceId := args[0]

		if !blueprintsInstancesDeleteBlueprintInstanceForce {
			fmt.Print("Are you sure? [y/N] ")
			var confirm string
			fmt.Scanln(&confirm)
			if confirm != "y" && confirm != "Y" {
				fmt.Println("Cancelled")
				return nil
			}
		}

		err := client.BlueprintsInstances.DeleteBlueprintInstance(ctx, blueprintInstanceId)
		if err != nil {
			return err
		}

		fmt.Println("Success")
		return nil
	},
}

func init() {
	blueprintsInstancesCmd.AddCommand(blueprintsInstancesDeleteBlueprintInstanceCmd)

	blueprintsInstancesDeleteBlueprintInstanceCmd.Flags().BoolVarP(&blueprintsInstancesDeleteBlueprintInstanceForce, "force", "f", false, "skip confirmation prompt")
}
