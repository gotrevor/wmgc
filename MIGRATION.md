# WMGC Website Migration Guide

## Overview

The wmgc directory contains the Western Mass Go Club website. Website files are committed to the `site` branch in git. FTP credentials are stored separately in `~/.secrets/wmgc/`.

## FTP Server Details

- **Host**: `wmgc.massgo.org`
- **Username**: `wmgcadmin`
- **Password**: Stored in `~/.secrets/wmgc/ftp-password`
- **Remote path**: `/wmgc.massgo.org` (not `/public_html`)

## Migration Steps

### 1. Clone Repository

```bash
cd ~/src
git clone -b site git@github.com:gotrevor/wmgc.git
cd wmgc
```

### 2. Set Up FTP Password

```bash
mkdir -p ~/.secrets/wmgc && chmod 700 ~/.secrets/wmgc
echo "YOUR_PASSWORD" > ~/.secrets/wmgc/ftp-password
chmod 600 ~/.secrets/wmgc/ftp-password
```

### 3. Test Connection

```bash
./bin/connect
```

This opens an interactive lftp session. The script uses a nix-shell shebang - lftp is fetched automatically via nix.

## Tools

### `bin/connect` - Interactive FTP

Opens an interactive lftp session for manual exploration/uploads.

```bash
./bin/connect
# Then: ls, cd, put, get, etc.
```

### `bin/deploy` - Automated Uploads

Upload specific files or directories to the site.

```bash
# Upload specific files
./bin/deploy index.html ne-open-2025.html

# Upload a directory
./bin/deploy ne-open-2025/

# Dry run (preview what would upload)
./bin/deploy --dry-run index.html
```

### `bin/pull` - Sync from Remote

Download files from the FTP server that are newer or missing locally.

```bash
# Preview what would download
./bin/pull --dry-run

# Download
./bin/pull
```

### `bin/diff` - Compare Local vs Remote

Show what files differ between local and remote.

```bash
./bin/diff
```

Output shows:
- 📤 Files that would be uploaded (local newer)
- 📥 Files that would be downloaded (remote newer)

## Images

Most images are gitignored to avoid repo bloat. Small essential images (favicon, logos, Go stones) are whitelisted in `.gitignore`.

To get images from the server:
```bash
./bin/pull
```

Large images stay local-only (not committed to git).

## Repository Info

- **Remote**: git@github.com:gotrevor/wmgc.git
- **Branch**: `site` (contains all website files)
- **Contents**: HTML files, tournament results, CSS
- **Images**: On FTP server, gitignored locally

## Known Issues

- Stray `ne-open-2025.html` in FTP root (should be in `/wmgc.massgo.org/`)
