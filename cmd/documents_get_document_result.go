package cmd

import (
	"context"

	"github.com/meibel-ai/meibel-cli/internal/output"
	sdk "github.com/meibel-ai/meibel-go"
	"github.com/spf13/cobra"
)

var (
	documentsGetDocumentResultFormat string
)

var documentsGetDocumentResultCmd = &cobra.Command{
	Use:   "get-result <job-id>",
	Short: "Get parsed document result",
	Long: `Download the parsed result of a completed document parsing job.

Arguments:
  job-id: required`,
	Args:    cobra.ExactArgs(1),
	Example: "meibel documents get-result <job-id> --format=<value>",
	RunE: func(cmd *cobra.Command, args []string) error {
		ctx := context.Background()

		jobId := args[0]

		opts := &sdk.GetDocumentResultOptions{}
		if documentsGetDocumentResultFormat != "" {
			opts.Format = &documentsGetDocumentResultFormat
		}

		result, err := client.Documents.GetDocumentResult(ctx, jobId, opts)
		if err != nil {
			return err
		}

		return output.Print(result)
	},
}

func init() {
	documentsCmd.AddCommand(documentsGetDocumentResultCmd)

	documentsGetDocumentResultCmd.Flags().StringVarP(&documentsGetDocumentResultFormat, "format", "", "markdown", "Result format: markdown, annotated, docling, json")
}
