package main

import (
	"sudoku/board"
	"time"
  "math/rand"
)

func main() {
	grid := [9][9]int{}

	rand.Seed(time.Now().UnixNano())
	board := board.Board{BoardArray: grid}
	board.GenerateBoard(0, 0)

	board.PrintBoard()

}
