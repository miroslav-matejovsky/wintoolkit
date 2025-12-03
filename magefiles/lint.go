package main

import (
	"fmt"

	"github.com/magefile/mage/sh"
)

// Lint runs golangci-lint on all packages.
func Lint() error {
	fmt.Println("linting...")
	for _, project := range projects {
		if err := LintProject(project); err != nil {
			return err
		}
	}
	fmt.Println("linting done")
	return nil
}

// LintProject runs golangci-lint on a specific project.
func LintProject(project string) error {
	path := "./" + project + "/..."
	fmt.Printf("linting %s...\n", project)
	if err := sh.RunV("golangci-lint", "run", path); err != nil {
		return err
	}
	fmt.Printf("lint done for %s\n", project)
	return nil
}
