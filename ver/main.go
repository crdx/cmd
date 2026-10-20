// Bash completion: ../cmdctl/completions/ver.bash

package main

import (
	"fmt"

	"crdx.org/col"
	"crdx.org/duckopt/v2"
	"crdx.org/logger"
)

func getUsage() string {
	return `
		Usage:
			$0 [options] (major | minor | patch) [--ref REF]
			$0 [options] <version> [--ref REF]
			$0

		Increment the current version and tag it.
		If <version> is specified (without 'v' prefix) then use that instead.

		Options:
			-r, --ref REF    Ref to tag [default: HEAD]
	`
}

type Opts struct {
	CmdMajor bool   `docopt:"major"`
	CmdMinor bool   `docopt:"minor"`
	CmdPatch bool   `docopt:"patch"`
	Version  string `docopt:"<version>"`
	Ref      string `docopt:"--ref"`
}

func main() {
	opts := duckopt.MustBind[Opts](getUsage(), "$0")
	logger.InitStderr()
	col.Init()

	status, err := workingTreeStatus(".")
	if err != nil {
		logger.Fatal("need to be in a git repository")
	}
	if status != "" {
		logger.Fatal("working tree is not clean:\n%s", status)
	}

	if opts.Version != "" {
		version, ok := parseVersion(opts.Version)
		if !ok {
			logger.Fatal("version should be <major>.<minor>.<patch> or <major>.<minor>")
		}

		newTag := version.Tag()
		ensureTagCanBeCreated(newTag, opts.Ref)
		fmt.Println(col.Yellow("New version: %s", newTag))
		execHook(preVersionHook, newTag, "")
		createTag(newTag, opts.Ref)
		execHook(postVersionHook, newTag, "")
	} else if opts.CmdMajor || opts.CmdMinor || opts.CmdPatch {
		oldTag, oldVersion, found := findLatestVersion()
		if !found {
			logger.Fatal("no existing version tags found")
		}

		if opts.CmdPatch && oldVersion.Scheme == SchemeTwoPart {
			logger.Fatal("found %s — patch only applies to 3-part versions", oldTag)
		}

		fmt.Println(col.Green("Old version: %s", oldTag))

		var newVersion Version

		switch {
		case opts.CmdMajor:
			newVersion = oldVersion.IncrementMajor()
		case opts.CmdMinor:
			newVersion = oldVersion.IncrementMinor()
		case opts.CmdPatch:
			newVersion = oldVersion.IncrementPatch()
		}

		newTag := newVersion.Tag()
		ensureTagCanBeCreated(newTag, opts.Ref)
		fmt.Println(col.Yellow("New version: %s", newTag))
		execHook(preVersionHook, newTag, oldTag)
		createTag(newTag, opts.Ref)
		execHook(postVersionHook, newTag, oldTag)
	} else {
		showLatest()
	}
}

func ensureTagCanBeCreated(tag string, ref string) {
	exists, err := tagExists(".", tag)
	if err != nil {
		logger.Fatal("could not inspect existing tags: %v", err)
	}
	if exists {
		logger.Fatal("tag %s already exists", tag)
	}
	if err := verifyRef(".", ref); err != nil {
		logger.Fatal("ref %s is not a commit", ref)
	}
}

// execHook prints a notice and runs a hook if it exists and is executable.
func execHook(name string, newTag string, oldTag string) {
	if willRunHook(name) {
		fmt.Println(col.Blue("==> Executing %s...", name))
	}
	runHook(name, newTag, oldTag)
}

func showLatest() {
	tag, _, found := findLatestVersion()
	if !found {
		logger.Fatal("no existing version tags found")
	}

	fmt.Printf("Latest version: %s\n", tag)
}
