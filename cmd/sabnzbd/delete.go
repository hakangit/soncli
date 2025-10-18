package sabnzbd

import (
	"fmt"

	"github.com/spf13/cobra"
	"sonarr-sabnzbd-cli/cmd"
)

// deleteCmd represents the delete command
var deleteCmd = &cobra.Command{
	Use:   "delete <nzo-id>",
	Short: "Delete job from queue",
	Long: `Delete a download job from the Sabnzbd queue using its NZO ID.

The NZO ID can be found using the 'sabnzbd queue' command.

Examples:
  sabnzbd delete SABnzbd_nzo_12345
  sabnzbd delete SABnzbd_nzo_67890`,
	Args: cobra.ExactArgs(1),
	RunE: func(command *cobra.Command, args []string) error {
		nzoID := args[0]

		// Delete from queue
		err := cmd.GetSabnzbdClient().DeleteFromQueue(nzoID)
		if err != nil {
			return fmt.Errorf("failed to delete job %s: %w", nzoID, err)
		}

		fmt.Printf("✅ Successfully deleted job %s from queue\n", nzoID)
		return nil
	},
}

func init() {
	sabnzbdCmd.AddCommand(deleteCmd)
}
