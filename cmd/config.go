package cmd

import (
	"fmt"

	"github.com/meibel-ai/meibel-cli/internal/config"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var configCmd = &cobra.Command{
	Use:   "config",
	Short: "Manage configuration",
	Long:  `Manage configuration settings for the Meibel CLI.`,
}

var configSetCmd = &cobra.Command{
	Use:   "set <key> <value>",
	Short: "Set a configuration value",
	Args:  cobra.ExactArgs(2),
	RunE: func(cmd *cobra.Command, args []string) error {
		cfg := config.New()
		if err := cfg.Load(""); err != nil {
			return err
		}

		key := args[0]
		value := args[1]

		cfg.Set(key, value)

		if err := cfg.Save(); err != nil {
			return fmt.Errorf("failed to save configuration: %w", err)
		}

		fmt.Printf("Configuration updated: %s = %s\n", key, value)
		return nil
	},
}

var configGetCmd = &cobra.Command{
	Use:   "get <key>",
	Short: "Get a configuration value",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		cfg := config.New()
		if err := cfg.Load(""); err != nil {
			return err
		}

		key := args[0]
		value := cfg.Get(key)

		if value != nil {
			fmt.Printf("%v\n", value)
		} else {
			fmt.Printf("Configuration key '%s' not found\n", key)
		}

		return nil
	},
}

var configListCmd = &cobra.Command{
	Use:   "list",
	Short: "List all configuration values",
	RunE: func(cmd *cobra.Command, args []string) error {
		cfg := config.New()
		if err := cfg.Load(""); err != nil {
			return err
		}

		settings := viper.AllSettings()
		for key, value := range settings {
			fmt.Printf("%s: %v\n", key, value)
		}

		return nil
	},
}

func init() {
	configCmd.AddCommand(configSetCmd)
	configCmd.AddCommand(configGetCmd)
	configCmd.AddCommand(configListCmd)
	// Don't add to rootCmd here - it's added in RegisterCustomCommands
}
