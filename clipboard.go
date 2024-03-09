package main

import (
	"github.com/atotto/clipboard"
)

func copyToClipboard(version string) error {
	err := clipboard.WriteAll(version)
	return err
}
