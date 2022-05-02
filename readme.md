## go-sudoku
![Go build](https://github.com/jacques-andre/go-sudoku/actions/workflows/go.yml/badge.svg)

### Features:
- **Backtracking algorithm:** Generates and solves Sudoku boards.
- **Undo/Redo move**: Implemented using a stack to push and pop the current board history.
- **Save/Load game**: Saves game history stack to a `.json` file. Allows for loading in stack and undoing moves back to `history[0]`

<img src=".github/static/demo.gif" height="700"/>

## Running the program:
Using the `makefile` included in the directory, run `make` to compile the program into a binary.

Run the compiled binary using `./bin/main`

Ensure that [Go](https://go.dev/doc/install) is installed. 

