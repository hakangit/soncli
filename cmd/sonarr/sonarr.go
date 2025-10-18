package sonarr

import (
	"github.com/spf13/cobra"
	"sonarr-sabnzbd-cli/cmd"
)

// sonarrCmd represents the sonarr command
var sonarrCmd = &cobra.Command{
	Use:   "sonarr",
	Short: "Manage Sonarr (TV show automation)",
	Long:  `Commands for managing your Sonarr instance including searching, adding, and monitoring TV series.`,
}

func init() {
	cmd.RootCmd().AddCommand(sonarrCmd)
}
