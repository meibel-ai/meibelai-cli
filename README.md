# Meibel CLI

A command-line interface for the [Meibel API](https://docs.meibel.ai).

## Installation

### Homebrew (macOS/Linux)

```bash
brew install meibel-ai/tap/meibel
```

### Go Install

```bash
go install github.com/meibel-ai/meibelai-cli@latest
```

### Download Binaries

Download from the [releases page](https://github.com/meibel-ai/meibelai-cli/releases).

### Linux Packages

```bash
# Debian/Ubuntu
sudo dpkg -i meibel_*_linux_amd64.deb

# RedHat/CentOS/Fedora
sudo rpm -i meibel_*_linux_amd64.rpm
```

## Quick Start

```bash
# Configure authentication (interactive wizard)
meibel config init

# Parse a document
meibel documents parse --file document.pdf

# Parse and wait for result
meibel documents parse --file document.pdf --wait

# List datasources
meibel datasources list

# Get help
meibel --help
meibel documents --help
```

## Authentication

The CLI reads your API key from the config file or environment variables.

### Interactive Setup

```bash
meibel config init
```

This prompts for your API key and environment (production, dev, local, or custom URL).

### Environment Variables

```bash
export MEIBEL_API_KEY="your-api-key"
export MEIBEL_BASE_URL="https://api.meibel.ai/v2"   # optional
export MEIBEL_TOKEN="your-bearer-token"              # alternative to API key
```

### Config File

Configuration is stored in `~/.meibel/config.yaml`:

```yaml
api_key: your-api-key
base_url: https://api.meibel.ai/v2
```

View your current config:

```bash
meibel config show
```

## Output Formats

```bash
# Table output (default)
meibel datasources list

# JSON output
meibel datasources list --json
```

## Shell Completions

```bash
# Bash
meibel completion bash > /etc/bash_completion.d/meibel

# Zsh
meibel completion zsh > "${fpath[1]}/_meibel"

# Fish
meibel completion fish > ~/.config/fish/completions/meibel.fish
```

## Documentation

- [API Reference](https://docs.meibel.ai/api-reference/overview)
- [CLI Guide](https://docs.meibel.ai/sdk/cli)

## License

MIT
