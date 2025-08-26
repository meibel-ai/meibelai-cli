# Meibel AI CLI

A command-line interface for interacting with the Meibel AI API, automatically generated from the OpenAPI specification.

## Features

- 🚀 Auto-generated commands from OpenAPI spec
- 🔐 Multiple authentication methods support
- 📊 Multiple output formats (JSON, YAML, Table)
- 🎨 Colorized JSON output
- ⚙️ Configuration profiles
- 🔍 Dry-run mode for previewing requests
- 💾 Request templates and interactive mode

## Installation

### Using Homebrew (macOS/Linux)

```bash
brew tap meibel-ai/tap
brew install meibel-cli
```

### Using Go Install

```bash
go install github.com/meibel-ai/meibel-cli@latest
```

### Download Pre-built Binaries

Download the latest release from the [releases page](https://github.com/meibel-ai/meibel-cli/releases).

### Build from Source

```bash
git clone https://github.com/meibel-ai/meibel-cli.git
cd meibel-cli
make build  # Generates commands and builds the binary
```

## Quick Start

### 1. Set up authentication

```bash
# Interactive mode
meibel auth login --interactive

# Or provide API key directly
meibel auth login --api-key "your-api-key"

# Check authentication status
meibel auth status
```

### 2. Explore available commands

```bash
# List all available commands
meibel --help

# List commands for a specific resource
meibel datasources --help
```

### 3. Common operations

```bash
# List datasources
meibel datasources list --output table

# Get a specific datasource
meibel datasources get <datasource-id> --output yaml

# Create a new datasource
meibel datasources create --data '{
  "name": "My Datasource",
  "type": "database",
  "config": {...}
}'

# Or from a file
meibel datasources create --from-file datasource.json

# Update a datasource
meibel datasources update <datasource-id> --data '{"name": "Updated Name"}'

# Delete a datasource
meibel datasources delete <datasource-id>
```

## Configuration

The CLI supports multiple configuration methods:

### Configuration file

Create `~/.meibel.yaml`:

```yaml
api_key: your-api-key
server: http://api.meibel.ai
output: json

profiles:
  staging:
    server: http://staging.api.meibel.ai
    api_key: staging-api-key
  
  development:
    server: http://localhost:8000
    api_key: dev-api-key
```

### Environment variables

```bash
export MEIBEL_API_KEY="your-api-key"
export MEIBEL_SERVER="http://api.meibel.ai"
export MEIBEL_OUTPUT="table"
```

### Command-line flags

```bash
meibel --api-key "your-api-key" --server "http://api.meibel.ai" datasources list
```

## Profiles

Use different configuration profiles:

```bash
# Use staging profile
meibel --profile staging datasources list

# Use development profile
meibel --profile development datasources list
```

## Output Formats

The CLI supports multiple output formats:

```bash
# JSON (default, with syntax highlighting)
meibel datasources list --output json

# YAML
meibel datasources list --output yaml

# Table (for list operations)
meibel datasources list --output table
```

## Advanced Usage

### Dry-run mode

Preview requests without executing them:

```bash
meibel --dry-run datasources create --data '{"name": "Test"}'
```

### Data Elements operations

```bash
# Add data elements
meibel data-elements add --datasource-id <id> --data '[
  {"key": "value1"},
  {"key": "value2"}
]'

# Retrieve data elements
meibel data-elements retrieve --datasource-id <id> --limit 10

# Update data elements
meibel data-elements update --element-id <id> --data '{"key": "new-value"}'

# Delete data elements
meibel data-elements delete --element-id <id>
```

### Blueprint Instance management

```bash
# List blueprint instances
meibel blueprint-instances list --output table

# Get instance status
meibel blueprint-instances get <instance-id>

# Update instance status
meibel blueprint-instances complete <instance-id>
meibel blueprint-instances fail <instance-id>
```

### TAG Configuration

```bash
# Configure tagging
meibel tag configure --data '{
  "datasource_id": "ds-123",
  "config": {...}
}'
```

### RAG Configuration

```bash
# Configure RAG settings
meibel rag configure --data '{
  "datasource_id": "ds-123",
  "strategy": "semantic",
  "parameters": {...}
}'
```

## Examples

### Working with pagination

```bash
# List with pagination
meibel datasources list --limit 20 --offset 0
meibel datasources list --limit 20 --offset 20
```

### Complex data operations

```bash
# Create datasource from file
cat > datasource.json << EOF
{
  "name": "Production Database",
  "type": "postgresql",
  "config": {
    "host": "db.example.com",
    "port": 5432,
    "database": "production"
  }
}
EOF

meibel datasources create --from-file datasource.json
```

### Batch operations

```bash
# Delete multiple data elements
for id in elem1 elem2 elem3; do
  meibel data-elements delete --element-id "$id"
done
```

## Development

### Building from Source

```bash
# Install dependencies
go mod download

# Generate commands from OpenAPI spec
make generate

# Build the CLI (includes generation)
make build

# Build without regenerating (faster for development)
make dev

# Run tests
go test ./...
```

### Understanding the Architecture

The Meibel CLI uses build-time code generation for optimal performance:

1. **Build-time Generation**: Commands are generated at build time from the OpenAPI spec, resulting in instant startup (<100ms)
2. **Hybrid Approach**: Supports both generated commands (from OpenAPI) and custom commands (auth, config, version)
3. **Smart Command Naming**: Operations are intelligently mapped to CLI commands using RESTful conventions

### Adding Custom Commands

To add custom commands that aren't generated from the OpenAPI spec:

1. Create a new file in the `cmd/` directory (e.g., `cmd/mycommand.go`)
2. Define your command using Cobra
3. Register it in `cmd/root.go` in the `RegisterCustomCommands` function

Example:
```go
// cmd/mycommand.go
package cmd

import (
    "github.com/spf13/cobra"
)

var myCmd = &cobra.Command{
    Use:   "mycommand",
    Short: "Description of my command",
    RunE: func(cmd *cobra.Command, args []string) error {
        // Command implementation
        return nil
    },
}

// In cmd/root.go, add to RegisterCustomCommands:
rootCmd.AddCommand(myCmd)
```

### Updating the OpenAPI Spec

The CLI generates commands from the OpenAPI spec at:
`https://storage.googleapis.com/meibel-api-docs/api.json`

To regenerate commands after the spec changes:

```bash
make generate
```

### Build Process

The Makefile provides several targets:

- `make generate` - Generate commands from the OpenAPI spec
- `make build` - Generate commands and build the binary
- `make dev` - Build without regenerating (for quick iteration)
- `make test` - Run all tests
- `make release` - Build release binaries using GoReleaser

## Troubleshooting

### Authentication issues

```bash
# Check current configuration
meibel auth status
meibel config list

# Re-authenticate
meibel auth login --interactive
```

### Connection issues

```bash
# Check server connectivity
meibel --server http://localhost:8000 datasources list

# Use different profile
meibel --profile development datasources list
```

### Debug mode

Set the `MEIBEL_DEBUG` environment variable for verbose output:

```bash
export MEIBEL_DEBUG=true
meibel datasources list
```

## License

MIT License

## Contributing

Contributions are welcome! Please feel free to submit a Pull Request.