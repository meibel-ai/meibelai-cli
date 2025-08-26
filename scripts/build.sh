#!/bin/bash
# Build script with version information

set -e

# Get version info
VERSION=$(git describe --tags --always --dirty 2>/dev/null || echo "dev")
COMMIT=$(git rev-parse --short HEAD 2>/dev/null || echo "unknown")
BUILD_DATE=$(date -u +"%Y-%m-%dT%H:%M:%SZ")

# Build flags
LDFLAGS="-X github.com/meibel-ai/meibel-cli/internal/version.Version=${VERSION}"
LDFLAGS="${LDFLAGS} -X github.com/meibel-ai/meibel-cli/internal/version.Commit=${COMMIT}"
LDFLAGS="${LDFLAGS} -X github.com/meibel-ai/meibel-cli/internal/version.BuildDate=${BUILD_DATE}"

echo "Building Meibel CLI..."
echo "  Version: ${VERSION}"
echo "  Commit: ${COMMIT}"
echo "  Build Date: ${BUILD_DATE}"

# Build binary
go build -ldflags "${LDFLAGS}" -o meibel

echo "Build complete!"
echo ""
echo "Run './meibel version --verbose' to see version information"