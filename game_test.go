package main

import (
	"fmt"
	"testing"
)

func TestValidInBoardRow(t *testing.T) {
	board := newPlayingBoard()

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
	board := newPlayingBoard()
	// Test vars
	valueN := 3
	rowN := 0

	got := board.validPosInRow(valueN, rowN)
	want := false

	fmt.Printf("Row:%v, checking:%d", board.BoardArray[rowN], valueN)

	if got != want {
		t.Errorf("Failed! Got:%t, want:%t", got, want)
	}
}

func TestValidInBoardCol(t *testing.T) {
	board := newPlayingBoard()
	// Test vars
	valueN := 7
	colN := 0

	got := board.validPosInCol(valueN, colN)
	want := true

	printBoard(board)
	fmt.Printf("Checking:%d, col:%d", valueN, colN)

	if got != want {
		t.Errorf("Failed! Got:%t, want:%t", got, want)
	}
}

func TestInValidInBoardCol(t *testing.T) {
	board := newPlayingBoard()

	// Test vars
	valueN := 9
	colN := 0

	got := board.validPosInCol(valueN, colN)
	want := false

	fmt.Printf("Checking:%d, col:%d", valueN, colN)
	printBoard(board)

	if got != want {
		t.Errorf("Failed! Got:%t, want:%t", got, want)
	}
}

func TestValidInBoardSubGrid(t *testing.T) {
	board := newPlayingBoard()

	// Test vars
	valueN := 9
	colN := 0
	rowN := 0

	got := board.validPosInSubGrid(valueN, colN, rowN)
	want := true

  fmt.Printf("Checking:%d,col:%d,row:%d \n", valueN, colN, rowN)
	printBoard(board)

	if got != want {
		t.Errorf("Failed! Got:%t, want:%t", got, want)
	}
}

func TestInValidInBoardSubGrid(t *testing.T) {
	board := newPlayingBoard()

	// Test vars
	valueN := 3
	colN := 0
	rowN := 0

	got := board.validPosInSubGrid(valueN, colN, rowN)
	want := false

  fmt.Printf("Checking:%d,col:%d,row:%d \n", valueN, colN, rowN)
	printBoard(board)

	if got != want {
		t.Errorf("Failed! Got:%t, want:%t", got, want)
	}
}
