package main

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/magefile/mage/mg"
	"github.com/magefile/mage/sh"
)

// Publish publishes the module to the Go registry by tagging and pushing the version.
// The version should be in the format x.y.z or vx.y.z.
func Publish(version string) error {
	// Validate version format: must be number.number.number or vnumber.number.number
	validVersion := regexp.MustCompile(`^v?\d+\.\d+\.\d+$`)
	if !validVersion.MatchString(version) {
		return fmt.Errorf("invalid version format: %s. Must be in the form x.y.z or vx.y.z", version)
	}

	// Determine the tag: add 'v' prefix if missing
	tag := version
	if !strings.HasPrefix(version, "v") {
		tag = "v" + version
	}

	mg.Deps(All)

	// Tag the current commit
	if err := sh.RunV("git", "tag", tag); err != nil {
		return fmt.Errorf("failed to create git tag: %w", err)
	}

	// Push the tag to the remote repository
	if err := sh.RunV("git", "push", "origin", tag); err != nil {
		return fmt.Errorf("failed to push git tag: %w", err)
	}

	fmt.Printf("Published version %s as a git tag\n", tag)
	return nil
}
