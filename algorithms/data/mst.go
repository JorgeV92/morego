package data

import "sort"

type Edge struct {
	U, V int
	W    int
}

func Kruskal(n int, edges []Edge) (int, []Edge) {
	// time O(M log N + N^2)
	cost := 0
	treeID := make([]int, n)
	result := make([]Edge, 0)

	for i := 0; i < n; i++ {
		treeID[i] = i
	}

	sort.Slice(edges, func(i, j int) bool {
		return edges[i].W < edges[j].W
	})

	for _, e := range edges {
		if treeID[e.U] != treeID[e.V] {
			cost += e.W
			result = append(result, e)
			oldID := treeID[e.U]
			newID := treeID[e.V]

			for i := 0; i < n; i++ {
				if treeID[i] == oldID {
					treeID[i] = newID
				}
			}
		}
	}
	return cost, result
}

type DSU struct {
	parent []int
	rank   []int
	size   []int
}

func NewDSU(n int) *DSU {
	d := &DSU{
		parent: make([]int, n),
		rank:   make([]int, n),
		size:   make([]int, n),
	}
	for i := 0; i < n; i++ {
		d.parent[i] = i
		d.rank[i] = 0
		d.size[i] = 1
	}
	return d
}

func (d *DSU) findSet(v int) int {
	if v == d.parent[v] {
		return v
	}
	d.parent[v] = d.findSet(d.parent[v])
	return d.parent[v]
}

func (d *DSU) unionSets(a, b int) {
	a = d.findSet(a)
	b = d.findSet(b)

	if a != b {
		if d.rank[a] < d.rank[b] {
			a, b = b, a
		}
		d.parent[b] = a
		d.size[a] += d.size[b]
		if d.rank[a] == d.rank[b] {
			d.rank[a]++
		}
	}
}

func KruskalDSU(n int, edges []Edge) (int, []Edge) {
	cost := 0
	result := make([]Edge, 0)
	dsu := NewDSU(n)

	sort.Slice(edges, func(i, j int) bool {
		return edges[i].W < edges[j].W
	})

	for _, e := range edges {
		if dsu.findSet(e.U) != dsu.findSet(e.V) {
			cost += e.W
			result = append(result, e)
			dsu.unionSets(e.U, e.V)
		}
	}
	return cost, result
}
