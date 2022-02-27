package main

import (
	"fmt"
	"math/rand"
	"sudoku/board"
	"time"
)

func main() {
	grid := [9][9]int{}

	rand.Seed(time.Now().UnixNano())
	board := board.Board{SolvedBoard: grid, UserBoard: grid}
	board.GenerateBoard(0, 0)
	board.GenerateUserBoard(2)

	fmt.Println("User:")
	printBoard(board.UserBoard)

	fmt.Println("Solved:")
	board.PrintBoard(board.SolvedBoard)

	for board.FreePos(board.UserBoard) != nil {
		// holds user selection
		var inputRow int
		var inputCol int
		var inputVal int

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
			// board.PrintBoard(board.UserBoard)
			printBoard(board.UserBoard)
		} else {
			fmt.Println("invalid")
		}
	}
	fmt.Println("Done!")

}

func printBoard(grid [9][9]int) {
	for row := 0; row < len(grid); row++ {
		fmt.Printf("row:%d \t", row)
		for col := 0; col < len(grid[row]); col++ {
			if grid[row][col] == 0 {
				fmt.Printf("|*%d*|", grid[row][col])
			} else {
				fmt.Printf("%d|", grid[row][col])
			}
		}
		fmt.Println()
	}
}
