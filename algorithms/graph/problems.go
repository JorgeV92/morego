package graph

import (
	"container/list"
	"sort"
)

type Problem string

const (
	SurroundingXO   Problem = "surrounding-xo"
	Topological     Problem = "topological-sort"
	LineChart       Problem = "line-chart"
	ObstacleRemoval Problem = "obstacle-removal"
	IncreasingPaths Problem = "count-paths"
)

func ProblemNames() []string {
	return []string{
		string(SurroundingXO),
		string(Topological),
		string(LineChart),
		string(ObstacleRemoval),
		string(IncreasingPaths),
	}
}

func SolveSurroundedRegions(grid [][]byte) {
	if len(grid) == 0 || len(grid[0]) == 0 {
		return
	}

	rows, cols := len(grid), len(grid[0])
	directions := [5]int{-1, 0, 1, 0, -1}

	var dfs func(int, int)
	dfs = func(row, col int) {
		if row < 0 || row >= rows || col < 0 || col >= cols || grid[row][col] != 'O' {
			return
		}

		grid[row][col] = '.'
		for i := 0; i < 4; i++ {
			dfs(row+directions[i], col+directions[i+1])
		}
	}

	for row := 0; row < rows; row++ {
		dfs(row, 0)
		dfs(row, cols-1)
	}

	for col := 1; col < cols-1; col++ {
		dfs(0, col)
		dfs(rows-1, col)
	}

	for row := 0; row < rows; row++ {
		for col := 0; col < cols; col++ {
			switch grid[row][col] {
			case '.':
				grid[row][col] = 'O'
			case 'O':
				grid[row][col] = 'X'
			}
		}
	}
}

func TopologicalSort(vertices int, edges [][2]int) []int {
	graph := make([][]int, vertices)
	inDegree := make([]int, vertices)

	for _, edge := range edges {
		from, to := edge[0], edge[1]
		graph[from] = append(graph[from], to)
		inDegree[to]++
	}

	queue := make([]int, 0, vertices)
	for vertex := 0; vertex < vertices; vertex++ {
		if inDegree[vertex] == 0 {
			queue = append(queue, vertex)
		}
	}

	order := make([]int, 0, vertices)
	for head := 0; head < len(queue); head++ {
		vertex := queue[head]
		order = append(order, vertex)

		for _, next := range graph[vertex] {
			inDegree[next]--
			if inDegree[next] == 0 {
				queue = append(queue, next)
			}
		}
	}

	return order
}

func MinimumLines(stockPrices [][]int) int {
	if len(stockPrices) <= 1 {
		return 0
	}

	prices := append([][]int(nil), stockPrices...)
	sort.Slice(prices, func(i, j int) bool {
		if prices[i][0] == prices[j][0] {
			return prices[i][1] < prices[j][1]
		}
		return prices[i][0] < prices[j][0]
	})

	lines := 0
	dx, dy := 0, 1

	for i := 1; i < len(prices); i++ {
		x0, y0 := prices[i-1][0], prices[i-1][1]
		x1, y1 := prices[i][0], prices[i][1]
		nextDX, nextDY := x1-x0, y1-y0

		if int64(nextDY)*int64(dx) != int64(nextDX)*int64(dy) {
			lines++
		}

		dx, dy = nextDX, nextDY
	}

	return lines
}

func MinimumObstacles(grid [][]int) int {
	if len(grid) == 0 || len(grid[0]) == 0 {
		return -1
	}

	rows, cols := len(grid), len(grid[0])
	inf := rows*cols + 1

	dist := make([][]int, rows)
	for row := range dist {
		dist[row] = make([]int, cols)
		for col := range dist[row] {
			dist[row][col] = inf
		}
	}

	type cell struct {
		row int
		col int
	}

	dist[0][0] = 0
	deque := list.New()
	deque.PushFront(cell{row: 0, col: 0})

	directions := [5]int{-1, 0, 1, 0, -1}

	for deque.Len() > 0 {
		front := deque.Front()
		current := front.Value.(cell)
		deque.Remove(front)

		if current.row == rows-1 && current.col == cols-1 {
			return dist[current.row][current.col]
		}

		for i := 0; i < 4; i++ {
			nextRow := current.row + directions[i]
			nextCol := current.col + directions[i+1]

			if nextRow < 0 || nextRow >= rows || nextCol < 0 || nextCol >= cols {
				continue
			}

			nextCost := dist[current.row][current.col] + grid[nextRow][nextCol]
			if nextCost >= dist[nextRow][nextCol] {
				continue
			}

			dist[nextRow][nextCol] = nextCost
			nextCell := cell{row: nextRow, col: nextCol}

			if grid[nextRow][nextCol] == 0 {
				deque.PushFront(nextCell)
			} else {
				deque.PushBack(nextCell)
			}
		}
	}

	return -1
}

func CountPaths(grid [][]int) (ans int) {
	const mod = 1e9 + 7
	m, n := len(grid), len(grid[0])
	f := make([][]int, m)
	for i := range f {
		f[i] = make([]int, n)
	}
	var dfs func(int, int) int
	dfs = func(i, j int) int {
		if f[i][j] != 0 {
			return f[i][j]
		}
		f[i][j] = 1
		dirs := [5]int{-1, 0, 1, 0, -1}
		for k := 0; k < 4; k++ {
			x, y := i+dirs[k], j+dirs[k+1]
			if x >= 0 && x < m && y >= 0 && y < n && grid[i][j] < grid[x][y] {
				f[i][j] = (f[i][j] + dfs(x, y)) % mod
			}
		}
		return f[i][j]
	}
	for i, row := range grid {
		for j := range row {
			ans = (ans + dfs(i, j)) % mod
		}
	}
	return
}
