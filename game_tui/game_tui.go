package game_tui

import (
	"strconv"
	"sudoku/board"

	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

type GameTUI struct {
	App   *tview.Application
	Grid  *tview.Grid
	Table *tview.Table
}

func NewGameTUI() *GameTUI {
	t := GameTUI{}

	// create an empty board
	grid := [9][9]int{}
	board := board.Board{grid, grid}

	// genearte board
	board.GenerateBoard(0, 0)
	board.GenerateUserBoard(35)

	// create new tview items
	t.App = tview.NewApplication()
	t.Table = tview.NewTable().SetBorders(true)
	t.Table.SetSelectable(true, true)
	t.Grid = tview.NewGrid().SetBorders(true)

	// add items to grid layout
	t.Grid.AddItem(t.Table, 0, 0, 2, 1, 1, 1, true)

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
			t.Table.SetCell(row, col, newCell)
		}
	}

	return &t
}
