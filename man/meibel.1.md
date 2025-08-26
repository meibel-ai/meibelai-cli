# meibel(1) -- CLI for interacting with Meibel AI API

## SYNOPSIS

`meibel` [*OPTIONS*] *COMMAND* [*ARGS*]

## DESCRIPTION

**meibel** is a command-line interface generated from the Meibel AI OpenAPI specification. It provides easy access to all Meibel AI API endpoints with support for multiple output formats, authentication methods, and configuration profiles.

## OPTIONS

`--api-key` *KEY*
: API key for authentication

`--config` *FILE*
: Config file location (default: $HOME/.meibel.yaml)

`--dry-run`
: Preview the request without executing

`--output` *FORMAT*
: Output format: json, yaml, table (default: json)

`--profile` *NAME*
: Configuration profile to use (default: default)

`--server` *URL*
: API server URL

`-h`, `--help`
: Show help information

`-v`, `--version`
: Show version information

## COMMANDS

### Authentication Commands

`auth login`
: Configure API authentication

`auth status`
: Check authentication status

### Resource Commands

`datasources`
: Manage datasources (create, list, get, update, delete)

`data-elements`
: Manage data elements (add, retrieve, update, delete)

`blueprint-instances`
: Manage blueprint instances

`tag`
: Configure tagging

`rag`
: Configure RAG (Retrieval-Augmented Generation)

`workflows`
: Manage workflows

### Utility Commands

`config`
: Manage configuration settings

`completion`
: Generate shell completion scripts

`version`
: Display version information

## EXAMPLES

### Authentication Setup

Set up authentication interactively:

    $ meibel auth login --interactive

Or provide API key directly:

    $ meibel auth login --api-key "your-api-key"

### Working with Datasources

List all datasources in table format:

    $ meibel datasources list --output table

Get a specific datasource:

    $ meibel datasources get <datasource-id>

Create a new datasource:

    $ meibel datasources create --data '{"name": "My DB", "type": "postgres"}'

### Using Different Output Formats

JSON output (default):

    $ meibel datasources list

YAML output:

    $ meibel datasources list --output yaml

Table output:

    $ meibel datasources list --output table

### Dry Run Mode

Preview a request without executing:

    $ meibel --dry-run datasources delete <id>

### Configuration Profiles

Use a different profile:

    $ meibel --profile staging datasources list

## CONFIGURATION

The CLI looks for configuration in the following order:

1. Command-line flags (highest priority)
2. Environment variables (MEIBEL_*)
3. Configuration file (~/.meibel.yaml)
4. Default values

### Configuration File Format

```yaml
api_key: your-api-key
server: http://api.meibel.ai
output: json
profile: default

profiles:
  staging:
    server: http://staging.api.meibel.ai
    api_key: staging-api-key
  
  development:
    server: http://localhost:8000
    api_key: dev-api-key
```

### Environment Variables

- `MEIBEL_API_KEY`: API key for authentication
- `MEIBEL_SERVER`: API server URL
- `MEIBEL_OUTPUT`: Default output format
- `MEIBEL_PROFILE`: Configuration profile to use

## FILES

`~/.meibel.yaml`
: Default configuration file

`~/.config/meibel/`
: Configuration directory

`/usr/share/bash-completion/completions/meibel`
: Bash completion script

`/usr/share/zsh/vendor-completions/_meibel`
: Zsh completion script

## EXIT STATUS

`0`
: Success

`1`
: General error

`2`
: Command line usage error

`3`
: Configuration error

`4`
: Authentication error

`5`
: Network error

## BUGS

Report bugs at: https://github.com/meibel-ai/meibel-cli/issues

## SEE ALSO

Project homepage: https://github.com/meibel-ai/meibel-cli

API documentation: https://api.meibel.ai/docs

## AUTHOR

Meibel AI Team <support@meibel.ai>

## COPYRIGHT

Copyright (C) 2024 Meibel AI. MIT License.