package main

import (
	"fmt"
	"math/rand"
	"sudoku/board"
	"sudoku/utils"
	"time"

	"github.com/fatih/color"
)

func main() {
	// best practice
	rand.Seed(time.Now().UnixNano())

	// First ask the user if they would like to play,
	// or replay a old game
	fmt.Println("Press 1 for a new game, 2 for replaying a game")

	// hold 1 or 2?
	var userIntAnswer int
	fmt.Scanln(&userIntAnswer)

	if userIntAnswer == 1 {
		newGame()
	}
	if userIntAnswer == 2 {
		fmt.Println("yes2")
	}

}
func newGame() {
	// blank 2d slice
	grid := [9][9]int{}

	// init new board struct, both vars use blank board
	mainBoard := board.Board{SolvedBoard: grid, UserBoard: grid}

	// generate initial board for SolvedBoard
	mainBoard.GenerateBoard(0, 0)

	// UserBoard will copy SolvedBoard with n amount of zeros
	mainBoard.GenerateUserBoard(30)

	moveMap := map[int][9][9]int{}
	moveMap[0] = mainBoard.UserBoard
	currentMove := 1

	freePos := mainBoard.FreePos(mainBoard.UserBoard)
	for freePos != nil {
		// DEBUG
		zeros := getZeros(mainBoard.UserBoard)
		for i, v := range zeros {
			// get the answer of where the blank cell is
			zeroRow := zeros[i][0]
			zeroCol := zeros[i][1]
			answer := mainBoard.SolvedBoard[zeroRow][zeroCol]
			_ = v

			fmt.Printf("DEBUG: Found blank pos at row:%d,col:%d,ans:%d,\n", zeroRow, zeroCol, answer)
		}

		// Print the current board
		fmt.Println("User Board:")
		mainBoard.PrintHighlightedSpacesBoard(mainBoard.UserBoard)

		// hold user input
		var inputRow int
		var inputCol int
		var inputValue int

		// ask for user input
		fmt.Println("Enter a row:")
		fmt.Scanln(&inputRow)
		fmt.Println("Enter a col:")
		fmt.Scanln(&inputCol)
		fmt.Println("Enter a value:")
		fmt.Scanln(&inputValue)

		// check input is valid
		for inputValue == 0 {
			fmt.Println("invalid value!")
			fmt.Println("Enter a value:")
			fmt.Scanln(&inputValue)
		}

		fmt.Printf("You choose: row:%d,col:%d,value:%d,POS:%d \n", inputRow, inputCol, inputValue, mainBoard.UserBoard[inputRow][inputCol])

		// user input was correct
		if mainBoard.ValidPos(inputValue, inputRow, inputCol, mainBoard.UserBoard) {
			// place on board
			mainBoard.UserBoard[inputRow][inputCol] = inputValue

			fmt.Println("Valid!")
			mainBoard.PrintCorrectSpaceBoard(mainBoard.UserBoard, inputRow, inputCol)
		}

		// update move
		moveMap[currentMove] = mainBoard.UserBoard
		currentMove++
		utils.WriteGame(moveMap)

		// update freePos
		freePos = mainBoard.FreePos(mainBoard.UserBoard)

	}

}

// Prints the board although rowN,colN will be highlighted red,
// shows user where there invalid selection was
func printInCorrectSpaceBoard(boardArray [9][9]int, rowN int, colN int, valueN int) {
	for row := 0; row < len(boardArray); row++ {
		for col := 0; col < len(boardArray[row]); col++ {
			if row == rowN && col == colN {
				red := color.New(color.FgRed).SprintFunc()
				fmt.Printf("%s|", red(valueN))
			} else if boardArray[row][col] == 0 {
				yellow := color.New(color.FgYellow).SprintFunc()
				fmt.Printf("%s|", yellow(boardArray[row][col]))
			} else {
				fmt.Printf("%d|", boardArray[row][col])
			}
		}
		fmt.Println()
	}
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
