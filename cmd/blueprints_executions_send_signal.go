package cmd

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/meibel-ai/meibel-cli/internal/output"
	"github.com/spf13/cobra"
)

var (
	blueprintsExecutionsSendSignalData        string
	blueprintsExecutionsSendSignalInteractive bool
)

var blueprintsExecutionsSendSignalCmd = &cobra.Command{
	Use:   "send-signal <blueprint-instance-id> <signal-name>",
	Short: "Send Signal",
	Long: `Send Signal

Arguments:
  blueprint-instance-id: Unique identifier for the workflow instance
  signal-name: Name of the signal to send`,
	Args:    cobra.ExactArgs(2),
	Example: "meibel blueprints executions send-signal <blueprint-instance-id> <signal-name>",
	RunE: func(cmd *cobra.Command, args []string) error {
		ctx := context.Background()

		blueprintInstanceId := args[0]
		signalName := args[1]

		var body interface{}

		if blueprintsExecutionsSendSignalData != "" {
			if err := json.Unmarshal([]byte(blueprintsExecutionsSendSignalData), &body); err != nil {
				return fmt.Errorf("invalid JSON data: %w", err)
			}
		} else {
			return fmt.Errorf("--data flag required (interactive form not available for this type)")
		}

		result, err := client.BlueprintsExecutions.SendSignal(ctx, blueprintInstanceId, signalName, &body)
		if err != nil {
			return err
		}

		return output.Print(result)
	},
}

func init() {
	blueprintsExecutionsCmd.AddCommand(blueprintsExecutionsSendSignalCmd)

	blueprintsExecutionsSendSignalCmd.Flags().StringVarP(&blueprintsExecutionsSendSignalData, "data", "d", "", "JSON data for the request body")
	blueprintsExecutionsSendSignalCmd.Flags().BoolVarP(&blueprintsExecutionsSendSignalInteractive, "interactive", "i", false, "use interactive form input")
}
