# OpenAPI CLI Generator Improvements Plan

## Overview
This document outlines comprehensive improvements for the Go OpenAPI-to-CLI generator, transforming it from a basic code generator into a production-ready, feature-rich CLI generation framework.

## 1. Architecture & Maintainability

### 1.1 Modular Package Structure
```
pkg/
  generator/
    command.go      // Command generation logic
    naming.go       // Naming strategies
    template.go     // Template management
    sse.go         // Server-Sent Events support
  parser/
    openapi.go      // OpenAPI parsing
    validation.go   // Spec validation
    extensions.go   // Custom extension handling
  codegen/
    golang.go       // Go-specific generation
    types.go        // Type mapping and generation
    templates/      // External template files
  console/
    formatter.go    // Rich console output
    interactive.go  // Interactive features
    sse_handler.go  // SSE streaming handlers
```

### 1.2 External Template System
- Move `codeTemplate` to external `.go.tmpl` files
- Support multiple template variants (basic, advanced, custom)
- Template inheritance and composition
- Hot-reloading for development

### 1.3 Configuration System
```yaml
# .meibel-cli.yaml
generator:
  template: "custom-template.go.tmpl"
  output_dir: "cmd/generated"
  naming:
    strategy: "rest-semantic"
    overrides:
      "createUser": "register"
      "getUserById": "show"
  features:
    auth: true
    sse: true
    pagination: true
    interactive: true
    rich_output: true
  console:
    colors: true
    tables: true
    progress: true
```

## 2. Enhanced Type System & Code Generation

### 2.1 Proper Go Type Generation
- Generate structs from OpenAPI schemas
- Handle complex nested objects, arrays, maps
- Support for `oneOf`, `anyOf`, `allOf` schemas
- Proper validation tags and JSON annotations
- Custom type mappings and overrides

### 2.2 Advanced Schema Support
```go
type SchemaGenerator interface {
    GenerateStruct(schema *openapi3.Schema) (*ast.StructType, error)
    GenerateValidation(schema *openapi3.Schema) ([]string, error)
    HandleComplexTypes(schema *openapi3.Schema) error
}
```

### 2.3 Better Request/Response Handling
- Strongly typed request/response objects
- Automatic serialization/deserialization
- Content-Type negotiation
- File upload/download support

## 3. Server-Sent Events (SSE) Support

### 3.1 SSE Detection in OpenAPI
```go
func isSSEEndpoint(op *openapi.Operation) bool {
    // Check for text/event-stream content type
    // Check for x-streaming extensions
    // Detect streaming patterns in operation ID
}
```

### 3.2 SSE-Specific Command Generation
- Real-time event streaming commands
- Interactive controls (pause, resume, filter)
- Reconnection logic with exponential backoff
- Event filtering and routing

### 3.3 Rich Console Output for SSE
```go
// Required libraries:
// github.com/charmbracelet/lipgloss - Modern styling
// github.com/charmbracelet/bubbles - Interactive components  
// github.com/fatih/color - Simple color output
// github.com/jedib0t/go-pretty/v6/table - Tables
// github.com/pterm/pterm - Rich terminal output
```

### 3.4 SSE Command Flags
```bash
--follow              # Keep connection open (default: true)
--reconnect          # Auto-reconnect on failure (default: true)
--event-types        # Filter specific event types
--since             # Resume from event ID/timestamp
--timeout           # Stop after duration
--max-events        # Stop after N events
--buffer-size       # Read buffer size
--output table|json|yaml|raw|pretty
--quiet             # Only event data
```

## 4. Smart Command Naming System

### 4.1 Multiple Naming Strategies
```go
type NamingStrategy interface {
    GenerateCommandName(op *Operation) string
    GenerateSubcommandName(op *Operation, parent string) string
}

// Strategies:
// - RESTSemantic: GET /users -> list, GET /users/123 -> get
// - OperationID: Use operationId directly
// - PathBased: Derive from URL path segments
// - Custom: User-defined mappings
```

### 4.2 Sub-resource Detection
- Automatic detection of nested resources
- Smart naming for sub-operations
- Parent-child command relationships
- Consistent naming patterns

### 4.3 Command Deduplication
- Detect and handle duplicate command names
- Merge similar operations intelligently
- Warning system for conflicts
- Alternative naming suggestions

## 5. Advanced Flag System

### 5.1 Smart Flag Generation
```go
type FlagConfig struct {
    Name         string
    Type         string  
    Default      interface{}
    Description  string
    Required     bool
    Enum         []string
    Validation   []string
    Aliases      []string
    Hidden       bool
    Deprecated   bool
}
```

### 5.2 Enhanced Flag Types
- Enum validation with autocomplete
- File path flags with validation
- Duration flags with parsing
- URL flags with validation
- Array flags with proper handling
- Object flags with JSON/YAML input

### 5.3 Conditional Flags
- Dependent flags (if A then B required)
- Mutually exclusive flags
- Context-sensitive flags
- Dynamic flag generation

## 6. Rich Console Experience

### 6.1 Output Formats
```go
type OutputFormatter interface {
    Format(data interface{}) error
    SupportsStreaming() bool
}

// Formats:
// - JSON (default)
// - YAML  
// - Table (with pagination)
// - CSV
// - Raw text
// - Custom templates
```

