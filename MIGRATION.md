# WMGC Website Migration Guide

## Overview
The wmgc directory contains the Western Mass Go Club website. All website files are committed to the `site` branch in git. The only file requiring manual setup is the FTP credentials file (excluded from git for security).

## Migration Steps

### 1. Clone Repository
```bash
cd ~/personal
git clone -b site git@github.com:gotrevor/wmgc.git
cd wmgc
```

### 2. Create FTP Credentials File

**Instructions for Claude Code:**

When performing this migration, create the `connect` file and prompt Trevor for the FTP password. Use this template:

```bash
# Ask Trevor for the FTP password first using AskUserQuestion
# Then create the file:

cat > ~/personal/wmgc/connect << 'EOF'
#!/usr/bin/env bash

ftp ftp://wmgcadmin:[PASSWORD]@wmgc.massgo.org
EOF

chmod +x ~/personal/wmgc/connect
```

**File Details:**
- **Filename**: `connect`
- **Purpose**: FTP connection script for uploading website updates
- **Format**: Bash script with embedded FTP credentials
- **Connection**: `ftp://wmgcadmin:[PASSWORD]@wmgc.massgo.org`
- **Username**: `wmgcadmin`
- **Password**: Ask Trevor (not stored in git)
- **Permissions**: Must be executable (`chmod +x`)

**Security Note**: This file is in `.gitignore` to prevent credential leakage. Never commit it to the repository.

## Repository Info
- **Remote**: git@github.com:gotrevor/wmgc.git
- **Branch**: `site` (contains all website files)
- **Contents**: HTML files, tournament results, images, CSS

## Future Improvement: Automated Deployment Tool

**Current State**: The `connect` script launches an interactive FTP session requiring manual file selection and upload commands. This is tedious for regular updates (tournament results, images, etc.).

**Goal**: Replace with a CLI utility for automated, selective deployment.

### Suggested Features

**Core Functionality:**
```bash
# Upload specific files
./deploy index.html ne-open-2025.html

# Upload specific directory
./deploy ne-open-2025/

# Upload everything (full sync)
./deploy --all

# Dry run (show what would be uploaded)
./deploy --dry-run index.html

# Upload to specific remote path
./deploy --remote-path /public_html/tournaments/ ma-2025/
```

**Nice-to-Have:**
- Only upload changed files (compare timestamps or checksums)
- Batch upload multiple files efficiently
- Progress indicators for large files
- Verify uploads completed successfully
- Support both individual files and directory recursion

### Implementation Options

**Option 1: lftp (recommended)**
- Modern FTP client with scripting support
- Built-in mirror/sync capabilities
- Can compare timestamps to skip unchanged files
- Handles batching efficiently

**Option 2: curl**
- Available on all systems
- Simple for single-file uploads
- Would need custom logic for directory sync

**Option 3: ncftp**
- Good batch operations
- Less common, may need installation

### Script Structure
```bash
#!/usr/bin/env bash
# deploy - Upload files to wmgc.massgo.org

FTP_USER="wmgcadmin"
FTP_PASS="[from environment or keychain]"
FTP_HOST="wmgc.massgo.org"
FTP_BASE="/public_html"

# Parse arguments, handle --all, --dry-run, etc.
# Use lftp or curl to perform uploads
# Show progress and confirm success
```

### Instructions for Future Claude Code

When Trevor asks to build this tool:
1. Check if `lftp` is installed (`which lftp`), offer to use it if available
2. Fall back to `curl` if lftp not available
3. Support the command patterns shown above
4. Keep FTP credentials secure (environment variable, macOS keychain, or prompt)
5. Test with the OGS image upload use case first
6. Make it robust enough for regular tournament result uploads

## After Migration
1. Test FTP connection: `./connect`
2. Verify tournament pages load correctly in a browser
3. Consider building the automated deployment tool (see above)
