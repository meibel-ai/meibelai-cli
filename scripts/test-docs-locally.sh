#!/bin/bash
# Script to test documentation generation locally

set -e

echo "Testing CLI documentation generation..."

# Build the CLI
echo "Building CLI..."
go build -o meibel

# Test command structure
echo ""
echo "=== Command Structure Test ==="
echo ""

# Test main help
echo "Main Help:"
./meibel --help | grep -A 20 "Available Commands:" | head -20

# Test data-elements commands
echo ""
echo "Data Elements Commands:"
./meibel data-elements --help | grep -A 20 "Available Commands:"

# Test specific command
echo ""
echo "Data Elements Search Command:"
./meibel data-elements search --help

# Generate sample documentation
echo ""
echo "=== Generating Sample Documentation ==="
echo ""

TEMP_DIR="temp-docs-test"
mkdir -p "$TEMP_DIR"

# Generate a sample MDX file
cat > "$TEMP_DIR/data-elements.mdx" << 'EOF'
---
title: "Data Elements"
description: "Manage data elements in your datasources"
---

## Overview

Data elements are the individual pieces of data stored in your datasources. The CLI provides commands to add, retrieve, update, search, and delete data elements.

## Commands

### List All Data Elements

List all data elements in a datasource:

```bash
meibel data-elements list <datasource_id> --limit 20
```

### Get a Specific Data Element

Retrieve a single data element by ID:

```bash
meibel data-elements get <datasource_id> <data_element_id>
```

### Search Data Elements

Search data elements using filters:

```bash
# Using inline filters
meibel data-elements search <datasource_id> \
  --regex_filter "pattern" \
  --media_type_filters "type1,type2" \
  --limit 10

# Using a filter file
meibel data-elements search <datasource_id> \
  --from-file filters.json
```

Example filter file (`filters.json`):
```json
{
  "regex_filter": ".*important.*",
  "media_type_filters": ["text", "image"],
  "date_range": {
    "start": "2024-01-01",
    "end": "2024-12-31"
  }
}
```

### Add a New Data Element

Add a data element to a datasource:

```bash
# Inline data
meibel data-elements add <datasource_id> \
  --data '{"content": "Hello World", "metadata": {"type": "text"}}'

# From file
meibel data-elements add <datasource_id> \
  --from-file element.json
```

### Update a Data Element

Update an existing data element:

```bash
meibel data-elements update <datasource_id> <data_element_id> \
  --data '{"content": "Updated content"}'
```

### Delete a Data Element

Delete a data element:

```bash
# Preview first (dry run)
meibel --dry-run data-elements delete <datasource_id> <data_element_id>

# Actual deletion
meibel data-elements delete <datasource_id> <data_element_id>
```

## Common Patterns

### Batch Operations

Process multiple data elements:

```bash
# Export all elements
meibel data-elements list <datasource_id> --limit 1000 > elements.json

# Process with jq
cat elements.json | jq '.[] | select(.metadata.type == "text")'

# Delete multiple elements
cat elements.json | jq -r '.[] | select(.metadata.obsolete == true) | .id' | \
  while read id; do
    meibel data-elements delete <datasource_id> "$id"
  done
```

### Pagination

Handle large datasets:

```bash
# First page
meibel data-elements list <datasource_id> --limit 50 --offset 0

# Next page  
meibel data-elements list <datasource_id> --limit 50 --offset 50
```

## Output Formats

All commands support multiple output formats:

```bash
# JSON (default)
meibel data-elements get <datasource_id> <element_id>

# YAML
meibel data-elements get <datasource_id> <element_id> --output yaml

# Table (for list commands)
meibel data-elements list <datasource_id> --output table
```
EOF

echo "Sample documentation generated in: $TEMP_DIR/data-elements.mdx"
echo ""
echo "=== Testing Documentation Commands ==="
echo ""

# Show how commands map to documentation
echo "Command to Documentation Mapping:"
echo "- 'meibel data-elements list' → List all data elements"
echo "- 'meibel data-elements get' → Get a specific element by ID"
echo "- 'meibel data-elements search' → Search with filters"
echo "- 'meibel data-elements add' → Add new element"
echo "- 'meibel data-elements update' → Update existing element"
echo "- 'meibel data-elements delete' → Delete element"

echo ""
echo "✅ Documentation test complete!"