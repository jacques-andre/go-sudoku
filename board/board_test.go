package board

import (
	"fmt"
	"testing"
)

// Test a VALID pos in row
func TestValidInBoardRow(t *testing.T) {
	// Blank board, should be able to place anywhere without failing
	grid := [9][9]int{}
	board := Board{SolvedBoard: grid}

	// Test vars
	valueN := 9
	rowN := 0

	got := board.ValidPosInRow(valueN, rowN, board.SolvedBoard)
	want := true

	fmt.Printf("Row:%v, checking:%d", board.SolvedBoard[rowN], valueN)

	if got != want {
		t.Errorf("Failed! Got:%t, want:%t", got, want)
	}
}

// Test a INVALID pos in row
func TestInValidInBoardRow(t *testing.T) {
	grid := [9][9]int{}

	// Make the first value 1,
	// Testing if we can overwrite this (false)
	grid[0] = [9]int{1, 0, 0, 0, 0, 0, 0, 0, 0}
	board := Board{SolvedBoard: grid}

	// Test vars
	valueN := 1
	rowN := 0

	got := board.ValidPosInRow(valueN, rowN, board.SolvedBoard)
	want := false

	fmt.Printf("Row:%v, checking:%d", board.SolvedBoard[rowN], valueN)

	if got != want {
		t.Errorf("Failed! Got:%t, want:%t", got, want)
	}
}

// Test a VALID pos in a col
func TestValidInBoardCol(t *testing.T) {
	// Blank board, should be able to place anywhere without failing
	grid := [9][9]int{}
	board := Board{SolvedBoard: grid}

	// Test vars
	valueN := 7
	colN := 0

	got := board.ValidPosInCol(valueN, colN, board.SolvedBoard)
	want := true

	board.PrintBoard(board.SolvedBoard)
	fmt.Printf("Checking:%d, col:%d", valueN, colN)

	if got != want {
		t.Errorf("Failed! Got:%t, want:%t", got, want)
	}
}

// Test a INVALID pos in a col
func TestInValidInBoardCol(t *testing.T) {
	grid := [9][9]int{}

	// Make the first value 1,
	// Testing if we can overwrite this (false)
	grid[0] = [9]int{1, 0, 0, 0, 0, 0, 0, 0, 0}
	board := Board{SolvedBoard: grid}

	// Test vars
	valueN := 1
	colN := 0

	got := board.ValidPosInCol(valueN, colN, board.SolvedBoard)
	want := false

	board.PrintBoard(board.SolvedBoard)
	fmt.Printf("Checking:%d, col:%d", valueN, colN)

	if got != want {
		t.Errorf("Failed! Got:%t, want:%t", got, want)
	}
}

// Test a VALID pos in subgrid (3x3)
func TestValidInBoardSubGrid(t *testing.T) {
	// Blank board, should be able to place anywhere without failing
	grid := [9][9]int{}
	board := Board{SolvedBoard: grid}

	// Test vars
	valueN := 9
	colN := 0
	rowN := 0

	got := board.ValidPosInSubGrid(valueN, colN, rowN, board.SolvedBoard)
	want := true

	board.PrintBoard(board.SolvedBoard)
	fmt.Printf("Checking:%d,col:%d,row:%d \n", valueN, colN, rowN)

	if got != want {
		t.Errorf("Failed! Got:%t, want:%t", got, want)
	}
}

// Test a INVALID pos in a subgrid (3x3)
func TestInValidInBoardSubGrid(t *testing.T) {
	grid := [9][9]int{}

	// Make the first value 1,
	// Testing if we can overwrite this (false)
	grid[0] = [9]int{1, 0, 0, 0, 0, 0, 0, 0, 0}
	board := Board{SolvedBoard: grid}

	// Test vars
	valueN := 1
	colN := 0
	rowN := 0

	board.PrintBoard(board.SolvedBoard)
	fmt.Printf("Checking:%d,col:%d,row:%d \n", valueN, colN, rowN)

	got := board.ValidPosInSubGrid(valueN, colN, rowN, board.SolvedBoard)
	want := false

	if got != want {
		t.Errorf("Failed! Got:%t, want:%t", got, want)
		board.PrintBoard(board.SolvedBoard)
		checkingPos := board.SolvedBoard[rowN][colN]
		fmt.Printf("checkingPos:%d", checkingPos)
	}
}
