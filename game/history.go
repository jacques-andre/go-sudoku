package game

import "sudoku/board"

type GameHistory struct {
	CurrentHistory Stack
	RedoHistory    Stack
}

func (g *GameHistory) AddMove(board board.Board) {
	g.CurrentHistory.Push(board)
	g.RedoHistory = nil
}

func (g *GameHistory) UndoMove() board.Board {
	lastItem, _ := g.CurrentHistory.Pop()
	g.RedoHistory.Push(lastItem)

	return lastItem
}
