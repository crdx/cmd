// Bash completion: ../cmdctl/completions/chronic.bash

package main

import (
	"bytes"
	"errors"
	"fmt"
	"os"
	"os/exec"
)

const usage = `Usage:
    chronic [options] command [args...]

Run a command quietly unless it fails. Standard output and standard error are
buffered and only shown if the command exits non-zero or is killed by a signal.

Options:
    -e, --stderr     Also trigger output if the command exits zero but wrote to stderr
    -v, --verbose    Label STDOUT/STDERR and report RETVAL when output is shown
`

func main() {
	triggerOnStderr := false
	verbose := false

	// TODO(x): This parser stops at the first non-option so the wrapped command keeps its own
	// flags. This will be replaced with crdx.org/duckopt once it supports native double-dash-free
	// arg pass-through.
	args := os.Args[1:]
	index := 0
	for index < len(args) {
		arg := args[index]
		if arg == "--" {
			index++
			break
		}
		if len(arg) > 2 && arg[0] == '-' && arg[1] == '-' {
			switch arg {
			case "--stderr":
				triggerOnStderr = true
			case "--verbose":
				verbose = true
			case "--help":
				fmt.Print(usage)
				os.Exit(0)
			default:
				fmt.Fprintf(os.Stderr, "Unknown option: %s\n", arg)
				os.Exit(255)
			}
			index++
			continue
		}
		if len(arg) < 2 || arg[0] != '-' {
			break
		}
		for _, option := range arg[1:] {
			switch option {
			case 'e':
				triggerOnStderr = true
			case 'v':
				verbose = true
			case 'h':
				fmt.Print(usage)
				os.Exit(0)
			default:
				fmt.Fprintf(os.Stderr, "Unknown option: %c\n", option)
				os.Exit(255)
			}
		}
		index++
	}

	command := args[index:]
	if len(command) == 0 {
		fmt.Fprint(os.Stderr, usage)
		os.Exit(255)
	}

	var stdout, stderr bytes.Buffer

	process := exec.Command(command[0], command[1:]...) //nolint:gosec
	process.Stdin = os.Stdin
	process.Stdout = &stdout
	process.Stderr = &stderr

	exitCode := 0
	wasSignaled := false

	if err := process.Run(); err != nil {
		if exitError, ok := errors.AsType[*exec.ExitError](err); ok {
			exitCode = exitError.ExitCode()
			wasSignaled = exitCode < 0
		} else {
			fmt.Fprintf(os.Stderr, "chronic: %v\n", err)
			os.Exit(127)
		}
	}

	displayCode := max(exitCode, 0)

	switch {
	case exitCode > 0:
		showOutput(stdout.Bytes(), stderr.Bytes(), verbose, displayCode) // bleh
		os.Exit(exitCode)
	case wasSignaled:
		showOutput(stdout.Bytes(), stderr.Bytes(), verbose, displayCode) // bleh
		os.Exit(1)
	case triggerOnStderr && stderr.Len() > 0:
		showOutput(stdout.Bytes(), stderr.Bytes(), verbose, displayCode) // bleh
		os.Exit(2)
	default:
		os.Exit(0)
	}
}

func showOutput(stdout, stderr []byte, verbose bool, retval int) {
	if verbose {
		_, _ = fmt.Fprint(os.Stdout, "STDOUT:\n")
	}
	_, _ = os.Stdout.Write(stdout)
	if verbose {
		_, _ = fmt.Fprint(os.Stdout, "\nSTDERR:\n")
	}
	_, _ = os.Stderr.Write(stderr)
	if verbose {
		_, _ = fmt.Fprintf(os.Stdout, "\nRETVAL: %d\n", retval)
	}
}
