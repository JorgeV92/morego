package design

import (
	"fmt"
	"testing"
)

func TestTaskManagerExecTopUsesPriorityAndTaskID(t *testing.T) {
	manager := Constructor([][]int{
		{1, 101, 5},
		{2, 202, 5},
		{3, 303, 7},
	})

	if got := manager.ExecTop(); got != 3 {
		t.Fatalf("unexpected first user: got %d want %d", got, 3)
	}

	if got := manager.ExecTop(); got != 2 {
		t.Fatalf("unexpected tie-break user: got %d want %d", got, 2)
	}

	if got := manager.ExecTop(); got != 1 {
		t.Fatalf("unexpected third user: got %d want %d", got, 1)
	}

	if got := manager.ExecTop(); got != -1 {
		t.Fatalf("expected empty task manager to return -1, got %d", got)
	}
}

func TestTaskManagerEditAndRemove(t *testing.T) {
	manager := Constructor([][]int{
		{1, 101, 2},
		{2, 202, 4},
	})

	manager.Edit(101, 6)
	manager.Rmv(202)

	if got := manager.ExecTop(); got != 1 {
		t.Fatalf("unexpected user after edit/remove: got %d want %d", got, 1)
	}

	if got := manager.ExecTop(); got != -1 {
		t.Fatalf("expected empty task manager to return -1, got %d", got)
	}
}

func TestTaskManagerIgnoresStaleHeapEntries(t *testing.T) {
	manager := Constructor(nil)

	manager.Add(1, 101, 3)
	manager.Edit(101, 8)
	manager.Edit(101, 5)

	if got := manager.ExecTop(); got != 1 {
		t.Fatalf("unexpected user for updated task: got %d want %d", got, 1)
	}

	if got := manager.ExecTop(); got != -1 {
		t.Fatalf("expected empty task manager to return -1, got %d", got)
	}
}

func ExampleTaskManager() {
	manager := Constructor([][]int{
		{1, 101, 5},
		{2, 202, 3},
	})

	manager.Add(3, 303, 6)
	fmt.Println(manager.ExecTop())

	manager.Edit(202, 7)
	fmt.Println(manager.ExecTop())

	manager.Rmv(101)
	fmt.Println(manager.ExecTop())

	// Output:
	// 3
	// 2
	// -1
}
