package main

import (
	"flag"
	"fmt"
	"os"
)

var (
	incType    = flag.String("type", "patch", "kind of increment; values: patch, minor, major")
	printTag   = flag.Bool("print-version", true, "whether to print the new version to stdout; values: true/false")
	printFancy = flag.Bool("print-fancy", true, "whether to print the new version in a nicer format; value: true/false")
	copyTag    = flag.Bool("copy-version", true, "whether to copy the new version to the system clipboard; values: true/false")
)

func die(msg string, args ...any) {
	fmt.Fprintf(os.Stderr, msg, args...)
	os.Exit(1)
}

func main() {
	// parse
	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "get the next semver for your repo from git tags\n\nFlags:\n")
		flag.PrintDefaults()
	}

	flag.Parse()

	var incT IncType
	switch *incType {
	case "patch":
		incT = IncPatch
	case "minor":
		incT = IncMinor
	case "major":
		incT = IncMajor
	default:
		die("unsupported type, check inc -help for possible values")
	}

	// get tag and validate
	tag, err := getLastTag()
	if err != nil {
		die("error reading tag: %s\n", err.Error())
	}
	isTagValid, err := isTagValid(tag)
	if err != nil {
		die("something went wrong trying to process the tag: %s\n", err.Error())
	}
	if !isTagValid {
		die("unsupported tag found: %s; run inc -help to see supported formats\n", tag)
	}

	semVer, err := getSemVer(tag)
	if err != nil {
		die("something went wrong trying to increment the tag: %s\n", err.Error())
	}

	// increment
	newTag := incSemVer(&semVer, incT)

	// output
	if *printTag {
		if *printFancy {
			fmt.Fprintf(os.Stdout, "%s 👉 %s\n", tag, newTag.String())
		} else {
			fmt.Fprintf(os.Stdout, "%s", newTag.String())
		}
	}
	if *copyTag {
		err = copyToClipboard(newTag.String())
		if err != nil {
			die("something went wrong copying the tag to clipboard: %s\n", err.Error())
		}
	}
}
