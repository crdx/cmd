// Bash completion: completions/cmdctl.bash

package main

import (
	"debug/buildinfo"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"

	"crdx.org/duckopt/v2"
	"crdx.org/logger"
)

type Completion struct {
	Name   string
	Script string
}

var bins = []Completion{
	{"cmdctl", cmdctlScript},
	{"uchar", ucharScript},
	{"unbuffer", unbufferScript},
	{"ver", verScript},
}

var bashrcPath = filepath.Join(os.Getenv("HOME"), ".bashrc")

const (
	evalLine    = `eval "$(cmdctl generate bash)"`
	packageName = "crdx.org/cmd"
)

func getUsage() string {
	return `
		Usage:
			cmdctl generate bash
			cmdctl install bash

		Commands:
			generate bash    Generate Bash completions for all tools on PATH
			install bash     Add generated script evaluation line to ~/.bashrc
	`
}

type Opts struct {
	Generate bool `docopt:"generate"`
	Bash     bool `docopt:"bash"`
	Install  bool `docopt:"install"`
}

func main() {
	logger.InitStderr()
	opts := duckopt.MustBind[Opts](getUsage())

	if opts.Generate && opts.Bash {
		generate()
	}

	if opts.Install && opts.Bash {
		install()
	}
}

func generate() {
	var builder strings.Builder
	for _, bin := range bins {
		path, err := exec.LookPath(bin.Name)
		if err != nil {
			continue
		}
		info, err := buildinfo.ReadFile(path)
		if err != nil || info.Main.Path != packageName {
			continue
		}
		builder.WriteString(bin.Script)
		if !strings.HasSuffix(bin.Script, "\n") {
			builder.WriteByte('\n')
		}
	}
	fmt.Print(builder.String())
}

func install() {
	bashrcBytes, err := os.ReadFile(bashrcPath) //nolint:gosec // path derived from $HOME
	if err != nil {
		logger.Fatal(err)
	}

	bashrc := string(bashrcBytes)

	for line := range strings.SplitSeq(bashrc, "\n") {
		if strings.TrimSpace(line) == evalLine {
			fmt.Fprintf(os.Stderr, "already installed in %s\n", bashrcPath)
			return
		}
	}

	if !strings.HasSuffix(bashrc, "\n") {
		bashrc += "\n"
	}
	bashrc += evalLine + "\n"

	if err := os.WriteFile(bashrcPath, []byte(bashrc), 0o644); err != nil { //nolint:gosec
		logger.Fatal(err)
	}

	fmt.Fprintf(os.Stderr, "added to %s\n", bashrcPath)
}
