// Bash completion: ../cmdctl/completions/unbuffer.bash

package main

import (
	"errors"
	"fmt"
	"io"
	"os"
	"os/exec"
	"os/signal"
	"path/filepath"
	"syscall"

	"github.com/creack/pty"
	"golang.org/x/sys/unix"
)

var name = filepath.Base(os.Args[0])

func main() {
	if len(os.Args) < 2 {
		fmt.Fprintf(os.Stderr, "Usage: %s <command> [args...]\n", name)
		os.Exit(1)
	}

	os.Exit(run())
}

func run() int {
	ptyMaster, tty, err := pty.Open()
	if err != nil {
		fmt.Fprintf(os.Stderr, "%s: failed to open pty: %v\n", name, err)
		return 1
	}

	defer func() { _ = ptyMaster.Close() }()

	termios, err := unix.IoctlGetTermios(int(tty.Fd()), unix.TCGETS) //nolint:gosec
	if err != nil {
		_ = tty.Close()
		fmt.Fprintf(os.Stderr, "%s: failed to get termios: %v\n", name, err)
		return 1
	}

	termios.Oflag &^= unix.OPOST

	if err := unix.IoctlSetTermios(int(tty.Fd()), unix.TCSETS, termios); err != nil { //nolint:gosec
		_ = tty.Close()
		fmt.Fprintf(os.Stderr, "%s: failed to set termios: %v\n", name, err)
		return 1
	}

	if windowSize, sizeErr := pty.GetsizeFull(os.Stdin); sizeErr == nil {
		_ = pty.Setsize(ptyMaster, windowSize)
	}

	command := exec.Command(os.Args[1], os.Args[2:]...) //nolint:gosec
	command.Stdin = tty
	command.Stdout = tty
	command.Stderr = tty
	command.SysProcAttr = &syscall.SysProcAttr{
		Setsid:  true,
		Setctty: true,
		Ctty:    0,
	}

	if err := command.Start(); err != nil {
		_ = tty.Close()
		fmt.Fprintf(os.Stderr, "%s: failed to start command: %v\n", name, err)
		return 1
	}

	_ = tty.Close()

	signals := make(chan os.Signal, 1)
	signal.Notify(signals, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		for sig := range signals {
			_ = command.Process.Signal(sig)
		}
	}()

	_, _ = io.Copy(os.Stdout, ptyMaster)

	err = command.Wait()

	signal.Stop(signals)
	close(signals)

	if err != nil {
		if exitError, ok := errors.AsType[*exec.ExitError](err); ok {
			return exitError.ExitCode()
		}

		return 1
	}

	return 0
}
