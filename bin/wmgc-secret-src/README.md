# wmgc-secret

Keychain helper for WMGC FTP credentials with caller verification.

## What it does

Retrieves the FTP password from macOS Keychain, but only when called by an authorized script. Prevents accidental credential leakage (e.g., from Claude Code or random shell commands).

## How it works

1. Walks the process tree via `ps`
2. Checks if any ancestor process cmdline contains an allowed script name (`bin/deploy`, `bin/connect`, `bin/diff`, `bin/pull`)
3. If allowed, reads from Keychain (service: `wmgc.massgo.org`, account: `wmgcadmin`)
4. Prints password to stdout (no newline) for shell substitution

## Commands

```bash
wmgc-secret get      # Print password (silent failure if unauthorized)
wmgc-secret check    # Verify access (prints diagnostic on failure)
```

On failure, both commands now print the process tree and allowed scripts list to stderr for debugging.

## Building

```bash
./build              # Uses nix develop + go build
```

Requires: Go, macOS Security framework (provided via nix flake).

## Keychain setup

```bash
# Add password
security add-generic-password -s "wmgc.massgo.org" -a "wmgcadmin" -w "YOUR_PASSWORD"

# Verify
security find-generic-password -s "wmgc.massgo.org" -a "wmgcadmin"
```

## Known issues

- Nix-shell shebangs strip absolute paths, so caller matching uses relative substrings (`bin/deploy` not full path)
- The `apple-sdk_15` nix package is required for the Security framework on modern nixpkgs
