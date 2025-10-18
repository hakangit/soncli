package sabnzbd

import (
	"fmt"

	"github.com/spf13/cobra"
	"sonarr-sabnzbd-cli/cmd"
)

// categoriesCmd represents the categories command
var categoriesCmd = &cobra.Command{
	Use:   "categories",
	Short: "List available categories",
	Long: `Display all available categories configured in Sabnzbd.

Examples:
  sabnzbd categories`,
	RunE: func(command *cobra.Command, args []string) error {
		// Get categories from Sabnzbd
		categories, err := cmd.GetSabnzbdClient().GetCategories()
		if err != nil {
			return fmt.Errorf("failed to get categories: %w", err)
		}

		if len(categories) == 0 {
			fmt.Println("No categories configured.")
			return nil
		}

		fmt.Printf("Available Categories (%d):\n\n", len(categories))

		for i, category := range categories {
			fmt.Printf("%d. %s\n", i+1, category)
		}

		return nil
	},
}

func init() {
	sabnzbdCmd.AddCommand(categoriesCmd)
}
