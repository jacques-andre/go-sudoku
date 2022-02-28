package main

import (
	"fmt"
	"math/rand"
	"sudoku/board"
	"time"
)

func main() {
	// best practice
	rand.Seed(time.Now().UnixNano())

	// blank board
	grid := [9][9]int{}
	// init new board struct, both vars use blank board
	board := board.Board{SolvedBoard: grid, UserBoard: grid}
	// generate initial board for SolvedBoard
	board.GenerateBoard(0, 0)
	// UserBoard will copy SolvedBoard with n amount of zeros
	board.GenerateUserBoard(2)

	fmt.Println("User Board:")
	board.PrintBoard(board.UserBoard)

	fmt.Println("Solved Board:")
	board.PrintBoard(board.SolvedBoard)

	// while there are no more 0's in the userboard,
	// keep prompting
	for board.FreePos(board.UserBoard) != nil {
		// holds user input
		var inputRow int
		var inputCol int
		var inputVal int

		// DEBUG: show zeros coords & answer on UserBoard
		// Ex: [0,4],[3,2]
		zeros := getZeros(board.UserBoard)
		for i, v := range zeros {
			// get the answer of where the blank cell is
			zeroRow := zeros[i][0]
			zeroCol := zeros[i][1]
			answer := board.SolvedBoard[zeroRow][zeroCol]

			fmt.Printf("DEBUG: Found blank pos at:%v,answer:%d \n", v, answer)
		}

		// take user input
		fmt.Printf("Choose a row:\n")
		fmt.Scanln(&inputRow)
		fmt.Printf("Choose a col:\n")
		fmt.Scanln(&inputCol)
		fmt.Printf("Choose a val:\n")
		fmt.Scanln(&inputVal)

		// check if the pos from user input is correct
		if board.ValidUserPos(inputVal, inputRow, inputCol, board.UserBoard) {
			fmt.Println("valid")
			board.UserBoard[inputRow][inputCol] = inputVal
			board.PrintBoard(board.UserBoard)

		} else {
			fmt.Println("invalid")
		}
	}
	fmt.Println("Done!")

}

// helper method: allows for finding zeros quickly on UserBoard
func getZeros(boardArray [9][9]int) [][]int {
	seen := [][]int{}

	for row := 0; row < len(boardArray); row++ {
		for col := 0; col < len(boardArray[row]); col++ {
			currentCoord := [][]int{{row, col}}

			if boardArray[row][col] == 0 {
				seen = append(seen, currentCoord...)
			}
		}
	}
	return seen
}
