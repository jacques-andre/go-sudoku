package main

import (
	"fmt"
	"testing"
)

func TestValidInBoardRow(t *testing.T) {
	grid := [9][9]int{}
	board := Board{BoardArray: grid}

	// Test vars
	valueN := 9
	rowN := 0

	got := board.validPosInRow(valueN, rowN)
	want := true

	fmt.Printf("Row:%v, checking:%d", board.BoardArray[rowN], valueN)

	if got != want {
		t.Errorf("Failed! Got:%t, want:%t", got, want)
	}
}

func TestInValidInBoardRow(t *testing.T) {
	grid := [9][9]int{}

	// Make the first value 1,
	// Testing if we can overwrite this (false)
	grid[0] = [9]int{1, 0, 0, 0, 0, 0, 0, 0, 0}
	board := Board{BoardArray: grid}

	// Test vars
	valueN := 1
	rowN := 0

	got := board.validPosInRow(valueN, rowN)
	want := false

	fmt.Printf("Row:%v, checking:%d", board.BoardArray[rowN], valueN)

	if got != want {
		t.Errorf("Failed! Got:%t, want:%t", got, want)
	}
}

func TestValidInBoardCol(t *testing.T) {
	grid := [9][9]int{}
	board := Board{BoardArray: grid}
	// Test vars
	valueN := 7
	colN := 0

	got := board.validPosInCol(valueN, colN)
	want := true

	board.printBoard()
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
	board := Board{BoardArray: grid}

	// Test vars
	valueN := 1
	colN := 0

	got := board.validPosInCol(valueN, colN)
	want := false

	board.printBoard()
	fmt.Printf("Checking:%d, col:%d", valueN, colN)

	if got != want {
		t.Errorf("Failed! Got:%t, want:%t", got, want)
	}
}

func TestValidInBoardSubGrid(t *testing.T) {
	grid := [9][9]int{}
	board := Board{BoardArray: grid}

	// Test vars
	valueN := 9
	colN := 0
	rowN := 0

	got := board.validPosInSubGrid(valueN, colN, rowN)
	want := true

	board.printBoard()
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
	board := Board{BoardArray: grid}

	// Test vars
	valueN := 1
	colN := 0
	rowN := 0

	got := board.validPosInSubGrid(valueN, colN, rowN)
	want := false

	if got != want {
		t.Errorf("Failed! Got:%t, want:%t", got, want)
		board.printBoard()
		checkingPos := board.BoardArray[rowN][colN]
		fmt.Printf("checkingPos:%d", checkingPos)
	}
}

// TODO
func TestValidBoard(t *testing.T) {
	grid := [9][9]int{}
	board := Board{BoardArray: grid}
	board.solveBoard(1, 0, 0)

	board.printBoard()

}
func contains(s []int, e int) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}
func uniqueInRow(row []int) bool {
	seenInRow := []int{}
	for i := 0; i < len(row); i++ {
		currentVal := row[i]

		if contains(seenInRow, currentVal) {
			return false
		} else {
			seenInRow = append(seenInRow, currentVal)
		}
	}
	return true
}
