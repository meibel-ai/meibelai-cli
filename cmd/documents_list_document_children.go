package cmd

import (
	"context"

	"github.com/meibel-ai/meibel-cli/internal/output"
	"github.com/spf13/cobra"
)

var documentsListDocumentChildrenCmd = &cobra.Command{
	Use:   "list-children <job-id>",
	Short: "List child documents",
	Long: `For container files (ZIP, TAR, EML), list the child documents extracted from the container.

Arguments:
  job-id: required`,
	Args:    cobra.ExactArgs(1),
	Example: "meibel documents list-children <job-id>",
	RunE: func(cmd *cobra.Command, args []string) error {
		ctx := context.Background()

		jobId := args[0]

		iter := client.Documents.ListDocumentChildren(ctx, jobId)

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
	documentsCmd.AddCommand(documentsListDocumentChildrenCmd)

}
