package main

import (
	"flag"
	"fmt"
	"os"

	"morego/algorithms/data"
)

const treapProblem = "treap"

func main() {
	problem := flag.String("problem", treapProblem, "data problem to run: treap")
	flag.Parse()

	if err := run(*problem); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func run(problem string) error {
	switch problem {
	case treapProblem:
		runTreapExample()
		return nil
	default:
		return fmt.Errorf("unknown data problem %q (available: %s)", problem, treapProblem)
	}
}

func runTreapExample() {
	var root *data.Node
	for _, value := range []int{5, 1, 3, 2, 4} {
		root = data.Insert(root, value)
	}

	fmt.Println("after insert:", inorderValues(root))
	fmt.Println("size:", data.GetSize(root))
	fmt.Println("sum:", data.GetSum(root))

	root = data.Erase(root, 3)
	fmt.Println("after erase 3:", inorderValues(root))

	left, right := data.SplitSz(root, 2)
	fmt.Println("left split:", inorderValues(left))
	fmt.Println("right split:", inorderValues(right))

	root = data.Merge(left, right)
	fmt.Println("merged again:", inorderValues(root))
}

func inorderValues(root *data.Node) []int {
	values := make([]int, 0)
	data.Inorder(root, &values)
	return values
}
