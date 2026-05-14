package main

import (
	"io"
	"os"

	"crdx.org/cmd/pkg/sh"
)

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
