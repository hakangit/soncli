package sabnzbd

import (
	"github.com/spf13/cobra"
	"sonarr-sabnzbd-cli/cmd"
)

// sabnzbdCmd represents the sabnzbd command
var sabnzbdCmd = &cobra.Command{
	Use:   "sabnzbd",
	Short: "Manage Sabnzbd (binary newsreader)",
	Long:  `Commands for managing your Sabnzbd instance including queue management, downloads, and history.`,
}

func init() {
	cmd.RootCmd().AddCommand(sabnzbdCmd)
}
