package main

import (
	"golang.design/x/clipboard"
)

func copyToClipboard(tag string) error {
	err := clipboard.Init()
	if err != nil {
		return err
	}
	clipboard.Write(clipboard.FmtText, []byte(tag))
	return nil
}
