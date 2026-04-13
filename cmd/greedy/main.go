package main

import (
	"flag"
	"fmt"
	"morego/algorithms/greedy"
	"os"
	"strings"
)

func main() {
	problem := flag.String(
		"problem",
		string(greedy.OneSwap),
		"greedy problem to run: "+strings.Join(greedy.ProblemNames(), ", "),
	)
	flag.Parse()

	if err := run(*problem); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func run(problem string) error {
	switch problem {
	case string(greedy.OneSwap):
		s := "768"
		fmt.Println(greedy.LargestSwap(s))
		return nil

	case string(greedy.LargestSubsK):
		s := "banana"
		k := 2
		fmt.Println(greedy.LargestSubsAtLeastK(s, k))
		return nil

	default:
		return fmt.Errorf(
			"unknown greedy problem %q (available: %s)",
			problem,
			strings.Join(greedy.ProblemNames(), ", "),
		)
	}
}
