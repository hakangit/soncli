package sonarr

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"sonarr-sabnzbd-cli/cmd"
	"sonarr-sabnzbd-cli/internal/ascii"
)

// seriesCmd represents the series command
var seriesCmd = &cobra.Command{
	Use:   "series",
	Short: "List all series in your library",
	Long: `Display all TV series currently in your Sonarr library.

Examples:
   sonarr series                    # List all series
   sonarr series --ascii            # List with ASCII art posters
   sonarr series | grep "Breaking"  # Filter for specific series`,
	RunE: func(command *cobra.Command, args []string) error {
		asciiOutput, _ := command.Flags().GetBool("ascii")
		jsonOutput, _ := command.Flags().GetBool("json")

		// Get all series from Sonarr
		series, err := cmd.GetSonarrClient().GetSeries()
		if err != nil {
			return fmt.Errorf("failed to get series: %w", err)
		}

		if len(series) == 0 {
			if jsonOutput {
				fmt.Println("[]")
			} else {
				fmt.Println("No series found in your library.")
			}
			return nil
		}

		// JSON output mode
		if jsonOutput {
			return json.NewEncoder(os.Stdout).Encode(series)
		}

		fmt.Printf("Your Library (%d series):\n\n", len(series))

		asciiConfig := ascii.DefaultConfig()
		asciiConfig.Width = 30 // Smaller width for series list

		for i, s := range series {
			status := "✓"
			if !s.Monitored {
				status = "○"
			}
			fmt.Printf("%d. %s %s (%d) - %s\n",
				i+1, status, s.Title, s.Year, s.Status)

			// Add ASCII art if requested
			if asciiOutput {
				asciiArt, err := ascii.GetSeriesPosterASCII(s, asciiConfig)
				if err != nil {
					fmt.Printf("   [Could not load poster: %v]\n", err)
				} else {
					fmt.Printf("   %s\n", asciiArt)
				}
				fmt.Println()
			}
		}

		return nil
	},
}

func init() {
	sonarrCmd.AddCommand(seriesCmd)
	seriesCmd.Flags().Bool("ascii", false, "Display ASCII art posters for series")
	seriesCmd.Flags().Bool("json", false, "Output results in JSON format")
}
