package main

import (
	"github.com/magefile/mage/mg"
)

var projects = []string{
	"magefiles",
	"fileinfo",
	"winservicedetail",
	"runtimesaudit",
}

// All runs all quality checks and tests.
func All() error {
	mg.Deps(Tidy, Fmt, Vet, Lint)
	if err := Test(); err != nil {
		return err
	}
	return nil
}
