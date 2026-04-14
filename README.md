# cmd

**cmd** is a go module for misc command-line tools.

_Generally_, you'll find each tool does one thing, and uses the universal medium of text. You know, the whole Unix philosophy deal.

Some of them might seem too basic to be their own program, but when Go builds binaries so rapidly, keeping them separate is not a problem. And by shoving them all in one go module, each tool can be installed individually while code can still be shared between them.

## Install

```bash
mise use -g go:crdx.org/cmd/$NAME
go install crdx.org/cmd/$NAME@latest
```

## Tools

- [uc](uc) — Describe Unicode characters.
- [unbuffer](unbuffer) — Force unbuffered, coloured output from commands.
