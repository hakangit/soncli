package sonarr

import (
	"fmt"
	"strconv"

	"github.com/spf13/cobra"
	"sonarr-sabnzbd-cli/cmd"
)

var (
	monitorState bool
)

// monitorCmd represents the monitor command
var monitorCmd = &cobra.Command{
	Use:   "monitor <series-id>",
	Short: "Toggle monitoring for series",
	Long: `Enable or disable monitoring for a TV series.

When monitoring is enabled, Sonarr will automatically search for and download
new episodes of the series. When disabled, Sonarr will ignore new episodes.

Examples:
  sonarr monitor 123 --enable    # Start monitoring series 123
  sonarr monitor 123 --disable   # Stop monitoring series 123`,
	Args: cobra.ExactArgs(1),
	RunE: func(command *cobra.Command, args []string) error {
		seriesID, err := strconv.Atoi(args[0])
		if err != nil {
			return fmt.Errorf("invalid series ID: %s", args[0])
		}

		// Get the current series
		series, err := cmd.GetSonarrClient().GetSeriesByID(seriesID)
		if err != nil {
			return fmt.Errorf("failed to get series: %w", err)
		}

		// Update monitoring state
		series.Monitored = monitorState

		// Update the series
		updatedSeries, err := cmd.GetSonarrClient().UpdateSeries(*series)
		if err != nil {
			return fmt.Errorf("failed to update series: %w", err)
		}

		action := "disabled"
		if monitorState {
			action = "enabled"
		}

		fmt.Printf("✅ Successfully %s monitoring for '%s'\n", action, updatedSeries.Title)
		return nil
	},
}

func init() {
	sonarrCmd.AddCommand(monitorCmd)
	monitorCmd.Flags().BoolVar(&monitorState, "enable", false, "Enable monitoring for the series")
	monitorCmd.Flags().BoolVar(&monitorState, "disable", true, "Disable monitoring for the series")
	monitorCmd.MarkFlagsMutuallyExclusive("enable", "disable")
}
