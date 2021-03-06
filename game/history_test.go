package game_test

import (
	"fmt"
	"os"
	"sudoku/board"
	"sudoku/game"
	"testing"
)

// Test able to add moves to gameHistory
func TestAddMove(t *testing.T) {
	// Create a empty board
	grid := [9][9]int{}
	board := board.Board{UserBoard: grid, SolvedBoard: grid}

	// Add empty board to history
	gameHistory := game.GameHistory{}
	gameHistory.AddMove(board)

	// Expecting one item to be added to the gameHistory stack
	got := len(gameHistory.CurrentHistory)
	want := 1

	if got != want {
		t.Errorf("Failed! Got:%d, want:%d", got, want)
	}
}

// Test able to undo board, get back correct item
func TestUndoMove(t *testing.T) {
	// Create a test board
	grid1 := [9][9]int{
		{9, 0, 0, 2, 0, 0, 0, 5, 0},
		{0, 7, 6, 0, 0, 8, 0, 4, 0},
		{0, 0, 0, 4, 0, 0, 0, 0, 3},
		{0, 6, 0, 1, 0, 0, 0, 0, 4},
		{0, 0, 4, 0, 9, 0, 5, 0, 0},
		{2, 0, 0, 0, 0, 6, 0, 7, 0},
		{3, 0, 0, 0, 0, 4, 0, 0, 0},
		{0, 2, 0, 8, 0, 0, 4, 3, 0},
		{0, 8, 0, 0, 0, 5, 0, 0, 2},
	}
	board1 := board.Board{UserBoard: grid1, SolvedBoard: grid1}

	// Create empty board
	grid2 := [9][9]int{}
	board2 := board.Board{UserBoard: grid2, SolvedBoard: grid2}

	// Add boards to history
	gameHistory := game.GameHistory{}
	gameHistory.AddMove(board1)
	gameHistory.AddMove(board2)

	// Undo move, expecting first board only in stack,
	// expecting 2nd board to be popped off
	undoBoard, _ := gameHistory.UndoMove()

	// Expecting to return to the test user board back,
	// as the blank board should be popped from the stack
	if undoBoard.UserBoard[0][0] != 9 {
		t.Errorf("Failed! Got wrong board in undo")
	}

	// Expecting testboard to be in the redo stack
	if gameHistory.RedoHistory[0].UserBoard[0][0] != 0 {
		t.Errorf("Failed! Got wrong board in redo history")
	}
}

// Test able to redo board, get back correct item
func TestRedoMove(t *testing.T) {
	// Create a test board
	grid1 := [9][9]int{
		{9, 0, 0, 2, 0, 0, 0, 5, 0},
		{0, 7, 6, 0, 0, 8, 0, 4, 0},
		{0, 0, 0, 4, 0, 0, 0, 0, 3},
		{0, 6, 0, 1, 0, 0, 0, 0, 4},
		{0, 0, 4, 0, 9, 0, 5, 0, 0},
		{2, 0, 0, 0, 0, 6, 0, 7, 0},
		{3, 0, 0, 0, 0, 4, 0, 0, 0},
		{0, 2, 0, 8, 0, 0, 4, 3, 0},
		{0, 8, 0, 0, 0, 5, 0, 0, 2},
	}
	board1 := board.Board{UserBoard: grid1, SolvedBoard: grid1}

	// Create empty board
	grid2 := [9][9]int{}
	board2 := board.Board{UserBoard: grid2, SolvedBoard: grid2}

	// Add boards to history
	gameHistory := game.GameHistory{}
	gameHistory.AddMove(board1)
	gameHistory.AddMove(board2)

	// Undo move, expecting first board only in stack,
	// expecting 2nd board to be popped off
	gameHistory.UndoMove()
	// Redo, expecting inital boards to be back in history
	gameHistory.RedoMove()

	// Redoing the move, expecting the first board to be back in history
	if len(gameHistory.CurrentHistory) != 2 {
		t.Errorf("Failed! CurrentHistory is does not include redone board")
	}

	// Check if currentboard has been redone
	if gameHistory.CurrentHistory[len(gameHistory.CurrentHistory)-1].UserBoard[0][0] != 0 {
		t.Errorf("Failed! CurrentHistory is does not include redone board")
	}
}

// Test able to save game history to file
func TestSaveGame(t *testing.T) {
	// Create a test board
	grid1 := [9][9]int{
		{9, 0, 0, 2, 0, 0, 0, 5, 0},
		{0, 7, 6, 0, 0, 8, 0, 4, 0},
		{0, 0, 0, 4, 0, 0, 0, 0, 3},
		{0, 6, 0, 1, 0, 0, 0, 0, 4},
		{0, 0, 4, 0, 9, 0, 5, 0, 0},
		{2, 0, 0, 0, 0, 6, 0, 7, 0},
		{3, 0, 0, 0, 0, 4, 0, 0, 0},
		{0, 2, 0, 8, 0, 0, 4, 3, 0},
		{0, 8, 0, 0, 0, 5, 0, 0, 2},
	}
	board1 := board.Board{UserBoard: grid1, SolvedBoard: grid1}

	// Create empty board
	grid2 := [9][9]int{}
	board2 := board.Board{UserBoard: grid2, SolvedBoard: grid2}

	// Add boards to history
	gameHistory := game.GameHistory{}
	gameHistory.AddMove(board1)
	gameHistory.AddMove(board2)

	gameHistory.SaveGame("test.json")

	// Check that the saved file exists
	if _, err := os.Stat("test.json"); os.IsNotExist(err) {
		t.Errorf("Failed! File does not exist!")
	}

}

// Test able to load game history from file into,
// board type
func TestLoadGame(t *testing.T) {
	TestSaveGame(t)

	// Load in the game
	gameHistory := game.GameHistory{}
	err := gameHistory.LoadGame("test.json")

	if err != nil {
		fmt.Printf("Error loading file! %s", err.Error())
	}

	// Test we are able to get the most up to date board,
	// Expecting to get boards used from TestSaveGame
	lastUserBoard := gameHistory.CurrentHistory[len(gameHistory.CurrentHistory)-1].UserBoard
	firstBoard := gameHistory.CurrentHistory[0].UserBoard

	if firstBoard[0][0] != 9 {
		t.Errorf("Failed! Getting wrong board in loaded file!")
	}

	if lastUserBoard[0][0] != 0 {
		t.Errorf("Failed! Getting wrong board in loaded file!")
	}

}
