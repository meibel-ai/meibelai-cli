package cmd

import (
	"context"
	"encoding/json"
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/charmbracelet/huh"
	"golang.org/x/term"
	sdk "github.com/meibel-ai/meibel-go"
)

var (
	blueprintsExecutionsSendChatMessageStreamData string
	blueprintsExecutionsSendChatMessageStreamInteractive bool
)

var blueprintsExecutionsSendChatMessageStreamCmd = &cobra.Command{
	Use:   "send-chat-message-stream <blueprint-instance-id>",
	Short: "Send a chat message and stream the response via SSE",
	Long:  `Send a chat message to a running chat agent workflow and stream the response as Server-Sent Events.

Arguments:
  blueprint-instance-id: required`,
	Args:  cobra.ExactArgs(1),
	Example: "meibel blueprints executions send-chat-message-stream <blueprint-instance-id>",
	RunE: func(cmd *cobra.Command, args []string) error {
		ctx := context.Background()

		blueprintInstanceId := args[0]

		var body sdk.ChatMessageRequest

		if blueprintsExecutionsSendChatMessageStreamData != "" {
			if err := json.Unmarshal([]byte(blueprintsExecutionsSendChatMessageStreamData), &body); err != nil {
				return fmt.Errorf("invalid JSON data: %w", err)
			}
		} else if blueprintsExecutionsSendChatMessageStreamInteractive || term.IsTerminal(int(os.Stdin.Fd())) {
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

		err := client.BlueprintsExecutions.SendChatMessageStream(ctx, blueprintInstanceId, body)
		if err != nil {
			return err
		}

		fmt.Println("Success")
		return nil
	},
}

func init() {
	blueprintsExecutionsCmd.AddCommand(blueprintsExecutionsSendChatMessageStreamCmd)

	blueprintsExecutionsSendChatMessageStreamCmd.Flags().StringVarP(&blueprintsExecutionsSendChatMessageStreamData, "data", "d", "", "JSON data for the request body")
	blueprintsExecutionsSendChatMessageStreamCmd.Flags().BoolVarP(&blueprintsExecutionsSendChatMessageStreamInteractive, "interactive", "i", false, "use interactive form input")
}
