package data

func Dfs(v int, vis []bool, adj [][]int) {
	vis[v] = true
	for _, u := range adj[v] {
		if !vis[u] {
			Dfs(u, vis, adj)
		}
	}
}

func Dfs2(v int, timeIn []int, timeOut []int, dfsTimer *int, color []int, adj [][]int) {

	timeIn[v] = *dfsTimer
	*dfsTimer = *dfsTimer + 1
	color[v] = 1
	for _, u := range adj[v] {
		if color[u] == 0 {
			Dfs2(u, timeIn, timeOut, dfsTimer, color, adj)
		}
	}
	color[v] = 2
	timeOut[v] = *dfsTimer
	*dfsTimer = *dfsTimer + 1
}
