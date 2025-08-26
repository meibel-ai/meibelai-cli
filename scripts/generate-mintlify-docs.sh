#!/bin/bash
# Script to generate Mintlify documentation from CLI commands

set -e

# Configuration
DOCS_DIR="${1:-../meibel-docs/cli}"  # Path to your Mintlify docs
CLI_BINARY="./meibel"
TEMP_DIR="temp-docs"

echo "Generating Mintlify documentation for CLI..."

# Build the CLI first
echo "Building CLI..."
go build -o meibel

# Create temp directory
mkdir -p "$TEMP_DIR"

# Function to generate MDX frontmatter
generate_frontmatter() {
    local title="$1"
    local description="$2"
    cat << EOF
---
title: "$title"
description: "$description"
---
EOF
}

# Function to convert command output to MDX
command_to_mdx() {
    local cmd="$1"
    local output_file="$2"
    local title="$3"
    
    echo "Generating docs for: $cmd"
    
    # Get command help
    local help_output=$($CLI_BINARY $cmd --help 2>&1 || true)
    
    # Extract description (first line after command name)
    local description=$(echo "$help_output" | head -n 2 | tail -n 1)
    
    # Generate MDX file
    {
        generate_frontmatter "$title" "$description"
        echo ""
        echo "## Usage"
        echo ""
        echo '```bash'
        echo "meibel $cmd [flags]"
        echo '```'
        echo ""
        echo "## Description"
        echo ""
        echo "$description"
        echo ""
        
        # Extract and format the help content
        echo "## Command Help"
        echo ""
        echo '```'
        echo "$help_output"
        echo '```'
        
        # Add examples if available
        if [ -f "examples/${cmd// /-}.md" ]; then
            echo ""
            echo "## Examples"
            echo ""
            cat "examples/${cmd// /-}.md"
        fi
    } > "$output_file"
}

# Create proper directory structure
mkdir -p "$TEMP_DIR/commands"
mkdir -p "$TEMP_DIR/config" 
mkdir -p "$TEMP_DIR/advanced"

# Generate overview page
echo "Generating CLI overview..."
{
    generate_frontmatter "Meibel CLI" "Command-line interface for the Meibel AI API"
    cat << 'EOF'

The Meibel CLI provides a powerful command-line interface for interacting with the Meibel AI API. It offers intuitive commands, multiple output formats, and advanced features like dry-run mode and configuration profiles.

## Features

- 🚀 Auto-generated from OpenAPI specification
- 🔐 Multiple authentication methods
- 📊 Multiple output formats (JSON, YAML, Table)
- 🎨 Syntax highlighting for JSON output
- ⚙️ Configuration profiles for different environments
- 🔍 Dry-run mode for safe testing
- 💾 Request templates and file input

## Quick Start

```bash
# Install
brew install meibel-ai/tap/meibel

# Configure authentication
meibel auth login

# Start using
meibel datasources list
```

## Next Steps

- [Installation Guide](./installation)
- [Authentication Setup](./authentication)
- [Command Reference](./commands/overview)
EOF
} > "$TEMP_DIR/overview.mdx"

# Generate installation page
echo "Generating installation guide..."
{
    generate_frontmatter "Installation" "How to install the Meibel CLI"
    cat << 'EOF'

## Homebrew (macOS/Linux)

```bash
brew tap meibel-ai/tap
brew install meibel
```

## Direct Download

```bash
# Install script
curl -sfL https://raw.githubusercontent.com/meibel-ai/meibel-cli/main/install.sh | sh

# Or download binary directly
curl -LO https://github.com/meibel-ai/meibel-cli/releases/latest/download/meibel_$(uname -s)_$(uname -m).tar.gz
tar -xzf meibel_*.tar.gz
sudo mv meibel /usr/local/bin/
```

## Package Managers

### Debian/Ubuntu

```bash
curl -LO https://github.com/meibel-ai/meibel-cli/releases/latest/download/meibel_linux_amd64.deb
sudo dpkg -i meibel_*.deb
```

### RedHat/Fedora

```bash
curl -LO https://github.com/meibel-ai/meibel-cli/releases/latest/download/meibel_linux_amd64.rpm
sudo rpm -i meibel_*.rpm
```

## Docker

```bash
docker pull meibelai/cli:latest
docker run --rm meibelai/cli:latest --help
```

## Verify Installation

```bash
meibel version
```
EOF
} > "$TEMP_DIR/installation.mdx"

