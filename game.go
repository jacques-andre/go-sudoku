package main

import (
	"fmt"
	"sudoku/board"
)

func main() {
	grid := [9][9]int{}

	board := board.Board{BoardArray: grid}
	board.SolveBoard(0, 0, 0)

	fmt.Println("Solved Board:")
	board.PrintBoard()
}
