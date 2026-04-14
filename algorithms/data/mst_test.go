package data

import (
	"reflect"
	"sort"
	"testing"
)

func TestMst(t *testing.T) {
	testKruskalImplementation(t, Kruskal)
}

func TestMstDSU(t *testing.T) {
	testKruskalImplementation(t, KruskalDSU)
}

func testKruskalImplementation(t *testing.T, kruskal func(int, []Edge) (int, []Edge)) {
	t.Helper()

	n := 4
	edges := []Edge{
		{0, 1, 10},
		{0, 2, 6},
		{0, 3, 5},
		{1, 3, 15},
		{2, 3, 4},
	}

	cost, mst := kruskal(n, append([]Edge(nil), edges...))

	wantCost := 19
	wantMST := []Edge{
		{2, 3, 4},
		{0, 3, 5},
		{0, 1, 10},
	}

	sortEdges(mst)
	sortEdges(wantMST)

	if cost != wantCost {
		t.Fatalf("unexpected MST cost: got %d want %d", cost, wantCost)
	}

	if !reflect.DeepEqual(mst, wantMST) {
		t.Fatalf("unexpected MST edges: got %+v want %+v", mst, wantMST)
	}
}

func sortEdges(edges []Edge) {
	sort.Slice(edges, func(i, j int) bool {
		if edges[i].W != edges[j].W {
			return edges[i].W < edges[j].W
		}
		if edges[i].U != edges[j].U {
			return edges[i].U < edges[j].U
		}
		return edges[i].V < edges[j].V
	})
}
