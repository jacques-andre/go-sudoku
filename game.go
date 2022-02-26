package main

import (
	"fmt"
	"math/rand"
	"sudoku/board"
	"time"
)

func main() {
	grid := [9][9]int{}

	rand.Seed(time.Now().UnixNano())
	board := board.Board{SolvedBoard: grid, UserBoard: grid}
	board.GenerateBoard(0, 0)
	board.GenerateUserBoard(70)

  // holds user selection
  var inputRow int
  var inputCol int


  for row := 0; row < len(board.UserBoard); row++{
    for col := 0; col < len(board.UserBoard[row]); col++{
      fmt.Printf("%d|", board.UserBoard[row][col])
    }
    fmt.Println()
  }



  fmt.Printf("Choose a row:\n")
  fmt.Scanln(&inputRow)

  fmt.Printf("Choose a col:\n")
  fmt.Scanln(&inputCol)

  fmt.Printf("You choose row:%v, col:%v, val:%d \n", inputRow, inputCol, board.UserBoard[inputRow][inputCol])
  board.PrintBoard()

}
