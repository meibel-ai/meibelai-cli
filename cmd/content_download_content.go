package cmd

import (
	"context"
	"fmt"

	"github.com/spf13/cobra"
)

var contentDownloadContentCmd = &cobra.Command{
	Use:   "download <datasource-id> <path>",
	Short: "Download Content",
	Long: `Download Content

Arguments:
  datasource-id: required
  path: required`,
	Args:    cobra.ExactArgs(2),
	Example: "meibel content download <datasource-id> <path>",
	RunE: func(cmd *cobra.Command, args []string) error {
		ctx := context.Background()

		datasourceId := args[0]
		path := args[1]

		err := client.Content.DownloadContent(ctx, datasourceId, path)
		if err != nil {
			return err
		}

		fmt.Println("Success")
		return nil
	},
}

func init() {
	contentCmd.AddCommand(contentDownloadContentCmd)

}
