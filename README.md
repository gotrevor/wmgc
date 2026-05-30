# WMGC Website

Western Mass Go Club website - [wmgc.massgo.org](http://wmgc.massgo.org)

## Quick Start

```bash
bin/deploy ma-state-2026/    # Upload a directory
bin/deploy index.html        # Upload specific files
bin/diff                     # Compare local vs remote
bin/pull                     # Download newer files from remote
bin/connect                  # Interactive FTP session
```

All scripts use nix-shell shebangs - `lftp` is fetched automatically.

## Credentials

FTP password is stored in macOS Keychain (service: `wmgc.massgo.org`, account: `wmgcadmin`).

The `bin/wmgc-secret` binary retrieves it with caller verification - only the `bin/` scripts can access it. See `bin/wmgc-secret-src/README.md` for details.

### First-time setup

```bash
# Add the FTP password to Keychain
security add-generic-password -s "wmgc.massgo.org" -a "wmgcadmin" -w "YOUR_PASSWORD"

# Verify it works
bin/wmgc-secret check
```

## Structure

```
index.html              # Main site
style.css               # Shared styles
ma-state-2026/          # 2026 MA State Championship
ma-2025/                # 2025 MA State Championship
ne-open-2024/           # New England Open 2024
ne-open-2025/           # New England Open 2025
bin/                    # Deploy tools (not uploaded)
```

## Deploying

```bash
# Upload a directory (mirrors all files)
bin/deploy ma-state-2026/

# Upload specific files
bin/deploy index.html style.css

# Preview first
bin/deploy --dry-run ma-state-2026/

# Or check what differs
bin/diff
```

## Repository

- **Remote**: `git@github.com:gotrevor/wmgc.git`
- **Branch**: `site`
- **FTP host**: `wmgc.massgo.org`
- **FTP path**: `/wmgc.massgo.org`
- **Images**: Large images are gitignored, live on FTP only. Use `bin/pull` to sync.
