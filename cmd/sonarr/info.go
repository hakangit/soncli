package sonarr

import (
	"fmt"

	"github.com/spf13/cobra"
	"sonarr-sabnzbd-cli/cmd"
)

// infoCmd represents the info command
var infoCmd = &cobra.Command{
	Use:   "info",
	Short: "Show Sonarr system info",
	Long: `Display system information and status for Sonarr.

Examples:
  sonarr info`,
	RunE: func(command *cobra.Command, args []string) error {
		// Get system status
		status, err := cmd.GetSonarrClient().GetSystemStatus()
		if err != nil {
			return fmt.Errorf("failed to get system status: %w", err)
		}

		fmt.Println("Sonarr System Information:")
		fmt.Println("==========================")
		fmt.Printf("Version: %s\n", status.Version)
		fmt.Printf("Build Time: %s\n", status.BuildTime)
		fmt.Printf("Is Production: %t\n", status.IsProduction)
		fmt.Printf("Is Admin: %t\n", status.IsAdmin)
		fmt.Printf("Is User Interactive: %t\n", status.IsUserInteractive)
		fmt.Printf("Startup Path: %s\n", status.StartupPath)
		fmt.Printf("App Data: %s\n", status.AppData)
		fmt.Printf("OS Name: %s\n", status.OsName)
		fmt.Printf("OS Version: %s\n", status.OsVersion)

		return nil
	},
}

func init() {
	sonarrCmd.AddCommand(infoCmd)
}
