package board

import (
	"fmt"
	"math/rand"
)

type Board struct {
	SolvedBoard [9][9]int
	UserBoard   [9][9]int
}

// Generates a playable board,
// With blanks amount of spaces
func (board *Board) GenerateUserBoard(blanks int) {
	if board.SolvedBoard[0][0] == 0 {
		panic("No solved board! Unable to generate UserBoard")
	}

	// set the UserBoard to the solved one,
	// blank out some spaces on this UserBoard
	// allows user to play the game
	board.UserBoard = board.SolvedBoard

	// go through userboard,
	// blank out random choice
	for i := 1; i <= blanks; {
		randRow := rand.Intn(9)
		randCol := rand.Intn(9)
		if board.UserBoard[randRow][randCol] != 0 {
			board.UserBoard[randRow][randCol] = 0
			i++
		}
	}
}

// Creates a valid random SolvedBoard
func (board *Board) GenerateBoard(row int, col int) bool {
	// get the next available pos on board
	freePos := board.FreePos(board.SolvedBoard)
	if freePos == nil {
		// no more positions, done
		return true
	}

	freeRow := freePos[0]
	freeCol := freePos[1]

	// generate random number
	randNum := rand.Intn(10-1) + 1

	if board.ValidPos(randNum, freeRow, freeCol, board.SolvedBoard) {
		// fmt.Printf("Valid: row:%d,col:%d,val:%d,mark:%d,\n", freeRow, freeCol, randNum)
		board.SolvedBoard[freeRow][freeCol] = randNum
		// board.PrintBoard()
		// fmt.Scanln()
		// fmt.Println("----")

		// re-generate a random number 8 times before back track
		for i := 0; i <= 8; i++ {
			// fmt.Printf("Checking: row:%d,col:%d,val:%d,i:%d\n", freeRow, freeCol, randNum, i)
			if board.GenerateBoard(freeRow, freeCol+1) {
				return true
			}
		}
		board.SolvedBoard[freeRow][freeCol] = 0
	}
	return false
}

// Solves UserBoard (existing board)
func (board *Board) SolveBoard(cellVal int, row int, col int) bool {
	// get the next available pos on board
	freePos := board.FreePos(board.UserBoard)
	if freePos == nil {
		// no more positions, done
		return true
	}

	freeRow := freePos[0]
	freeCol := freePos[1]

	// check every possibility at current square
	for i := 1; i <= 9; i++ {
		// board.PrintBoard()
		// fmt.Printf("Checking: row:%d,col:%d,val:%d \n", freeRow, freeCol, i)
		if board.ValidPos(i, freeRow, freeCol, board.UserBoard) {
			// fmt.Printf("Valid! row:%d, col:%d, val:%d \n", freeRow, freeCol, i)
			board.UserBoard[freeRow][freeCol] = i
			if board.SolveBoard(i, freeRow, freeCol+1) {
				return true
			}
			board.UserBoard[freeRow][freeCol] = 0
		}
	}
	return false
}

// Checks all valid pos funcs
// Returns true if valid
func (board *Board) ValidPos(cellVal int, row int, col int, boardArray [9][9]int) bool {
	isValidInRow := board.ValidPosInRow(cellVal, row, boardArray)
	isValidInCol := board.ValidPosInCol(cellVal, col, boardArray)
	isValidInSubGrid := board.ValidPosInSubGrid(cellVal, col, row, boardArray)

	// fmt.Printf("DEBUG: row:%v, col:%v, subgrid:%v \n", isValidInRow, isValidInCol, isValidInSubGrid)

	if isValidInRow && isValidInCol && isValidInSubGrid {
		return true
	}
	return false
}

func (board *Board) ValidUserPos(cellVal int, row int, col int, boardArray [9][9]int) bool {
	// answer to the cell the user chose
	validInputCell := board.SolvedBoard[row][col]

	if board.ValidPos(cellVal, row, col, board.UserBoard) {
		if cellVal == validInputCell {
			return true
		}
	}
	return false
}

// Checks next free pos on board
// Returns [row,col] of free pos
func (board *Board) FreePos(boardArray [9][9]int) []int {
	for row := 0; row < len(boardArray); row++ {
		for col := 0; col < len(boardArray[row]); col++ {
			if boardArray[row][col] == 0 {
				validPos := []int{row, col}
				return validPos
			}
		}
	}
	return nil
}

// Check if cellVal can be placed in the row rowN
// Returns true if valid pos
func (board *Board) ValidPosInRow(cellVal int, row int, boardArray [9][9]int) bool {
	for i := 0; i < len(boardArray[row]); i++ {
		currentValue := boardArray[row][i]
		if currentValue == cellVal {
			return false
		}
	}
	return true
}

// Check if cellVal can be placed in the column colN
// Returns true if valid pos
func (board *Board) ValidPosInCol(cellVal int, col int, boardArray [9][9]int) bool {
	for i := 0; i < len(boardArray[col]); i++ {
		currentValue := boardArray[i][col]
		if currentValue == cellVal {
			return false
		}
	}
	return true
}

// Check if cellVal can be placed in the 3x3 subgrid
// Returns true if valid pos
func (board *Board) ValidPosInSubGrid(cellVal int, colN int, rowN int, boardArray [9][9]int) bool {
	rowStart := (rowN / 3) * 3
	colStart := (colN / 3) * 3

	// Go through the subgrid,
	// check if any values are equal to cellVal
	for row := rowStart; row < rowStart+3; row++ {
		for col := colStart; col < colStart+3; col++ {
			subGridCell := boardArray[row][col]
			if subGridCell == cellVal {
				return false
			}
		}
	}
	return true

}

// Helper Method:
// Prints board nicely
func (board *Board) PrintBoard(boardArray [9][9]int) {
	for row := 0; row < len(boardArray); row++ {
		fmt.Printf("row:%d \t", row)
		for col := 0; col < len(boardArray[row]); col++ {
			fmt.Printf("%d|", boardArray[row][col])
		}
		fmt.Println()
	}
}
