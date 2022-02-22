package board

import (
	"fmt"
	"math/rand"
	"time"
)

type Board struct {
	BoardArray [9][9]int
}

// Generates a valid board with some blank spaces,
// user playable board
func (board *Board) SolveBoard(cellVal int, row int, col int) bool {
	// get the next available pos on board
	freePos := board.FreePos()
	if freePos == nil {
		// no more positions, done
		return true
	}

	freeRow := freePos[0]
	freeCol := freePos[1]

	// generate random number
	randNum := rand.Intn(10-1) + 1
	rand.Seed(time.Now().UnixNano())

	for i := 0; i <= 8; i++ {
		// board.printBoard()
		// fmt.Printf("Checking: row:%d,col:%d,val:%d \n", freeRow, freeCol, randNum)
		if board.ValidPos(randNum, freeRow, freeCol) {
			// fmt.Printf("Valid! row:%d, col:%d, val:%d \n", freeRow, freeCol, randNum)
			board.BoardArray[freeRow][freeCol] = randNum
			if board.SolveBoard(randNum, freeRow, freeCol) {
				return true
			}
			board.BoardArray[freeRow][freeCol] = 0
		}
	}
	return false
}

// Checks all valid pos funcs, returns if validPos
func (board *Board) ValidPos(cellVal int, row int, col int) bool {
	isValidInRow := board.ValidPosInRow(cellVal, row)
	isValidInCol := board.ValidPosInCol(cellVal, col)
	isValidInSubGrid := board.ValidPosInSubGrid(cellVal, col, row)

	if isValidInRow && isValidInCol && isValidInSubGrid {
		return true
	}
	return false
}

func (board *Board) FreePos() []int {
	for row := 0; row < len(board.BoardArray); row++ {
		for col := 0; col < len(board.BoardArray[row]); col++ {
			if board.BoardArray[row][col] == 0 {
				validPos := []int{row, col}
				return validPos
			}
		}
	}
	return nil
}

// Check if cellVal can be placed in the row rowN
// Returns true if valid pos
func (board *Board) ValidPosInRow(cellVal int, rowN int) bool {
	for i := 0; i < len(board.BoardArray[rowN]); i++ {
		currentValue := board.BoardArray[rowN][i]
		if currentValue == cellVal {
			return false
		}
	}
	return true
}

// Check if cellVal can be placed in the column colN
// Returns true if valid pos
func (board *Board) ValidPosInCol(cellVal int, colN int) bool {
	for i := 0; i < len(board.BoardArray[colN]); i++ {
		currentValue := board.BoardArray[i][colN]
		if currentValue == cellVal {
			return false
		}
	}
	return true
}

// Check if cellVal can be placed in the 3x3 subgrid
// Returns true if valid pos
func (board *Board) ValidPosInSubGrid(cellVal int, colN int, rowN int) bool {
	rowStart := (rowN / 3) * 3
	colStart := (colN / 3) * 3

	// Go through the subgrid,
	// check if any values are equal to,
	// cellVal > invalid
	for row := rowStart; row < rowStart+3; row++ {
		for col := colStart; col < colStart+3; col++ {
			subGridCell := board.BoardArray[row][col]
			if subGridCell == cellVal {
				return false
			}
		}
	}
	return true

}

// Helper Method:
// Prints board nicely
func (board *Board) PrintBoard() {
	for _, v := range board.BoardArray {
		fmt.Println(v)
	}
}

