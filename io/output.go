package io

import (
	"fmt"
	"strconv"

	"github.com/kvrhdn/tfe-run/gha"
)

// Output data generated by tfe-run
type Output struct {
	RunURL     string
	HasChanges bool
}

// WriteOutput writes output.
func WriteOutput(output *Output) {
	if gha.InGitHubActions() {
		gha.WriteOutput("run-url", output.RunURL)
		gha.WriteOutput("has-changes", strconv.FormatBool(output.HasChanges))
	} else {
		fmt.Printf("Output:\n")
		fmt.Printf(" - run-url:     %s\n", output.RunURL)
		fmt.Printf(" - has-changes: %v\n", output.HasChanges)
	}
}