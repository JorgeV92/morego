package data

import (
	"math/rand"
	"reflect"
	"testing"
)

func TestTreapInsertMaintainsSortedUniqueValues(t *testing.T) {
	rand.Seed(1)

	var root *Node
	for _, value := range []int{5, 1, 3, 3, 2, 4} {
		root = Insert(root, value)
	}

	got := inorderValues(root)
	want := []int{1, 2, 3, 4, 5}

	if !reflect.DeepEqual(got, want) {
		t.Fatalf("unexpected inorder values: got %v want %v", got, want)
	}

	if size := GetSize(root); size != 5 {
		t.Fatalf("unexpected treap size: got %d want %d", size, 5)
	}

	if sum := GetSum(root); sum != 15 {
		t.Fatalf("unexpected treap sum: got %d want %d", sum, 15)
	}
}

func TestTreapEraseRemovesValue(t *testing.T) {
	rand.Seed(1)

	var root *Node
	for _, value := range []int{1, 2, 3, 4, 5} {
		root = Insert(root, value)
	}

	root = erase(root, 3)

	got := inorderValues(root)
	want := []int{1, 2, 4, 5}

	if !reflect.DeepEqual(got, want) {
		t.Fatalf("unexpected values after erase: got %v want %v", got, want)
	}

	if size := GetSize(root); size != 4 {
		t.Fatalf("unexpected size after erase: got %d want %d", size, 4)
	}

	if sum := GetSum(root); sum != 12 {
		t.Fatalf("unexpected sum after erase: got %d want %d", sum, 12)
	}
}

func TestTreapSplitAndMergeByValue(t *testing.T) {
	rand.Seed(1)

	var root *Node
	for _, value := range []int{1, 2, 3, 4, 5} {
		root = Insert(root, value)
	}

	left, right := Split(root, 3)

	if got := inorderValues(left); !reflect.DeepEqual(got, []int{1, 2}) {
		t.Fatalf("unexpected left split values: got %v want %v", got, []int{1, 2})
	}

	if got := inorderValues(right); !reflect.DeepEqual(got, []int{3, 4, 5}) {
		t.Fatalf("unexpected right split values: got %v want %v", got, []int{3, 4, 5})
	}

	merged := Merge(left, right)
	if got := inorderValues(merged); !reflect.DeepEqual(got, []int{1, 2, 3, 4, 5}) {
		t.Fatalf("unexpected merged values: got %v want %v", got, []int{1, 2, 3, 4, 5})
	}
}

func TestTreapSplitBySize(t *testing.T) {
	rand.Seed(1)

	var root *Node
	for _, value := range []int{1, 2, 3, 4, 5} {
		root = Insert(root, value)
	}

	left, right := SplitSz(root, 2)

	if got := inorderValues(left); !reflect.DeepEqual(got, []int{1, 2}) {
		t.Fatalf("unexpected left split-by-size values: got %v want %v", got, []int{1, 2})
	}

	if got := inorderValues(right); !reflect.DeepEqual(got, []int{3, 4, 5}) {
		t.Fatalf("unexpected right split-by-size values: got %v want %v", got, []int{3, 4, 5})
	}
}

func inorderValues(root *Node) []int {
	values := make([]int, 0)
	Inorder(root, &values)
	return values
}
