package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"sonarr-sabnzbd-cli/internal/api/sabnzbd"
	"sonarr-sabnzbd-cli/internal/api/sonarr"
	"sonarr-sabnzbd-cli/internal/config"
	"sonarr-sabnzbd-cli/internal/models"
)

var (
	cfg           *models.Config
	sonarrClient  *sonarr.Client
	sabnzbdClient *sabnzbd.Client
)

var rootCmd = &cobra.Command{
	Use:   "sonarr-sabnzbd-cli",
	Short: "CLI tool for managing Sonarr and Sabnzbd",
	Long: `A command-line interface for managing Sonarr (TV show automation)
and Sabnzbd (binary newsreader) with unified commands and interactive features.`,
	PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
		var err error
		cfg, err = config.LoadConfig()
		if err != nil {
			return fmt.Errorf("failed to load config: %w", err)
		}

		// Initialize API clients
		sonarrClient = sonarr.NewClient(
			cfg.Sonarr.Host,
			cfg.Sonarr.Port,
			cfg.Sonarr.APIKey,
			cfg.Sonarr.Timeout,
		)

		sabnzbdClient = sabnzbd.NewClient(
			cfg.Sabnzbd.Host,
			cfg.Sabnzbd.Port,
			cfg.Sabnzbd.APIKey,
			cfg.Sabnzbd.Username,
			cfg.Sabnzbd.Password,
			cfg.Sabnzbd.Timeout,
		)

		return nil
	},
}

func Execute() error {
	return rootCmd.Execute()
}

// GetConfig returns the loaded configuration
func GetConfig() *models.Config {
	return cfg
}

// GetSonarrClient returns the Sonarr API client
func GetSonarrClient() *sonarr.Client {
	return sonarrClient
}

// GetSabnzbdClient returns the Sabnzbd API client
func GetSabnzbdClient() *sabnzbd.Client {
	return sabnzbdClient
}

// RootCmd returns the root command
func RootCmd() *cobra.Command {
	return rootCmd
}

// statusCmd represents the status command
var statusCmd = &cobra.Command{
	Use:   "status",
	Short: "Check the status of Sonarr and Sabnzbd services",
	Long:  `Check the connectivity and status of both Sonarr and Sabnzbd services.`,
	RunE: func(command *cobra.Command, args []string) error {
		fmt.Println("Checking service status...")

		// Check Sonarr status
		fmt.Print("Sonarr: ")
		sonarrStatus, err := sonarrClient.GetSystemStatus()
		if err != nil {
			fmt.Printf("❌ Error - %v\n", err)
		} else {
			fmt.Printf("✅ Connected - Version %s\n", sonarrStatus.Version)
		}

		// Check Sabnzbd status
		fmt.Print("Sabnzbd: ")
		sabnzbdVersion, err := sabnzbdClient.GetVersion()
		if err != nil {
			fmt.Printf("❌ Error - %v\n", err)
		} else {
			fmt.Printf("✅ Connected - Version %s\n", sabnzbdVersion)
		}

		return nil
	},
}

// completionCmd represents the completion command
var completionCmd = &cobra.Command{
	Use:   "completion [bash|zsh|fish|powershell]",
	Short: "Generate completion script",
	Long: `To load completions:

Bash:

  $ source <(sonarr-sabnzbd-cli completion bash)

  # To load completions for each session, execute once:
  # Linux:
  $ sonarr-sabnzbd-cli completion bash > /etc/bash_completion.d/sonarr-sabnzbd-cli
  # macOS:
  $ sonarr-sabnzbd-cli completion bash > /usr/local/etc/bash_completion.d/sonarr-sabnzbd-cli

Zsh:

  # If shell completion is not already enabled in your environment,
  # you will need to enable it.  You can execute the following once:

  $ echo "autoload -U compinit; compinit" >> ~/.zshrc

  # To load completions for each session, execute once:
  $ sonarr-sabnzbd-cli completion zsh > "${fpath[1]}/_sonarr-sabnzbd-cli"

  # You will need to start a new shell for this setup to take effect.

fish:

  $ sonarr-sabnzbd-cli completion fish | source

  # To load completions for each session, execute once:
  $ sonarr-sabnzbd-cli completion fish > ~/.config/fish/completions/sonarr-sabnzbd-cli.fish

PowerShell:

  PS> sonarr-sabnzbd-cli completion powershell | Out-String | Invoke-Expression

  # To load completions for every new session, run:
  PS> sonarr-sabnzbd-cli completion powershell > sonarr-sabnzbd-cli.ps1
  # and source this file from your PowerShell profile.
`,
	DisableFlagsInUseLine: true,
	ValidArgs:             []string{"bash", "zsh", "fish", "powershell"},
	Args:                  cobra.ExactValidArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		switch args[0] {
		case "bash":
			cmd.Root().GenBashCompletion(cmd.OutOrStdout())
		case "zsh":
			cmd.Root().GenZshCompletion(cmd.OutOrStdout())
		case "fish":
			cmd.Root().GenFishCompletion(cmd.OutOrStdout(), true)
		case "powershell":
			cmd.Root().GenPowerShellCompletionWithDesc(cmd.OutOrStdout())
		}
	},
}

func init() {
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	rootCmd.AddCommand(statusCmd)
	rootCmd.AddCommand(completionCmd)
}