### 6.2 Interactive Features
```go
// Use github.com/AlecAivazis/survey/v2 for prompts
type InteractivePrompt interface {
    PromptMissingRequired(flags []Flag) error
    ConfirmDangerous(operation string) (bool, error)
    SelectFromEnum(options []string) (string, error)
}
```

### 6.3 Progress & Status
```go
// Use github.com/schollz/progressbar/v3
// Use github.com/briandowns/spinner  
type ProgressReporter interface {
    StartOperation(name string)
    UpdateProgress(current, total int64)
    Finish(success bool, message string)
}
```

## 7. Error Handling & Resilience

### 7.1 Structured Error Response
```go
type APIError struct {
    StatusCode int                    `json:"status_code"`
    Message    string                 `json:"message"`
    Details    map[string]interface{} `json:"details,omitempty"`
    RequestID  string                 `json:"request_id,omitempty"`
    Timestamp  time.Time              `json:"timestamp"`
}
```

### 7.2 Retry Logic
- Configurable retry policies
- Exponential backoff with jitter
- Circuit breaker pattern
- Timeout handling

### 7.3 User-Friendly Error Messages
- HTTP status code explanations
- Suggestion system for common errors
- Link to documentation
- Debug mode for detailed errors

## 8. Authentication & Security

### 8.1 Multiple Auth Methods
```go
type AuthProvider interface {
    Authenticate(req *http.Request) error
    RefreshToken() error
    IsExpired() bool
}

// Support:
// - API Keys (header/query)
// - Bearer tokens
// - OAuth2 flows
// - Basic auth
// - Custom auth schemes
```

### 8.2 Credential Management
- Secure credential storage
- Multiple profiles/environments
- Token refresh automation
- Credential validation

## 9. Testing & Quality Assurance

### 9.1 Comprehensive Test Suite
```go
// Test categories:
// - Unit tests for each naming strategy
// - Integration tests with real OpenAPI specs  
// - Snapshot tests for generated code
// - Performance benchmarks
// - SSE streaming tests
// - Error handling tests
```

### 9.2 Code Quality
- Generated code validation (syntax check)
- Linting with golangci-lint
- Formatting with gofmt
- Import organization

### 9.3 Spec Validation
- OpenAPI 3.0/3.1 compliance checking
- Warning for unsupported features
- Best practices recommendations
- Breaking change detection

## 10. Developer Experience

### 10.1 Development Tools
```bash
# Watch mode for rapid iteration
meibel-cli generate --watch --spec api.json

# Dry run to preview changes
meibel-cli generate --dry-run --spec api.json

# Debug mode with verbose output
meibel-cli generate --debug --spec api.json

# Diff mode to see changes
meibel-cli generate --diff --spec api.json
```

### 10.2 Documentation Generation
- Auto-generate command help from OpenAPI descriptions
- Man page generation
- Markdown documentation export
- Usage examples from OpenAPI

### 10.3 Plugin Architecture
```go
type Plugin interface {
    Name() string
    PreGenerate(spec *openapi3.T) error
    PostGenerate(commands []Command) error
    CustomizeCommand(cmd *Command) error
}
```

## 11. Performance & Scalability

### 11.1 Incremental Generation
- Change detection in OpenAPI specs
- Only regenerate modified endpoints
- Dependency tracking
- Build caching

### 11.2 Memory Efficiency
- Streaming parser for large specs
- Lazy loading of schemas
- Memory pooling for large operations
- Garbage collection optimization

### 11.3 Parallel Processing
- Concurrent command generation
- Parallel template rendering
- Async I/O operations
- Worker pool pattern

## 12. Advanced Features

### 12.1 Mock & Stub Generation
- Generate mock servers for testing
- Stub implementations for development
- Request/response examples
- Integration with testing frameworks

### 12.2 API Discovery
- Auto-discovery of OpenAPI endpoints
- Version detection and handling
- Schema evolution tracking
- Backward compatibility checking

### 12.3 Analytics & Metrics
- Command usage statistics
- Performance metrics collection
- Error rate tracking
- User behavior insights

## Implementation Priority

### Phase 1 (High Priority)
1. Modular architecture refactoring
2. External template system
3. Basic SSE support
4. Rich console output
5. Configuration system

### Phase 2 (Medium Priority)
1. Advanced type system
2. Interactive features
3. Better error handling
4. Authentication support
5. Testing framework

### Phase 3 (Low Priority)
1. Plugin architecture
2. Performance optimizations
3. Advanced features
4. Analytics
5. Documentation generation

## Required Dependencies

```go
// Core libraries
github.com/getkin/kin-openapi/openapi3
github.com/spf13/cobra
github.com/spf13/viper

// Rich console output
github.com/charmbracelet/lipgloss
github.com/charmbracelet/bubbles
github.com/fatih/color
github.com/jedib0t/go-pretty/v6/table
github.com/pterm/pterm

// Interactive features
github.com/AlecAivazis/survey/v2
github.com/schollz/progressbar/v3
github.com/briandowns/spinner

// Utilities
github.com/stretchr/testify
gopkg.in/yaml.v3
```

## Success Metrics

- **Code Quality**: Generated code passes all lints and builds successfully
- **User Experience**: Intuitive command structure and helpful error messages
- **Performance**: Sub-second generation for typical API specs
- **Maintainability**: Clean, modular codebase with >90% test coverage
- **Feature Completeness**: Support for all major OpenAPI 3.0+ features
- **SSE Support**: Real-time streaming with rich console output