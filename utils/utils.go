package utils

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

func LoadGame() map[int][9][9]int {

	// read in file, store in map
	result := map[int][9][9]int{}
	file, _ := ioutil.ReadFile("moves.json")
	json.Unmarshal(file, &result)

	// go through all key & values
	for key, value := range result {
		// empty grid
		grid := [9][9]int{}

		fmt.Printf("key:%v\n", key)

		// loop through row & col
		for row := 0; row < len(value); row++ {
			for col := 0; col < len(value[row]); col++ {
				fmt.Printf("row:%d,col:%d,value:%d,\n", row, col, value[row][col])
				grid[row][col] = value[row][col]
			}
		}
	}
	return result
}

// Write map to file
func WriteGame(gameMap map[int][9][9]int) {
	file, err := json.MarshalIndent(gameMap, "", " ")

	if err != nil {
		fmt.Printf("Error: %s", err.Error())
	} else {
		_ = ioutil.WriteFile("moves.json", file, 0644)
	}

}
