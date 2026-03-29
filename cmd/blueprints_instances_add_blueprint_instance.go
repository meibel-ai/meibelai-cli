package cmd

import (
	"context"
	"encoding/json"
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/charmbracelet/huh"
	"golang.org/x/term"
	"github.com/meibel-ai/meibel-cli/internal/output"
	sdk "github.com/meibel-ai/meibel-go"
)

var (
	blueprintsInstancesAddBlueprintInstanceData string
	blueprintsInstancesAddBlueprintInstanceInteractive bool
)

var blueprintsInstancesAddBlueprintInstanceCmd = &cobra.Command{
	Use:   "add",
	Short: "Add Blueprint Instance",
	Long:  `Add Blueprint Instance`,
	Example: "meibel blueprints instances add",
	RunE: func(cmd *cobra.Command, args []string) error {
		ctx := context.Background()

		var body sdk.AddBlueprintInstanceRequest

		if blueprintsInstancesAddBlueprintInstanceData != "" {
			if err := json.Unmarshal([]byte(blueprintsInstancesAddBlueprintInstanceData), &body); err != nil {
				return fmt.Errorf("invalid JSON data: %w", err)
			}
		} else if blueprintsInstancesAddBlueprintInstanceInteractive || term.IsTerminal(int(os.Stdin.Fd())) {
			// Interactive form
			form := huh.NewForm(
				huh.NewGroup(
				),
			)

			if err := form.Run(); err != nil {
				return err
			}
		} else {
			return fmt.Errorf("--data flag required in non-interactive mode")
		}

		result, err := client.BlueprintsInstances.AddBlueprintInstance(ctx, body)
		if err != nil {
			return err
		}

		return output.Print(result)
	},
}

func init() {
	blueprintsInstancesCmd.AddCommand(blueprintsInstancesAddBlueprintInstanceCmd)

	blueprintsInstancesAddBlueprintInstanceCmd.Flags().StringVarP(&blueprintsInstancesAddBlueprintInstanceData, "data", "d", "", "JSON data for the request body")
	blueprintsInstancesAddBlueprintInstanceCmd.Flags().BoolVarP(&blueprintsInstancesAddBlueprintInstanceInteractive, "interactive", "i", false, "use interactive form input")
}
