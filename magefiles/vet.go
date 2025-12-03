package main

import (
	"fmt"

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
	path := "./" + project + "/..."
	fmt.Printf("vetting %s...\n", project)
	if err := sh.RunV("go", "vet", path); err != nil {
		return err
	}
	fmt.Printf("vet done for %s\n", project)
	return nil
}
