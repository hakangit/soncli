package sonarr

import (
	"fmt"
	"strconv"

	"github.com/spf13/cobra"
	"sonarr-sabnzbd-cli/cmd"
	"sonarr-sabnzbd-cli/internal/models"
)

// addCmd represents the add command
var addCmd = &cobra.Command{
	Use:   "add <tvdb-id>",
	Short: "Add a TV series to your library by TVDB ID",
	Long: `Add a TV series to your Sonarr library using its TVDB ID.

You can find TVDB IDs using the 'sonarr search' command.

Examples:
  sonarr add 81189    # Add Breaking Bad
  sonarr add 78804    # Add The Office (US)`,
	Args: cobra.ExactArgs(1),
	RunE: func(command *cobra.Command, args []string) error {
		tvdbIDStr := args[0]

		// Parse TVDB ID
		tvdbID, err := strconv.Atoi(tvdbIDStr)
		if err != nil {
			return fmt.Errorf("invalid TVDB ID '%s': must be a number", tvdbIDStr)
		}

		fmt.Printf("Adding series with TVDB ID: %d\n", tvdbID)

		// Get quality profiles
		profiles, err := cmd.GetSonarrClient().GetQualityProfiles()
		if err != nil {
			return fmt.Errorf("failed to get quality profiles: %w", err)
		}

		// Get root folders
		rootFolders, err := cmd.GetSonarrClient().GetRootFolders()
		if err != nil {
			return fmt.Errorf("failed to get root folders: %w", err)
		}

		// Use first profile and first root folder
		if len(profiles) == 0 {
			return fmt.Errorf("no quality profiles configured in Sonarr")
		}
		if len(rootFolders) == 0 {
			return fmt.Errorf("no root folders configured in Sonarr")
		}

		qualityProfile := profiles[0]
		rootFolder := rootFolders[0]

		fmt.Printf("Using quality profile: %s\n", qualityProfile.Name)
		fmt.Printf("Using root folder: %s\n", rootFolder.Path)

		// Create a minimal series object for adding
		series := models.Series{
			TVDBID: tvdbID,
		}

		// Add the series
		addedSeries, err := cmd.GetSonarrClient().AddSeries(series, rootFolder, qualityProfile)
		if err != nil {
			return fmt.Errorf("failed to add series: %w", err)
		}

		fmt.Printf("✅ Successfully added %s (ID: %d)\n", addedSeries.Title, addedSeries.ID)
		return nil
	},
}

func init() {
	sonarrCmd.AddCommand(addCmd)
}
