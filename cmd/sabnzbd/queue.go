package sabnzbd

import (
	"encoding/json"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/spf13/cobra"
	"sonarr-sabnzbd-cli/cmd"
)

// queueCmd represents the queue command
var queueCmd = &cobra.Command{
	Use:   "queue",
	Short: "View current download queue",
	Long: `Display all downloads currently in your Sabnzbd queue.

Examples:
   sabnzbd queue                    # View all queued downloads
   sabnzbd queue --json             # Output in JSON format
   sabnzbd queue | head -10         # View first 10 downloads`,
	RunE: func(command *cobra.Command, args []string) error {
		jsonOutput, _ := command.Flags().GetBool("json")

		// Get queue from Sabnzbd
		queue, err := cmd.GetSabnzbdClient().GetQueue()
		if err != nil {
			return fmt.Errorf("failed to get queue: %w", err)
		}

		// JSON output mode
		if jsonOutput {
			return json.NewEncoder(os.Stdout).Encode(queue)
		}

		if len(queue.Slots) == 0 {
			fmt.Println("Download queue is empty.")
			return nil
		}

		fmt.Printf("📥 Download Queue (%d active)\n", len(queue.Slots))
		fmt.Println(strings.Repeat("─", 80))

		for i, slot := range queue.Slots {
			status := getStatusIcon(slot.Status)
			statusText := slot.Status
			if statusText == "" {
				statusText = "Queued"
			}

			// Progress bar
			progressBar := ""
			if slot.Percentage != "" && slot.Percentage != "0" {
				percentage, _ := strconv.Atoi(slot.Percentage)
				progressBar = createProgressBar(percentage, 20)
			}

			fmt.Printf("%d. %s %s\n", i+1, status, slot.Name)
			fmt.Printf("   📏 Size: %s | ⏱️  ETA: %s\n", slot.Size, slot.ETA)

			if progressBar != "" {
				fmt.Printf("   %s %s%%\n", progressBar, slot.Percentage)
			}

			if slot.Category != "" && slot.Category != "*" {
				fmt.Printf("   📂 Category: %s\n", slot.Category)
			}
			fmt.Println()
		}

		// Show overall queue status with cool formatting
		fmt.Println(strings.Repeat("─", 80))
		status := "🚀 Downloading"
		if queue.Paused {
			status = "⏸️  Paused"
		}
		fmt.Printf("📊 Queue Status: %s\n", status)
		fmt.Printf("⚡ Speed: %s\n", queue.Speed)
		fmt.Printf("⏰ Time Left: %s\n", queue.TimeLeft)
		fmt.Printf("💾 Size Left: %s\n", queue.SizeLeft)

		return nil
	},
}

func init() {
	sabnzbdCmd.AddCommand(queueCmd)
}

// getStatusIcon returns an appropriate icon for the download status
func getStatusIcon(status string) string {
	switch strings.ToLower(status) {
	case "downloading":
		return "⬇️"
	case "queued":
		return "⏳"
	case "paused":
		return "⏸️"
	case "failed":
		return "❌"
	case "completed":
		return "✅"
	default:
		return "📄"
	}
}

// createProgressBar creates a visual progress bar
func createProgressBar(percentage int, width int) string {
	if percentage < 0 {
		percentage = 0
	}
	if percentage > 100 {
		percentage = 100
	}

	filled := percentage * width / 100
	empty := width - filled

	bar := strings.Repeat("█", filled) + strings.Repeat("░", empty)
	return fmt.Sprintf("[%s]", bar)
}
