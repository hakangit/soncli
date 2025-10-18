package main

import (
	"os"

	"sonarr-sabnzbd-cli/cmd"
	_ "sonarr-sabnzbd-cli/cmd/sabnzbd" // Import for side effects (init functions)
	_ "sonarr-sabnzbd-cli/cmd/shared"  // Import for side effects (init functions)
	_ "sonarr-sabnzbd-cli/cmd/sonarr"  // Import for side effects (init functions)
)

func main() {
	if err := cmd.Execute(); err != nil {
		os.Exit(1)
	}
}
