package main

import (
	"fmt"
	"math/rand"
	"sudoku/game"
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

	if userIntAnswer == 1 {
		game.NewGame()
	}
	if userIntAnswer == 2 {
		t := game.GameHistory{}
		t.LoadGame("test.json")
	}

}
