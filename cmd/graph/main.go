package main

import (
	"flag"
	"fmt"
	"os"
	"strings"

	"morego/algorithms/graph"
)

func main() {
	problem := flag.String(
		"problem",
		string(graph.ObstacleRemoval),
		"graph problem to run: "+strings.Join(graph.ProblemNames(), ", "),
	)
	flag.Parse()

	if err := run(*problem); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func run(problem string) error {
	switch problem {
	case string(graph.SurroundingXO):
		grid := [][]byte{
			{'X', 'O', 'X', 'O', 'X', 'X'},
			{'X', 'O', 'X', 'X', 'O', 'X'},
			{'X', 'X', 'X', 'O', 'X', 'X'},
			{'O', 'X', 'X', 'X', 'X', 'X'},
			{'X', 'X', 'X', 'O', 'X', 'O'},
			{'O', 'O', 'X', 'O', 'O', 'O'},
		}

		graph.SolveSurroundedRegions(grid)
		printByteGrid(grid)
		return nil

	case string(graph.Topological):
		edges := [][2]int{
			{0, 1},
			{1, 2},
			{2, 3},
			{4, 5},
			{5, 1},
			{5, 2},
		}

		fmt.Println(graph.TopologicalSort(6, edges))
		return nil

	case string(graph.LineChart):
		stockPrices := [][]int{
			{1, 7},
			{2, 6},
			{3, 5},
			{4, 4},
			{5, 4},
			{6, 3},
			{7, 2},
			{8, 1},
		}

		fmt.Println(graph.MinimumLines(stockPrices))
		return nil

	case string(graph.ObstacleRemoval):
		grid1 := [][]int{
			{0, 1, 1},
			{1, 1, 0},
			{1, 1, 0},
		}

		grid2 := [][]int{
			{0, 1, 0, 0, 0},
			{0, 1, 0, 1, 0},
			{0, 0, 0, 1, 0},
		}

		fmt.Println(graph.MinimumObstacles(grid1))
		fmt.Println(graph.MinimumObstacles(grid2))
		return nil

	default:
		return fmt.Errorf(
			"unknown graph problem %q (available: %s)",
			problem,
			strings.Join(graph.ProblemNames(), ", "),
		)
	}
}

func printByteGrid(grid [][]byte) {
	for _, row := range grid {
		for _, value := range row {
			fmt.Printf("%c ", value)
		}
		fmt.Println()
	}
}
