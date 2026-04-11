package arrays

import (
	"reflect"
	"testing"
)

func TestThreeSum(t *testing.T) {
	nums := []int{-1, 0, 1, 2, -1, -4}
	want := [][]int{
		{-1, -1, 2},
		{-1, 0, 1},
	}

	got := ThreeSum(nums)
	if !reflect.DeepEqual(got, want) {
		t.Fatalf("unexpected triplets: got %v want %v", got, want)
	}
}

func TestThreeSumTargetExists(t *testing.T) {
	nums := []int{1, 4, 45, 6, 10, 8}

	if !ThreeSumTargetExists(nums, 13) {
		t.Fatal("expected target 13 to exist")
	}

	if ThreeSumTargetExists(nums, 100) {
		t.Fatal("did not expect target 100 to exist")
	}
}

func TestArrayChange(t *testing.T) {
	nums := []int{1, 2, 4, 6}
	operations := [][2]int{
		{1, 3},
		{4, 7},
		{6, 1},
	}

	want := []int{3, 2, 7, 1}
	got := ArrayChange(nums, operations)

	if !reflect.DeepEqual(got, want) {
		t.Fatalf("unexpected array after changes: got %v want %v", got, want)
	}
}
