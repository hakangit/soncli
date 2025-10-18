package sabnzbd

import (
	"fmt"

	"github.com/spf13/cobra"
	"sonarr-sabnzbd-cli/cmd"
)

// speedCmd represents the speed command
var speedCmd = &cobra.Command{
	Use:   "speed <limit>",
	Short: "Set download speed limit",
	Long: `Set the download speed limit for Sabnzbd.

The limit can be specified as:
- A percentage (e.g., "50" for 50%)
- An absolute speed (e.g., "10M" for 10 MB/s, "1G" for 1 GB/s)

Examples:
  sabnzbd speed 50      # Set speed to 50%
  sabnzbd speed 10M     # Set speed to 10 MB/s
  sabnzbd speed 0       # Remove speed limit`,
	Args: cobra.ExactArgs(1),
	RunE: func(command *cobra.Command, args []string) error {
		limit := args[0]

		// Set speed limit
		err := cmd.GetSabnzbdClient().SetSpeedLimit(limit)
		if err != nil {
			return fmt.Errorf("failed to set speed limit: %w", err)
		}

		if limit == "0" {
			fmt.Println("✅ Successfully removed speed limit")
		} else {
			fmt.Printf("✅ Successfully set speed limit to %s\n", limit)
		}

		return nil
	},
}

func init() {
	sabnzbdCmd.AddCommand(speedCmd)
}
