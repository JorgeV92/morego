package data

func findCycle(adj [][]int) ([]int, bool) {
	n := len(adj)

	color := make([]int, n)
	parent := make([]int, n)
	for i := range parent {
		parent[i] = -1
	}

	cycleStart, cycleEnd := -1, -1

	var dfs func(v int) bool
	dfs = func(v int) bool {
		color[v] = 1

		for _, u := range adj[v] {
			if color[u] == 0 {
				parent[u] = v
				if dfs(u) {
					return true
				}
			} else if color[u] == 1 {
				cycleStart = u
				cycleEnd = v
				return true
			}
		}

		color[v] = 2
		return false
	}

	for v := 0; v < n; v++ {
		if color[v] == 0 && dfs(v) {
			break
		}
	}

	if cycleStart == -1 {
		return nil, false
	}

	cycle := []int{cycleStart}
	for v := cycleEnd; v != cycleStart; v = parent[v] {
		cycle = append(cycle, v)
	}
	cycle = append(cycle, cycleStart)

	// reverse
	for i, j := 0, len(cycle)-1; i < j; i, j = i+1, j-1 {
		cycle[i], cycle[j] = cycle[j], cycle[i]
	}

	return cycle, true
}
