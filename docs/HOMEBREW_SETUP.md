# Homebrew Distribution Setup

This guide explains how to set up Homebrew distribution for the Meibel CLI.

## Prerequisites

1. A GitHub account with access to create repositories under `meibel-ai`
2. GoReleaser configured (already done)
3. A GitHub Personal Access Token with `repo` scope

## Setup Steps

### 1. Create the Homebrew Tap Repository

Create a new repository named `homebrew-tap` under the `meibel-ai` organization:

```bash
# Create the repository on GitHub
# Repository name: meibel-ai/homebrew-tap
# Make it public
# Initialize with README
```

### 2. Create Initial Tap Structure

Clone the tap repository and create the initial structure:

```bash
git clone https://github.com/meibel-ai/homebrew-tap.git
cd homebrew-tap

# Create Formula directory
mkdir -p Formula

# Create README
cat > README.md << 'EOF'
# Meibel Homebrew Tap

This is the official Homebrew tap for Meibel AI tools.

## Installation

```bash
brew tap meibel-ai/tap
brew install meibel
```

## Available Formulas

- `meibel` - CLI for interacting with Meibel AI API

## Updating

```bash
brew update
brew upgrade meibel
```
EOF

# Create .gitignore
cat > .gitignore << 'EOF'
*.swp
.DS_Store
EOF

# Commit and push
git add .
git commit -m "Initial tap setup"
git push origin main
```

### 3. Generate GitHub Token for GoReleaser

1. Go to GitHub Settings → Developer settings → Personal access tokens
2. Generate a new token with `repo` scope
3. Name it `HOMEBREW_TAP_GITHUB_TOKEN`
4. Copy the token

### 4. Add Token to Repository Secrets

In your `meibel-cli` repository:

1. Go to Settings → Secrets and variables → Actions
2. Add a new repository secret:
   - Name: `HOMEBREW_TAP_GITHUB_TOKEN`
   - Value: [paste your token]

### 5. Test the Setup

Create a test release:

```bash
# In the meibel-cli directory
# Create a tag
git tag -a v0.1.0 -m "Initial release"
git push origin v0.1.0

# Or use goreleaser in snapshot mode
make snapshot
```

### 6. Create Docker Hub Account (Optional)

For Docker distribution:

1. Create account at https://hub.docker.com
2. Create organization: `meibelai`
3. Add Docker Hub credentials to GitHub secrets:
   - `DOCKER_USERNAME`
   - `DOCKER_PASSWORD`

## Release Process

### Automated Release

When you push a tag, the GitHub Action will:

1. Build binaries for all platforms
2. Create .deb, .rpm, .apk packages
3. Update Homebrew formula automatically
4. Push Docker images
5. Create GitHub release with installers

```bash
# Bump version
./scripts/version.sh minor

# Push tag (triggers release)
git push origin main --tags
```

### Manual Testing

Test GoReleaser locally:

```bash
# Dry run (no publishing)
goreleaser release --snapshot --clean

# Check outputs
ls -la dist/
```

## Installation Methods After Setup

Users can install via multiple methods:

### Homebrew (macOS/Linux)
```bash
brew install meibel-ai/tap/meibel
```

### Direct Script
```bash
curl -sfL https://raw.githubusercontent.com/meibel-ai/meibel-cli/main/install.sh | sh
```

### APT (Debian/Ubuntu)
```bash
# Download and install .deb
curl -LO https://github.com/meibel-ai/meibel-cli/releases/latest/download/meibel_linux_amd64.deb
sudo dpkg -i meibel_linux_amd64.deb
```

### YUM (RHEL/CentOS/Fedora)
```bash
# Download and install .rpm
curl -LO https://github.com/meibel-ai/meibel-cli/releases/latest/download/meibel_linux_amd64.rpm
sudo rpm -i meibel_linux_amd64.rpm
```

### Docker
```bash
docker run --rm meibelai/cli:latest --help
```

### Snap
```bash
snap install meibel
```

## Troubleshooting

### Formula Not Updating

Check that:
1. The `HOMEBREW_TAP_GITHUB_TOKEN` has push access to the tap repo
2. The tap repository exists and is accessible
3. GoReleaser logs in the GitHub Action

### Docker Push Failing

Ensure:
1. Docker Hub credentials are correct
2. The `meibelai` organization exists on Docker Hub
3. You have push permissions

### Version Conflicts

If users report version conflicts:
```bash
# Users should run
brew update
brew uninstall meibel
brew install meibel-ai/tap/meibel
```

## Maintenance

### Updating the Formula Manually

If automatic updates fail:

```bash
# In homebrew-tap repo
cd Formula
# Edit meibel.rb with new version and checksums
git add meibel.rb
git commit -m "Update meibel to v1.2.3"
git push
```

### Deprecating Old Versions

Add to formula:
```ruby
deprecate! date: "2025-01-01", because: "is outdated"
```

## Best Practices

1. **Semantic Versioning**: Always use semantic versioning (v1.2.3)
2. **Test Releases**: Use `make snapshot` to test before releasing
3. **Release Notes**: Write clear release notes for users
4. **Breaking Changes**: Announce breaking changes prominently
5. **Deprecation**: Give users time before removing features

## Additional Resources

- [Homebrew Formula Cookbook](https://docs.brew.sh/Formula-Cookbook)
- [GoReleaser Documentation](https://goreleaser.com/documentation/homebrew/)
- [GitHub Actions for Go](https://github.com/actions/setup-go)