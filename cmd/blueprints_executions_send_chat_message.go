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
	blueprintsExecutionsSendChatMessageData        string
	blueprintsExecutionsSendChatMessageInteractive bool
)

var blueprintsExecutionsSendChatMessageCmd = &cobra.Command{
	Use:   "send-chat-message <blueprint-instance-id>",
	Short: "Send Chat Message",
	Long: `Send Chat Message

Arguments:
  blueprint-instance-id: required`,
	Args:    cobra.ExactArgs(1),
	Example: "meibel blueprints executions send-chat-message <blueprint-instance-id>",
	RunE: func(cmd *cobra.Command, args []string) error {
		ctx := context.Background()

		blueprintInstanceId := args[0]

		var body sdk.ChatMessageRequest

		if blueprintsExecutionsSendChatMessageData != "" {
			if err := json.Unmarshal([]byte(blueprintsExecutionsSendChatMessageData), &body); err != nil {
				return fmt.Errorf("invalid JSON data: %w", err)
			}
		} else if blueprintsExecutionsSendChatMessageInteractive || term.IsTerminal(int(os.Stdin.Fd())) {
			// Interactive form
			form := huh.NewForm(
				huh.NewGroup(
					huh.NewInput().Title("UserMessage").Description("The user's chat message").Value(&body.UserMessage),
				),
			)

			if err := form.Run(); err != nil {
				return err
			}
		} else {
			return fmt.Errorf("--data flag required in non-interactive mode")
		}

		result, err := client.BlueprintsExecutions.SendChatMessage(ctx, blueprintInstanceId, body)
		if err != nil {
			return err
		}

		return output.Print(result)
	},
}

func init() {
	blueprintsExecutionsCmd.AddCommand(blueprintsExecutionsSendChatMessageCmd)

	blueprintsExecutionsSendChatMessageCmd.Flags().StringVarP(&blueprintsExecutionsSendChatMessageData, "data", "d", "", "JSON data for the request body")
	blueprintsExecutionsSendChatMessageCmd.Flags().BoolVarP(&blueprintsExecutionsSendChatMessageInteractive, "interactive", "i", false, "use interactive form input")
}
