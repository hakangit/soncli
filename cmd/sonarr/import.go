package sonarr

import (
	"fmt"

	"github.com/spf13/cobra"
	"sonarr-sabnzbd-cli/cmd"
)

// importCmd represents the import command
var importCmd = &cobra.Command{
	Use:   "import <path>",
	Short: "Import downloaded files",
	Long: `Scan a directory for downloaded episode files and import them into Sonarr.

This command tells Sonarr to scan the specified path for episode files that
may have been downloaded outside of Sonarr's normal process.

Examples:
  sonarr import "/downloads/complete/TV Shows"
  sonarr import "/tmp/episodes"`,
	Args: cobra.ExactArgs(1),
	RunE: func(command *cobra.Command, args []string) error {
		path := args[0]

		// Import downloads
		err := cmd.GetSonarrClient().ImportDownloads(path)
		if err != nil {
			return fmt.Errorf("failed to import downloads: %w", err)
		}

		fmt.Printf("✅ Successfully initiated import scan for: %s\n", path)
		fmt.Println("Check Sonarr logs for import progress and results.")
		return nil
	},
}

func init() {
	sonarrCmd.AddCommand(importCmd)
}
