package shared

import (
	"fmt"

	"github.com/spf13/cobra"
	"sonarr-sabnzbd-cli/cmd"
)

// docsCmd represents the docs command
var docsCmd = &cobra.Command{
	Use:   "docs",
	Short: "Show comprehensive help and usage examples",
	Long:  `Display detailed help documentation with examples for all commands.`,
	Run: func(cmd *cobra.Command, args []string) {
		showComprehensiveHelp()
	},
}

func init() {
	cmd.RootCmd().AddCommand(docsCmd)
}

func showComprehensiveHelp() {
	fmt.Print(`# Sonarr-Sabnzbd CLI Tool

A unified command-line interface for managing Sonarr (TV show automation) and Sabnzbd (binary newsreader).

## Quick Start

` + "```" + `bash
# Check service status
sonarr-sabnzbd-cli status

# Search for a TV series
sonarr-sabnzbd-cli sonarr search "Breaking Bad"

# Add a series automatically
sonarr-sabnzbd-cli sonarr search "Breaking Bad" --add

# View your series library
sonarr-sabnzbd-cli sonarr series
` + "```" + `

## Configuration

The tool automatically creates a config file at ` + "`" + `~/.config/sonarr-sabnzbd-cli/config.yaml` + "`" + `:

` + "```" + `yaml
sonarr:
  host: "10.84.30.100"
  port: 8989
  api_key: "your-sonarr-api-key"
  timeout: "30s"

sabnzbd:
  host: "10.84.30.100"
  port: 8081
  api_key: "your-sabnzbd-api-key"
  username: ""
  password: ""
  timeout: "30s"

ui:
  colors: true
  max_results: 10
` + "```" + `

## Commands

### Global Commands

#### ` + "`" + `status` + "`" + `
Check connectivity and status of both Sonarr and Sabnzbd services.

` + "```" + `bash
sonarr-sabnzbd-cli status
# Output:
# Checking service status...
# Sonarr: ✅ Connected - Version 3.0.10.1567
# Sabnzbd: ✅ Connected - Version 4.3.2
` + "```" + `

### Sonarr Commands

#### ` + "`" + `sonarr search <query>` + "`" + `
Search for TV series in TheTVDB.

**Options:**
- ` + "`" + `--add` + "`" + `: Automatically add the first search result
- ` + "`" + `--json` + "`" + `: Output results in JSON format

**Examples:**
` + "```" + `bash
# Search for series
sonarr-sabnzbd-cli sonarr search "Breaking Bad"

# Add first result automatically
sonarr-sabnzbd-cli sonarr search "The Office" --add

# Get JSON output for scripting
sonarr-sabnzbd-cli sonarr search "Stranger Things" --json
` + "```" + `

#### ` + "`" + `sonarr series` + "`" + `
List all series in your library.

` + "```" + `bash
sonarr-sabnzbd-cli sonarr series
` + "```" + `

#### ` + "`" + `sonarr add <tvdb-id>` + "`" + `
Add a series to your library by TVDB ID.

` + "```" + `bash
sonarr-sabnzbd-cli sonarr add 81189
` + "```" + `

#### ` + "`" + `sonarr episodes <series-id>` + "`" + `
View episodes for a specific series.

` + "```" + `bash
sonarr-sabnzbd-cli sonarr episodes 123
` + "```" + `

#### ` + "`" + `sonarr monitor <series-id>` + "`" + `
Toggle monitoring for a series.

` + "```" + `bash
sonarr-sabnzbd-cli sonarr monitor 123
` + "```" + `

#### ` + "`" + `sonarr import <path>` + "`" + `
Import downloaded episodes from a directory.

` + "```" + `bash
sonarr-sabnzbd-cli sonarr import "/path/to/downloads"
` + "```" + `

#### ` + "`" + `sonarr profiles` + "`" + `
List available quality profiles.

` + "```" + `bash
sonarr-sabnzbd-cli sonarr profiles
` + "```" + `

#### ` + "`" + `sonarr root-folders` + "`" + `
List configured root folders.

` + "```" + `bash
sonarr-sabnzbd-cli sonarr root-folders
` + "```" + `

### Sabnzbd Commands

#### ` + "`" + `sabnzbd queue` + "`" + `
View current download queue.

` + "```" + `bash
sonarr-sabnzbd-cli sabnzbd queue
` + "```" + `

#### ` + "`" + `sabnzbd history` + "`" + `
View download history.

` + "```" + `bash
sonarr-sabnzbd-cli sabnzbd history
` + "```" + `

#### ` + "`" + `sabnzbd add <url>` + "`" + `
Add an NZB file by URL.

` + "```" + `bash
sonarr-sabnzbd-cli sabnzbd add "https://example.com/file.nzb"
` + "```" + `

#### ` + "`" + `sabnzbd pause` + "`" + `
Pause all downloads.

` + "```" + `bash
sonarr-sabnzbd-cli sabnzbd pause
` + "```" + `

#### ` + "`" + `sabnzbd resume` + "`" + `
Resume all downloads.

` + "```" + `bash
sonarr-sabnzbd-cli sabnzbd resume
` + "```" + `

#### ` + "`" + `sabnzbd speed <limit>` + "`" + `
Set download speed limit (e.g., "50", "100K", "2M").

` + "```" + `bash
sonarr-sabnzbd-cli sabnzbd speed 100K
` + "```" + `

#### ` + "`" + `sabnzbd categories` + "`" + `
List available categories.

` + "```" + `bash
sonarr-sabnzbd-cli sabnzbd categories
` + "```" + `

## Workflow Examples

### Adding a New Series
` + "```" + `bash
# 1. Search for the series
sonarr-sabnzbd-cli sonarr search "Breaking Bad"

# 2. Note the TVDB ID from the output (e.g., 81189)

# 3. Add the series
sonarr-sabnzbd-cli sonarr add 81189
` + "```" + `

### Automated Series Addition
` + "```" + `bash
# Search and add in one command
sonarr-sabnzbd-cli sonarr search "The Mandalorian" --add
` + "```" + `

### Download Management
` + "```" + `bash
# Check current downloads
sonarr-sabnzbd-cli sabnzbd queue

# Add a new download
sonarr-sabnzbd-cli sabnzbd add "https://nzb.site/download.nzb"

# Pause downloads temporarily
sonarr-sabnzbd-cli sabnzbd pause

# Resume when ready
sonarr-sabnzbd-cli sabnzbd resume
` + "```" + `

### System Monitoring
` + "```" + `bash
# Check both services
sonarr-sabnzbd-cli status

# View series library
sonarr-sabnzbd-cli sonarr series

# Check download history
sonarr-sabnzbd-cli sabnzbd history
` + "```" + `

## Scripting Examples

### JSON Output for Scripts
` + "```" + `bash
# Get search results as JSON
RESULTS=$(sonarr-sabnzbd-cli sonarr search "Game of Thrones" --json)

# Extract first result ID
TVDB_ID=$(echo $RESULTS | jq '.[0].tvdbId')

# Add the series
sonarr-sabnzbd-cli sonarr add $TVDB_ID
` + "```" + `

### Batch Operations
` + "```" + `bash
# Add multiple series
for series in "Breaking Bad" "The Office" "Stranger Things"; do
  sonarr-sabnzbd-cli sonarr search "$series" --add
done
` + "```" + `

## Troubleshooting

### Common Issues

**"API Key Required"**
- Check that your API keys are correctly set in the config file
- Verify the keys work in the web interfaces

**"Connection refused"**
- Ensure Sonarr/Sabnzbd are running
- Check the host and port settings in config
- Verify network connectivity

**"No series found"**
- Try different search terms
- Check if the series exists in TheTVDB

### Getting API Keys

**Sonarr:**
1. Open Sonarr web interface
2. Go to Settings → General
3. Copy the "API Key" field

**Sabnzbd:**
1. Open Sabnzbd web interface
2. Go to Config → General
3. Copy the "API Key" field

## Development

This tool is built with Go and uses the following libraries:
- Cobra: CLI framework
- Viper: Configuration management
- Sonarr/Sabnzbd REST APIs

## License

GPL-3.0

---

For more information, visit the [GitHub repository](https://github.com/username/sonarr-sabnzbd-cli) or run ` + "`" + `sonarr-sabnzbd-cli --help` + "`" + `.
`)
}
