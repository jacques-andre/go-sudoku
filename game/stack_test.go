package game

import (
	"sudoku/board"
	"testing"
)

// Test able to add to stack
func TestAddToStack(t *testing.T) {
	grid := [9][9]int{}
	board := board.Board{UserBoard: grid, SolvedBoard: grid}

	// Create a emptyStack
	testStack := Stack{}
	// Push board to stack
	testStack.Push(board)

	// Expecting one item to be added to the stack
	got := len(testStack)
	want := 1

	if got != want {
		t.Errorf("Failed! Got:%d, want:%d", got, want)
	}
}

// Test able to pop from the stack
func TestPopStack(t *testing.T) {
	grid := [9][9]int{}
	board := board.Board{UserBoard: grid, SolvedBoard: grid}

	// Create a emptyStack
	testStack := Stack{}
	// Push board to stack
	testStack.Push(board)
	testStack.Pop()

	// Expecting no items to be added to the stack, (popped from stack)
	got := len(testStack)
	want := 0

	if got != want {
		t.Errorf("Failed! Got:%d, want:%d", got, want)
	}
}

// Test IsEmpty
func TestIsEmpty(t *testing.T) {
	grid := [9][9]int{}
	board := board.Board{UserBoard: grid, SolvedBoard: grid}

	// Create a emptyStack
	testStack := Stack{}
	// Push board to stack
	testStack.Push(board)
	testStack.Pop()

	// Expecting no items to be added to the stack, (popped from stack)
	got := testStack.IsEmpty()
	want := true

	if got != want {
		t.Errorf("Failed! Got:%v, want:%v", got, want)
	}
}
