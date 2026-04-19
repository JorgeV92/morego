package graph

import (
	"container/list"
	"math/bits"
	"slices"
	"sort"
)

type Problem string

const (
	SurroundingXO         Problem = "surrounding-xo"
	Topological           Problem = "topological-sort"
	LineChart             Problem = "line-chart"
	ObstacleRemoval       Problem = "obstacle-removal"
	IncreasingPaths       Problem = "count-paths"
	ParallelCourse2       Problem = "parallel-course"
	MinimizeMalwareSpread Problem = "min-mal-spread"
	ReachDestinantion     Problem = "reach-destination"
	NumberOfGoodPaths     Problem = "num-of-good-paths"
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

func MinNumOfSemesters(n int, realtions [][]int, k int) int {
	d := make([]int, n+1)
	for _, e := range realtions {
		d[e[1]] |= 1 << e[0]
	}
	type pair struct{ v, t int }
	q := []pair{{0, 0}}
	vis := map[int]bool{0: true}
	for len(q) > 0 {
		p := q[0]
		q = q[1:]
		cur, t := p.v, p.t
		if cur == (1<<(n+1) - 2) {
			return t
		}
		nxt := 0
		for i := 1; i <= n; i++ {
			if (cur & d[i]) == d[i] {
				nxt |= 1 << i
			}
		}
		nxt ^= cur
		if bits.OnesCount(uint(nxt)) <= k {
			if !vis[nxt|cur] {
				vis[nxt|cur] = true
				q = append(q, pair{nxt | cur, t + 1})
			}
		} else {
			x := nxt
			for nxt > 0 {
				if bits.OnesCount(uint(nxt)) == k && !vis[nxt|cur] {
					vis[nxt|cur] = true
					q = append(q, pair{nxt | cur, t + 1})
				}
				nxt = (nxt - 1) & x
			}
		}
	}
	return 0
}

type dsu struct {
	p, size []int
}

func newDSU(n int) *dsu {
	p := make([]int, n)
	size := make([]int, n)
	for i := range p {
		p[i] = i
		size[i] = 1
	}
	return &dsu{p, size}
}

func (d *dsu) find(x int) int {
	if d.p[x] != x {
		d.p[x] = d.find(d.p[x])
	}
	return d.p[x]
}

func (d *dsu) unite(x, y int) bool {
	x, y = d.find(x), d.find(y)
	if x == y {
		return false
	}
	if d.size[x] < d.size[y] {
		x, y = y, x
	}
	d.p[y] = x
	d.size[x] += d.size[y]
	return true
}

func (d *dsu) sz(x int) int {
	return d.size[x]
}

func MinMalwareSpread(graph [][]int, initial []int) int {
	n := len(graph)
	d := newDSU(n)
	for i := range graph {
		for j := 0; j < n; j++ {
			if graph[i][j] == 1 {
				d.unite(i, j)
			}
		}
	}
	cnt := make([]int, n)
	for _, x := range initial {
		cnt[d.find(x)]++
	}
	ans, mx := n, 0
	for _, x := range initial {
		root := d.find(x)
		if cnt[root] == 1 {
			z := d.sz(root)
			if z > mx || z == mx && ans > x {
				ans = x
				mx = z
			}
		}
	}
	if ans == n {
		return slices.Min(initial)
	}
	return ans
}

func MinCost(maxTime int, edges [][]int, passingFees []int) int {
	m, n := maxTime, len(passingFees)
	inf := 1 << 30
	f := make([][]int, m+1)
	for i := range f {
		f[i] = make([]int, n)
		for j := range f[i] {
			f[i][j] = inf
		}
	}
	f[0][0] = passingFees[0]
	for i := 1; i <= m; i++ {
		for _, e := range edges {
			x, y, t := e[0], e[1], e[2]
			if t <= i {
				f[i][x] = min(f[i][x], f[i-t][y]+passingFees[x])
				f[i][y] = min(f[i][y], f[i-t][x]+passingFees[y])
			}
		}
	}
	ans := inf
	for i := 1; i <= m; i++ {
		ans = min(ans, f[i][n-1])
	}
	if ans == inf {
		return -1
	}
	return ans
}

func NumOfGoodPaths(vals []int, edges [][]int) int {
	n := len(vals)
	pr := make([]int, n)
	size := map[int]map[int]int{}
	type pair struct{ v, i int }
	arr := make([]pair, n)
	var find func(x int) int
	find = func(x int) int {
		if x != pr[x] {
			pr[x] = find(pr[x])
		}
		return pr[x]
	}
	for i, v := range vals {
		pr[i] = i
		if size[i] == nil {
			size[i] = map[int]int{}
		}
		size[i][v] = 1
		arr[i] = pair{v, i}
	}
	sort.Slice(arr, func(i, j int) bool { return arr[i].v < arr[j].v })
	g := make([][]int, n)
	for _, e := range edges {
		a, b := e[0], e[1]
		g[a] = append(g[a], b)
		g[b] = append(g[b], a)
	}
	ans := n
	for _, e := range arr {
		v, x := e.v, e.i
		for _, y := range g[x] {
			if vals[y] > v {
				continue
			}
			x, y = find(x), find(y)
			if x != y {
				ans += size[x][v] * size[y][v]
				pr[y] = x
				size[x][v] += size[y][x]
			}

		}
	}
	return ans
}
