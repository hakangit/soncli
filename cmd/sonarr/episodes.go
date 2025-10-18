package sonarr

import (
	"fmt"
	"strconv"

	"github.com/spf13/cobra"
	"sonarr-sabnzbd-cli/cmd"
)

// episodesCmd represents the episodes command
var episodesCmd = &cobra.Command{
	Use:   "episodes <series-id>",
	Short: "View episodes for a series",
	Long: `Display all episodes for a specific TV series.

The series ID can be found using the 'sonarr series' command.

Examples:
  sonarr episodes 123
  sonarr episodes 456 | grep -i "pilot"`,
	Args: cobra.ExactArgs(1),
	RunE: func(command *cobra.Command, args []string) error {
		seriesID, err := strconv.Atoi(args[0])
		if err != nil {
			return fmt.Errorf("invalid series ID: %s", args[0])
		}

		// Get episodes for the series
		episodes, err := cmd.GetSonarrClient().GetEpisodes(seriesID)
		if err != nil {
			return fmt.Errorf("failed to get episodes: %w", err)
		}

		if len(episodes) == 0 {
			fmt.Printf("No episodes found for series ID %d.\n", seriesID)
			return nil
		}

		fmt.Printf("Episodes for Series ID %d (%d episodes):\n\n", seriesID, len(episodes))

		for _, episode := range episodes {
			status := "Missing"
			if episode.HasFile {
				status = "Downloaded"
			} else if episode.Monitored {
				status = "Monitored"
			} else {
				status = "Unmonitored"
			}

			fmt.Printf("S%02dE%02d - %s\n", episode.SeasonNumber, episode.EpisodeNumber, episode.Title)
			fmt.Printf("   Status: %s | Air Date: %s\n", status, episode.AirDate)
			if episode.Overview != "" {
				// Truncate overview if too long
				overview := episode.Overview
				if len(overview) > 100 {
					overview = overview[:97] + "..."
				}
				fmt.Printf("   Overview: %s\n", overview)
			}
			fmt.Println()
		}

		return nil
	},
}

func init() {
	sonarrCmd.AddCommand(episodesCmd)
}
