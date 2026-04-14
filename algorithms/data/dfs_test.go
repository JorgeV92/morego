package data

import (
	"reflect"
	"testing"
)

func TestDfsMarksReachableVertices(t *testing.T) {
	adj := [][]int{
		{1, 2},
		{3},
		{},
		{},
		{5},
		{},
	}
	vis := make([]bool, len(adj))

	Dfs(0, vis, adj)

	want := []bool{true, true, true, true, false, false}
	if !reflect.DeepEqual(vis, want) {
		t.Fatalf("unexpected visited state: got %v want %v", vis, want)
	}
}

func TestDfs2TracksTimesAndColors(t *testing.T) {
	adj := [][]int{
		{1, 2},
		{3},
		{},
		{},
		{},
	}

	timeIn := []int{-1, -1, -1, -1, -1}
	timeOut := []int{-1, -1, -1, -1, -1}
	color := make([]int, len(adj))
	timer := 0

	Dfs2(0, timeIn, timeOut, &timer, color, adj)

	wantTimeIn := []int{0, 1, 5, 2, -1}
	wantTimeOut := []int{7, 4, 6, 3, -1}
	wantColor := []int{2, 2, 2, 2, 0}

	if !reflect.DeepEqual(timeIn, wantTimeIn) {
		t.Fatalf("unexpected entry times: got %v want %v", timeIn, wantTimeIn)
	}

	if !reflect.DeepEqual(timeOut, wantTimeOut) {
		t.Fatalf("unexpected exit times: got %v want %v", timeOut, wantTimeOut)
	}

	if !reflect.DeepEqual(color, wantColor) {
		t.Fatalf("unexpected colors: got %v want %v", color, wantColor)
	}

	if timer != 8 {
		t.Fatalf("unexpected timer value: got %d want %d", timer, 8)
	}
}
