#!/bin/bash
# Script to generate CLI documentation

set -e

DOCS_DIR="docs/commands"
CLI_BINARY="./meibel"

echo "Generating CLI documentation..."

# Build the CLI first
echo "Building CLI..."
go build -o meibel

# Create docs directory
mkdir -p "$DOCS_DIR"

# Generate main command documentation
echo "# Meibel CLI Command Reference" > "$DOCS_DIR/README.md"
echo "" >> "$DOCS_DIR/README.md"
echo "This document contains the complete command reference for the Meibel CLI." >> "$DOCS_DIR/README.md"
echo "" >> "$DOCS_DIR/README.md"
echo "Generated on: $(date)" >> "$DOCS_DIR/README.md"
echo "" >> "$DOCS_DIR/README.md"

# Function to generate docs for a command
generate_command_docs() {
    local cmd="$1"
    local output_file="$2"
    
    echo "Generating docs for: $cmd"
    
    {
        echo "# $cmd"
        echo ""
        $CLI_BINARY $cmd --help 2>/dev/null || true
    } > "$output_file"
}

# Get all top-level commands
COMMANDS=$($CLI_BINARY --help | grep -A 100 "Available Commands:" | grep -B 100 "Flags:" | grep "^  " | awk '{print $1}' | grep -v "^$")

# Generate docs for each command
echo "## Available Commands" >> "$DOCS_DIR/README.md"
echo "" >> "$DOCS_DIR/README.md"

for cmd in $COMMANDS; do
    if [ "$cmd" != "completion" ] && [ "$cmd" != "help" ]; then
        echo "- [$cmd](./$cmd.md)" >> "$DOCS_DIR/README.md"
        generate_command_docs "$cmd" "$DOCS_DIR/$cmd.md"
        
        # Get subcommands
        SUBCMDS=$($CLI_BINARY $cmd --help 2>/dev/null | grep -A 100 "Available Commands:" | grep -B 100 "Flags:" | grep "^  " | awk '{print $1}' | grep -v "^$" || true)
        
        if [ -n "$SUBCMDS" ]; then
            mkdir -p "$DOCS_DIR/$cmd"
            echo "" >> "$DOCS_DIR/$cmd.md"
            echo "## Subcommands" >> "$DOCS_DIR/$cmd.md"
            echo "" >> "$DOCS_DIR/$cmd.md"
            
            for subcmd in $SUBCMDS; do
                echo "- [$subcmd](./$cmd/$subcmd.md)" >> "$DOCS_DIR/$cmd.md"
                generate_command_docs "$cmd $subcmd" "$DOCS_DIR/$cmd/$subcmd.md"
            done
        fi
    fi
done

# Generate examples documentation
echo "" >> "$DOCS_DIR/README.md"
echo "## Common Examples" >> "$DOCS_DIR/README.md"
echo "" >> "$DOCS_DIR/README.md"

cat >> "$DOCS_DIR/README.md" << 'EOF'
### Authentication

```bash
# Set up authentication
meibel auth login --interactive

# Check authentication status
meibel auth status
```

### Working with Datasources

```bash
# List all datasources
meibel datasources list --output table

# Get a specific datasource
meibel datasources get <datasource-id>

# Create a new datasource
meibel datasources create --data '{"name": "My Datasource", "type": "postgres"}'

# Update a datasource
meibel datasources update <datasource-id> --data '{"name": "Updated Name"}'

# Delete a datasource
meibel datasources delete <datasource-id>
```

### Working with Data Elements

```bash
# Add data elements
meibel data-elements add --datasource-id <id> --data '[{"key": "value"}]'

# Retrieve data elements
meibel data-elements retrieve --datasource-id <id> --limit 10

# Update a data element
meibel data-elements update --element-id <id> --data '{"key": "new-value"}'

# Delete a data element
meibel data-elements delete --element-id <id>
```

### Output Formats

```bash
# JSON output (default)
meibel datasources list

# YAML output
meibel datasources list --output yaml

# Table output
meibel datasources list --output table
```

### Dry Run Mode

```bash
# Preview a request without executing it
meibel --dry-run datasources delete <datasource-id>
```
EOF

echo ""
echo "Documentation generated in $DOCS_DIR"
echo "Total commands documented: $(find "$DOCS_DIR" -name "*.md" | wc -l)"