# Meibel CLI Distribution Guide

This document summarizes all the distribution methods configured for the Meibel CLI.

## Distribution Channels

### 1. Homebrew (macOS/Linux)

**Installation:**
```bash
brew tap meibel-ai/tap
brew install meibel
```

**Update:**
```bash
brew upgrade meibel
```

**Requirements:**
- Create `meibel-ai/homebrew-tap` repository
- Add `HOMEBREW_TAP_GITHUB_TOKEN` secret

### 2. Direct Download

**Install Script:**
```bash
curl -sfL https://raw.githubusercontent.com/meibel-ai/meibel-cli/main/install.sh | sh
```

**Manual Download:**
- Visit: https://github.com/meibel-ai/meibel-cli/releases
- Download appropriate binary for your platform
- Extract and move to PATH

### 3. Package Managers

**APT (Debian/Ubuntu):**
```bash
# Download latest .deb
curl -LO https://github.com/meibel-ai/meibel-cli/releases/latest/download/meibel_linux_amd64.deb
sudo dpkg -i meibel_linux_amd64.deb
```

**YUM (RHEL/CentOS/Fedora):**
```bash
# Download latest .rpm  
curl -LO https://github.com/meibel-ai/meibel-cli/releases/latest/download/meibel_linux_amd64.rpm
sudo rpm -i meibel_linux_amd64.rpm
```

**APK (Alpine):**
```bash
# Download latest .apk
curl -LO https://github.com/meibel-ai/meibel-cli/releases/latest/download/meibel_linux_amd64.apk
apk add --allow-untrusted meibel_linux_amd64.apk
```

### 4. Docker

**Run directly:**
```bash
docker run --rm -it meibelai/cli:latest --help
```

**Use in Dockerfile:**
```dockerfile
FROM meibelai/cli:latest
# Your configuration
```

**Requirements:**
- Docker Hub account: `meibelai`
- Add `DOCKER_USERNAME` and `DOCKER_PASSWORD` secrets

### 5. Snap Store

**Installation:**
```bash
snap install meibel
```

**Requirements:**
- Snapcraft account
- Configure snap store credentials

### 6. Go Install

**For Go developers:**
```bash
go install github.com/meibel-ai/meibel-cli@latest
```

## Platform Support

| Platform | Architecture | Package Types |
|----------|-------------|---------------|
| macOS | amd64, arm64 | tar.gz, Homebrew |
| Linux | amd64, arm64, arm, 386 | tar.gz, deb, rpm, apk, snap |
| Windows | amd64, 386 | zip |
| Docker | amd64, arm64 | Multi-arch images |

## Release Process

### Automated Release

1. **Update version:**
   ```bash
   ./scripts/version.sh minor
   ```

2. **Push tag:**
   ```bash
   git push origin main --tags
   ```

3. **GoReleaser handles:**
   - Building all binaries
   - Creating packages (.deb, .rpm, .apk)
   - Updating Homebrew formula
   - Pushing Docker images
   - Creating GitHub release
   - Generating changelog

### Manual Testing

**Test release locally:**
```bash
# Create snapshot (no publishing)
make snapshot

# Check outputs
ls -la dist/
```

## File Locations

After installation, files are typically located at:

**Binary:**
- `/usr/local/bin/meibel` (Homebrew, manual)
- `/usr/bin/meibel` (package managers)

**Completions:**
- Bash: `/usr/share/bash-completion/completions/meibel`
- Zsh: `/usr/share/zsh/vendor-completions/_meibel`
- Fish: `/usr/share/fish/vendor_completions.d/meibel.fish`

**Documentation:**
- `/usr/share/doc/meibel/`
- `/usr/share/man/man1/meibel.1`

**Configuration:**
- `~/.meibel.yaml` (user config)
- `~/.config/meibel/` (config directory)

## Post-Installation

Users should:

1. **Configure authentication:**
   ```bash
   meibel auth login
   ```

2. **Enable completions:**
   ```bash
   # Bash
   source <(meibel completion bash)
   
   # Zsh
   meibel completion zsh > ~/.zsh/completions/_meibel
   ```

3. **Verify installation:**
   ```bash
   meibel version --verbose
   ```

## Maintenance

### Updating GoReleaser Config

Edit `.goreleaser.yml` to:
- Add new platforms
- Change package contents
- Update Docker configuration
- Modify release notes

### Testing Changes

```bash
# Validate configuration
goreleaser check

# Test build
goreleaser build --snapshot --clean
```

### Monitoring Releases

Check:
- GitHub Actions logs
- GoReleaser output
- Homebrew tap updates
- Docker Hub images
- User feedback

## Troubleshooting

### Common Issues

1. **Homebrew formula not updating:**
   - Check GitHub token permissions
   - Verify tap repository access
   - Review GoReleaser logs

2. **Docker push failing:**
   - Verify Docker Hub credentials
   - Check organization permissions
   - Ensure image names are correct

3. **Package installation errors:**
   - Check package dependencies
   - Verify file permissions
   - Test on target platform

### Debug Commands

```bash
# Check installed version
meibel version --verbose

# Verify binary location
which meibel

# Test configuration
meibel config list

# Check completion
meibel completion bash
```

## Future Enhancements

1. **APT/YUM Repositories:**
   - Host dedicated package repositories
   - Automatic updates via package manager

2. **Windows Package Managers:**
   - Chocolatey support
   - Scoop manifest
   - Windows Store

3. **CI/CD Integration:**
   - GitHub Actions
   - GitLab CI
   - Jenkins plugins

4. **Cloud Platforms:**
   - AWS Systems Manager
   - Azure CLI extensions
   - Google Cloud SDK