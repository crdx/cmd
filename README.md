# cmd

A collection of command-line tools.

## Tools

### [unbuffer](unbuffer)

```bash
mise use go:crdx.org/cmd/unbuffer
go install crdx.org/cmd/unbuffer@latest
```

Run a command with its output connected to a pseudo-terminal, forcing programs that check for a TTY to produce unbuffered, coloured output.

```
unbuffer <command> [args...]
```
