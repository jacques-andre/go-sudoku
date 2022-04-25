#!/bin/bash

docker build -t go-sudoku . && clear && docker run --name sudoku-container -it go-sudoku
