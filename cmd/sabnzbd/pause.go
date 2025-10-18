package sabnzbd

import (
	"fmt"

	"github.com/spf13/cobra"
	"sonarr-sabnzbd-cli/cmd"
)

// pauseCmd represents the pause command
var pauseCmd = &cobra.Command{
	Use:   "pause",
	Short: "Pause all downloads",
	Long: `Pause all downloads in the Sabnzbd queue.

Examples:
  sabnzbd pause`,
	RunE: func(command *cobra.Command, args []string) error {
		// Pause the queue
		err := cmd.GetSabnzbdClient().PauseQueue()
		if err != nil {
			return fmt.Errorf("failed to pause queue: %w", err)
		}

		fmt.Println("✅ Successfully paused all downloads")
		return nil
	},
}

func init() {
	sabnzbdCmd.AddCommand(pauseCmd)
}
