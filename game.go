package main

import (
	"math/rand"
	"sudoku/game_tui"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())

	t := game_tui.NewGameTUI()

	// run the app
	err := t.App.SetRoot(t.Grid, true).EnableMouse(false).Run()
	if err != nil {
		panic(err)
	}
}
