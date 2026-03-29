# Claude Code Instructions: OpenAPI-Based Ergonomic CLI Generator

## Project Overview
Build a production-ready, ergonomic CLI tool in Go using Cobra that automatically generates commands from an OpenAPI specification. Focus on developer experience, intuitive usage patterns, and robust error handling.

## Architecture Requirements

### Core Components
1. **OpenAPI Parser**: Parse and validate OpenAPI 3.x specifications
2. **Command Generator**: Generate Cobra commands from OpenAPI operations
3. **Request Builder**: Build HTTP requests with proper validation
4. **Response Handler**: Format and display responses elegantly
5. **Authentication Manager**: Handle various auth schemes
6. **Configuration System**: Support multiple config sources

### Project Structure
```
├── cmd/
│   ├── root.go              # Root command setup and registration
│   ├── generated_commands.go # Auto-generated commands from OpenAPI
│   ├── auth.go              # Authentication commands
│   ├── config.go            # Configuration commands
│   ├── version.go           # Version command
│   └── generate/            
│       └── main.go          # Build-time code generator
├── internal/
│   ├── client/              # HTTP client with auth
│   ├── config/              # Configuration management
│   ├── formatter/           # Response formatting (JSON, YAML, Table)
│   ├── openapi/             # OpenAPI types and utilities
│   ├── validator/           # Input validation
│   └── version/             # Version information
├── docs/                    # Documentation
├── scripts/                 # Build and release scripts
├── Makefile                 # Build automation
├── .goreleaser.yaml         # GoReleaser configuration
└── main.go                  # Entry point
```

## Key Features to Implement

### 1. OpenAPI Specification Support
- **Full OpenAPI 3.x compatibility**
- Support for JSON and YAML formats
- Remote spec fetching with caching
- Spec validation and error reporting
- Multiple spec sources (file, URL, embedded)
OPENAPI LOCATION: https://storage.googleapis.com/meibel-api-docs/api.json

### 2. Dynamic Command Generation
- Generate commands from OpenAPI operations
- Smart command naming and grouping
- Operation tags → command groups
- Path parameters → command arguments
- Query parameters → flags
- Request body → structured input

### 3. Ergonomic CLI Design
- **Intuitive command structure**: `/users/{id}` → `users get <id>`
- **Smart flag naming**: Convert `camelCase` to `kebab-case`
- **Rich help system**: Include operation descriptions and examples
- **Auto-completion**: Generate completion for bash/zsh/fish
- **Progressive disclosure**: Show relevant options contextually

### 4. Advanced Input Handling
- **JSON/YAML input**: Support both inline and file-based
- **Interactive prompts**: For required parameters when missing
- **Template support**: Pre-built request templates
- **Validation**: Client-side validation using OpenAPI schema
- **Type coercion**: Smart parameter type conversion

### 5. Response Management
- **Multiple output formats**: JSON, YAML, table, custom templates
- **Streaming support**: Handle large responses and SSE
- **Error formatting**: Human-readable error messages
- **Response caching**: Optional caching for GET requests
- **Pagination**: Auto-handle paginated responses

### 6. Authentication & Security
- **Multiple auth schemes**: API Key, Bearer, Basic, OAuth2
- **Credential management**: Secure storage and rotation
- **Environment integration**: Support env vars and config files
- **Token refresh**: Automatic token renewal
- **Security scheme detection**: Auto-configure from spec

## Implementation Guidelines

### Command Generation Strategy
```go
// Example command structure for POST /users/{id}/posts
cli posts create --user-id 123 --title "Hello" --body "World"
cli posts create --user-id 123 --from-file post.json
cli posts create --user-id 123 --interactive
```

### Configuration Hierarchy
1. Command-line flags (highest priority)
2. Environment variables
3. Configuration files (.cli.yaml, .cli.json)
4. OpenAPI spec defaults (lowest priority)

### Error Handling Standards
- **Validation errors**: Show specific field errors with suggestions
- **HTTP errors**: Display status codes with human-readable messages
- **Network errors**: Provide retry suggestions and connectivity hints
- **Auth errors**: Clear instructions for credential setup
- **Spec errors**: Detailed parsing errors with line numbers

### Output Formatting
- **Default**: Pretty JSON with syntax highlighting
- **Table**: Structured data in tabular format
- **Template**: Custom Go templates for specific formats
- **Raw**: Direct server response
- **Quiet**: Minimal output for scripting

## Code Quality Requirements

### Go Best Practices
- Follow effective Go patterns and idioms
- Use context.Context for cancellation and timeouts
- Implement proper error wrapping with fmt.Errorf
- Use structured logging (slog or logrus)
- Include comprehensive unit tests (>80% coverage)

### Dependencies
- **Required**: `github.com/spf13/cobra`, `github.com/spf13/viper`
- **OpenAPI**: `github.com/getkin/kin-openapi` or `github.com/go-openapi/spec`
- **HTTP**: Enhanced `net/http` with retry logic
- **Config**: Support for YAML, JSON, TOML
- **Testing**: `github.com/stretchr/testify`

### Performance Considerations
- Build-time command generation (no runtime parsing)
- Instant startup time (<100ms)
- Efficient JSON parsing and streaming
- Connection pooling for HTTP requests
- Minimal memory footprint (~50MB)

