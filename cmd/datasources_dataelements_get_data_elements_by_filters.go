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
	datasourcesDataelementsGetDataElementsByFiltersRegexFilter      string
	datasourcesDataelementsGetDataElementsByFiltersMediaTypeFilters []string
	datasourcesDataelementsGetDataElementsByFiltersData             string
	datasourcesDataelementsGetDataElementsByFiltersInteractive      bool
)

var datasourcesDataelementsGetDataElementsByFiltersCmd = &cobra.Command{
	Use:   "get-data-elements-by-filters <datasource-id>",
	Short: "Get Data Elements By Filters",
	Long: `Get Data Elements By Filters

Arguments:
  datasource-id: required`,
	Args:    cobra.ExactArgs(1),
	Example: "meibel datasources dataelements get-data-elements-by-filters <datasource-id> --regex-filter=<value> --media-type-filters=<value>",
	RunE: func(cmd *cobra.Command, args []string) error {
		ctx := context.Background()

		datasourceId := args[0]

		var body sdk.DataElementFilterRequest

		if datasourcesDataelementsGetDataElementsByFiltersData != "" {
			if err := json.Unmarshal([]byte(datasourcesDataelementsGetDataElementsByFiltersData), &body); err != nil {
				return fmt.Errorf("invalid JSON data: %w", err)
			}
		} else if datasourcesDataelementsGetDataElementsByFiltersInteractive || term.IsTerminal(int(os.Stdin.Fd())) {
			// Interactive form
			form := huh.NewForm(
				huh.NewGroup(),
			)

			if err := form.Run(); err != nil {
				return err
			}
		} else {
			return fmt.Errorf("--data flag required in non-interactive mode")
		}

		opts := &sdk.GetDataElementsByFiltersOptions{}
		if datasourcesDataelementsGetDataElementsByFiltersRegexFilter != "" {
			opts.RegexFilter = &datasourcesDataelementsGetDataElementsByFiltersRegexFilter
		}

		result, err := client.DatasourcesDataelements.GetDataElementsByFilters(ctx, datasourceId, &body, opts)
		if err != nil {
			return err
		}

		return output.Print(result)
	},
}

func init() {
	datasourcesDataelementsCmd.AddCommand(datasourcesDataelementsGetDataElementsByFiltersCmd)

	datasourcesDataelementsGetDataElementsByFiltersCmd.Flags().StringVarP(&datasourcesDataelementsGetDataElementsByFiltersRegexFilter, "regex-filter", "", "", "The regex-filter parameter")
	datasourcesDataelementsGetDataElementsByFiltersCmd.Flags().StringSliceVarP(&datasourcesDataelementsGetDataElementsByFiltersMediaTypeFilters, "media-type-filters", "", nil, "The media-type-filters parameter")
	datasourcesDataelementsGetDataElementsByFiltersCmd.Flags().StringVarP(&datasourcesDataelementsGetDataElementsByFiltersData, "data", "d", "", "JSON data for the request body")
	datasourcesDataelementsGetDataElementsByFiltersCmd.Flags().BoolVarP(&datasourcesDataelementsGetDataElementsByFiltersInteractive, "interactive", "i", false, "use interactive form input")
}
