package cmd

import (
	"fmt"
	"os"

	"github.com/meibel-ai/meibel-cli/internal/client"
	"github.com/meibel-ai/meibel-cli/internal/config"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	cfgFile string
	rootCmd = &cobra.Command{
		Use:   "meibel",
		Short: "CLI for interacting with Meibel AI API",
		Long:  `A command-line interface for the Meibel AI API with both generated and custom commands.`,
	}
)

func Execute() error {
	return rootCmd.Execute()
}

func init() {
	cobra.OnInitialize(initConfig)

	// Global flags
	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.meibel.yaml)")
	rootCmd.PersistentFlags().String("api-key", "", "API key for authentication")
	rootCmd.PersistentFlags().String("server", "", "API server URL")
	rootCmd.PersistentFlags().String("output", "json", "Output format (json, yaml, table)")
	rootCmd.PersistentFlags().Bool("dry-run", false, "Preview the request without executing")
	rootCmd.PersistentFlags().String("profile", "default", "Configuration profile to use")

	// Bind flags to viper
	viper.BindPFlag("api_key", rootCmd.PersistentFlags().Lookup("api-key"))
	viper.BindPFlag("server", rootCmd.PersistentFlags().Lookup("server"))
	viper.BindPFlag("output", rootCmd.PersistentFlags().Lookup("output"))
	viper.BindPFlag("profile", rootCmd.PersistentFlags().Lookup("profile"))

	// Initialize commands
	httpClient := client.New()

	// Register custom commands (non-generated)
	RegisterCustomCommands(rootCmd, httpClient)

	// Register generated commands from OpenAPI
	RegisterGeneratedCommands(rootCmd, httpClient)
}

func initConfig() {
	cfg := config.New()
	if err := cfg.Load(cfgFile); err != nil {
		fmt.Fprintf(os.Stderr, "Error loading config: %v\n", err)
	}
}

// RegisterCustomCommands adds all hand-written commands
func RegisterCustomCommands(rootCmd *cobra.Command, httpClient *client.Client) {
	// Add auth command
	rootCmd.AddCommand(authCmd)

	// Add config command
	rootCmd.AddCommand(configCmd)

	// Add version command
	rootCmd.AddCommand(versionCmd)

	// Add any other custom commands here
	// Examples:
	// - batch processing commands
	// - export/import commands
	// - interactive setup commands
	// - debug/troubleshooting commands
}
