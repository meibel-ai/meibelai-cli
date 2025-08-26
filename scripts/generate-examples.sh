#!/bin/bash
# Generate interactive examples for Mintlify

set -e

OUTPUT_DIR="${1:-examples}"
mkdir -p "$OUTPUT_DIR"

# Function to create example MDX files
create_example() {
    local name="$1"
    local file="$2"
    
    cat > "$OUTPUT_DIR/$file" << 'EOF'
---
title: "'"$name"' Examples"
description: "Interactive examples for '"$name"' commands"
---

EOF
}

# Authentication Examples
create_example "Authentication" "auth-examples.mdx"
cat >> "$OUTPUT_DIR/auth-examples.mdx" << 'EOF'
## Initial Setup

<Steps>
  <Step title="Install the CLI">
    <CodeGroup>
    ```bash macOS
    brew install meibel-ai/tap/meibel
    ```
    
    ```bash Linux
    curl -sfL https://raw.githubusercontent.com/meibel-ai/meibel-cli/main/install.sh | sh
    ```
    
    ```bash Docker
    docker pull meibelai/cli:latest
    ```
    </CodeGroup>
  </Step>
  
  <Step title="Configure Authentication">
    <Tabs>
      <Tab title="Interactive">
        ```bash
        meibel auth login --interactive
        ```
        
        You'll be prompted to enter your API key securely.
      </Tab>
      
      <Tab title="Direct">
        ```bash
        meibel auth login --api-key "your-api-key-here"
        ```
      </Tab>
      
      <Tab title="Environment Variable">
        ```bash
        export MEIBEL_API_KEY="your-api-key-here"
        meibel datasources list
        ```
      </Tab>
    </Tabs>
  </Step>
  
  <Step title="Verify Authentication">
    ```bash
    meibel auth status
    ```
    
    Expected output:
    ```
    Profile: default
    Server: http://api.meibel.ai
    API Key: sk-abc...***
    ```
  </Step>
</Steps>

## Multiple Profiles

<Accordion title="Working with Multiple Environments">
  ```bash
  # Production
  meibel auth login --profile production --api-key "prod-key"
  
  # Staging
  meibel auth login --profile staging --api-key "staging-key" --server "https://staging.api.meibel.ai"
  
  # Development
  meibel auth login --profile dev --api-key "dev-key" --server "http://localhost:8000"
  
  # Use a specific profile
  meibel --profile staging datasources list
  ```
</Accordion>
EOF

# Datasource Examples
create_example "Datasources" "datasource-examples.mdx"
cat >> "$OUTPUT_DIR/datasource-examples.mdx" << 'EOF'
## Common Operations

<CardGroup cols={2}>
  <Card title="List Datasources" icon="list">
    ```bash
    meibel datasources list --output table
    ```
  </Card>
  
  <Card title="Get Details" icon="eye">
    ```bash
    meibel datasources get <id> --output yaml
    ```
  </Card>
  
  <Card title="Create New" icon="plus">
    ```bash
    meibel datasources create --data '{
      "name": "Production DB",
      "type": "postgresql"
    }'
    ```
  </Card>
  
  <Card title="Update Existing" icon="pen">
    ```bash
    meibel datasources update <id> --data '{
      "name": "Updated Name"
    }'
    ```
  </Card>
</CardGroup>

## Advanced Examples

### Batch Operations

<CodeGroup>
```bash Create from File
# datasource.json
{
  "name": "Analytics DB",
  "type": "postgresql",
  "config": {
    "host": "db.example.com",
    "port": 5432,
    "database": "analytics"
  }
}

# Create datasource
meibel datasources create --from-file datasource.json
```

```bash Batch Update
# Update multiple datasources
for id in ds-1 ds-2 ds-3; do
  meibel datasources update $id --data '{"status": "active"}'
done
```

```bash Export All
# Export all datasources
meibel datasources list --output json > datasources-backup.json
```
</CodeGroup>

### Filtering and Pagination

<Tabs>
  <Tab title="Pagination">
    ```bash
    # First page
    meibel datasources list --limit 20 --offset 0
    
    # Second page
    meibel datasources list --limit 20 --offset 20
    ```
  </Tab>
  
  <Tab title="With jq">
    ```bash
    # Filter by type
    meibel datasources list --output json | jq '.[] | select(.type == "postgresql")'
    
    # Count by type
    meibel datasources list --output json | jq 'group_by(.type) | map({type: .[0].type, count: length})'
    ```
  </Tab>
</Tabs>

<Warning>
  Always use `--dry-run` when testing destructive operations:
  
  ```bash
  meibel --dry-run datasources delete <id>
  ```
</Warning>
EOF

# Create more example files...
echo "Examples generated in: $OUTPUT_DIR"