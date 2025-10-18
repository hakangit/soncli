package sonarr

import (
	"fmt"

	"github.com/spf13/cobra"
	"sonarr-sabnzbd-cli/cmd"
)

// profilesCmd represents the profiles command
var profilesCmd = &cobra.Command{
	Use:   "profiles",
	Short: "List quality profiles",
	Long: `Display all available quality profiles configured in Sonarr.

Quality profiles determine the quality and format preferences for downloads.

Examples:
  sonarr profiles`,
	RunE: func(command *cobra.Command, args []string) error {
		// Get quality profiles
		profiles, err := cmd.GetSonarrClient().GetQualityProfiles()
		if err != nil {
			return fmt.Errorf("failed to get quality profiles: %w", err)
		}

		if len(profiles) == 0 {
			fmt.Println("No quality profiles found.")
			return nil
		}

		fmt.Printf("Quality Profiles (%d):\n\n", len(profiles))

		for i, profile := range profiles {
			fmt.Printf("%d. %s (ID: %d)\n", i+1, profile.Name, profile.ID)
		}

		return nil
	},
}

func init() {
	sonarrCmd.AddCommand(profilesCmd)
}
