package main

import (
  "fmt"
)

type Board struct {
  BoardArray [9][9]int
}

// Generates a valid board with some blank spaces,
// user playable board
func newPlayingBoard() *Board{
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
  return &Board{
    BoardArray: validBoard,
  }
}

func printBoard(board *Board){
  for _, v := range board.BoardArray{
    fmt.Println(v)
  }
}


func main(){
  board := newPlayingBoard()
  printBoard(board)
}
