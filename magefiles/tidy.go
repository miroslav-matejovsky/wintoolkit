package main

import (
	"fmt"
	"os"

	"github.com/magefile/mage/sh"
)

// Tidy cleans up go.mod and go.sum files in all projects.
func Tidy() error {
	if err := sh.RunV("go", "work", "sync"); err != nil {
		return err
	}
	for _, project := range projects {
		if err := TidyProject(project); err != nil {
			return err
		}
	}
	fmt.Println("tidy done")
	return nil
}

// TidyProject runs go mod tidy on a specific project.
func TidyProject(project string) error {
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
