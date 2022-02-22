package main

import (
	"fmt"
	"sudoku/board"
)

func main() {
	// grid := [9][9]int{
	// {0, 0, 6, 7, 0, 0, 0, 8, 5},
	// {0, 8, 0, 0, 0, 0, 9, 0, 0},
	// {0, 0, 0, 0, 5, 0, 1, 0, 0},
	// {3, 0, 0, 9, 0, 0, 0, 0, 0},
	// {0, 0, 0, 0, 0, 0, 0, 0, 1},
	// {0, 4, 0, 0, 3, 0, 0, 5, 6},
	// {0, 6, 0, 0, 9, 0, 0, 4, 3},
	// {0, 0, 0, 0, 0, 0, 8, 0, 0},
	// {0, 0, 7, 0, 0, 2, 0, 0, 0},
	// }
	grid := [9][9]int{}

	board := board.Board{BoardArray: grid}
	board.GenerateBoard(0, 0, 0)

	fmt.Println("Solved Board:")
	board.PrintBoard()
}
