package main

import "fmt"

const (
	versionPrefix = "v"
)

type IncType uint

const (
	IncPatch IncType = iota
	IncMinor
	IncMajor
)

type SemVer struct {
	prefix bool
	major  int
	minor  int
	patch  int
}

func (s SemVer) String() string {
	var prefixStr string
	if s.prefix {
		prefixStr = versionPrefix
	}
	return fmt.Sprintf("%s%d.%d.%d", prefixStr, s.major, s.minor, s.patch)
}
