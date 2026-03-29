package cmd

import (
	"context"

	"github.com/spf13/cobra"
	"github.com/meibel-ai/meibel-cli/internal/output"
	sdk "github.com/meibel-ai/meibel-go"
)

var (
	confidenceScoringGetAllScoringJobsAgentName string
	confidenceScoringGetAllScoringJobsAgentVersion string
	confidenceScoringGetAllScoringJobsAgentExecutionId string
	confidenceScoringGetAllScoringJobsAgentWorkflowName string
	confidenceScoringGetAllScoringJobsAgentWorkflowVersion string
	confidenceScoringGetAllScoringJobsAgentWorkflowExecutionId string
	confidenceScoringGetAllScoringJobsToolId string
	confidenceScoringGetAllScoringJobsToolInstanceId string
	confidenceScoringGetAllScoringJobsToolExecutionId string
)

var confidenceScoringGetAllScoringJobsCmd = &cobra.Command{
	Use:   "get-all-jobs",
	Short: "Get All Scoring Jobs",
	Long:  `Get all scoring jobs for the caller's customer.`,
	Example: "meibel confidence-scoring get-all-jobs --agent-name=<value> --agent-version=<value>",
	RunE: func(cmd *cobra.Command, args []string) error {
		ctx := context.Background()

		opts := &sdk.GetAllScoringJobsOptions{}
		if confidenceScoringGetAllScoringJobsAgentName != "" {
			opts.AgentName = &confidenceScoringGetAllScoringJobsAgentName
		}
		if confidenceScoringGetAllScoringJobsAgentVersion != "" {
			opts.AgentVersion = &confidenceScoringGetAllScoringJobsAgentVersion
		}
		if confidenceScoringGetAllScoringJobsAgentExecutionId != "" {
			opts.AgentExecutionId = &confidenceScoringGetAllScoringJobsAgentExecutionId
		}
		if confidenceScoringGetAllScoringJobsAgentWorkflowName != "" {
			opts.AgentWorkflowName = &confidenceScoringGetAllScoringJobsAgentWorkflowName
		}
		if confidenceScoringGetAllScoringJobsAgentWorkflowVersion != "" {
			opts.AgentWorkflowVersion = &confidenceScoringGetAllScoringJobsAgentWorkflowVersion
		}
		if confidenceScoringGetAllScoringJobsAgentWorkflowExecutionId != "" {
			opts.AgentWorkflowExecutionId = &confidenceScoringGetAllScoringJobsAgentWorkflowExecutionId
		}
		if confidenceScoringGetAllScoringJobsToolId != "" {
			opts.ToolId = &confidenceScoringGetAllScoringJobsToolId
		}
		if confidenceScoringGetAllScoringJobsToolInstanceId != "" {
			opts.ToolInstanceId = &confidenceScoringGetAllScoringJobsToolInstanceId
		}
		if confidenceScoringGetAllScoringJobsToolExecutionId != "" {
			opts.ToolExecutionId = &confidenceScoringGetAllScoringJobsToolExecutionId
		}

		result, err := client.ConfidenceScoring.GetAllScoringJobs(ctx, opts)
		if err != nil {
			return err
		}

		return output.Print(result)
	},
}

func init() {
	confidenceScoringCmd.AddCommand(confidenceScoringGetAllScoringJobsCmd)

	confidenceScoringGetAllScoringJobsCmd.Flags().StringVarP(&confidenceScoringGetAllScoringJobsAgentName, "agent-name", "", "", "The agent-name parameter")
	confidenceScoringGetAllScoringJobsCmd.Flags().StringVarP(&confidenceScoringGetAllScoringJobsAgentVersion, "agent-version", "", "", "The agent-version parameter")
	confidenceScoringGetAllScoringJobsCmd.Flags().StringVarP(&confidenceScoringGetAllScoringJobsAgentExecutionId, "agent-execution-id", "", "", "The agent-execution-id parameter")
	confidenceScoringGetAllScoringJobsCmd.Flags().StringVarP(&confidenceScoringGetAllScoringJobsAgentWorkflowName, "agent-workflow-name", "", "", "The agent-workflow-name parameter")
	confidenceScoringGetAllScoringJobsCmd.Flags().StringVarP(&confidenceScoringGetAllScoringJobsAgentWorkflowVersion, "agent-workflow-version", "", "", "The agent-workflow-version parameter")
	confidenceScoringGetAllScoringJobsCmd.Flags().StringVarP(&confidenceScoringGetAllScoringJobsAgentWorkflowExecutionId, "agent-workflow-execution-id", "", "", "The agent-workflow-execution-id parameter")
	confidenceScoringGetAllScoringJobsCmd.Flags().StringVarP(&confidenceScoringGetAllScoringJobsToolId, "tool-id", "", "", "The tool-id parameter")
	confidenceScoringGetAllScoringJobsCmd.Flags().StringVarP(&confidenceScoringGetAllScoringJobsToolInstanceId, "tool-instance-id", "", "", "The tool-instance-id parameter")
	confidenceScoringGetAllScoringJobsCmd.Flags().StringVarP(&confidenceScoringGetAllScoringJobsToolExecutionId, "tool-execution-id", "", "", "The tool-execution-id parameter")
}
