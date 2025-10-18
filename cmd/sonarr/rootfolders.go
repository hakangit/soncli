package sonarr

import (
	"fmt"

	"github.com/spf13/cobra"
	"sonarr-sabnzbd-cli/cmd"
)

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

// rootfoldersCmd represents the root-folders command
var rootfoldersCmd = &cobra.Command{
	Use:   "root-folders",
	Short: "List root folders",
	Long: `Display all configured root folders in Sonarr.

Root folders are the base directories where Sonarr stores TV series.

Examples:
  sonarr root-folders`,
	RunE: func(command *cobra.Command, args []string) error {
		// Get root folders
		folders, err := cmd.GetSonarrClient().GetRootFolders()
		if err != nil {
			return fmt.Errorf("failed to get root folders: %w", err)
		}

		if len(folders) == 0 {
			fmt.Println("No root folders configured.")
			return nil
		}

		fmt.Printf("Root Folders (%d):\n\n", len(folders))

		for i, folder := range folders {
			fmt.Printf("%d. %s\n", i+1, folder.Path)
			fmt.Printf("   Free Space: %s\n", formatBytes(folder.FreeSpace))
			fmt.Println()
		}

		return nil
	},
}

func init() {
	sonarrCmd.AddCommand(rootfoldersCmd)
}
