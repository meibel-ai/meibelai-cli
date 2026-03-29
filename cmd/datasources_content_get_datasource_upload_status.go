package cmd

import (
	"context"

	"github.com/meibel-ai/meibel-cli/internal/output"
	"github.com/spf13/cobra"
)

var datasourcesContentGetDatasourceUploadStatusCmd = &cobra.Command{
	Use:   "get-upload-status <datasource-id> <upload-id>",
	Short: "Get upload status",
	Long: `Get the current status of a content upload operation

Arguments:
  datasource-id: required
  upload-id: required`,
	Args:    cobra.ExactArgs(2),
	Example: "meibel datasources content get-upload-status <datasource-id> <upload-id>",
	RunE: func(cmd *cobra.Command, args []string) error {
		ctx := context.Background()

		datasourceId := args[0]
		uploadId := args[1]

		result, err := client.DatasourcesContent.GetDatasourceUploadStatus(ctx, datasourceId, uploadId)
		if err != nil {
			return err
		}

		return output.Print(result)
	},
}

func init() {
	datasourcesContentCmd.AddCommand(datasourcesContentGetDatasourceUploadStatusCmd)

}
