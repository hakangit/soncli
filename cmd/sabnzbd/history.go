package sabnzbd

import (
	"fmt"
	"strings"
	"time"

	"github.com/spf13/cobra"
	"sonarr-sabnzbd-cli/cmd"
)

// historyCmd represents the history command
var historyCmd = &cobra.Command{
	Use:   "history",
	Short: "View download history",
	Long: `Display completed downloads from your Sabnzbd history.

Examples:
  sabnzbd history                    # View recent download history
  sabnzbd history | head -20         # View last 20 downloads`,
	RunE: func(command *cobra.Command, args []string) error {
		// Get history from Sabnzbd
		history, err := cmd.GetSabnzbdClient().GetHistory()
		if err != nil {
			return fmt.Errorf("failed to get history: %w", err)
		}

		if len(history.Slots) == 0 {
			fmt.Println("Download history is empty.")
			return nil
		}

		fmt.Printf("📚 Download History (%d completed)\n", len(history.Slots))
		fmt.Println(strings.Repeat("─", 80))

		for i, slot := range history.Slots {
			// Format completion time
			completionTime := "Unknown"
			if slot.Completed > 0 {
				t := time.Unix(slot.Completed, 0)
				completionTime = t.Format("2006-01-02 15:04:05")
			}

			status := getHistoryStatusIcon(slot.Status)

			fmt.Printf("%d. %s %s\n", i+1, status, slot.Name)
			fmt.Printf("   📏 Size: %s | ⏰ Completed: %s\n",
				formatBytes(slot.Bytes), completionTime)
			if slot.Category != "" && slot.Category != "*" {
				fmt.Printf("   📂 Category: %s\n", slot.Category)
			}
			fmt.Println()
		}

		return nil
	},
}

func init() {
	sabnzbdCmd.AddCommand(historyCmd)
}

// getHistoryStatusIcon returns an appropriate icon for history status
func getHistoryStatusIcon(status string) string {
	switch strings.ToLower(status) {
	case "completed":
		return "✅"
	case "failed":
		return "❌"
	case "repairing":
		return "🔧"
	case "extracting":
		return "📦"
	case "verifying":
		return "🔍"
	default:
		return "📄"
	}
}

// formatBytes formats bytes into human readable format
func formatBytes(bytes int64) string {
	const unit = 1024
	if bytes < unit {
		return fmt.Sprintf("%d B", bytes)
	}
	div, exp := int64(unit), 0
	for n := bytes / unit; n >= unit; n /= unit {
		div *= unit
		exp++
	}
	return fmt.Sprintf("%.1f %cB", float64(bytes)/float64(div), "KMGTPE"[exp])
}
