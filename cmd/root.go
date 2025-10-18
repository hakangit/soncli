package cmd

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"

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
// runSetupWizard runs an interactive setup wizard
func runSetupWizard() error {
	fmt.Println("🚀 Sonarr-Sabnzbd CLI Setup Wizard")
	fmt.Println("===================================")
	fmt.Println()

	reader := bufio.NewReader(os.Stdin)

	// Sonarr setup
	fmt.Println("📺 Sonarr Configuration")
	fmt.Println("----------------------")

	sonarrHost := promptWithDefault(reader, "Sonarr Host", "localhost")
	sonarrPortStr := promptWithDefault(reader, "Sonarr Port", "8989")
	sonarrAPIKey := promptRequired(reader, "Sonarr API Key (get from Settings > General > API Key)")

	sonarrPort, err := strconv.Atoi(sonarrPortStr)
	if err != nil {
		return fmt.Errorf("invalid port number: %w", err)
	}

	// Sabnzbd setup
	fmt.Println()
	fmt.Println("📥 Sabnzbd Configuration")
	fmt.Println("-----------------------")

	sabnzbdHost := promptWithDefault(reader, "Sabnzbd Host", "localhost")
	sabnzbdPortStr := promptWithDefault(reader, "Sabnzbd Port", "8080")
	sabnzbdAPIKey := promptRequired(reader, "Sabnzbd API Key (get from Config > General > API Key)")

	sabnzbdPort, err := strconv.Atoi(sabnzbdPortStr)
	if err != nil {
		return fmt.Errorf("invalid port number: %w", err)
	}

	// Optional authentication for Sabnzbd
	sabnzbdUsername := promptOptional(reader, "Sabnzbd Username (leave empty if no auth)")
	sabnzbdPassword := ""
	if sabnzbdUsername != "" {
		sabnzbdPassword = promptOptional(reader, "Sabnzbd Password")
	}

	// Create config
	cfg := &models.Config{
		Sonarr: models.SonarrConfig{
			Host:    sonarrHost,
			Port:    sonarrPort,
			APIKey:  sonarrAPIKey,
			Timeout: 30 * time.Second,
		},
		Sabnzbd: models.SabnzbdConfig{
			Host:     sabnzbdHost,
			Port:     sabnzbdPort,
			APIKey:   sabnzbdAPIKey,
			Username: sabnzbdUsername,
			Password: sabnzbdPassword,
			Timeout:  30 * time.Second,
		},
		UI: models.UIConfig{
			Colors:     true,
			MaxResults: 10,
		},
	}

	// Test connections
	fmt.Println()
	fmt.Println("🔍 Testing Connections...")
	fmt.Println("------------------------")

	// Test Sonarr
	fmt.Print("📺 Testing Sonarr connection... ")
	sonarrClient := sonarr.NewClient(sonarrHost, sonarrPort, sonarrAPIKey, 30*time.Second)
	if _, err := sonarrClient.GetSystemStatus(); err != nil {
		fmt.Printf("❌ Failed: %v\n", err)
		fmt.Println("⚠️  Configuration saved but Sonarr connection failed. Check your settings.")
	} else {
		fmt.Println("✅ Success!")
	}

	// Test Sabnzbd
	fmt.Print("📥 Testing Sabnzbd connection... ")
	sabnzbdClient := sabnzbd.NewClient(sabnzbdHost, sabnzbdPort, sabnzbdAPIKey, sabnzbdUsername, sabnzbdPassword, 30*time.Second)
	if _, err := sabnzbdClient.GetVersion(); err != nil {
		fmt.Printf("❌ Failed: %v\n", err)
		fmt.Println("⚠️  Configuration saved but Sabnzbd connection failed. Check your settings.")
	} else {
		fmt.Println("✅ Success!")
	}

	// Save config
	fmt.Println()
	fmt.Print("💾 Saving configuration... ")
	if err := config.SaveConfig(cfg); err != nil {
		return fmt.Errorf("failed to save config: %w", err)
	}
	fmt.Println("✅ Saved!")

	fmt.Println()
	fmt.Println("🎉 Setup complete! You can now use the CLI:")
	fmt.Println("   sonarr-sabnzbd-cli status    # Check service status")
	fmt.Println("   sonarr search \"Breaking Bad\" --add 1    # Search and add shows")
	fmt.Println("   sabnzbd queue               # View download queue")

	return nil
}

// promptWithDefault prompts user with a default value
func promptWithDefault(reader *bufio.Reader, prompt, defaultValue string) string {
	fmt.Printf("%s [%s]: ", prompt, defaultValue)
	input, _ := reader.ReadString('\n')
	input = strings.TrimSpace(input)
	if input == "" {
		return defaultValue
	}
	return input
}

// promptRequired prompts user for required input
func promptRequired(reader *bufio.Reader, prompt string) string {
	for {
		fmt.Printf("%s: ", prompt)
		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)
		if input != "" {
			return input
		}
		fmt.Println("This field is required. Please enter a value.")
	}
}

// promptOptional prompts user for optional input
func promptOptional(reader *bufio.Reader, prompt string) string {
	fmt.Printf("%s (optional): ", prompt)
	input, _ := reader.ReadString('\n')
	return strings.TrimSpace(input)
}

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

// setupCmd represents the setup command
var setupCmd = &cobra.Command{
	Use:   "setup",
	Short: "Interactive setup wizard",
	Long: `Run an interactive setup wizard to configure Sonarr and Sabnzbd connections.

This command will guide you through:
- Configuring your Sonarr server connection
- Configuring your Sabnzbd server connection
- Testing the connections
- Saving the configuration

Examples:
  sonarr-sabnzbd-cli setup`,
	RunE: func(cmd *cobra.Command, args []string) error {
		return runSetupWizard()
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
	rootCmd.AddCommand(setupCmd)
	rootCmd.AddCommand(completionCmd)
}
