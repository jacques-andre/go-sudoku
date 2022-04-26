package game

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
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

// Save the passed in board to fileName (json)
func (g *GameHistory) SaveGame(board board.Board, fileName string) error {
	// Map move name to board postion
	movesMap := make(map[string][9][9]int)

	// Go through CurrentHistory, append to map
	for i := 0; i < len(g.CurrentHistory); i++ {
		moveName := fmt.Sprintf("move-%d", i)
		currentBoard := g.CurrentHistory[i].UserBoard

		movesMap[moveName] = currentBoard
	}
	file, err := json.MarshalIndent(movesMap, "", " ")

	if err != nil {
		fmt.Printf("Error!: %s \n", err.Error())
		return err
	}
	err = ioutil.WriteFile(fileName, file, 0644)
	if err != nil {
		fmt.Printf("Error!: %s \n", err.Error())
		return err
	} else {
		fmt.Printf("Successfully wrote %d moves to %s \n", len(g.CurrentHistory), fileName)
	}

	return nil
}

func (g *GameHistory) LoadGame(fileName string) error {
	// Map move name to board postion
	movesMap := make(map[string][9][9]int)

	jsonFile, err := ioutil.ReadFile(fileName)

	if err != nil {
		return err
	}
	json.Unmarshal(jsonFile, &movesMap)

	// Convert map into CurrentHistory
	for _, value := range movesMap {
		currentBoard := board.Board{UserBoard: value, SolvedBoard: value}
		g.CurrentHistory.Push(currentBoard)
	}
	for i := range g.CurrentHistory {
		fmt.Printf("Move-%d\n", i)
		g.CurrentHistory[i].PrintBoard(g.CurrentHistory[i].UserBoard)
		fmt.Println("----")
	}
	return nil

}