# Generate authentication page
echo "Generating authentication guide..."
{
    generate_frontmatter "Authentication" "Setting up authentication for the Meibel CLI"
    cat << 'EOF'

## Configuration Methods

### Interactive Setup

The easiest way to configure authentication:

```bash
meibel auth login --interactive
```

### Direct Configuration

```bash
meibel auth login --api-key "your-api-key-here"
```

### Environment Variable

```bash
export MEIBEL_API_KEY="your-api-key-here"
```

### Configuration File

Create `~/.meibel.yaml`:

```yaml
api_key: your-api-key-here
server: https://api.meibel.ai
```

## Multiple Profiles

Configure different environments:

```bash
# Production
meibel auth login --profile production --api-key "prod-key"

# Staging  
meibel auth login --profile staging --api-key "staging-key" --server "https://staging.api.meibel.ai"

# Use specific profile
meibel --profile staging datasources list
```

## Verify Authentication

```bash
meibel auth status
```

## Security Best Practices

1. Never commit API keys to version control
2. Use environment variables in CI/CD
3. Rotate keys regularly
4. Use restricted keys when possible
EOF
} > "$TEMP_DIR/authentication.mdx"

# Generate commands overview
echo "Generating commands overview..."
{
    generate_frontmatter "Commands Overview" "Available commands in the Meibel CLI"
    echo ""
    echo "## Command Structure"
    echo ""
    echo '```'
    echo "meibel <resource> <action> [arguments] [flags]"
    echo '```'
    echo ""
    echo "## Available Resources"
    echo ""
    # List main command groups
    for cmd in $COMMANDS; do
        if [ "$cmd" != "auth" ] && [ "$cmd" != "config" ] && [ "$cmd" != "version" ]; then
            echo "- [$cmd](./$(echo $cmd | tr '[:upper:]' '[:lower:]'))"
        fi
    done
} > "$TEMP_DIR/commands/overview.mdx"

# Generate docs for each command group
for cmd in $COMMANDS; do
    if [ -n "$cmd" ] && [ "$cmd" != "auth" ] && [ "$cmd" != "config" ] && [ "$cmd" != "version" ]; then
        # Create safe filename
        filename=$(echo "$cmd" | tr '[:upper:]' '[:lower:]' | tr ' ' '-')
        
        echo "Generating docs for command group: $cmd"
        
        # Generate command group overview
        {
            # Get help text
            help_output=$($CLI_BINARY $cmd --help 2>&1 || true)
            description=$(echo "$help_output" | head -n 2 | tail -n 1)
            
            generate_frontmatter "${cmd^} Commands" "$description"
            echo ""
            echo "## Available Commands"
            echo ""
            
            # Extract subcommands
            SUBCMDS=$($CLI_BINARY $cmd --help 2>/dev/null | grep -A 100 "Available Commands:" | grep -B 100 "Flags:" | grep "^  " | awk '{print $1}' | grep -v "^$" || true)
            
            if [ -n "$SUBCMDS" ]; then
                echo "| Command | Description |"
                echo "|---------|-------------|"
                
                for subcmd in $SUBCMDS; do
                    desc=$($CLI_BINARY $cmd --help 2>/dev/null | grep "^  $subcmd" | sed "s/^  $subcmd *//" | head -n 1)
                    echo "| \`$subcmd\` | $desc |"
                done
                
                echo ""
                echo "## Examples"
                echo ""
                
                # Add command-specific examples
                case "$cmd" in
                    "datasources")
                        cat << 'EOF'
### List all datasources
```bash
meibel datasources list --output table
```

### Get a specific datasource
```bash
meibel datasources get <datasource_id>
```

### Create a new datasource
```bash
meibel datasources create --data '{"name": "Production DB", "type": "postgres"}'
```
EOF
                        ;;
                    "data-elements")
                        cat << 'EOF'
