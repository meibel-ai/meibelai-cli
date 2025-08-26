# Development Guide

This guide covers best practices for maintaining and updating the Meibel CLI, particularly when the OpenAPI specification changes.

## Table of Contents

- [Architecture Overview](#architecture-overview)
- [Updating from API Changes](#updating-from-api-changes)
- [Development Workflow](#development-workflow)
- [Testing Strategy](#testing-strategy)
- [CI/CD Pipeline](#cicd-pipeline)
- [Versioning Strategy](#versioning-strategy)
- [Troubleshooting](#troubleshooting)

## Architecture Overview

The CLI follows a modular architecture that makes it easy to update when the API changes:

```
┌─────────────────┐
│  OpenAPI Spec   │ (https://storage.googleapis.com/meibel-api-docs/api.json)
└────────┬────────┘
         │ Fetched/Cached
┌────────▼────────┐
│  OpenAPI Parser │ (internal/openapi)
└────────┬────────┘
         │ Parsed Operations
┌────────▼────────┐
│Command Generator│ (internal/generator)
└────────┬────────┘
         │ Cobra Commands
┌────────▼────────┐
│   CLI Runtime   │ (cmd/root.go)
└─────────────────┘
```

## Updating from API Changes

### Automated Update Process

1. **Set up a scheduled check for API changes:**

```bash
#!/bin/bash
# scripts/check-api-updates.sh

CURRENT_SPEC="api.json"
NEW_SPEC="api-new.json"
SPEC_URL="https://storage.googleapis.com/meibel-api-docs/api.json"

# Download latest spec
curl -s "$SPEC_URL" > "$NEW_SPEC"

# Check if spec has changed
if ! diff -q "$CURRENT_SPEC" "$NEW_SPEC" > /dev/null 2>&1; then
    echo "API specification has changed!"
    
    # Generate changelog
    echo "## API Changes" > api-changes.md
    echo "" >> api-changes.md
    
    # Use jq to extract changes (install: brew install jq)
    echo "### New Endpoints:" >> api-changes.md
    jq -r '.paths | keys[]' "$NEW_SPEC" | sort > new-endpoints.txt
    jq -r '.paths | keys[]' "$CURRENT_SPEC" | sort > old-endpoints.txt
    comm -13 old-endpoints.txt new-endpoints.txt >> api-changes.md
    
    echo "" >> api-changes.md
    echo "### Removed Endpoints:" >> api-changes.md
    comm -23 old-endpoints.txt new-endpoints.txt >> api-changes.md
    
    # Update the spec
    mv "$NEW_SPEC" "$CURRENT_SPEC"
    
    # Trigger rebuild
    make update-cli
else
    echo "No API changes detected"
    rm "$NEW_SPEC"
fi
```

2. **Create a Makefile for common tasks:**

```makefile
# Makefile

.PHONY: update-cli test build release

# Update CLI from latest API spec
update-cli:
	@echo "Updating from latest API spec..."
	@curl -s https://storage.googleapis.com/meibel-api-docs/api.json > api.json
	@go generate ./...
	@go build -o meibel
	@echo "Update complete!"

# Run tests
test:
	@go test -v ./...
	@go test -race ./...

# Build for all platforms
build:
	@echo "Building for multiple platforms..."
	@GOOS=darwin GOARCH=amd64 go build -o dist/meibel-darwin-amd64
	@GOOS=darwin GOARCH=arm64 go build -o dist/meibel-darwin-arm64
	@GOOS=linux GOARCH=amd64 go build -o dist/meibel-linux-amd64
	@GOOS=windows GOARCH=amd64 go build -o dist/meibel-windows-amd64.exe

# Create a new release
release: test build
	@echo "Creating release..."
	@./scripts/create-release.sh

# Check for API updates
check-updates:
	@./scripts/check-api-updates.sh

# Run linters
lint:
	@golangci-lint run
	@go mod tidy

# Generate test coverage
coverage:
	@go test -coverprofile=coverage.out ./...
	@go tool cover -html=coverage.out -o coverage.html
```

### Manual Update Process

When the API specification changes:

1. **Download the latest spec:**
```bash
curl -o api.json https://storage.googleapis.com/meibel-api-docs/api.json
```

2. **Review changes:**
```bash
# Compare with previous version (if using git)
git diff api.json

# Or use a JSON diff tool
jq -S . api.json > api-formatted.json
jq -S . api-old.json > api-old-formatted.json
diff api-old-formatted.json api-formatted.json
```

3. **Test the changes:**
```bash
# Build and test
go build -o meibel
./meibel --help

# Run integration tests
go test -tags=integration ./...
```

4. **Update documentation:**
```bash
# Generate command documentation
./scripts/generate-docs.sh
```

## Development Workflow

### 1. Feature Development

```bash
# Create a feature branch
git checkout -b feature/api-v2-support

# Make changes
# ... edit files ...

# Test locally
go test ./...
go build -o meibel
./meibel --dry-run datasources list

# Commit changes
git add .
git commit -m "feat: add support for API v2 endpoints"
```

### 2. Adding New Features

When adding features beyond API updates:

```go
// Example: Adding a new formatter in internal/formatter/csv.go
package formatter

import (
    "encoding/csv"
    "bytes"
)

type CSVFormatter struct{}

func (f *CSVFormatter) Format(data []byte) error {
    // Implementation
    return nil
}
```

### 3. Handling Breaking Changes

For breaking API changes, implement version detection:

```go
// internal/openapi/version.go
package openapi

func DetectAPIVersion(spec *openapi3.T) string {
    version := spec.Info.Version
    // Parse version and return major version
    return parseMajorVersion(version)
}

// internal/generator/generator.go
func (g *Generator) GenerateCommands() ([]*cobra.Command, error) {
    version := openapi.DetectAPIVersion(g.spec)
    
    switch version {
    case "1":
        return g.generateV1Commands()
    case "2":
        return g.generateV2Commands()
    default:
        return g.generateV2Commands() // Use latest as default
    }
}
```

## Testing Strategy

### 1. Unit Tests

Create comprehensive unit tests for each module:

```go
// internal/openapi/parser_test.go
package openapi_test

import (
    "testing"
    "github.com/stretchr/testify/assert"
    "github.com/meibel-ai/meibel-cli/internal/openapi"
)

func TestLoadSpecFromFile(t *testing.T) {
    spec, err := openapi.LoadSpecFromFile("testdata/simple-api.json")
    assert.NoError(t, err)
    assert.NotNil(t, spec)
    assert.Equal(t, "0.1.0", spec.Info.Version)
}
```

### 2. Integration Tests

Test against a mock server:

```go
// internal/client/client_integration_test.go
// +build integration

package client_test

import (
    "net/http/httptest"
    "testing"
)

func TestClientIntegration(t *testing.T) {
    server := httptest.NewServer(mockHandler())
    defer server.Close()
    
    // Test against mock server
}
```

### 3. End-to-End Tests

```bash
#!/bin/bash
# scripts/e2e-test.sh

# Start a test server
docker run -d -p 8080:8080 meibel/api-mock

# Run CLI commands
./meibel --server http://localhost:8080 auth login --api-key test-key
./meibel --server http://localhost:8080 datasources list

# Cleanup
docker stop $(docker ps -q --filter ancestor=meibel/api-mock)
```

## CI/CD Pipeline

### GitHub Actions Workflow

```yaml
# .github/workflows/ci.yml
name: CI

on:
  push:
    branches: [ main ]
  pull_request:
    branches: [ main ]
  schedule:
    # Check for API updates daily
    - cron: '0 0 * * *'

jobs:
  test:
    runs-on: ubuntu-latest
    steps:
    - uses: actions/checkout@v3
    
    - name: Set up Go
      uses: actions/setup-go@v4
      with:
        go-version: 1.21
    
    - name: Cache Go modules
      uses: actions/cache@v3
      with:
        path: ~/go/pkg/mod
        key: ${{ runner.os }}-go-${{ hashFiles('**/go.sum') }}
    
    - name: Download dependencies
      run: go mod download
    
    - name: Run tests
      run: go test -v ./...
    
    - name: Run linter
      uses: golangci/golangci-lint-action@v3
    
    - name: Build
      run: go build -o meibel

  check-api-updates:
    runs-on: ubuntu-latest
    if: github.event_name == 'schedule'
    steps:
    - uses: actions/checkout@v3
    
    - name: Check for API updates
      run: |
        ./scripts/check-api-updates.sh
        if [ -f api-changes.md ]; then
          echo "API_CHANGED=true" >> $GITHUB_ENV
        fi
    
    - name: Create PR for API updates
      if: env.API_CHANGED == 'true'
      uses: peter-evans/create-pull-request@v5
      with:
        commit-message: 'chore: update API specification'
        title: 'API Specification Update'
        body-path: api-changes.md
        branch: api-update-${{ github.run_id }}
```

### Release Workflow

```yaml
# .github/workflows/release.yml
name: Release

on:
  push:
    tags:
      - 'v*'

jobs:
  release:
    runs-on: ubuntu-latest
    steps:
    - uses: actions/checkout@v3
    
    - name: Set up Go
      uses: actions/setup-go@v4
      with:
        go-version: 1.21
    
    - name: Build releases
      run: make build
    
    - name: Create Release
      uses: softprops/action-gh-release@v1
      with:
        files: dist/*
        generate_release_notes: true
```

## Versioning Strategy

### Semantic Versioning

Follow semantic versioning for the CLI:

- **MAJOR**: Breaking changes in CLI interface
- **MINOR**: New features, new API endpoints support
- **PATCH**: Bug fixes, documentation updates

### Version Compatibility Matrix

Maintain a compatibility matrix:

| CLI Version | API Version | Notes |
|-------------|-------------|-------|
| v1.0.x      | 0.1.0       | Initial release |
| v1.1.x      | 0.2.0       | Added RAG endpoints |
| v2.0.x      | 1.0.0       | Breaking changes in API |

### Handling Multiple API Versions

```go
// internal/config/api_version.go
package config

type APIVersion struct {
    Version string
    BaseURL string
}

var APIVersions = map[string]APIVersion{
    "v1": {
        Version: "0.1.0",
        BaseURL: "http://api.meibel.ai",
    },
    "v2": {
        Version: "1.0.0",
        BaseURL: "http://api.v2.meibel.ai",
    },
}
```

## Troubleshooting

### Common Issues

1. **Spec Validation Errors**
   ```bash
   # Skip validation temporarily
   # Edit internal/openapi/parser.go to comment out validation
   ```

2. **Command Name Conflicts**
   ```go
   // Add conflict resolution in generator
   func (g *Generator) resolveCommandConflicts(commands map[string]*cobra.Command) {
       // Implementation
   }
   ```

3. **Authentication Issues**
   ```bash
   # Debug authentication
   MEIBEL_DEBUG=true ./meibel auth status
   ```

### Debug Mode

Enable debug logging:

```go
// internal/debug/debug.go
package debug

import (
    "log"
    "os"
)

func Log(format string, args ...interface{}) {
    if os.Getenv("MEIBEL_DEBUG") == "true" {
        log.Printf("[DEBUG] "+format, args...)
    }
}
```

### Performance Profiling

```bash
# CPU profiling
go test -cpuprofile=cpu.prof -bench=.
go tool pprof cpu.prof

# Memory profiling
go test -memprofile=mem.prof -bench=.
go tool pprof mem.prof
```

## Best Practices

1. **Always maintain backward compatibility** when possible
2. **Document all breaking changes** in CHANGELOG.md
3. **Keep the OpenAPI spec cached locally** for offline development
4. **Use feature flags** for experimental features
5. **Implement graceful degradation** for older API versions
6. **Add comprehensive logging** for debugging
7. **Monitor API deprecation notices** in responses
8. **Maintain integration tests** against real API
9. **Version lock dependencies** for reproducible builds
10. **Automate as much as possible** to reduce human error

## Release Process

### Overview

The Meibel CLI uses GoReleaser for automated releases across multiple platforms. When you push a tag, the release process automatically:

- Builds binaries for all platforms (macOS, Linux, Windows)
- Creates Linux packages (.deb, .rpm, .apk)
- Updates the Homebrew formula
- Pushes Docker images
- Publishes to Snap store
- Generates changelog
- Creates GitHub release with all assets

### Prerequisites

1. **Create Homebrew Tap Repository**
   - Create `meibel-ai/homebrew-tap` repository on GitHub
   - Make it public
   - See [HOMEBREW_SETUP.md](docs/HOMEBREW_SETUP.md) for detailed setup

2. **Configure GitHub Secrets**
   - `HOMEBREW_TAP_GITHUB_TOKEN`: Personal access token with repo scope
   - `DOCKER_USERNAME`: Docker Hub username (optional)
   - `DOCKER_PASSWORD`: Docker Hub password (optional)

3. **Install GoReleaser**
   ```bash
   brew install goreleaser
   ```

### Creating a Release

#### 1. Update Version

```bash
# Bump patch version (0.1.0 → 0.1.1)
./scripts/version.sh patch

# Bump minor version (0.1.0 → 0.2.0)
./scripts/version.sh minor

# Bump major version (0.1.0 → 1.0.0)
./scripts/version.sh major
```

#### 2. Test Release Locally

```bash
# Create a snapshot release (doesn't publish)
make snapshot

# Check the output
ls -la dist/
```

#### 3. Create Official Release

```bash
# Push the tag to trigger release
git push origin main --tags
```

The GitHub Action will automatically run GoReleaser, which will:
1. Build binaries for all platforms
2. Create archives and packages
3. Update Homebrew formula
4. Push Docker images
5. Create GitHub release

### Distribution Channels

After release, users can install via:

#### Homebrew (macOS/Linux)
```bash
brew install meibel-ai/tap/meibel
```

#### Direct Installation Script
```bash
curl -sfL https://raw.githubusercontent.com/meibel-ai/meibel-cli/main/install.sh | sh
```

#### Package Managers

**Debian/Ubuntu (APT):**
```bash
curl -LO https://github.com/meibel-ai/meibel-cli/releases/latest/download/meibel_$(lsb_release -cs)_amd64.deb
sudo dpkg -i meibel_*.deb
```

**RHEL/Fedora (YUM):**
```bash
curl -LO https://github.com/meibel-ai/meibel-cli/releases/latest/download/meibel_linux_amd64.rpm
sudo rpm -i meibel_*.rpm
```

#### Docker
```bash
docker run --rm meibelai/cli:latest --help
```

#### Snap Store
```bash
snap install meibel
```

### Platform Support Matrix

| Platform | Architectures | Package Formats |
|----------|--------------|-----------------|
| macOS    | amd64, arm64 | .tar.gz, Homebrew |
| Linux    | amd64, arm64, arm, 386 | .tar.gz, .deb, .rpm, .apk, Snap |
| Windows  | amd64, 386   | .zip |
| Docker   | amd64, arm64 | Multi-arch images |

### Release Configuration

The release process is configured in:

- `.goreleaser.yml`: Main GoReleaser configuration
- `.github/workflows/release.yml`: GitHub Actions workflow
- `scripts/version.sh`: Version bumping script
- `install.sh`: Universal installation script

### Troubleshooting Releases

#### GoReleaser Validation
```bash
# Check configuration
goreleaser check

# Test build without publishing
goreleaser build --snapshot --clean
```

#### Common Issues

1. **Homebrew Formula Not Updating**
   - Verify `HOMEBREW_TAP_GITHUB_TOKEN` has push access
   - Check GoReleaser logs in GitHub Actions
   - Manually check the tap repository

2. **Docker Push Failures**
   - Verify Docker Hub credentials
   - Ensure `meibelai` organization exists
   - Check Docker Hub rate limits

3. **Package Building Errors**
   - Review GoReleaser output
   - Test locally with `make snapshot`
   - Check platform-specific requirements

### Manual Release (Emergency)

If automated release fails:

```bash
# Build manually
make build-all

# Create GitHub release manually
gh release create v1.2.3 dist/* --title "v1.2.3" --notes "Release notes"

# Update Homebrew formula manually
# Edit Formula/meibel.rb in homebrew-tap repo
```

### Post-Release Checklist

- [ ] Verify GitHub release is created
- [ ] Test Homebrew installation: `brew install meibel-ai/tap/meibel`
- [ ] Test install script: `curl -sfL .../install.sh | sh`
- [ ] Check Docker images: `docker pull meibelai/cli:latest`
- [ ] Update documentation if needed
- [ ] Announce release (if applicable)

## Contributing

See [CONTRIBUTING.md](CONTRIBUTING.md) for guidelines on contributing to the project.