package cmd

import (
	"fmt"
	"runtime"

	"github.com/meibel-ai/meibel-cli/internal/version"
	"github.com/spf13/cobra"
)

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print version information",
	Long:  `Print detailed version information about the Meibel CLI.`,
	Run: func(cmd *cobra.Command, args []string) {
		verbose, _ := cmd.Flags().GetBool("verbose")

		if verbose {
			fmt.Printf("Meibel CLI %s\n", version.Version)
			fmt.Printf("  Commit: %s\n", version.Commit)
			fmt.Printf("  Built: %s\n", version.BuildDate)
			fmt.Printf("  Go version: %s\n", runtime.Version())
			fmt.Printf("  OS/Arch: %s/%s\n", runtime.GOOS, runtime.GOARCH)
		} else {
			fmt.Printf("meibel version %s\n", version.Version)
		}
	},
}

func init() {
	versionCmd.Flags().BoolP("verbose", "v", false, "Show verbose version information")
	// Don't add to rootCmd here - it's added in RegisterCustomCommands
}
