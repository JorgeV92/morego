package main

import (
	"strings"
	"testing"
)

func TestSolveRoadReparationSample(t *testing.T) {
	input := strings.NewReader(`5 6
								1 2 3
								2 3 5
								2 4 2
								3 4 8
								5 1 7
								5 4 4
								`)

	var output strings.Builder
	if err := solveRoadReparation(input, &output); err != nil {
		t.Fatalf("solveRoadReparation returned error: %v", err)
	}

	if got := output.String(); got != "14\n" {
		t.Fatalf("unexpected output: got %q want %q", got, "14\n")
	}
}

func TestSolveRoadReparationImpossible(t *testing.T) {
	input := strings.NewReader(`4 2
1 2 5
3 4 7
`)

	var output strings.Builder
	if err := solveRoadReparation(input, &output); err != nil {
		t.Fatalf("solveRoadReparation returned error: %v", err)
	}

	if got := output.String(); got != "IMPOSSIBLE\n" {
		t.Fatalf("unexpected output: got %q want %q", got, "IMPOSSIBLE\n")
	}
}
