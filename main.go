package main

import (
	"fmt"
	"math/rand"
	"sudoku/board"

	// "sudoku/game"
	"time"
)

func main() {
	// best practice
	rand.Seed(time.Now().UnixNano())

	// First ask the user if they would like to play,
	// or replay a old game
	fmt.Println("Press 1 for a new game, 2 for replaying a game")

	testGrid := [9][9]int{
		{9, 0, 0, 2, 0, 0, 0, 5, 0},
		{0, 7, 6, 0, 0, 8, 0, 4, 0},
		{0, 0, 0, 4, 0, 0, 0, 0, 3},
		{0, 6, 0, 1, 0, 0, 0, 0, 4},
		{0, 0, 4, 0, 9, 0, 5, 0, 0},
		{2, 0, 0, 0, 0, 6, 0, 7, 0},
		{3, 0, 0, 0, 0, 4, 0, 0, 0},
		{0, 2, 0, 8, 0, 0, 4, 3, 0},
		{0, 8, 0, 0, 0, 5, 0, 0, 2},
	}

	board := board.Board{UserBoard: testGrid, SolvedBoard: testGrid}

	board.PrintBoard(board.UserBoard)
	fmt.Println("-------")

	board.SolveBoard(0, 0, &board.UserBoard)
	board.PrintBoard(board.UserBoard)

	// // hold 1 or 2?
	// var userIntAnswer int
	// fmt.Scanln(&userIntAnswer)

	// if userIntAnswer == 1 {
	// 	gameHistory := game.GameHistory{}
	// 	game.NewGame(gameHistory)
	// }
	// if userIntAnswer == 2 {
	// 	gameHistory := game.GameHistory{}
	// 	gameHistory.LoadGame("test.json")

	// 	game.NewGame(gameHistory)
	// }

}
