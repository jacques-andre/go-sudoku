package game

import (
	"errors"
	"fmt"
	"sudoku/board"
)

type GameHistory struct {
	CurrentHistory Stack
	RedoHistory    Stack
}

// Updates current history to include the passed in board,
// this also clears the redo history
func (g *GameHistory) AddMove(board board.Board) {
	g.CurrentHistory.Push(board)
	g.RedoHistory = nil
}

// Pops the last item off the stack, adds to redo history,
// returns the last version of the board
func (g *GameHistory) UndoMove() (board.Board, error) {

	// Cant undo when only original state of the board is left
	if len(g.CurrentHistory) == 1 {
		return board.Board{}, errors.New("No undo stack")
	}

	// get the last item on the stack,
	// this is the most up to date board
	lastItem, _ := g.CurrentHistory.Pop()

	// DEBUG
	fmt.Println("DEBUG: Popped off:")
	lastItem.PrintBoard(lastItem.UserBoard)

	// Add this undone move to the redo stack
	g.RedoHistory.Push(lastItem)

	// Return the last state of the board
	return g.CurrentHistory[len(g.CurrentHistory)-1], nil
}

// Pops the last item off the stack, adds to undo history,
// returns the last version of the board
func (g *GameHistory) RedoMove() (board.Board, error) {
	// If no redo history, throw err
	if len(g.RedoHistory) == 0 {
		return board.Board{}, errors.New("No redo stack")
	}

	// Get the last redo item
	lastItem, _ := g.RedoHistory.Pop()

	fmt.Println("DEBUG: Popped off:")
	lastItem.PrintBoard(lastItem.UserBoard)

	// We are adding back the redo item to CurrentHistory
	g.CurrentHistory.Push(lastItem)

	return lastItem, nil
}

func (g *GameHistory) SaveGame(board board.Board) error {
	for i := 0; i < len(g.CurrentHistory); i++ {
		moveName := fmt.Sprintf("move-%d", i)
		fmt.Println(moveName)
		g.CurrentHistory[i].PrintBoard(g.CurrentHistory[i].UserBoard)
	}
	return nil
}
