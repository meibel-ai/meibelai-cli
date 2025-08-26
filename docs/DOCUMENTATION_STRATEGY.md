# CLI Documentation Strategy for Mintlify

This document outlines how we keep the Meibel CLI documentation in sync with the actual CLI implementation.

## Overview

Our documentation strategy ensures that:
1. CLI command references are always up-to-date
2. Examples are tested and working
3. API coverage is transparent
4. Changes are automatically detected and documented

## Architecture

```
┌─────────────────┐     ┌──────────────────┐     ┌─────────────────┐
│   CLI Source    │────▶│ Doc Generation   │────▶│ Mintlify Docs   │
│   (this repo)   │     │    Scripts       │     │  (docs repo)    │
└─────────────────┘     └──────────────────┘     └─────────────────┘
         │                       │                         │
         │                       │                         │
         └───────────────────────┴─────────────────────────┘
                        GitHub Actions
```

## Automated Sync Process

### 1. Trigger Points

Documentation regeneration is triggered by:
- Push to `main` branch (affecting `cmd/` or `internal/`)
- New release publication
- API specification updates
- Manual workflow dispatch

### 2. Generation Process

```bash
# On every trigger:
1. Build the CLI
2. Extract command help text
3. Generate MDX files with Mintlify components
4. Create API coverage matrix
5. Generate interactive examples
6. Create PR to docs repository
```

### 3. Files Generated

For each CLI command, we generate:

```
docs/cli/
├── cli-reference.mdx           # Main reference page
├── cli-auth.mdx               # Auth commands
├── cli-datasources.mdx        # Datasource commands
├── datasources/
│   ├── create.mdx            # Subcommand docs
│   ├── list.mdx
│   └── ...
├── api-operations.mdx         # API coverage matrix
└── examples/
    ├── auth-examples.mdx
    └── datasource-examples.mdx
```

## Mintlify Integration

### Navigation Structure

Add to your `mint.json`:

```json
{
  "navigation": [
    {
      "group": "CLI Reference",
      "pages": [
        "cli/cli-reference",
        {
          "group": "Commands",
          "pages": [
            "cli/cli-auth",
            "cli/cli-datasources",
            "cli/cli-config"
          ]
        },
        {
          "group": "Examples",
          "pages": [
            "cli/examples/auth-examples",
            "cli/examples/datasource-examples"
          ]
        },
        "cli/api-operations"
      ]
    }
  ]
}
```

### Mintlify Components Used

- **`<Steps>`** - For multi-step tutorials
- **`<Tabs>`** - For alternative approaches
- **`<CodeGroup>`** - For language/platform variants
- **`<Card>`** & **`<CardGroup>`** - For command overviews
- **`<Accordion>`** - For expandable sections
- **`<Note>`**, **`<Warning>`**, **`<Tip>`** - For callouts

## Manual Documentation

Some documentation requires manual maintenance:

1. **Conceptual Guides** - Best practices, architecture
2. **Tutorials** - Step-by-step workflows
3. **Troubleshooting** - Common issues and solutions
4. **Migration Guides** - Version upgrade paths

## CI/CD Integration

### GitHub Actions Workflow

The `.github/workflows/sync-docs.yml` workflow:
1. Runs on CLI changes
2. Generates documentation
3. Creates PR to docs repo
4. Notifies team for review

### Local Testing

```bash
# Generate docs locally
./scripts/generate-mintlify-docs.sh ../path/to/mintlify/docs/cli

# Preview in Mintlify
cd ../path/to/mintlify
mintlify dev
```

## Quality Checks

### Automated Checks

- Command help text completeness
- Example validity
- Link verification
- API coverage percentage

### Manual Review

PR reviewers should check:
- [ ] Command descriptions are clear
- [ ] Examples are practical
- [ ] Navigation makes sense
- [ ] No sensitive information exposed

## Versioning Strategy

### Documentation Versions

- **Latest** - Tracks `main` branch
- **Stable** - Tracks latest release
- **Version-specific** - For major versions

### Version Notices

Add version compatibility notices:

```mdx
<Note>
  Available in v1.2.0 and later
</Note>
```

## Best Practices

### 1. Command Help Text

Write help text in commands with documentation in mind:

```go
var createCmd = &cobra.Command{
    Use:   "create",
    Short: "Create a new datasource",
    Long: `Create a new datasource in your Meibel account.

This command allows you to create datasources from JSON data provided
directly or from a file. The datasource configuration should match
the schema defined in the API documentation.`,
    Example: `  # Create from inline JSON
  meibel datasources create --data '{"name": "My DB", "type": "postgres"}'
  
  # Create from file
  meibel datasources create --from-file datasource.json`,
}
```

### 2. Structured Examples

Create example files in `examples/` directory:

```markdown
# examples/datasources-create.md

## Basic Creation
\`\`\`bash
meibel datasources create --data '{"name": "Production DB"}'
\`\`\`

## From File
\`\`\`bash
meibel datasources create --from-file config.json
\`\`\`
```

### 3. API Sync Transparency

Always document:
- Which API version the CLI supports
- Any operations not yet implemented
- Workarounds for missing features

## Troubleshooting

### Common Issues

1. **Docs not updating**
   - Check GitHub Action logs
   - Verify DOCS_REPO_TOKEN permissions
   - Ensure PR was created

2. **Broken examples**
   - Test examples locally first
   - Use `--dry-run` in examples
   - Include error handling

3. **Navigation issues**
   - Verify mint.json structure
   - Check file paths match
   - Test with `mintlify dev`

## Command Naming Convention

### Problem
OpenAPI specs often have multiple endpoints that could result in duplicate command names (e.g., multiple GET operations).

### Solution
Our command generator uses intelligent naming:

| OpenAPI Operation | Generated Command | Example |
|------------------|------------------|---------|
| `get_data_elements` | `list` | `meibel data-elements list` |
| `get_data_element` | `get` | `meibel data-elements get <id>` |
| `get_data_elements_by_filters` | `search` | `meibel data-elements search` |
| `add_data_element` | `add` | `meibel data-elements add` |

### Naming Rules

1. **Plural GET** → `list`
2. **Singular GET** → `get`
3. **GET with filters** → `search`
4. **POST (non-search)** → `create` or `add`
5. **POST with filters** → `search`
6. **PUT** → `update`
7. **DELETE** → `delete`

## Future Enhancements

1. **Interactive playground** - Embed terminal emulator
2. **Video tutorials** - Record common workflows
3. **API mocking** - Safe practice environment
4. **Localization** - Multi-language support
5. **Search integration** - Command fuzzy search

## Maintenance Checklist

Weekly:
- [ ] Review documentation PRs
- [ ] Test random examples
- [ ] Check for broken links

Monthly:
- [ ] Audit API coverage
- [ ] Update troubleshooting guide
- [ ] Review user feedback

Quarterly:
- [ ] Major documentation review
- [ ] Update conceptual guides
- [ ] Plan new tutorials