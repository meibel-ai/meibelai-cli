# Meibel CLI

A command-line interface for the [Meibel API](https://docs.meibel.ai), generated from the OpenAPI specification using [Forge](https://github.com/meibel-ai/forge).

## Installation

### Homebrew (macOS/Linux)

```bash
brew install meibel-ai/tap/meibel
```

### Go Install

```bash
go install github.com/meibel-ai/meibel-cli@latest
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
# Configure authentication
meibel config set api_key "your-api-key"
meibel config set base_url "https://api.meibel.ai/v2"

# Parse a document
meibel documents parse-document --file document.pdf

# Process a document synchronously
meibel documents process-document --file document.pdf

# List datasources
meibel datasources list-datasources --json

# Get help
meibel --help
meibel documents --help
```

## Configuration

Configuration is stored in `~/.meibel/config.yaml`.

```bash
# Set values
meibel config set api_key "your-key"
meibel config set base_url "https://api.meibel.ai/v2"

# View config
meibel config
```

Or use environment variables:

```bash
export MEIBEL_API_KEY="your-api-key"
export MEIBEL_BASE_URL="https://api.meibel.ai/v2"
```

## Output Formats

```bash
# JSON output (default)
meibel datasources list-datasources --json

# Table output
meibel datasources list-datasources
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
- [SDK Guides](https://docs.meibel.ai/sdk/go)

## License

MIT
