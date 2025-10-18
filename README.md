# Sonarr-Sabnzbd CLI

A powerful command-line interface for managing Sonarr (TV shows) and Sabnzbd (downloads) with beautiful formatting, ASCII art, and shell completions.

## ✨ Features

- **🎬 Sonarr Integration**: Search, add, and manage TV series
- **📥 Sabnzbd Integration**: Monitor downloads with progress bars and rich formatting
- **🎨 ASCII Art**: Scalarr-style 8x8 colored ASCII art posters
- **🔧 Shell Completions**: Bash, Zsh, Fish, and PowerShell support
- **📊 JSON Output**: Perfect for scripting and automation
- **🚀 Rich UI**: Emojis, progress bars, and beautiful formatting

## 🚀 Installation

### Build from source
```bash
git clone <repository-url>
cd sonarr-cli
go build -o sonarr-sabnzbd-cli main.go
sudo mv sonarr-sabnzbd-cli /usr/local/bin/
```

### Shell Completions

#### Zsh
```bash
# Generate completion script
sonarr-sabnzbd-cli completion zsh > ~/.zsh/_sonarr-sabnzbd-cli

# Add to ~/.zshrc if not present
echo "autoload -Uz compinit && compinit" >> ~/.zshrc

# Reload
source ~/.zshrc
```

#### Bash
```bash
sonarr-sabnzbd-cli completion bash > /etc/bash_completion.d/sonarr-sabnzbd-cli
```

## 📖 Usage

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
# Check service status
sonarr-sabnzbd-cli status

# Generate shell completions
sonarr-sabnzbd-cli completion [bash|zsh|fish|powershell]
```

## ⚙️ Configuration

Create a config file at `~/.config/sonarr-sabnzbd-cli/config.yaml`:

```yaml
sonarr:
  host: "localhost"
  port: 8989
  apikey: "your-sonarr-api-key"
  timeout: 30

sabnzbd:
  host: "localhost"
  port: 8080
  apikey: "your-sabnzbd-api-key"
  username: ""  # optional
  password: ""  # optional
  timeout: 30
```

## 🎨 ASCII Art

Display beautiful 8x8 colored ASCII art posters just like Scalarr:

```bash
sonarr search "Breaking Bad" --ascii
sonarr series --ascii
```

## 📊 JSON Output

All commands support JSON output for scripting:

```bash
sonarr series --json | jq '.[] | select(.monitored) | .title'
sabnzbd queue --json | jq '.slots[0].percentage'
```

## 🛠️ Development

```bash
# Install dependencies
go mod download

# Run tests
go test ./...

# Build
go build -o sonarr-sabnzbd-cli main.go

# Lint
go vet ./...
```

## 📝 License

MIT License - see LICENSE file for details.

## 🤝 Contributing

1. Fork the repository
2. Create a feature branch
3. Make your changes
4. Add tests if applicable
5. Submit a pull request

## 🙏 Acknowledgments

- [Cobra](https://github.com/spf13/cobra) - CLI framework
- [Scalarr](https://github.com/zemmyang/scalarr) - Inspiration for ASCII art
- [Sonarr](https://sonarr.tv/) - TV show management
- [Sabnzbd](https://sabnzbd.org/) - Binary newsreader</content>
</xai:function_call">README.md