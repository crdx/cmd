# cmdctl

Generate a bash completion script for the currently installed tools.

Run `cmdctl generate bash` to generate a script containing completion functions for all the installed tools, concatenated together. Qualifying tools are found by searching PATH for binaries that match by name, and checking that they're Go binaries with the correct Go package name.

Run `cmdctl install` to add the eval line to `~/.bashrc`, or add it manually:

```bash
eval "$(cmdctl generate bash)"
```

## Install

```bash
mise use -g go:crdx.org/cmd/cmdctl
go install crdx.org/cmd/cmdctl@latest
```

## Usage

```
Usage:
    cmdctl generate bash
    cmdctl install bash

Commands:
    generate bash    Generate bash completions for crdx.org/cmd tools on PATH
    install bash       Add eval line to ~/.bashrc
```
