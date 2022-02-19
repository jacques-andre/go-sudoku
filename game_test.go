package main

import (
	"fmt"
	"testing"
)

func TestValidInBoardRow(t *testing.T){
  board := newPlayingBoard()

  // Test vars
  valueN := 9
  rowN := 0

  got := validPosInRow(board, valueN, rowN)
  want := true
  
  fmt.Printf("Row:%v, checking:%d", board.BoardArray[rowN], valueN)


  if got != want {
    t.Errorf("Failed! Got:%t, want:%t", got, want)
  }
}

func TestInValidInBoardRow(t *testing.T){
  board := newPlayingBoard()
  // Test vars
  valueN := 3
  rowN := 0

  got := validPosInRow(board, valueN, rowN)
  want := false
  
  fmt.Printf("Row:%v, checking:%d", board.BoardArray[rowN], valueN)


  if got != want {
    t.Errorf("Failed! Got:%t, want:%t", got, want)
  }
}

func TestValidInBoardCol(t *testing.T){
  board := newPlayingBoard()
  // Test vars
  valueN := 7
  colN := 0

  got := validPosInCol(board, valueN, colN)
  want := true
  
  printBoard(board)
  fmt.Printf("Checking:%d, col:%d", valueN, colN)


  if got != want {
    t.Errorf("Failed! Got:%t, want:%t", got, want)
  }
}
