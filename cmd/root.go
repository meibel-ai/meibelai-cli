package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/meibel-ai/meibel-cli/internal/config"
	"github.com/meibel-ai/meibel-cli/internal/output"
	sdk "github.com/meibel-ai/meibel-go"
)

var (
	cfgFile string
	jsonOutput bool
	debug bool
	client *sdk.MeibelgoClient
)

var rootCmd = &cobra.Command{
	Use:   "meibel",
	Short: "meibel CLI",
	Long: `The Meibel API provides document parsing, datasource management, and AI agent orchestration.`,
	PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
		// Initialize configuration
		if err := config.Init(cfgFile); err != nil {
			return err
		}

		// Initialize SDK client
		opts := []sdk.ClientOption{}

		if baseURL := viper.GetString("base_url"); baseURL != "" {
			opts = append(opts, sdk.WithBaseURL(baseURL))
		}

		if apiKey := viper.GetString("api_key"); apiKey != "" {
			opts = append(opts, sdk.WithAPIKey(apiKey))
		}

		if token := viper.GetString("token"); token != "" {
			opts = append(opts, sdk.WithBearerToken(token))
		}

		client = sdk.NewClient(opts...)

		// Set output format
		if jsonOutput {
			output.SetFormat(output.FormatJSON)
		}

		return nil
	},
}

// Execute runs the root command.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func init() {
	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.meibel/config.yaml)")
	rootCmd.PersistentFlags().BoolVar(&jsonOutput, "json", false, "output as JSON")
	rootCmd.PersistentFlags().BoolVar(&debug, "debug", false, "enable debug logging")

	// Bind flags to viper
	viper.BindPFlag("json", rootCmd.PersistentFlags().Lookup("json"))
	viper.BindPFlag("debug", rootCmd.PersistentFlags().Lookup("debug"))
}
