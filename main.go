package main

import (
	"fmt"
	"math/rand"
	"sudoku/game"

	// "sudoku/game"
	"time"
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

	// Start a new game
	if userIntAnswer == 1 {
		gameHistory := game.GameHistory{}
		game.NewGame(gameHistory)
	}
	// Load in game from file
	if userIntAnswer == 2 {
		fmt.Println("What is the name of the file you would like to load in?")

		// Hold user file input
		var userFileInput string
		fmt.Scanln(&userFileInput)

		// Create a gameHistory with file input
		gameHistory := game.GameHistory{}
		err := gameHistory.LoadGame(userFileInput)

		if err != nil {
			panic(err.Error())
		}

		game.NewGame(gameHistory)
	}

}
