package main

import (
	"fmt"
	"math/rand"
	"sudoku/board"
	"time"

	"github.com/fatih/color"
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
	// board.PrintBoard(board.UserBoard)
	printHighlightedSpacesBoard(board.UserBoard)

	fmt.Println("Solved Board:")
	board.PrintBoard(board.SolvedBoard)

	// while there are more 0's in the userboard,
	// keep prompting
	for board.FreePos(board.UserBoard) != nil {
		// holds user input
		var inputRow int
		var inputCol int
		var inputVal int

		// DEBUG: show zeros coords & answer on UserBoard
		// Ex: [0,4],[3,2]
		zeros := getZeros(board.UserBoard)
		count := 0
		for i, v := range zeros {
			// get the answer of where the blank cell is
			zeroRow := zeros[i][0]
			zeroCol := zeros[i][1]
			answer := board.SolvedBoard[zeroRow][zeroCol]

			count++

			fmt.Printf("DEBUG: Found blank pos at:%v,answer:%d \n", v, answer)
		}
		fmt.Printf("DEBUG: spaces left:%d \n", count)

		// take user input
		fmt.Printf("Choose a row:\n")
		fmt.Scanln(&inputRow)
		fmt.Printf("Choose a col:\n")
		fmt.Scanln(&inputCol)
		fmt.Printf("Choose a val:\n")
		fmt.Scanln(&inputVal)

		// check if the pos from user input is correct
		if board.ValidUserPos(inputVal, inputRow, inputCol, board.UserBoard) {
			color.Green("Valid! Placed this value on the board:")
			// Update UserBoard with newly validated value
			board.UserBoard[inputRow][inputCol] = inputVal
			// board.PrintBoard(board.UserBoard)
			printCorrectSpaceBoard(board.UserBoard, inputRow, inputCol)
		} else {
			color.Red("Invalid!")
		}
		fmt.Println("-----") // formatting
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

// Prints the board although 0's will be highlighted,
// yellow. Makes it easier to find 0's on board
func printHighlightedSpacesBoard(boardArray [9][9]int) {
	for row := 0; row < len(boardArray); row++ {
		for col := 0; col < len(boardArray[row]); col++ {
			if boardArray[row][col] == 0 {
				yellow := color.New(color.FgYellow).SprintFunc()
				fmt.Printf("%s|", yellow(boardArray[row][col]))
			} else {
				fmt.Printf("%d|", boardArray[row][col])
			}
		}
		fmt.Println()
	}
}

// Prints the board although rowN,colN will be highlighted green,
// shows user where there valid selection was
func printCorrectSpaceBoard(boardArray [9][9]int, rowN int, colN int) {
	for row := 0; row < len(boardArray); row++ {
		for col := 0; col < len(boardArray[row]); col++ {
			if row == rowN && col == colN {
				green := color.New(color.FgGreen).SprintFunc()
				fmt.Printf("%s|", green(boardArray[row][col]))
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
