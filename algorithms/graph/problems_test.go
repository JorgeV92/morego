package graph

import "testing"

func TestSolveSurroundedRegions(t *testing.T) {
	grid := [][]byte{
		{'X', 'O', 'X', 'O', 'X', 'X'},
		{'X', 'O', 'X', 'X', 'O', 'X'},
		{'X', 'X', 'X', 'O', 'X', 'X'},
		{'O', 'X', 'X', 'X', 'X', 'X'},
		{'X', 'X', 'X', 'O', 'X', 'O'},
		{'O', 'O', 'X', 'O', 'O', 'O'},
	}

	expected := [][]byte{
		{'X', 'O', 'X', 'O', 'X', 'X'},
		{'X', 'O', 'X', 'X', 'X', 'X'},
		{'X', 'X', 'X', 'X', 'X', 'X'},
		{'O', 'X', 'X', 'X', 'X', 'X'},
		{'X', 'X', 'X', 'O', 'X', 'O'},
		{'O', 'O', 'X', 'O', 'O', 'O'},
	}

	SolveSurroundedRegions(grid)

	if !sameByteGrid(grid, expected) {
		t.Fatalf("unexpected grid result: got %q want %q", grid, expected)
	}
}

func TestTopologicalSort(t *testing.T) {
	edges := [][2]int{
		{0, 1},
		{1, 2},
		{2, 3},
		{4, 5},
		{5, 1},
		{5, 2},
	}

	order := TopologicalSort(6, edges)

	if len(order) != 6 {
		t.Fatalf("unexpected order length: got %d want 6", len(order))
	}

	positions := make(map[int]int, len(order))
	for index, vertex := range order {
		positions[vertex] = index
	}

	for _, edge := range edges {
		from, to := edge[0], edge[1]
		if positions[from] >= positions[to] {
			t.Fatalf("invalid topological order %v for edge %v", order, edge)
		}
	}
}

func TestMinimumLines(t *testing.T) {
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

	if got := MinimumLines(stockPrices); got != 3 {
		t.Fatalf("unexpected number of lines: got %d want 3", got)
	}
}

func TestMinimumObstacles(t *testing.T) {
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

	if got := MinimumObstacles(grid1); got != 2 {
		t.Fatalf("unexpected obstacle result for grid1: got %d want 2", got)
	}

	if got := MinimumObstacles(grid2); got != 0 {
		t.Fatalf("unexpected obstacle result for grid2: got %d want 0", got)
	}
}

func sameByteGrid(left, right [][]byte) bool {
	if len(left) != len(right) {
		return false
	}

	for row := range left {
		if len(left[row]) != len(right[row]) {
			return false
		}

		for col := range left[row] {
			if left[row][col] != right[row][col] {
				return false
			}
		}
	}

	return true
}
