package cmd

import (
	"github.com/spf13/cobra"
)

var confidenceScoringCmd = &cobra.Command{
	Use:   "confidence-scoring",
	Short: "Manage confidence_scoring",
	Long:  `Commands for managing confidence_scoring resources.`,
}

func init() {
	rootCmd.AddCommand(confidenceScoringCmd)
}
