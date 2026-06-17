# chronic

Run a command quietly unless it fails. Standard output and standard error are buffered and only displayed if the command fails (i.e. exits non-zero or is killed). This is useful for cron jobs where you only want output if something goes wrong.

This is a Go port of [moreutils](https://joeyh.name/code/moreutils/) `chronic`, with long options added.

## Install

```bash
mise use -g go:crdx.org/cmd/chronic
go install crdx.org/cmd/chronic@latest
```

## Usage

```
Usage:
    chronic [options] command [args...]

Run a command quietly unless it fails. Standard output and standard error are
buffered and only shown if the command exits non-zero or is killed by a signal.

Options:
    -e, --stderr     Also trigger output if the command exits zero but wrote to stderr
    -v, --verbose    Label STDOUT/STDERR and report RETVAL when output is shown
```

By default nothing is shown unless the command fails. If invoked with the `-e`/`--stderr` argument, output is also shown when the command exits zero but wrote to standard error, and in that case exits with status 2.

If invoked with the `-v`/`--verbose` argument, the replayed output is labelled with `STDOUT` and `STDERR` headers and followed by a `RETVAL` line reporting the exit status.

Option parsing stops at the first non-option argument, so flags belonging to the wrapped command are passed through untouched. A `--` terminator also ends option parsing.
