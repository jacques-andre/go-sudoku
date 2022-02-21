package main

import (
	"fmt"
)

type Board struct {
	BoardArray [9][9]int
}

// Generates a valid board with some blank spaces,
// user playable board
func newPlayingBoard() *Board {
	// Create an inital test board
	validBoard := [9][9]int{
		{3, 0, 6, 5, 0, 8, 4, 0, 0},
		{5, 2, 0, 0, 0, 0, 0, 0, 0},
		{0, 8, 7, 0, 0, 0, 0, 3, 1},
		{0, 0, 3, 0, 1, 0, 0, 8, 0},
		{9, 0, 0, 8, 6, 3, 0, 0, 5},
		{0, 5, 0, 0, 9, 0, 6, 0, 0},
		{1, 3, 0, 0, 0, 0, 2, 5, 0},
		{0, 0, 0, 0, 0, 0, 0, 7, 4},
		{0, 0, 5, 2, 0, 6, 3, 0, 0},
	}
	// Return new Board struct with board ^ as value
	return &Board{
		BoardArray: validBoard,
	}
}

// Check if cellVal can be placed in the row rowN
// Returns true if valid pos
func (board *Board) validPosInRow(cellVal int, rowN int) bool {
	for i := 0; i < len(board.BoardArray[rowN]); i++ {
		currentValue := &board.BoardArray[rowN][i]
		if *currentValue == cellVal {
			return false
		}
	}
	return true
}

// Check if cellVal can be placed in the column colN
// Returns true if valid pos
func (board *Board) validPosInCol(cellVal int, colN int) bool {
	for i := 0; i < len(board.BoardArray[colN]); i++ {
		currentValue := &board.BoardArray[i][colN]
		if *currentValue == cellVal {
			return false
		}
	}
	return true
}

// Check if cellVal can be placed in the 3x3 subgrid
// Returns true if valid pos
func (board *Board) validPosInSubGrid(cellVal int, colN int, rowN int) bool {
	rowStart := (rowN / 3) * 3
	colStart := (colN / 3) * 3

	for i := colStart; i < colStart+3; i++ {
		for j := rowStart; j < rowStart+3; j++ {
			subGridCell := &board.BoardArray[i][j]
			if *subGridCell == cellVal {
				return false
			}
		}
	}
	return true

}

// Helper Method:
// Prints board nicely
func printBoard(board *Board) {
	for _, v := range board.BoardArray {
		fmt.Println(v)
	}
}

func main() {
	board := newPlayingBoard()

	// place(board)
	printBoard(board)
}
