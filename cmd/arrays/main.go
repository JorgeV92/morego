package main

import (
	"flag"
	"fmt"
	"os"
	"strings"

	"morego/algorithms/arrays"
)

func main() {
	problem := flag.String(
		"problem",
		string(arrays.ArrayChangeProblem),
		"array problem to run: "+strings.Join(arrays.ProblemNames(), ", "),
	)
	flag.Parse()

	if err := run(*problem); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func run(problem string) error {
	switch problem {
	case string(arrays.ThreeSumProblem):
		nums := []int{-1, 0, 1, 2, -1, -4}
		for _, triplet := range arrays.ThreeSum(nums) {
			fmt.Println(triplet)
		}

		targetNums := []int{1, 4, 45, 6, 10, 8}
		fmt.Println(arrays.ThreeSumTargetExists(targetNums, 13))
		return nil

	case string(arrays.ArrayChangeProblem):
		nums := []int{1, 2, 4, 6}
		operations := [][2]int{
			{1, 3},
			{4, 7},
			{6, 1},
		}

		fmt.Println(arrays.ArrayChange(nums, operations))
		return nil

	default:
		return fmt.Errorf(
			"unknown array problem %q (available: %s)",
			problem,
			strings.Join(arrays.ProblemNames(), ", "),
		)
	}
}
