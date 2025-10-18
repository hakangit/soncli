package sabnzbd

import (
	"fmt"
	"strings"

	"github.com/spf13/cobra"
	"sonarr-sabnzbd-cli/cmd"
)

// infoCmd represents the info command
var infoCmd = &cobra.Command{
	Use:   "info",
	Short: "Show Sabnzbd system info",
	Long: `Display system information and status for Sabnzbd.

Examples:
  sabnzbd info`,
	RunE: func(command *cobra.Command, args []string) error {
		client := cmd.GetSabnzbdClient()

		// Get version
		version, err := client.GetVersion()
		if err != nil {
			return fmt.Errorf("failed to get version: %w", err)
		}

		// Get queue status
		queue, err := client.GetQueue()
		if err != nil {
			return fmt.Errorf("failed to get queue: %w", err)
		}

		// Cool header
		fmt.Println("🚀 Sabnzbd System Information")
		fmt.Println(strings.Repeat("═", 50))

		// Status with icon
		statusIcon := "🟢"
		if queue.Paused {
			statusIcon = "⏸️"
		}

		fmt.Printf("📦 Version: %s\n", version)
		fmt.Printf("📊 Status: %s %s\n", statusIcon, queue.Status)
		fmt.Printf("⚡ Speed: %s\n", queue.Speed)

		if queue.SpeedLimit != "" && queue.SpeedLimit != "100" {
			fmt.Printf("🎛️  Speed Limit: %s%%\n", queue.SpeedLimit)
		}

		fmt.Printf("📥 Active Downloads: %d\n", len(queue.Slots))

		if queue.TimeLeft != "" && queue.TimeLeft != "0:00:00" {
			fmt.Printf("⏰ Time Left: %s\n", queue.TimeLeft)
		}

		if queue.SizeLeft != "" && queue.SizeLeft != "0 B" {
			fmt.Printf("💾 Size Left: %s\n", queue.SizeLeft)
		}

		// Queue size info
		if queue.Size != "" {
			fmt.Printf("📊 Total Queue Size: %s\n", queue.Size)
		}

		return nil
	},
}

func init() {
	sabnzbdCmd.AddCommand(infoCmd)
}
