package main

import (
	"io"
	"os"
	"strings"

	"crdx.org/cmd/pkg/sh"
)

func workingTreeStatus(directory string) (string, error) {
	output, err := sh.Cmd("git", "status", "--porcelain", "--untracked-files=all").Dir(directory).Output()
	return strings.TrimRight(string(output), "\r\n"), err
}

func tagExists(directory string, tag string) (bool, error) {
	output, err := sh.Cmd("git", "tag", "--list", tag).Dir(directory).Output()
	return len(output) > 0, err
}

func verifyRef(directory string, ref string) error {
	return sh.Cmd("git", "rev-parse", "--verify", "--quiet", ref+"^{commit}").Dir(directory).Run()
}

// findLatestVersion finds the latest version tag reachable from HEAD.
func findLatestVersion() (string, Version, bool) {
	lines := sh.Cmd("git", "tag", "--merged", "HEAD", "--sort=-v:refname").Stderr(io.Discard).JustOutputLines()

	for _, tag := range lines {
		raw := removePrefix(tag)
		if version, ok := parseVersion(raw); ok {
			return tag, version, true
		}
	}

	return "", Version{}, false
}

func createTag(tag string, ref string) {
	if err := sh.Cmd("git", "tag", tag, "-a", "-m", tag, ref).Run(); err != nil {
		os.Exit(1)
	}
}
