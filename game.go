package main

import (
	"fmt"
	"math/rand"
	"sudoku/board"
	"time"

	"strconv"

	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

func main() {
	rand.Seed(time.Now().UnixNano())

	grid := [9][9]int{}
	board := board.Board{SolvedBoard: grid, UserBoard: grid}
	board.GenerateBoard(0, 0)
	board.GenerateUserBoard(30)

	app := tview.NewApplication()
	table := tview.NewTable().SetBorders(true)
	test := tview.NewTextView()

	tviewGrid := tview.NewGrid().SetBorders(true)

	tviewGrid.AddItem(table, 0, 0, 2, 1, 1, 1, true)
	tviewGrid.AddItem(test, 0, 1, 2, 2, 2, 1, false)

	// render the board as a table
	for row := 0; row < len(board.UserBoard); row++ {
		for col := 0; col < len(board.UserBoard[row]); col++ {
			currentBoardValue := board.UserBoard[row][col]

			// Create a new cell with board value
			newCell := tview.NewTableCell(strconv.Itoa(currentBoardValue))
			// if cell is 0 show in different color
			if currentBoardValue == 0 {
				newCell = tview.NewTableCell(strconv.Itoa(currentBoardValue)).SetTextColor(tcell.ColorYellow)
			}

			// add new cell to table
			table.SetCell(row, col, newCell)
		}
	}

	table.SetSelectable(true, true)

	// when cell is selected, change color to red
	table.SetSelectedFunc(func(row int, col int) {
		selectedCell := table.GetCell(row, col)
		selectedCell.SetTextColor(tcell.ColorRed)

		test.SetText(fmt.Sprintf("You choose: row:%d,col:%d,value:%d", row, col, board.UserBoard[row][col]))
	})

	// run the app
	err := app.SetRoot(tviewGrid, true).EnableMouse(true).Run()
	if err != nil {
		panic(err)
	}

}
