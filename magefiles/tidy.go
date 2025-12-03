//go:build mage

package main

import (
	"fmt"
	"os"

	"github.com/magefile/mage/sh"
)

// Tidy cleans up go.mod and go.sum files.
func Tidy() error {
	for _, project := range projects {
		if err := tidy(project); err != nil {
			return err
		}
	}
	fmt.Println("tidy done")
	return nil
}

func tidy(project string) error {
	cwd, err := os.Getwd()
	if err != nil {
		return err
	}
	defer func() {
		_ = os.Chdir(cwd)
	}()
	if err := os.Chdir(project); err != nil {
		return err
	}
	if err := sh.RunWith(nil, "go", "mod", "tidy"); err != nil {
		return err
	}
	fmt.Printf("tidy done for %s\n", project)
	return nil
}
