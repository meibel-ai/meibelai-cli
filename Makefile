.PHONY: generate build clean test install dev

# Variables
BINARY_NAME=meibel
GENERATE_TOOL=./cmd/generate/main.go
GENERATED_FILE=cmd/generated_commands.go
OPENAPI_URL=https://storage.googleapis.com/meibel-api-docs/api.json

# Default target
all: generate build

# Generate commands from OpenAPI spec
generate:
	@echo "Generating commands from OpenAPI spec..."
	@go run $(GENERATE_TOOL) -spec $(OPENAPI_URL) -output $(GENERATED_FILE)

# Build the CLI binary
build: generate
	@echo "Building $(BINARY_NAME)..."
	@go build -o $(BINARY_NAME) main.go

# Development build (skip generation, use existing generated file)
dev:
	@echo "Building $(BINARY_NAME) (dev mode, no regeneration)..."
	@go build -o $(BINARY_NAME) main.go

# Install the CLI to $GOPATH/bin
install: generate
	@echo "Installing $(BINARY_NAME)..."
	@go install

# Run tests
test:
	@echo "Running tests..."
	@go test ./... -v

# Clean build artifacts
clean:
	@echo "Cleaning..."
	@rm -f $(BINARY_NAME)
	@rm -f $(GENERATED_FILE)

# Update dependencies
deps:
	@echo "Updating dependencies..."
	@go mod tidy
	@go mod download

# Generate and show what changed
diff-generate: 
	@cp $(GENERATED_FILE) $(GENERATED_FILE).bak 2>/dev/null || true
	@$(MAKE) generate
	@diff -u $(GENERATED_FILE).bak $(GENERATED_FILE) || true
	@rm -f $(GENERATED_FILE).bak

# Force regenerate (useful if OpenAPI spec changed)
force-generate:
	@rm -f $(GENERATED_FILE)
	@$(MAKE) generate

# Run the CLI
run: build
	@./$(BINARY_NAME)

# Development workflow - watch for changes and rebuild
watch:
	@echo "Watching for changes..."
	@which reflex > /dev/null || (echo "Installing reflex..." && go install github.com/cespare/reflex@latest)
	@reflex -r '\.go$$' -s -- sh -c 'make dev'

# Help
help:
	@echo "Meibel CLI Makefile"
	@echo ""
	@echo "Usage:"
	@echo "  make              - Generate commands and build the CLI"
	@echo "  make generate     - Generate commands from OpenAPI spec"
	@echo "  make build        - Generate and build the CLI"
	@echo "  make dev          - Build without regenerating (faster for development)"
	@echo "  make install      - Generate and install to GOPATH/bin"
	@echo "  make test         - Run tests"
	@echo "  make clean        - Remove build artifacts"
	@echo "  make deps         - Update dependencies"
	@echo "  make force-generate - Force regenerate commands"
	@echo "  make watch        - Watch for changes and rebuild"
	@echo "  make help         - Show this help message"