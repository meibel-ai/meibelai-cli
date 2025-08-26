#!/bin/sh
# Meibel CLI Installation Script
# This script downloads and installs the Meibel CLI

set -e

# Default values
BINARY_NAME="meibel"
GITHUB_REPO="meibel-ai/meibel-cli"
INSTALL_DIR="/usr/local/bin"
VERSION=""

# Colors for output
RED='\033[0;31m'
GREEN='\033[0;32m'
YELLOW='\033[1;33m'
NC='\033[0m' # No Color

# Functions
info() {
    printf "${GREEN}[INFO]${NC} %s\n" "$1"
}

error() {
    printf "${RED}[ERROR]${NC} %s\n" "$1" >&2
    exit 1
}

warning() {
    printf "${YELLOW}[WARNING]${NC} %s\n" "$1"
}

# Detect OS and architecture
detect_platform() {
    OS=$(uname -s | tr '[:upper:]' '[:lower:]')
    ARCH=$(uname -m)
    
    case $OS in
        darwin)
            OS="Darwin"
            ;;
        linux)
            OS="Linux"
            ;;
        *)
            error "Unsupported operating system: $OS"
            ;;
    esac
    
    case $ARCH in
        x86_64|amd64)
            ARCH="x86_64"
            ;;
        aarch64|arm64)
            ARCH="arm64"
            ;;
        armv7l|armv7)
            ARCH="armv7"
            ;;
        i386|i686)
            ARCH="i386"
            ;;
        *)
            error "Unsupported architecture: $ARCH"
            ;;
    esac
    
    PLATFORM="${OS}_${ARCH}"
    info "Detected platform: $PLATFORM"
}

# Get the latest version from GitHub
get_latest_version() {
    if [ -z "$VERSION" ]; then
        info "Fetching latest version..."
        VERSION=$(curl -sfL "https://api.github.com/repos/${GITHUB_REPO}/releases/latest" | grep '"tag_name":' | sed -E 's/.*"([^"]+)".*/\1/')
        if [ -z "$VERSION" ]; then
            error "Failed to fetch latest version"
        fi
    fi
    info "Installing version: $VERSION"
}

# Download the binary
download_binary() {
    DOWNLOAD_URL="https://github.com/${GITHUB_REPO}/releases/download/${VERSION}/${BINARY_NAME}_${PLATFORM}.tar.gz"
    TEMP_DIR=$(mktemp -d)
    ARCHIVE_PATH="${TEMP_DIR}/${BINARY_NAME}.tar.gz"
    
    info "Downloading from: $DOWNLOAD_URL"
    
    if ! curl -sfL "$DOWNLOAD_URL" -o "$ARCHIVE_PATH"; then
        error "Failed to download binary"
    fi
    
    info "Extracting archive..."
    tar -xzf "$ARCHIVE_PATH" -C "$TEMP_DIR"
    
    if [ ! -f "${TEMP_DIR}/${BINARY_NAME}" ]; then
        error "Binary not found in archive"
    fi
    
    BINARY_PATH="${TEMP_DIR}/${BINARY_NAME}"
}

# Verify checksum
verify_checksum() {
    info "Verifying checksum..."
    CHECKSUMS_URL="https://github.com/${GITHUB_REPO}/releases/download/${VERSION}/checksums.txt"
    CHECKSUMS_PATH="${TEMP_DIR}/checksums.txt"
    
    if ! curl -sfL "$CHECKSUMS_URL" -o "$CHECKSUMS_PATH"; then
        warning "Failed to download checksums, skipping verification"
        return
    fi
    
    # Extract expected checksum
    EXPECTED=$(grep "${BINARY_NAME}_${PLATFORM}.tar.gz" "$CHECKSUMS_PATH" | cut -d' ' -f1)
    if [ -z "$EXPECTED" ]; then
        warning "Checksum not found for this platform, skipping verification"
        return
    fi
    
    # Calculate actual checksum
    if command -v sha256sum >/dev/null 2>&1; then
        ACTUAL=$(sha256sum "$ARCHIVE_PATH" | cut -d' ' -f1)
    elif command -v shasum >/dev/null 2>&1; then
        ACTUAL=$(shasum -a 256 "$ARCHIVE_PATH" | cut -d' ' -f1)
    else
        warning "No checksum utility found, skipping verification"
        return
    fi
    
    if [ "$EXPECTED" != "$ACTUAL" ]; then
        error "Checksum verification failed"
    fi
    
    info "Checksum verified successfully"
}

# Install the binary
install_binary() {
    # Check if we need sudo
    if [ -w "$INSTALL_DIR" ]; then
        SUDO=""
    else
        SUDO="sudo"
        info "Root access required to install to $INSTALL_DIR"
    fi
    
    info "Installing $BINARY_NAME to $INSTALL_DIR"
    
    # Make binary executable
    chmod +x "$BINARY_PATH"
    
    # Install binary
    $SUDO mv "$BINARY_PATH" "$INSTALL_DIR/$BINARY_NAME"
    
    # Clean up
    rm -rf "$TEMP_DIR"
    
    info "Installation complete!"
}

# Verify installation
verify_installation() {
    if ! command -v "$BINARY_NAME" >/dev/null 2>&1; then
        warning "$BINARY_NAME not found in PATH"
        warning "Add $INSTALL_DIR to your PATH or specify a different install directory"
        return
    fi
    
    VERSION_OUTPUT=$("$BINARY_NAME" version 2>&1)
    info "Successfully installed: $VERSION_OUTPUT"
    
    echo ""
    info "Get started with:"
    echo "  ${BINARY_NAME} --help"
    echo "  ${BINARY_NAME} auth login"
}

# Parse command line arguments
while [ $# -gt 0 ]; do
    case $1 in
        -v|--version)
            VERSION="$2"
            shift 2
            ;;
        -d|--dir)
            INSTALL_DIR="$2"
            shift 2
            ;;
        -h|--help)
            echo "Meibel CLI Installation Script"
            echo ""
            echo "Usage: $0 [options]"
            echo ""
            echo "Options:"
            echo "  -v, --version VERSION    Install specific version (default: latest)"
            echo "  -d, --dir DIRECTORY      Install directory (default: /usr/local/bin)"
            echo "  -h, --help              Show this help message"
            echo ""
            echo "Examples:"
            echo "  # Install latest version"
            echo "  curl -sfL https://install.meibel.ai | sh"
            echo ""
            echo "  # Install specific version"
            echo "  curl -sfL https://install.meibel.ai | sh -s -- -v v1.0.0"
            echo ""
            echo "  # Install to custom directory"
            echo "  curl -sfL https://install.meibel.ai | sh -s -- -d ~/.local/bin"
            exit 0
            ;;
        *)
            error "Unknown option: $1"
            ;;
    esac
done

# Main installation flow
main() {
    info "Starting Meibel CLI installation..."
    
    detect_platform
    get_latest_version
    download_binary
    verify_checksum
    install_binary
    verify_installation
    
    echo ""
    info "For more information, visit: https://github.com/${GITHUB_REPO}"
}

# Run main function
main