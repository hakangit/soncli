package sabnzbd

import (
	"fmt"

	"github.com/spf13/cobra"
	"sonarr-sabnzbd-cli/cmd"
)

// resumeCmd represents the resume command
var resumeCmd = &cobra.Command{
	Use:   "resume",
	Short: "Resume all downloads",
	Long: `Resume all downloads in the Sabnzbd queue.

Examples:
  sabnzbd resume`,
	RunE: func(command *cobra.Command, args []string) error {
		// Resume the queue
		err := cmd.GetSabnzbdClient().ResumeQueue()
		if err != nil {
			return fmt.Errorf("failed to resume queue: %w", err)
		}

		fmt.Println("✅ Successfully resumed all downloads")
		return nil
	},
}

func init() {
	sabnzbdCmd.AddCommand(resumeCmd)
}
