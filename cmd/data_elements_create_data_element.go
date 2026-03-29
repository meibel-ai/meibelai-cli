package cmd

import (
	"context"
	"encoding/json"
	"fmt"
	"os"

	"github.com/charmbracelet/huh"
	"github.com/meibel-ai/meibel-cli/internal/output"
	sdk "github.com/meibel-ai/meibel-go"
	"github.com/spf13/cobra"
	"golang.org/x/term"
)

var (
	dataElementsCreateDataElementData        string
	dataElementsCreateDataElementInteractive bool
)

var dataElementsCreateDataElementCmd = &cobra.Command{
	Use:   "create <datasource-id>",
	Short: "Create Data Element",
	Long: `Create Data Element

Arguments:
  datasource-id: required`,
	Args:    cobra.ExactArgs(1),
	Example: "meibel data-elements create <datasource-id>",
	RunE: func(cmd *cobra.Command, args []string) error {
		ctx := context.Background()

		datasourceId := args[0]

		var body sdk.CreateDataElementRequest

		if dataElementsCreateDataElementData != "" {
			if err := json.Unmarshal([]byte(dataElementsCreateDataElementData), &body); err != nil {
				return fmt.Errorf("invalid JSON data: %w", err)
			}
		} else if dataElementsCreateDataElementInteractive || term.IsTerminal(int(os.Stdin.Fd())) {
			// Interactive form
			form := huh.NewForm(
				huh.NewGroup(
					huh.NewInput().Title("Name").Description("Data element name").Value(&body.Name),
				),
			)

			if err := form.Run(); err != nil {
				return err
			}
		} else {
			return fmt.Errorf("--data flag required in non-interactive mode")
		}

		result, err := client.DataElements.CreateDataElement(ctx, datasourceId, body)
		if err != nil {
			return err
		}

		return output.Print(result)
	},
}

func init() {
	dataElementsCmd.AddCommand(dataElementsCreateDataElementCmd)

	dataElementsCreateDataElementCmd.Flags().StringVarP(&dataElementsCreateDataElementData, "data", "d", "", "JSON data for the request body")
	dataElementsCreateDataElementCmd.Flags().BoolVarP(&dataElementsCreateDataElementInteractive, "interactive", "i", false, "use interactive form input")
}
