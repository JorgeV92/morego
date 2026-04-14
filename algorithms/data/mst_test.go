package data

import (
	"io"
	"os"
	"reflect"
	"sort"
	"strings"
	"testing"
)

func TestMst(t *testing.T) {
	testKruskalImplementation(t, Kruskal)
}

func TestMstDSU(t *testing.T) {
	testKruskalImplementation(t, KruskalDSU)
}

func TestPrimPrintsMSTForConnectedGraph(t *testing.T) {
	adj := [][]int{
		{0, 10, 6, 5},
		{10, 0, INF, 15},
		{6, INF, 0, 4},
		{5, 15, 4, 0},
	}

	got := captureStdout(t, func() {
		Prim(adj)
	})

	want := "3 0\n2 3\n1 0\n19\n"
	if got != want {
		t.Fatalf("unexpected Prim output: got %q want %q", got, want)
	}
}

func TestPrimReportsNoMSTForDisconnectedGraph(t *testing.T) {
	adj := [][]int{
		{0, 7, INF},
		{7, 0, INF},
		{INF, INF, 0},
	}

	got := captureStdout(t, func() {
		Prim(adj)
	})

	if !strings.Contains(got, "No MST\n") {
		t.Fatalf("expected Prim to report missing MST, got %q", got)
	}
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

func captureStdout(t *testing.T, fn func()) string {
	t.Helper()

	oldStdout := os.Stdout
	reader, writer, err := os.Pipe()
	if err != nil {
		t.Fatalf("failed to create stdout pipe: %v", err)
	}

	os.Stdout = writer
	t.Cleanup(func() {
		os.Stdout = oldStdout
	})

	fn()

	if err := writer.Close(); err != nil {
		t.Fatalf("failed to close stdout writer: %v", err)
	}

	output, err := io.ReadAll(reader)
	if err != nil {
		t.Fatalf("failed to read stdout: %v", err)
	}

	if err := reader.Close(); err != nil {
		t.Fatalf("failed to close stdout reader: %v", err)
	}

	os.Stdout = oldStdout
	return string(output)
}
