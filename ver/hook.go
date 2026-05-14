package main

import (
	"os"

	"crdx.org/cmd/pkg/sh"
	"crdx.org/logger"
)

const hookDir = "hooks"

const (
	preVersionHook  = "pre-version"
	postVersionHook = "post-version"
)

// hookPath returns the file path for a named hook.
func hookPath(name string) string {
	return hookDir + "/" + name
}

// shouldRunHook checks whether a hook script exists and is executable.
func shouldRunHook(path string) bool {
	info, err := os.Stat(path)
	if err != nil {
		return false
	}
	return info.Mode()&0o111 != 0
}

// willRunHook checks whether a hook with the given name exists and is executable.
func willRunHook(name string) bool {
	return shouldRunHook(hookPath(name))
}

// buildHookArgs builds the argument list for a hook script. Versions have the prefix stripped.
func buildHookArgs(path string, newTag string, oldTag string) []string {
	args := []string{path, removePrefix(newTag)}
	if oldTag != "" {
		args = append(args, removePrefix(oldTag))
	}
	return args
}

// runHook executes a named hook script if it exists and is executable.
func runHook(name string, newTag string, oldTag string) {
	path := hookPath(name)

	if !shouldRunHook(path) {
		return
	}

	args := buildHookArgs(path, newTag, oldTag)
	if err := sh.Cmd(args...).Passthrough().Run(); err != nil {
		logger.Fatal("skipping tag as %s failed", name)
	}
}
