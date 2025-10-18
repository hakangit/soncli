# Soncli

[![Go](https://img.shields.io/badge/Go-1.21+-blue.svg)](https://golang.org)
[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://opensource.org/licenses/MIT)

A command-line interface for managing Sonarr (TV shows) and Sabnzbd (downloads).

## Features

- Sonarr integration: Search, add, and manage TV series
- Sabnzbd integration: Monitor downloads with progress bars
- ASCII art: Display 8x8 colored ASCII art posters for TV shows
- Shell completions: Support for Bash, Zsh, Fish, and PowerShell
- JSON output: Structured output for scripting and automation
- Interactive setup: Configuration wizard for first-time setup

## Installation

### Quick Setup

1. Clone and build:
```bash
git clone https://github.com/hakangit/soncli.git
cd soncli
go build -o soncli main.go
sudo mv soncli /usr/local/bin/
```

2. Run setup:
```bash
soncli setup
```

The setup wizard will guide you through configuring your connections.

### Manual Installation

If you prefer manual configuration, build from source and create the config file manually.

### Shell Completions

#### Zsh
```bash
soncli completion zsh > ~/.zsh/_soncli
echo "autoload -Uz compinit && compinit" >> ~/.zshrc
source ~/.zshrc
```

#### Bash
```bash
soncli completion bash > /etc/bash_completion.d/soncli
```

## Usage

### Sonarr Commands

```bash
# Search for TV series
sonarr search "Breaking Bad"

# Add a specific series by number
sonarr search "Breaking Bad" --add 1

# View your library with ASCII art
sonarr series --ascii

# Get series information
sonarr info
```

### Sabnzbd Commands

```bash
# View download queue with progress bars
sabnzbd queue

# View download history
sabnzbd history

# Get system information
sabnzbd info

# Add NZB by URL
sabnzbd add "https://example.com/file.nzb"

# Control downloads
sabnzbd pause
sabnzbd resume
sabnzbd speed 50
```

### General

```bash
# Interactive setup wizard (first-time setup)
soncli setup

# Check service status
soncli status

# Generate shell completions
soncli completion [bash|zsh|fish|powershell]
```

## Configuration

### Setup

Run the setup wizard:
```bash
soncli setup
```

This will guide you through configuring your connections.

### Manual Configuration

If you prefer manual setup, create a config file at `~/.config/soncli/config.yaml`:

```yaml
sonarr:
  host: "localhost"          # Your Sonarr server IP/hostname
  port: 8989                 # Default Sonarr port
  apikey: "your-api-key-here"  # Get from Sonarr Settings > General > API Key
  timeout: 30

sabnzbd:
  host: "localhost"          # Your Sabnzbd server IP/hostname
  port: 8080                 # Default Sabnzbd port
  apikey: "your-api-key-here"  # Get from Sabnzbd Config > General > API Key
  username: ""               # Optional: only if authentication is enabled
  password: ""               # Optional: only if authentication is enabled
  timeout: 30

# UI preferences
ui:
  colors: true               # Enable colored output
  max_results: 10            # Maximum search results to show
```

## ASCII Art

Display 8x8 colored ASCII art posters:

```bash
sonarr search "Breaking Bad" --ascii
sonarr series --ascii
```

## JSON Output

Commands support JSON output for scripting:

```bash
sonarr series --json | jq '.[] | select(.monitored) | .title'
sabnzbd queue --json | jq '.slots[0].percentage'
```

## Development

```bash
# Install dependencies
go mod download

# Run tests
go test ./...

# Build
go build -o soncli main.go

# Lint
go vet ./...
```

## License

MIT License - see LICENSE file for details.

## Contributing

1. Fork the repository
2. Create a feature branch
3. Make your changes
4. Add tests if applicable
5. Submit a pull request

## Acknowledgments

- [Cobra](https://github.com/spf13/cobra) - CLI framework
- [Scalarr](https://github.com/zemmyang/scalarr) - Inspiration for ASCII art
- [Sonarr](https://sonarr.tv/) - TV show management
- [Sabnzbd](https://sabnzbd.org/) - Binary newsreader
