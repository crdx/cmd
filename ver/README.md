# ver

Increment and tag semver versions in a git repository. Supports both two-part (`1.2`) and three-part (`1.2.3`) version schemes.

When run without arguments, shows the latest version tag.

Supports pre-version and post-version hook scripts (`hooks/pre-version` and `hooks/post-version`). A non-zero exit code stops the release process.

## Install

```bash
$ mise use -g go:crdx.org/cmd/ver
$ go install crdx.org/cmd/ver@latest
```

## Usage

```
ver                          # show latest version
ver major [--ref REF]        # increment major version
ver minor [--ref REF]        # increment minor version
ver patch [--ref REF]        # increment patch version
ver <version> [--ref REF]    # tag an explicit version (without 'v' prefix)
```

## Examples

```bash
$ ver
Latest version: v1.11.0

$ ver patch
Old version: v1.11.0
New version: v1.11.1

$ ver minor
Old version: v1.11.0
New version: v1.12.0

$ ver 2.0.0
New version: v2.0.0

$ ver patch --ref abc1234
Old version: v1.11.0
New version: v1.11.1
```
