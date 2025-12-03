package main

import (
	"fmt"
	"os"

	"github.com/magefile/mage/sh"
)

// Vet runs go vet on all packages.
func Vet() error {
	for _, project := range projects {
		if err := VetProject(project); err != nil {
			return err
		}
	}
	fmt.Println("vet done")
	return nil
}

// VetProject runs go vet on a specific project.
func VetProject(project string) error {
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
	if err := sh.RunV("go", "vet", "./..."); err != nil {
		return err
	}
	fmt.Printf("vet done for %s\n", project)
	return nil
}
