package sh

import (
	"errors"
	"io"
	"os"
	"os/exec"
	"strings"

	"github.com/samber/lo"
)

// Command is a fluent builder for executing external commands.
type Command struct {
	cmd *exec.Cmd
}

// Cmd creates a new command builder.
func Cmd(cmd ...string) *Command {
	return &Command{cmd: exec.Command(cmd[0], cmd[1:]...)} //nolint:gosec // args are trusted
}

// Dir sets the working directory for the command.
func (self *Command) Dir(dir string) *Command {
	self.cmd.Dir = dir
	return self
}

// Env sets the environment for the command.
func (self *Command) Env(env []string) *Command {
	self.cmd.Env = env
	return self
}

// Stdin sets the stdin reader for the command.
func (self *Command) Stdin(r io.Reader) *Command {
	self.cmd.Stdin = r
	return self
}

// Stdout sets the stdout writer for the command.
func (self *Command) Stdout(w io.Writer) *Command {
	self.cmd.Stdout = w
	return self
}

// Stderr sets the stderr writer for the command.
func (self *Command) Stderr(w io.Writer) *Command {
	self.cmd.Stderr = w
	return self
}

// Interactive connects stdin, stdout, and stderr to os.Stdin, os.Stdout, and os.Stderr.
func (self *Command) Interactive() *Command {
	self.cmd.Stdin = os.Stdin
	self.cmd.Stdout = os.Stdout
	self.cmd.Stderr = os.Stderr
	return self
}

// Passthrough connects stdout and stderr to os.Stdout and os.Stderr.
func (self *Command) Passthrough() *Command {
	self.cmd.Stdout = os.Stdout
	self.cmd.Stderr = os.Stderr
	return self
}

// Quiet connects stdin and stderr to os.Stdin and os.Stderr, but not stdout.
func (self *Command) Quiet() *Command {
	self.cmd.Stdin = os.Stdin
	self.cmd.Stderr = os.Stderr
	return self
}

// Run runs the command and returns an error if it fails.
func (self *Command) Run() error {
	return self.cmd.Run()
}

// Success runs the command and returns true if it exits with code 0.
func (self *Command) Success() bool {
	return self.Run() == nil
}

// ExitCode runs the command and returns its exit code.
// Returns -1 if the command failed to start.
func (self *Command) ExitCode() int {
	err := self.Run()
	if err == nil {
		return 0
	}
	if exitErr, ok := errors.AsType[*exec.ExitError](err); ok {
		return exitErr.ExitCode()
	}
	return -1
}

// Output runs the command and returns its stdout.
func (self *Command) Output() ([]byte, error) {
	return self.cmd.Output()
}

// CombinedOutput runs the command and returns its combined stdout and stderr.
func (self *Command) CombinedOutput() ([]byte, error) {
	return self.cmd.CombinedOutput()
}

// MustRun runs the command and panics if it fails.
func (self *Command) MustRun() {
	lo.Must0(self.Run())
}

// MustOutput runs the command and returns its stdout, panicking if it fails.
func (self *Command) MustOutput() []byte {
	return lo.Must(self.Output())
}

// MustCombinedOutput runs the command and returns its combined stdout and stderr, panicking if it fails.
func (self *Command) MustCombinedOutput() []byte {
	return lo.Must(self.CombinedOutput())
}

// OutputString runs the command and returns its stdout as a trimmed string.
func (self *Command) OutputString() (string, error) {
	b, err := self.Output()
	if err != nil {
		return "", err
	}
	return strings.TrimSpace(string(b)), nil
}

// MustOutputString runs the command and returns its stdout as a trimmed string, panicking if it fails.
func (self *Command) MustOutputString() string {
	return lo.Must(self.OutputString())
}

// OutputLines runs the command and returns its stdout as a slice of lines.
func (self *Command) OutputLines() ([]string, error) {
	s, err := self.OutputString()
	if err != nil {
		return nil, err
	}
	return strings.Split(s, "\n"), nil
}

// MustOutputLines runs the command and returns its stdout as a slice of lines, panicking if it fails.
func (self *Command) MustOutputLines() []string {
	return lo.Must(self.OutputLines())
}

// JustOutputString runs the command and returns its stdout as a trimmed string, or an empty string if the command fails.
func (self *Command) JustOutputString() string {
	s, err := self.OutputString()
	if err != nil {
		return ""
	}
	return s
}

// JustOutputLines runs the command and returns its stdout as a slice of lines, or an empty slice if the command fails.
func (self *Command) JustOutputLines() []string {
	lines, err := self.OutputLines()
	if err != nil {
		return nil
	}
	return lines
}
