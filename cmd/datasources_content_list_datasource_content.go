package cmd

import (
	"context"

	"github.com/spf13/cobra"
	"github.com/meibel-ai/meibel-cli/internal/output"
	sdk "github.com/meibel-ai/meibel-go"
)

var (
	datasourcesContentListDatasourceContentPrefix string
	datasourcesContentListDatasourceContentContinuationToken string
	datasourcesContentListDatasourceContentLimit int64
)

var datasourcesContentListDatasourceContentCmd = &cobra.Command{
	Use:   "list <datasource-id>",
	Short: "List datasource content",
	Long:  `List files and directories in a datasource with optional filtering and pagination

Arguments:
  datasource-id: required`,
	Args:  cobra.ExactArgs(1),
	Example: "meibel datasources content list <datasource-id> --prefix=<value> --continuation-token=<value>",
	RunE: func(cmd *cobra.Command, args []string) error {
		ctx := context.Background()

		datasourceId := args[0]

		opts := &sdk.ListDatasourceContentOptions{}
		if datasourcesContentListDatasourceContentPrefix != "" {
			opts.Prefix = &datasourcesContentListDatasourceContentPrefix
		}
		if datasourcesContentListDatasourceContentContinuationToken != "" {
			opts.ContinuationToken = &datasourcesContentListDatasourceContentContinuationToken
		}
		if datasourcesContentListDatasourceContentLimit != 0 {
			opts.Limit = &datasourcesContentListDatasourceContentLimit
		}

		iter := client.DatasourcesContent.ListDatasourceContent(ctx, datasourceId, opts)

		var items []interface{}
		for iter.Next(ctx) {
			items = append(items, iter.Item())
		}
		if err := iter.Err(); err != nil {
			return err
		}

		return output.Print(items)
	},
}

func init() {
	datasourcesContentCmd.AddCommand(datasourcesContentListDatasourceContentCmd)

	datasourcesContentListDatasourceContentCmd.Flags().StringVarP(&datasourcesContentListDatasourceContentPrefix, "prefix", "", "", "Filter content by path prefix")
	datasourcesContentListDatasourceContentCmd.Flags().StringVarP(&datasourcesContentListDatasourceContentContinuationToken, "continuation-token", "", "", "Token for pagination to get next page of results")
	datasourcesContentListDatasourceContentCmd.Flags().Int64VarP(&datasourcesContentListDatasourceContentLimit, "limit", "", 1000, "Maximum number of items to return (1-10000)")
}
