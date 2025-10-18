package sonarr

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"sonarr-sabnzbd-cli/cmd"
	"sonarr-sabnzbd-cli/internal/ascii"
	"sonarr-sabnzbd-cli/internal/models"
)

// searchCmd represents the search command
var searchCmd = &cobra.Command{
	Use:   "search <query>",
	Short: "Search for TV series to add to your library",
	Long: `Search TheTVDB for TV series and display results. Use --add <number> to add a specific result.

Examples:
  sonarr search "Breaking Bad"              # Display search results
  sonarr search "Breaking Bad" --add 1      # Add first result
  sonarr search "Breaking Bad" --add 3      # Add third result
  sonarr search "The Office" --json         # Output in JSON format
  sonarr search "Stranger Things" --ascii   # Display with ASCII art posters`,
	Args: cobra.ExactArgs(1),
	RunE: func(command *cobra.Command, args []string) error {
		query := args[0]
		addIndex, _ := command.Flags().GetInt("add")
		jsonOutput, _ := command.Flags().GetBool("json")
		asciiOutput, _ := command.Flags().GetBool("ascii")

		// Search for series
		results, err := cmd.GetSonarrClient().LookupSeries(query)
		if err != nil {
			return fmt.Errorf("failed to search series: %w", err)
		}

		if len(results) == 0 {
			if jsonOutput {
				fmt.Println("[]")
			} else {
				fmt.Println("No series found matching your query.")
			}
			return nil
		}

		// JSON output mode
		if jsonOutput {
			return outputJSON(results)
		}

		// Display results
		fmt.Printf("Found %d series:\n\n", len(results))

		asciiConfig := ascii.DefaultConfig()
		// Keep default 8x8 size for compact display

		for i, series := range results {
			status := "✓"
			if series.Status != "Continuing" {
				status = "○"
			}

			fmt.Printf("%d. %s %s (%d) - %s\n",
				i+1, status, series.Title, series.Year, series.Status)

			// Add ASCII art if requested
			if asciiOutput {
				asciiArt, err := ascii.GetSeriesPosterASCII(series, asciiConfig)
				if err != nil {
					fmt.Printf("   [Could not load poster: %v]\n", err)
				} else {
					fmt.Printf("   %s\n", asciiArt)
				}
				fmt.Println()
			}
		}

		// Auto-add result by index if requested
		if addIndex > 0 {
			if addIndex > len(results) {
				return fmt.Errorf("invalid series number %d (only %d results found)", addIndex, len(results))
			}
			return addSeries(results[addIndex-1]) // Convert to 0-based index
		}

		fmt.Printf("\nUse --add <number> to add a specific series, or run:\n")
		fmt.Printf("  sonarr add %d\n", results[0].TVDBID)

		return nil
	},
}

func init() {
	sonarrCmd.AddCommand(searchCmd)
	searchCmd.Flags().Int("add", 0, "Add the series at the specified number (1-based)")
	searchCmd.Flags().Bool("json", false, "Output results in JSON format")
	searchCmd.Flags().Bool("ascii", false, "Display ASCII art posters for search results")
}

// outputJSON outputs results in JSON format
func outputJSON(results []models.Series) error {
	return json.NewEncoder(os.Stdout).Encode(results)
}

// addSeries adds a series to Sonarr
func addSeries(series models.Series) error {
	fmt.Printf("Adding series: %s (%d)\n", series.Title, series.Year)

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
	if len(profiles) == 0 || len(rootFolders) == 0 {
		return fmt.Errorf("no quality profiles or root folders configured in Sonarr")
	}

	qualityProfile := profiles[0]
	rootFolder := rootFolders[0]

	fmt.Printf("Using quality profile: %s\n", qualityProfile.Name)
	fmt.Printf("Using root folder: %s\n", rootFolder.Path)

	// Add the series
	addedSeries, err := cmd.GetSonarrClient().AddSeries(series, rootFolder, qualityProfile)
	if err != nil {
		return fmt.Errorf("failed to add series: %w", err)
	}

	fmt.Printf("✅ Successfully added %s (ID: %d)\n", addedSeries.Title, addedSeries.ID)
	return nil
}
