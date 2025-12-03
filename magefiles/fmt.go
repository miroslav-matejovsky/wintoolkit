package main

import (
	"fmt"

	"github.com/magefile/mage/sh"
)

// Fmt formats Go source files with gofmt.
func Fmt() error {
	if err := sh.RunV("gofmt", "-s", "-w", "."); err != nil {
		return err
	}
	fmt.Println("fmt done")
	return nil
}
