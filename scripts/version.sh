#!/bin/bash
# Script to manage version bumping

set -e

# Get current version from git tags
CURRENT_VERSION=$(git describe --tags --abbrev=0 2>/dev/null || echo "v0.0.0")
CURRENT_VERSION=${CURRENT_VERSION#v}

# Parse version components
IFS='.' read -r -a VERSION_PARTS <<< "$CURRENT_VERSION"
MAJOR="${VERSION_PARTS[0]}"
MINOR="${VERSION_PARTS[1]}"
PATCH="${VERSION_PARTS[2]}"

# Function to bump version
bump_version() {
    case "$1" in
        major)
            MAJOR=$((MAJOR + 1))
            MINOR=0
            PATCH=0
            ;;
        minor)
            MINOR=$((MINOR + 1))
            PATCH=0
            ;;
        patch)
            PATCH=$((PATCH + 1))
            ;;
        *)
            echo "Usage: $0 {major|minor|patch}"
            exit 1
            ;;
    esac
    
    NEW_VERSION="v${MAJOR}.${MINOR}.${PATCH}"
    echo "Bumping version from v${CURRENT_VERSION} to ${NEW_VERSION}"
    
    # Update version in code if needed
    if [ -f "internal/version/version.go" ]; then
        cat > internal/version/version.go << EOF
package version

// Version is the current version of the CLI
const Version = "${NEW_VERSION}"

// Commit is the git commit hash (set by build)
var Commit = "unknown"

// BuildDate is the build date (set by build)
var BuildDate = "unknown"
EOF
    fi
    
    # Commit changes
    git add .
    git commit -m "chore: bump version to ${NEW_VERSION}" || true
    
    # Create tag
    git tag -a "${NEW_VERSION}" -m "Release ${NEW_VERSION}"
    
    echo ""
    echo "Version bumped to ${NEW_VERSION}"
    echo "To push the changes and trigger a release:"
    echo "  git push origin main"
    echo "  git push origin ${NEW_VERSION}"
}

# Main
if [ $# -eq 0 ]; then
    echo "Current version: v${CURRENT_VERSION}"
    echo "Usage: $0 {major|minor|patch}"
    exit 0
fi

bump_version "$1"