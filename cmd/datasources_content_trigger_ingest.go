package cmd

import (
	"context"

	"github.com/spf13/cobra"
	"github.com/meibel-ai/meibel-cli/internal/output"
)

var datasourcesContentTriggerIngestCmd = &cobra.Command{
	Use:   "trigger-ingest <datasource-id>",
	Short: "Trigger ingest",
	Long:  `Trigger ingestion for a datasource

Arguments:
  datasource-id: required`,
	Args:  cobra.ExactArgs(1),
	Example: "meibel datasources content trigger-ingest <datasource-id>",
	RunE: func(cmd *cobra.Command, args []string) error {
		ctx := context.Background()

		datasourceId := args[0]

		result, err := client.DatasourcesContent.TriggerIngest(ctx, datasourceId)
		if err != nil {
			return err
		}

		return output.Print(result)
	},
}

func init() {
	datasourcesContentCmd.AddCommand(datasourcesContentTriggerIngestCmd)

}
