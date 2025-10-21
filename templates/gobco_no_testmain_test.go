//go:build ignore
// +build ignore

package main

// This file is used if the code to be instrumented does not define its own
// TestMain function.

import (
	"os"
	"testing"

	"main/gobco_test"
)

func TestMain(m *testing.M) {
	gobco_test.GobcoCounts.Load(gobco_test.GobcoCounts.Filename())
	exitCode := m.Run()
	gobco_test.GobcoCounts.Persist()
	os.Exit(exitCode)
}
