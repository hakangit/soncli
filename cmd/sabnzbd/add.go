package sabnzbd

import (
	"fmt"

	"github.com/spf13/cobra"
	"sonarr-sabnzbd-cli/cmd"
)

var (
	addCategory string
)

// addCmd represents the add command
var addCmd = &cobra.Command{
	Use:   "add <nzb-url>",
	Short: "Add NZB by URL",
	Long: `Add an NZB file to the download queue by providing its URL.

Examples:
  sabnzbd add "https://example.com/file.nzb"
  sabnzbd add "https://example.com/file.nzb" --category tv
  sabnzbd add "https://example.com/file.nzb" --category movies`,
	Args: cobra.ExactArgs(1),
	RunE: func(command *cobra.Command, args []string) error {
		nzbURL := args[0]

		// Add NZB to Sabnzbd
		nzoIDs, err := cmd.GetSabnzbdClient().AddNZB(nzbURL, addCategory)
		if err != nil {
			return fmt.Errorf("failed to add NZB: %w", err)
		}

		fmt.Printf("✅ Successfully added NZB to queue\n")
		fmt.Printf("NZB IDs: %v\n", nzoIDs)

		if addCategory != "" {
			fmt.Printf("Category: %s\n", addCategory)
		}

		return nil
	},
}

func init() {
	sabnzbdCmd.AddCommand(addCmd)
	addCmd.Flags().StringVarP(&addCategory, "category", "c", "", "Category for the download")
}
