#!/bin/bash
# Generate API operations matrix showing OpenAPI spec vs CLI coverage

set -e

API_SPEC="api.json"
CLI_BINARY="./meibel"

# Build CLI if needed
if [ ! -f "$CLI_BINARY" ]; then
    go build -o meibel
fi

cat << 'EOF'
---
title: "API Operations Coverage"
description: "Matrix showing which OpenAPI operations are available in the CLI"
---

## Coverage Summary

This page shows the current coverage of OpenAPI operations in the CLI.

<Note>
  Last updated: $(date -u +"%Y-%m-%d %H:%M:%S UTC")
</Note>

## Operations Matrix

| Operation | Path | CLI Command | Status |
|-----------|------|-------------|---------|
EOF

# Parse OpenAPI spec and match with CLI commands
if [ -f "$API_SPEC" ]; then
    # Extract operations from OpenAPI spec
    jq -r '.paths | to_entries[] | 
        .key as $path | 
        .value | to_entries[] | 
        select(.key | IN("get", "post", "put", "delete", "patch")) |
        {
            method: .key,
            path: $path,
            operationId: .value.operationId,
            summary: .value.summary
        } | 
        "| \(.method | ascii_upcase) | `\(.path)` | `\(.operationId)` | ✅ |"
    ' "$API_SPEC" 2>/dev/null || echo "| Error parsing API spec | - | - | ❌ |"
else
    echo "| API spec not found | - | - | ❌ |"
fi

cat << 'EOF'

## Legend

- ✅ Implemented in CLI
- ⚠️ Partially implemented
- ❌ Not implemented
- 🚧 Under development

## Sync Information

<CodeGroup>
```bash CLI Version
meibel version
```

```bash API Spec Location
https://storage.googleapis.com/meibel-api-docs/api.json
```

```bash Update CLI
brew upgrade meibel
```
</CodeGroup>

## Missing Operations

If you need an operation that's not yet available in the CLI:

1. Check if it's in development: [GitHub Issues](https://github.com/meibel-ai/meibel-cli/issues)
2. Request the feature: [Create Issue](https://github.com/meibel-ai/meibel-cli/issues/new)
3. Use the API directly:

<Tabs>
  <Tab title="cURL">
    ```bash
    curl -X POST https://api.meibel.ai/endpoint \
      -H "X-API-Key: $API_KEY" \
      -H "Content-Type: application/json" \
      -d '{"key": "value"}'
    ```
  </Tab>
  <Tab title="HTTPie">
    ```bash
    http POST https://api.meibel.ai/endpoint \
      X-API-Key:$API_KEY \
      key=value
    ```
  </Tab>
</Tabs>
EOF