### List data elements
```bash
meibel data-elements list <datasource_id> --limit 20
```

### Search data elements
```bash
meibel data-elements search <datasource_id> --regex_filter "important"
```

### Add a data element
```bash
meibel data-elements add <datasource_id> --from-file element.json
```
EOF
                        ;;
                esac
            fi
        } > "$TEMP_DIR/commands/${filename}.mdx"
    fi
done

# Generate configuration pages
echo "Generating configuration documentation..."

# Profiles page
{
    generate_frontmatter "Configuration Profiles" "Using multiple configuration profiles"
    cat << 'EOF'

## Overview

Profiles allow you to manage multiple configurations for different environments.

## Creating Profiles

```bash
# Set up production profile
meibel config set api_key "prod-key" --profile production
meibel config set server "https://api.meibel.ai" --profile production

# Set up staging profile  
meibel config set api_key "staging-key" --profile staging
meibel config set server "https://staging.api.meibel.ai" --profile staging
```

## Using Profiles

```bash
# Use specific profile for a command
meibel --profile staging datasources list

# Set default profile
meibel config set profile staging
```

## Profile Configuration File

```yaml
# ~/.meibel.yaml
api_key: default-key
server: https://api.meibel.ai

profiles:
  production:
    api_key: prod-key
    server: https://api.meibel.ai
    
  staging:
    api_key: staging-key  
    server: https://staging.api.meibel.ai
    
  development:
    api_key: dev-key
    server: http://localhost:8000
```
EOF
} > "$TEMP_DIR/config/profiles.mdx"

# Environment variables page
{
    generate_frontmatter "Environment Variables" "Configuring the CLI with environment variables"
    cat << 'EOF'

## Available Variables

| Variable | Description | Example |
|----------|-------------|---------|
| `MEIBEL_API_KEY` | API key for authentication | `sk_live_xxxxx` |
| `MEIBEL_SERVER` | API server URL | `https://api.meibel.ai` |
| `MEIBEL_OUTPUT` | Default output format | `json`, `yaml`, `table` |
| `MEIBEL_PROFILE` | Default profile to use | `production` |

## Usage

```bash
# Set for current session
export MEIBEL_API_KEY="your-key"
export MEIBEL_OUTPUT="table"

# Set for single command
MEIBEL_PROFILE=staging meibel datasources list
```

## Priority Order

Configuration is loaded in this order (highest to lowest priority):

1. Command-line flags
2. Environment variables
3. Configuration file
4. Default values
EOF
} > "$TEMP_DIR/config/environment.mdx"

# Output formats page
{
    generate_frontmatter "Output Formats" "Available output formats for CLI commands"
    cat << 'EOF'

## Supported Formats

### JSON (Default)

Pretty-printed JSON with syntax highlighting:

```bash
meibel datasources get <id>
```

### YAML

Human-readable YAML format:

```bash
meibel datasources get <id> --output yaml
```

### Table

Formatted tables for list operations:

```bash
meibel datasources list --output table
```

## Setting Default Format

```bash
# Via config
meibel config set output table

# Via environment
export MEIBEL_OUTPUT=yaml
```

## Format Examples

<CodeGroup>

```json JSON
{
  "id": "ds-123",
  "name": "Production Database",
  "type": "postgresql",
  "created_at": "2024-01-15T10:30:00Z"
}
```

```yaml YAML
id: ds-123
name: Production Database
type: postgresql
created_at: '2024-01-15T10:30:00Z'
```

```text Table
ID      | Name                | Type       | Created
--------|---------------------|------------|------------------------
ds-123  | Production Database | postgresql | 2024-01-15 10:30:00
```

</CodeGroup>
EOF
} > "$TEMP_DIR/config/output-formats.mdx"

# Generate advanced pages
echo "Generating advanced documentation..."

