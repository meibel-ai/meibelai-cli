package cmd

import (
	"fmt"
	"strings"
	"syscall"

	"github.com/meibel-ai/meibel-cli/internal/config"
	"github.com/spf13/cobra"
	"golang.org/x/term"
)

var authCmd = &cobra.Command{
	Use:   "auth",
	Short: "Manage authentication",
	Long:  `Configure authentication for the Meibel AI API.`,
}

var authLoginCmd = &cobra.Command{
	Use:   "login",
	Short: "Set up API authentication",
	RunE: func(cmd *cobra.Command, args []string) error {
		interactive, _ := cmd.Flags().GetBool("interactive")
		apiKey, _ := cmd.Flags().GetString("api-key")

		if !interactive && apiKey == "" {
			return fmt.Errorf("provide --api-key or use --interactive mode")
		}

		if interactive {
			fmt.Print("Enter your API key: ")
			byteKey, err := term.ReadPassword(int(syscall.Stdin))
			if err != nil {
				return fmt.Errorf("failed to read API key: %w", err)
			}
			fmt.Println()
			apiKey = string(byteKey)
		}

		cfg := config.New()
		if err := cfg.Load(""); err != nil {
			return err
		}

		cfg.Set("api_key", apiKey)

		server, _ := cmd.Flags().GetString("server")
		if server != "" {
			cfg.Set("server", server)
		}

		if err := cfg.Save(); err != nil {
			return fmt.Errorf("failed to save configuration: %w", err)
		}

		fmt.Println("Authentication configured successfully!")
		return nil
	},
}

var authStatusCmd = &cobra.Command{
	Use:   "status",
	Short: "Check authentication status",
	RunE: func(cmd *cobra.Command, args []string) error {
		cfg := config.New()
		if err := cfg.Load(""); err != nil {
			return err
		}

		apiKey := cfg.GetString("api_key")
		server := cfg.GetString("server")
		profile := cfg.GetString("profile")

		fmt.Printf("Profile: %s\n", profile)
		fmt.Printf("Server: %s\n", server)
		if apiKey != "" {
			fmt.Printf("API Key: %s...%s\n", apiKey[:min(8, len(apiKey))], strings.Repeat("*", max(0, len(apiKey)-8)))
		} else {
			fmt.Println("API Key: Not configured")
		}

		return nil
	},
}

var authLogoutCmd = &cobra.Command{
	Use:   "logout",
	Short: "Clear stored credentials",
	RunE: func(cmd *cobra.Command, args []string) error {
		cfg := config.New()
		if err := cfg.Load(""); err != nil {
			return err
		}

		cfg.Set("api_key", "")

		if err := cfg.Save(); err != nil {
			return fmt.Errorf("failed to save configuration: %w", err)
		}

		fmt.Println("Logged out successfully")
		return nil
	},
}

func init() {
	authLoginCmd.Flags().StringP("api-key", "k", "", "API key")
	authLoginCmd.Flags().StringP("server", "s", "", "API server URL")
	authLoginCmd.Flags().BoolP("interactive", "i", false, "Interactive mode")

	authCmd.AddCommand(authLoginCmd)
	authCmd.AddCommand(authStatusCmd)
	authCmd.AddCommand(authLogoutCmd)
	// Don't add to rootCmd here - it's added in RegisterCustomCommands
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
