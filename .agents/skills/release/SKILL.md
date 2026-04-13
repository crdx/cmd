---
name: release
description: Release a new version of cmd. Use when the user asks to release, deploy, cut a release, or bump the version.
---

# Release

## Overview

The repo contains multiple Go binaries in subdirectories. A single semver tag versions the whole repo. A GitHub Action creates the release and warms the Go proxy. Users install individual binaries via `mise use go:crdx.org/cmd/<name>`.

## Workflow

### 1. Identify changes

Compare HEAD to the last tag to see what changed:

```bash
git log --oneline $(git describe --tags --abbrev=0)..HEAD
```

Read the diffs to understand which binaries were affected and what the changes are.

### 2. Suggest version number

Based on the changes, suggest the next version using semver:

- **patch** — bug fixes, minor tweaks
- **minor** — new features, new binaries
- **major** — breaking changes

Use the `ask` tool to confirm the version number with the user. Do NOT proceed until confirmed.

### 3. Update CHANGELOG.md

Add a new `## vX.Y.Z` section at the top of the changelog (below `# Changelog`). Format:

```markdown
## vX.Y.Z

- binary: description of change
- binary: another change
```

Group changes by binary. Every entry must be prefixed with the binary it affects (etied to a specific binary, use `meta:` as the prefix. Keep descriptions concise.

After editing, use the `ask` tool to check with the user:

- "Yes, looks good" → proceed to step 4.
- "No, I made changes" → re-read CHANGELOG.md to pick up the user's edits, then proceed to step 4.

### 4. Commit and hand off

Give the user the commands to commit, tag and push:

```bash
git add -A && git commit -m 'Release vX.Y.Z' && \
    semver minor && git push && git push --tags && \
    mise cache clear
```

Adjust `major`/`minor`/`patch` to match the agreed version. If a specific version is needed, use `semver X.Y.Z` (without `v` prefix).

**Stop here.** The user will run the commands above and confirm when done.

### 5. Verify (after user confirms push is done)

The GitHub Action handles creating the release and warming the Go proxy. The release is not done until the new version is confirmed visible on the Go module proxy. Check it directly:

```bash
curl -sf https://proxy.golang.org/crdx.org/cmd/@v/vX.Y.Z.info
```

If it returns JSON with the version, the release is complete. If it 404s, wait 30 seconds and retry. Keep retrying until it succeeds (use incremental backoff). Only then declare the release done.
