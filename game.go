package main

import (
	"sudoku/board"
)

func main() {
	grid := [9][9]int{}

  board := board.Board{BoardArray: grid}
  board.GenerateBoard(0, 0, 0)

  board.PrintBoard()

}
