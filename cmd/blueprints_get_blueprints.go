package cmd

import (
	"context"

	"github.com/meibel-ai/meibel-cli/internal/output"
	"github.com/spf13/cobra"
)

var blueprintsGetBlueprintsCmd = &cobra.Command{
	Use:     "list",
	Short:   "Get Blueprints",
	Long:    `Get Blueprints`,
	Example: "meibel blueprints list",
	RunE: func(cmd *cobra.Command, args []string) error {
		ctx := context.Background()

		result, err := client.Blueprints.GetBlueprints(ctx)
		if err != nil {
			return err
		}

		return output.Print(result)
	},
}

func init() {
	blueprintsCmd.AddCommand(blueprintsGetBlueprintsCmd)

}
