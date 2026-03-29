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
	datasourcesDataelementsAddDataElementData        string
	datasourcesDataelementsAddDataElementInteractive bool
)

var datasourcesDataelementsAddDataElementCmd = &cobra.Command{
	Use:   "add-data-element <datasource-id>",
	Short: "Add Data Element",
	Long: `Add Data Element

Arguments:
  datasource-id: required`,
	Args:    cobra.ExactArgs(1),
	Example: "meibel datasources dataelements add-data-element <datasource-id>",
	RunE: func(cmd *cobra.Command, args []string) error {
		ctx := context.Background()

		datasourceId := args[0]

		var body sdk.AddDataElementRequest

		if datasourcesDataelementsAddDataElementData != "" {
			if err := json.Unmarshal([]byte(datasourcesDataelementsAddDataElementData), &body); err != nil {
				return fmt.Errorf("invalid JSON data: %w", err)
			}
		} else if datasourcesDataelementsAddDataElementInteractive || term.IsTerminal(int(os.Stdin.Fd())) {
			// Interactive form
			form := huh.NewForm(
				huh.NewGroup(
					huh.NewInput().Title("Name").Description("").Value(&body.Name),
					huh.NewInput().Title("Path").Description("").Value(&body.Path),
					huh.NewInput().Title("MediaType").Description("").Value(&body.MediaType),
				),
			)

			if err := form.Run(); err != nil {
				return err
			}
		} else {
			return fmt.Errorf("--data flag required in non-interactive mode")
		}

		result, err := client.DatasourcesDataelements.AddDataElement(ctx, datasourceId, body)
		if err != nil {
			return err
		}

		return output.Print(result)
	},
}

func init() {
	datasourcesDataelementsCmd.AddCommand(datasourcesDataelementsAddDataElementCmd)

	datasourcesDataelementsAddDataElementCmd.Flags().StringVarP(&datasourcesDataelementsAddDataElementData, "data", "d", "", "JSON data for the request body")
	datasourcesDataelementsAddDataElementCmd.Flags().BoolVarP(&datasourcesDataelementsAddDataElementInteractive, "interactive", "i", false, "use interactive form input")
}
