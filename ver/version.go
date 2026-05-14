package main

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

type Scheme int

const (
	SchemeTwoPart   Scheme = 2
	SchemeThreePart Scheme = 3
)

type Version struct {
	Major  int
	Minor  int
	Patch  int
	Scheme Scheme
}

var (
	threePartPattern = regexp.MustCompile(`^\d+\.\d+\.\d+$`)
	twoPartPattern   = regexp.MustCompile(`^\d+\.\d+$`)
)

// removePrefix strips any non-numeric prefix (normally "v").
func removePrefix(tag string) string {
	for i, char := range tag {
		if char >= '0' && char <= '9' {
			return tag[i:]
		}
	}
	return ""
}

// parseVersion parses a version string like "1.2.3" or "1.2" (without prefix).
func parseVersion(raw string) (Version, bool) {
	// Strip any suffix like "-rc1".
	base, _, _ := strings.Cut(raw, "-")

	mustInt := func(s string) int {
		n, err := strconv.Atoi(s)
		if err != nil {
			panic(err)
		}
		return n
	}

	if threePartPattern.MatchString(base) {
		parts := strings.Split(base, ".")
		return Version{
			Major:  mustInt(parts[0]),
			Minor:  mustInt(parts[1]),
			Patch:  mustInt(parts[2]),
			Scheme: SchemeThreePart,
		}, true
	}

	if twoPartPattern.MatchString(base) {
		parts := strings.Split(base, ".")
		return Version{
			Major:  mustInt(parts[0]),
			Minor:  mustInt(parts[1]),
			Scheme: SchemeTwoPart,
		}, true
	}

	return Version{}, false
}

func (self Version) String() string {
	if self.Scheme == SchemeThreePart {
		return fmt.Sprintf("%d.%d.%d", self.Major, self.Minor, self.Patch)
	}
	return fmt.Sprintf("%d.%d", self.Major, self.Minor)
}

func (self Version) Tag() string {
	return "v" + self.String()
}

func (self Version) IncrementMajor() Version {
	if self.Scheme == SchemeThreePart {
		return Version{Major: self.Major + 1, Minor: 0, Patch: 0, Scheme: SchemeThreePart}
	}
	return Version{Major: self.Major + 1, Minor: 0, Scheme: SchemeTwoPart}
}

func (self Version) IncrementMinor() Version {
	if self.Scheme == SchemeThreePart {
		return Version{Major: self.Major, Minor: self.Minor + 1, Patch: 0, Scheme: SchemeThreePart}
	}
	return Version{Major: self.Major, Minor: self.Minor + 1, Scheme: SchemeTwoPart}
}

func (self Version) IncrementPatch() Version {
	return Version{Major: self.Major, Minor: self.Minor, Patch: self.Patch + 1, Scheme: SchemeThreePart}
}
