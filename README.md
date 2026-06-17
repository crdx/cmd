# cmd

**cmd** is a go module for misc command-line tools.

_Generally_, you'll find each tool does one thing, and uses the universal medium of text. You know, the whole Unix philosophy deal.

Some of them might seem too basic to be their own program, but when Go builds binaries so rapidly, keeping them separate is not a problem. And by shoving them all in one go module, each tool can be installed individually while code can still be shared between them.

## Install

```bash
mise use -g go:crdx.org/cmd/$NAME
go install crdx.org/cmd/$NAME@latest
```

## Completions

Install [cmdctl](cmdctl) and run `cmdctl install` to add the eval line to `~/.bashrc`, or add it manually:

```bash
eval "$(cmdctl generate bash)"
```

Bash completions are then provided for any tools on `PATH`. Completions activate for newly installed tools on the next shell startup.

See the [cmdctl](cmdctl) README for more information about how shell completions are generated.

## Tools

The README in the root directory of each tool has tool-specific documentation.

- [chronic](chronic) — Run a command quietly unless it fails.
- [cmdctl](cmdctl) — Self-referential management tool.
- [uchar](uchar) — Describe Unicode characters.
- [unbuffer](unbuffer) — Force unbuffered, coloured output from commands.
- [ver](ver) — Bump and tag (sem)versions.
