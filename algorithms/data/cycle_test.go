package data

import (
	"reflect"
	"testing"
)

func TestFindCycleReturnsCycle(t *testing.T) {
	adj := [][]int{
		{1},
		{2},
		{0, 3},
		{},
	}

	got, ok := findCycle(adj)
	want := []int{0, 1, 2, 0}

	if !ok {
		t.Fatal("expected a cycle, but none was found")
	}

	if !reflect.DeepEqual(got, want) {
		t.Fatalf("unexpected cycle: got %v want %v", got, want)
	}
}

func TestFindCycleReturnsFalseForDAG(t *testing.T) {
	adj := [][]int{
		{1, 2},
		{3},
		{3},
		{},
	}

	got, ok := findCycle(adj)
	if ok {
		t.Fatalf("expected no cycle, but got %v", got)
	}

	if got != nil {
		t.Fatalf("expected nil cycle for DAG, got %v", got)
	}
}

func TestFindCycleInDisconnectedGraph(t *testing.T) {
	adj := [][]int{
		{1},
		{},
		{3},
		{4},
		{2},
	}

	got, ok := findCycle(adj)
	want := []int{2, 3, 4, 2}

	if !ok {
		t.Fatal("expected a cycle in the second component, but none was found")
	}

	if !reflect.DeepEqual(got, want) {
		t.Fatalf("unexpected cycle in disconnected graph: got %v want %v", got, want)
	}
}
