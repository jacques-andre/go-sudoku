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

		// show zeros coords on UserBoard, helps with debugging
		// Ex: [0,4],[3,2]
		zeros := getZeros(board.UserBoard)
		for i, v := range zeros {
			fmt.Printf("zeros:i:%d,v:%v \n", i, v)
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
			coord := [][]int{{row, col}}
			if boardArray[row][col] == 0 {
				seen = append(seen, coord...)
			}
		}
	}
	return seen
}