# Dry-run page
{
    generate_frontmatter "Dry Run Mode" "Preview commands without executing them"
    cat << 'EOF'

## Overview

Dry-run mode allows you to preview API requests without actually executing them. This is useful for testing commands, debugging, and learning how the CLI works.

## Usage

Add the `--dry-run` flag to any command:

```bash
meibel --dry-run datasources create --data '{"name": "Test"}'
```

## Output

Dry-run mode shows:
- HTTP method
- Full URL
- Request headers (sanitized)
- Request body

Example:

```
=== DRY RUN MODE ===
Method: POST
URL: https://api.meibel.ai/datasource
Headers:
  Content-Type: application/json
  X-API-Key: sk_live_*****
Body:
{
  "name": "Test"
}
===================
```

## Use Cases

### Testing Complex Commands

```bash
# Preview a complex search
meibel --dry-run data-elements search ds-123 \
  --regex_filter "important" \
  --media_type_filters "text,image" \
  --limit 50
```

### Learning API Structure

```bash
# See how CLI commands map to API calls
meibel --dry-run datasources update ds-123 \
  --data '{"status": "active"}'
```

### Debugging

```bash
# Check if authentication is set correctly
meibel --dry-run auth status
```
EOF
} > "$TEMP_DIR/advanced/dry-run.mdx"

# Scripting page
{
    generate_frontmatter "Scripting" "Using the CLI in scripts and automation"
    cat << 'EOF'

## Exit Codes

The CLI uses standard exit codes:

| Code | Meaning |
|------|---------|
| 0 | Success |
| 1 | General error |
| 2 | Invalid arguments |
| 3 | Authentication error |
| 4 | Network error |

## Scripting Examples

### Error Handling

```bash
#!/bin/bash
if meibel datasources list > /dev/null 2>&1; then
    echo "Connection successful"
else
    echo "Failed to connect" >&2
    exit 1
fi
```

### Processing Output

```bash
# Get datasource IDs
ids=$(meibel datasources list --output json | jq -r '.[].id')

# Process each datasource
for id in $ids; do
    echo "Processing $id..."
    meibel datasources get "$id" --output json | jq '.name'
done
```

### Batch Operations

```bash
# Add multiple data elements
cat elements.jsonl | while read line; do
    meibel data-elements add ds-123 --data "$line"
done
```

## Best Practices

1. **Always check exit codes**
   ```bash
   meibel datasources create --data "$json" || {
       echo "Failed to create datasource" >&2
       exit 1
   }
   ```

2. **Use `--output json` for parsing**
   ```bash
   result=$(meibel datasources list --output json)
   count=$(echo "$result" | jq length)
   ```

3. **Handle pagination**
   ```bash
   offset=0
   limit=100
   while true; do
       result=$(meibel data-elements list ds-123 \
           --limit $limit --offset $offset --output json)
       
       # Process results
       echo "$result" | jq -r '.[]'
       
       # Check if more results
       if [ $(echo "$result" | jq length) -lt $limit ]; then
           break
       fi
       offset=$((offset + limit))
   done
   ```

4. **Use environment variables for auth**
   ```bash
   export MEIBEL_API_KEY="$SECRET_API_KEY"
   meibel datasources list
   ```
EOF
} > "$TEMP_DIR/advanced/scripting.mdx"

