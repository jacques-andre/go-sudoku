package game

import (
	"fmt"
	"sudoku/board"

	"github.com/fatih/color"
)

func NewGame() {
	// blank 2d slice
	gameHistory := GameHistory{}
	mainBoard := board.Board{SolvedBoard: [9][9]int{}, UserBoard: [9][9]int{}}

	// generate initial board for SolvedBoard
	mainBoard.GenerateBoard(0, 0)

	// UserBoard will copy SolvedBoard with n amount of zeros
	mainBoard.GenerateUserBoard(2)

	gameHistory.AddMove(mainBoard)

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

		// ask user what action?
		fmt.Println("Undo (u), Redo(r), Place(p)?")
		var userChoice string
		fmt.Scanln(&userChoice)

		if userChoice == "p" {
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
				gameHistory.AddMove(mainBoard)
				mainBoard.UserBoard[inputRow][inputCol] = inputValue

				fmt.Println("Valid!")
				mainBoard.PrintCorrectSpaceBoard(mainBoard.UserBoard, inputRow, inputCol)
			}

			// update freePos
			freePos = mainBoard.FreePos(mainBoard.UserBoard)
		} else if userChoice == "u" {
			lastMove := gameHistory.CurrentHistory[len(gameHistory.CurrentHistory)-1].UserBoard
			mainBoard.UserBoard = lastMove
		}

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
