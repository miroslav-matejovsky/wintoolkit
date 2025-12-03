package main

import (
	"fmt"

	"github.com/magefile/mage/sh"
)

// Test runs all tests and outputs results in testdox format using gotestsum.
func Test() error {
	fmt.Println("testing...")
	for _, project := range projects {
		if err := TestProject(project); err != nil {
			return err
		}
	}
	fmt.Println("testing done")
	return nil
}

// TestProject runs tests for a specific project and outputs results in testdox format using gotestsum.
func TestProject(project string) error {
	fmt.Printf("testing %s...\n", project)
	path := "./" + project + "/..."
	if err := sh.RunV("gotestsum", "--format", "testdox", path); err != nil {
		return err
	}
	fmt.Printf("tests done for %s\n", project)
	return nil
}
