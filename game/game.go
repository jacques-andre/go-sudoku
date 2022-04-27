package game

import (
	"fmt"
	"sudoku/board"
)

// Run the game program, keeps track of history and the current state of the board,
func NewGame(gameHistory GameHistory) {
	// Initliaze
	mainBoard := board.Board{SolvedBoard: [9][9]int{}, UserBoard: [9][9]int{}}

	// Check if there is already a game history,
	// use that game if so
	if len(gameHistory.CurrentHistory) >= 1 {
		lastBoard := gameHistory.CurrentHistory[len(gameHistory.CurrentHistory)-1].UserBoard
		mainBoard.UserBoard = lastBoard
		mainBoard.SolvedBoard = lastBoard

		// solve the board that has been passed
		mainBoard.SolveBoard(0, 0, &mainBoard.SolvedBoard)

		fmt.Printf("Successfully loaded in game, stack size: %d \n", len(gameHistory.CurrentHistory))

		// DEBUG
		// mainBoard.PrintBoard(mainBoard.UserBoard)
		// fmt.Println("Solved:")
		// mainBoard.PrintBoard(mainBoard.SolvedBoard)

	} else {
		// Not loaded in, start new game
		const USER_BLANK_SPACES = 6
		// Generates a SolvedBoard
		mainBoard.GenerateBoard(0, 0)
		// Generates a user playable board with n blank spaces
		mainBoard.GenerateUserBoard(USER_BLANK_SPACES)

		// Add the first state of the board to history
		gameHistory.AddMove(mainBoard)
	}

	// While we have a free postion on the board,
	// continue the game for the user
	freePos := mainBoard.FreePos(mainBoard.UserBoard)
	for freePos != nil {
		fmt.Println("------") // Formatting

		// DEBUG
		zeros := getZeros(mainBoard.UserBoard)
		count := 0
		for i, v := range zeros {
			// get the answer of where the blank cell is
			zeroRow := zeros[i][0]
			zeroCol := zeros[i][1]
			answer := mainBoard.SolvedBoard[zeroRow][zeroCol]
			_ = v
			count++

			fmt.Printf("DEBUG: Found blank pos at row:%d,col:%d,ans:%d,\n", zeroRow, zeroCol, answer)
		}
		fmt.Printf("DEBUG: Total blanks: %d\n", count)

		// DEBUG
		fmt.Printf("CurrentHistory Stack size: %v\n", len(gameHistory.CurrentHistory))
		fmt.Printf("Redo Stack size: %v\n", len(gameHistory.RedoHistory))

		// Print the current board
		fmt.Println("Current Board:")
		mainBoard.PrintHighlightedSpacesBoard(mainBoard.UserBoard)

		// ask user what action?
		fmt.Println("Undo (u), Redo(r), Place(p), Save(s)?")
		var userChoice string
		fmt.Scanln(&userChoice)

		// User is placing on the board
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

			// TODO VALIDATE USER INPUT
			// could be a seperate func

			// check input is valid
			for inputValue == 0 {
				fmt.Println("invalid value!")
				fmt.Println("Enter a value:")
				fmt.Scanln(&inputValue)
			}

			fmt.Printf("You chose: row:%d,col:%d,value:%d,\n", inputRow, inputCol, inputValue)

			// If the user input was correct,
			// show it on the board with color formatting,
			// update board with this value & history
			if mainBoard.SolvedBoard[inputRow][inputCol] == inputValue {
				mainBoard.UserBoard[inputRow][inputCol] = inputValue // place the validated user input value on the board
				gameHistory.AddMove(mainBoard)
				fmt.Println("Correct!")
				mainBoard.PrintCorrectSpaceBoard(mainBoard.UserBoard, inputRow, inputCol)
			} else {
				fmt.Println("Incorrect!")
				mainBoard.PrintInCorrectSpaceBoard(mainBoard.UserBoard, inputRow, inputCol, inputValue)
			}

		}

		// User is undoing a place on the board,
		// first get the last done move from gameHistory,
		// reapply this move to the mainBoard
		if userChoice == "u" {
			lastMove, err := gameHistory.UndoMove()
			if err == nil {
				mainBoard.UserBoard = lastMove.UserBoard
			} else {
				// Display the error msg to the user
				fmt.Printf("Error! %s \n", err.Error())
			}
		}
		// User is redoing a place on the board,
		if userChoice == "r" {
			lastMove, err := gameHistory.RedoMove()

			if err == nil {
				mainBoard.UserBoard = lastMove.UserBoard
			} else {
				// Display the error msg to the user
				fmt.Printf("Error! %s \n", err.Error())
			}
		}
		// User is saving the current board state to file
		if userChoice == "s" {
			fmt.Println("File name:")

			// Take in file name
			var userFileInput string
			fmt.Scanln(&userFileInput)

			gameHistory.SaveGame(userFileInput)
		}
		// Update freePos after potential placing
		freePos = mainBoard.FreePos(mainBoard.UserBoard)
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