## User Experience Focus

### Discovery and Learning
- `cli --help` shows available API groups
- `cli users --help` shows all user operations
- `cli users create --help` shows parameter details and examples
- Built-in `describe` command for exploring the API

### Common Usage Patterns
```bash
# Initialize with OpenAPI spec
cli init --spec https://api.example.com/openapi.json

# Explore the API
cli describe
cli describe users

# Authentication setup
cli auth login --interactive
cli config set api-key "sk-..."

# Common operations
cli users list --limit 10 --format table
cli users create --interactive
cli users get 123 --output yaml
cli users update 123 --from-file user.json

# Advanced features
cli users list --watch  # Stream updates
cli --dry-run users delete 123  # Preview without executing
cli --profile staging users list  # Use different environment
```

### Developer Experience
- **Fast startup**: < 100ms for simple commands
- **Helpful errors**: Suggest corrections for typos
- **Examples everywhere**: Show usage examples in help
- **IDE integration**: Generate completion scripts
- **Extensibility**: Plugin system for custom formatters

## Testing Strategy

### Unit Tests
- OpenAPI parsing with various specs
- Command generation logic
- Request building and validation
- Response formatting
- Authentication flows

### Integration Tests
- End-to-end CLI workflows
- Real API interactions (with mocking)
- Configuration management
- Error scenarios

### Example Test Cases
```go
func TestCommandGeneration(t *testing.T) {
    spec := loadTestSpec("petstore.yaml")
    generator := NewGenerator(spec)
    
    commands := generator.GenerateCommands()
    
    assert.Contains(t, commands, "pets")
    assert.Contains(t, commands["pets"].Commands, "list")
    assert.Contains(t, commands["pets"].Commands, "create")
}
```

## Documentation Requirements

### Generated Documentation
- CLI reference (markdown/man pages)
- API operation mapping
- Configuration options
- Authentication setup guides

### Examples and Tutorials
- Getting started guide
- Common workflow examples
- Advanced usage patterns
- Troubleshooting guide

## Build and Distribution

### Build Configuration
- Multi-platform builds (Linux, macOS, Windows)
- Static binaries with embedded specs
- Version information and build metadata
- Reproducible builds

### Distribution
- GitHub releases with pre-built binaries
- Package managers (Homebrew, apt, yum)
- Docker images for containerized usage
- Shell completion scripts

## Success Criteria

### Functional Requirements
- ✅ Parse any valid OpenAPI 3.x specification
- ✅ Generate intuitive CLI commands for all operations
- ✅ Handle complex parameter types and validation
- ✅ Support all common authentication schemes
- ✅ Provide multiple output formats
- ✅ Maintain excellent error messages

### Non-Functional Requirements
- ✅ Startup time under 100ms for cached specs
- ✅ Memory usage under 50MB for typical usage
- ✅ Support specs with 1000+ operations
- ✅ Comprehensive test coverage (>80%)
- ✅ Clear, maintainable code structure

### User Experience Goals
- ✅ New users can get started in under 5 minutes
- ✅ Discoverable without reading documentation
- ✅ Consistent with common CLI tool patterns
- ✅ Powerful enough for complex automation

---

## Implementation Architecture

### Build-Time Code Generation

The CLI uses build-time code generation for optimal performance:

1. **Generator Tool** (`cmd/generate/main.go`):
   - Parses OpenAPI specification at build time
   - Generates Go code for all API operations
   - Creates `cmd/generated_commands.go` with all commands
   - Supports intelligent command naming and sub-resource handling

2. **Hybrid Command System**:
   - Generated commands from OpenAPI spec
   - Custom commands (auth, config, version) in separate files
   - Both types registered in `cmd/root.go`

3. **Build Process**:
   ```bash
   make generate  # Run the code generator
   make build     # Generate and build
   make dev       # Quick build without regeneration
   ```

### Key Implementation Details

#### Command Generation Strategy
- Operations are mapped to intuitive CLI commands
- Smart naming based on HTTP method and path patterns
- Sub-resources properly nested (e.g., `blueprint-instances <id> events create`)
- Deduplication of operation names while preserving functionality

#### Performance Characteristics
- Instant startup time (<100ms) due to build-time generation
- No runtime OpenAPI parsing overhead
- Minimal memory footprint
- Static command tree for fast command resolution

#### Extensibility
- Easy to add custom commands alongside generated ones
- Clear separation of concerns between generated and custom code
- Generator can be enhanced without affecting runtime

## Implementation Status

The current implementation successfully achieves:
- ✅ Build-time command generation from OpenAPI spec
- ✅ Instant CLI startup with no runtime parsing
- ✅ Support for both generated and custom commands
- ✅ Smart command naming with operation ID parsing
- ✅ Multiple output formats (JSON, YAML, table)
- ✅ Comprehensive authentication management
- ✅ Configuration profiles and environment support
- ✅ GoReleaser integration for distribution
- ✅ Homebrew tap support
- ✅ Colorized JSON output for better readability
- ✅ Dry-run mode for request preview
- ✅ Secure credential storage with masked display

Next improvements should focus on:
- Enhanced sub-resource command structure
- Better handling of nested resources in the generator
- Template support for complex requests
- Batch operation capabilities
- Interactive mode for complex workflows