package cmd

import (
	"context"
	"encoding/json"
	"fmt"
	"os"
	"os/signal"
	"path/filepath"

	"github.com/charmbracelet/huh"
	"github.com/meibel-ai/meibel-cli/internal/config"
	"github.com/meibel-ai/meibel-cli/internal/output"
	"github.com/meibel-ai/meibel-cli/internal/tui"
	sdk "github.com/meibel-ai/meibel-go"
	"github.com/spf13/cobra"
)

var (
	documentsProcessDocumentFormat  string
	documentsProcessDocumentFile    string
	documentsProcessDocumentTrace   bool
	documentsProcessDocumentBrowser bool
)

var documentsProcessDocumentCmd = &cobra.Command{
	Use:     "process",
	Short:   "Parse a document (sync)",
	Long:    `Upload a document and block until parsing is complete. Returns the full parsed result.`,
	Example: "meibel documents process --format=<value>",
	Hidden:  true,
	RunE: func(cmd *cobra.Command, args []string) error {
		ctx := context.Background()

		if documentsProcessDocumentFile == "" {
			home, _ := os.UserHomeDir()
			if home == "" {
				home, _ = os.Getwd()
			}
			picker := huh.NewFilePicker().
				Title("Select a file").
				CurrentDirectory(home).
				FileAllowed(true).
				DirAllowed(false).
				ShowHidden(false).
				ShowSize(true).
				ShowPermissions(false).
				Height(15).
				Value(&documentsProcessDocumentFile)
			if err := huh.NewForm(huh.NewGroup(picker)).Run(); err != nil {
				return err
			}
			if documentsProcessDocumentFile == "" {
				return fmt.Errorf("no file selected")
			}
		}

		f, err := os.Open(documentsProcessDocumentFile)
		if err != nil {
			return fmt.Errorf("failed to open file: %w", err)
		}
		defer f.Close()

		result, err := client.Documents.ProcessDocument(ctx, f, filepath.Base(documentsProcessDocumentFile), opts)
		if err != nil {
			return err
		}

		type jobResult struct {
			JobID string `json:"job_id"`
		}
		var jr jobResult
		b, _ := json.Marshal(result)
		json.Unmarshal(b, &jr)

		if documentsProcessDocumentBrowser && jr.JobID != "" {
			consoleURL := deriveConsoleURL(config.GetString("base_url"))
			projectID := config.GetString("project_id")
			if consoleURL != "" && projectID != "" {
				url := fmt.Sprintf("%s/projects/%s/documents/%s", consoleURL, projectID, jr.JobID)
				openBrowser(url)
			}
		}

		if documentsProcessDocumentTrace && jr.JobID != "" {
			output.Print(result)

			ctx, cancel := signal.NotifyContext(ctx, os.Interrupt)
			defer cancel()

			stream, err := client.Documents.StreamDocumentTrace(ctx, jr.JobID)
			if err != nil {
				return err
			}
			defer stream.Close()

			return tui.StreamEvents(ctx, stream)
		}

		return output.Print(result)
	},
}

func init() {
	documentsCmd.AddCommand(documentsProcessDocumentCmd)

	documentsProcessDocumentCmd.Flags().StringVarP(&documentsProcessDocumentFormat, "format", "", "markdown", "Result format: markdown, annotated, docling, json")
	documentsProcessDocumentCmd.Flags().StringVarP(&documentsProcessDocumentFile, "file", "f", "", "path to file to upload (interactive picker if omitted)")
	documentsProcessDocumentCmd.MarkFlagFilename("file")
	documentsProcessDocumentCmd.Flags().BoolVar(&documentsProcessDocumentTrace, "trace", false, "stream parsing trace after upload")
	documentsProcessDocumentCmd.Flags().BoolVar(&documentsProcessDocumentBrowser, "browser", false, "open trace in console")
}
