package main

import (
	"regexp"
	"strconv"
	"strings"
)

func isTagValid(tag string) (bool, error) {
	regexPattern := `^?\d+\.\d+\.\d+$`
	return regexp.MatchString(regexPattern, tag)
}

func getSemVer(versionStr string) (SemVer, error) {
	var hasPrefix bool
	switch versionStr[0] {
	case versionPrefix[0]:
		hasPrefix = true
	default:
		hasPrefix = false
	}
	versionEls := strings.Split(versionStr, ".")
	var majorStr string
	if hasPrefix {
		majorStr = versionEls[0][1:]
	} else {
		majorStr = versionEls[0]
	}

	major, err := strconv.Atoi(majorStr)
	if err != nil {
		return SemVer{}, err
	}
	minor, err := strconv.Atoi(versionEls[1])
	if err != nil {
		return SemVer{}, err
	}
	patch, err := strconv.Atoi(versionEls[2])
	if err != nil {
		return SemVer{}, err
	}
	return SemVer{hasPrefix, major, minor, patch}, nil
}

func incSemVer(semVer *SemVer, incType IncType) *SemVer {
	switch incType {
	case IncPatch:
		semVer.patch++
	case IncMinor:
		semVer.minor++
		semVer.patch = 0
	case IncMajor:
		semVer.major++
		semVer.minor = 0
		semVer.patch = 0
	}

	return semVer
}
