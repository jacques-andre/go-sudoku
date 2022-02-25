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
  board.GenerateUserBoard(20)

  count := 0

  fmt.Printf("User board:  \n")
  for row := 0; row < len(board.UserBoard); row++{
    for col := 0; col < len(board.UserBoard[row]); col++{
      fmt.Printf("%d|", board.UserBoard[row][col])
      if board.UserBoard[row][col] == 0{
        count++
      }
    }
    fmt.Println()
  }
  fmt.Printf("zeros:%d \n", count)

  fmt.Printf("Solved board: \n")
  board.PrintBoard()

}
