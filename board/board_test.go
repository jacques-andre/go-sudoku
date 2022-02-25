package board

import (
	"fmt"
	"testing"
)

func TestValidInBoardRow(t *testing.T) {
	grid := [9][9]int{}
	board := Board{SolvedBoard: grid}

	// Test vars
	valueN := 9
	rowN := 0

	got := board.ValidPosInRow(valueN, rowN)
	want := true

	fmt.Printf("Row:%v, checking:%d", board.SolvedBoard[rowN], valueN)

	if got != want {
		t.Errorf("Failed! Got:%t, want:%t", got, want)
	}
}

func TestInValidInBoardRow(t *testing.T) {
	grid := [9][9]int{}

	// Make the first value 1,
	// Testing if we can overwrite this (false)
	grid[0] = [9]int{1, 0, 0, 0, 0, 0, 0, 0, 0}
	board := Board{SolvedBoard: grid}

	// Test vars
	valueN := 1
	rowN := 0

	got := board.ValidPosInRow(valueN, rowN)
	want := false

	fmt.Printf("Row:%v, checking:%d", board.SolvedBoard[rowN], valueN)

	if got != want {
		t.Errorf("Failed! Got:%t, want:%t", got, want)
	}
}

func TestValidInBoardCol(t *testing.T) {
	grid := [9][9]int{}
	board := Board{SolvedBoard: grid}
	// Test vars
	valueN := 7
	colN := 0

	got := board.ValidPosInCol(valueN, colN)
	want := true

	board.PrintBoard()
	fmt.Printf("Checking:%d, col:%d", valueN, colN)

	if got != want {
		t.Errorf("Failed! Got:%t, want:%t", got, want)
	}
}

func TestInValidInBoardCol(t *testing.T) {
	grid := [9][9]int{}

	// Make the first value 1,
	// Testing if we can overwrite this (false)
	grid[0] = [9]int{1, 0, 0, 0, 0, 0, 0, 0, 0}
	board := Board{SolvedBoard: grid}

	// Test vars
	valueN := 1
	colN := 0

	got := board.ValidPosInCol(valueN, colN)
	want := false

	board.PrintBoard()
	fmt.Printf("Checking:%d, col:%d", valueN, colN)

	if got != want {
		t.Errorf("Failed! Got:%t, want:%t", got, want)
	}
}

func TestValidInBoardSubGrid(t *testing.T) {
	grid := [9][9]int{}
	board := Board{SolvedBoard: grid}

	// Test vars
	valueN := 9
	colN := 0
	rowN := 0

	got := board.ValidPosInSubGrid(valueN, colN, rowN)
	want := true

	board.PrintBoard()
	fmt.Printf("Checking:%d,col:%d,row:%d \n", valueN, colN, rowN)

	if got != want {
		t.Errorf("Failed! Got:%t, want:%t", got, want)
	}
}

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

	board.PrintBoard()
	fmt.Printf("Checking:%d,col:%d,row:%d \n", valueN, colN, rowN)

	got := board.ValidPosInSubGrid(valueN, colN, rowN)
	want := false

	if got != want {
		t.Errorf("Failed! Got:%t, want:%t", got, want)
		board.PrintBoard()
		checkingPos := board.SolvedBoard[rowN][colN]
		fmt.Printf("checkingPos:%d", checkingPos)
	}
}

// TODO
// func TestValidBoard(t *testing.T) {
// 	grid := [9][9]int{}
// 	board := Board{SolvedBoard: grid}
// 	board.GenerateBoard(1, 0)

// 	// Go through all rows check if valid row
// 	for i := 0; i < len(board.SolvedBoard); i++ {
// 		row := board.SolvedBoard[i]
// 		if !uniqueInRow(row) {
// 			t.Errorf("Invalid In row: %v", row)
// 		}
// 	}

// 	board.PrintBoard()

// }
func contains(s []int, e int) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}

func uniqueInRow(row [9]int) bool {
	seenInRow := []int{}
	for i := 0; i < len(row); i++ {
		currentVal := row[i]

		if contains(seenInRow, currentVal) {
			return false
		} else {
			seenInRow = append(seenInRow, currentVal)
		}
	}
	fmt.Printf("Seen: %v \n \n", seenInRow)
	return true
}

func uniqueInCol(col [9]int) bool {
	seenInCol := []int{}
	for i := 0; i < len(col); i++ {
		currentVal := col[i]

		if contains(seenInCol, currentVal) {
			return false
		} else {
			seenInCol = append(seenInCol, currentVal)
		}
	}
	return true
}