# Troubleshooting page
{
    generate_frontmatter "Troubleshooting" "Common issues and solutions"
    cat << 'EOF'

## Common Issues

### Authentication Errors

**Problem**: `Error: API error (401): Unauthorized`

**Solutions**:
1. Check API key is set:
   ```bash
   meibel auth status
   ```

2. Verify key is correct:
   ```bash
   meibel auth login --interactive
   ```

3. Check environment:
   ```bash
   echo $MEIBEL_API_KEY
   ```

### Network Errors

**Problem**: `Error: request failed: connection refused`

**Solutions**:
1. Check server URL:
   ```bash
   meibel config get server
   ```

2. Test connectivity:
   ```bash
   curl -I https://api.meibel.ai
   ```

3. Use custom server:
   ```bash
   meibel --server http://localhost:8000 datasources list
   ```

### Command Not Found

**Problem**: `command not found: meibel`

**Solutions**:
1. Check installation:
   ```bash
   which meibel
   ```

2. Add to PATH:
   ```bash
   export PATH="$PATH:/usr/local/bin"
   ```

3. Reinstall:
   ```bash
   brew reinstall meibel
   ```

## Debug Mode

Enable verbose logging:

```bash
export MEIBEL_DEBUG=true
meibel datasources list
```

## Getting Help

### Built-in Help

```bash
# General help
meibel --help

# Command help
meibel datasources --help

# Subcommand help
meibel datasources create --help
```

### Version Information

```bash
# Simple version
meibel version

# Detailed version
meibel version --verbose
```

### Support Resources

- GitHub Issues: [github.com/meibel-ai/meibel-cli/issues](https://github.com/meibel-ai/meibel-cli/issues)
- API Documentation: [docs.meibel.ai](https://docs.meibel.ai)
- Email Support: support@meibel.ai
EOF
} > "$TEMP_DIR/advanced/troubleshooting.mdx"

# Generate navigation structure
echo ""
echo "Generating navigation structure for docs.json..."
echo ""
echo '{'
echo '  "group": "CLI",'
echo '  "pages": ['
echo '    "sdk/cli/overview",'
echo '    "sdk/cli/installation",'
echo '    "sdk/cli/authentication",'
echo '    {'
echo '      "group": "Commands",'
echo '      "pages": ['
echo '        "sdk/cli/commands/overview",'

for cmd in $COMMANDS; do
    if [ "$cmd" != "auth" ] && [ "$cmd" != "config" ] && [ "$cmd" != "version" ]; then
        filename=$(echo "$cmd" | tr '[:upper:]' '[:lower:]' | tr ' ' '-')
        echo '        "sdk/cli/commands/'$filename'",'
    fi
done

echo '      ]'
echo '    },'
echo '    {'
echo '      "group": "Configuration",'
echo '      "pages": ['
echo '        "sdk/cli/config/profiles",'
echo '        "sdk/cli/config/environment",'
echo '        "sdk/cli/config/output-formats"'
echo '      ]'
echo '    },'
echo '    {'
echo '      "group": "Advanced",'
echo '      "pages": ['
echo '        "sdk/cli/advanced/dry-run",'
echo '        "sdk/cli/advanced/scripting",'
echo '        "sdk/cli/advanced/troubleshooting"'
echo '      ]'
echo '    }'
echo '  ]'
echo '}'

# Generate OpenAPI sync status
echo "Generating API sync status..."
{
    generate_frontmatter "API Sync Status" "Current synchronization status with OpenAPI specification"
    echo ""
    echo "## API Specification"
    echo ""
    echo "- **Source**: https://storage.googleapis.com/meibel-api-docs/api.json"
    echo "- **Last Sync**: $(date -u +"%Y-%m-%d %H:%M:%S UTC")"
    echo "- **CLI Version**: $($CLI_BINARY version 2>&1 | grep -oE 'v[0-9]+\.[0-9]+\.[0-9]+' || echo "unknown")"
    echo ""
    
    # List all available operations
    echo "## Available Operations"
    echo ""
    echo "The following operations are available in the CLI:"
    echo ""
    
    for cmd in $COMMANDS; do
        echo "### $cmd"
        $CLI_BINARY $cmd --help 2>/dev/null | grep -A 20 "Available Commands:" | grep -B 20 "Flags:" | grep "^  " | grep -v "^$" | sed 's/^/- /' || true
        echo ""
    done
} > "$TEMP_DIR/api-sync-status.mdx"

# Copy to Mintlify docs directory
if [ -d "$DOCS_DIR" ]; then
    echo "Copying to Mintlify docs directory..."
    mkdir -p "$DOCS_DIR"
    cp -r "$TEMP_DIR"/* "$DOCS_DIR/"
    
    echo ""
    echo "✅ Documentation generated successfully!"
    echo "📁 Output directory: $DOCS_DIR"
    echo ""
    echo "Add this to your mint.json navigation:"
    echo ""
    echo '  {
    "group": "CLI Reference",
    "pages": '$NAV_ITEMS'
  }'
else
    echo ""
    echo "⚠️  Mintlify docs directory not found: $DOCS_DIR"
    echo "📁 Documentation generated in: $TEMP_DIR"
    echo ""
    echo "To use these docs:"
    echo "1. Copy the files to your Mintlify docs/cli directory"
    echo "2. Add the navigation structure to your mint.json"
fi

# Clean up
# rm -rf "$TEMP_DIR"