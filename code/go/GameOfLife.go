package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"strconv"
)

func rules(grid,row,col) int {
	var countNeighbors = 0
	neighbors := [...][2]int{{row + 1, col}, {row - 1, col}, {row, col + 1}, {row, col - 1}, {row - 1, col - 1}, {row + 1, col + 1}, {row - 1, col + 1}, {row + 1, col - 1}}
	for i := 0; i < len(neighbors); i++ {
		var currentRow int = neighbors[i][0]
		var currentColumn int = neighbors[i][1]
		if grid[currentRow][currentColumn] == 1 {
			countNeighbors++
		}
	}
	if countNeighbors <= 1 && grid[row][col] == 1 {
		return 0
	} else if countNeighbors >= 4 && countNeighbors <= 8 && grid[row][col] == 1 {
		return 0
	} else if countNeighbors >= 2 && countNeighbors <= 3 && grid[row][col] == 1 {
		return 1
	} else if countNeighbors == 3 && grid[row][col] == 0 {
		return 1
	} else {
		return 0
	}
}

func display(grid) {
	for i := 0; i < 8; i++ {
		for j := 0; j < 8; j++ {
			fmt.Print(strconv.Itoa(grid[i][j]) + " ")
		}
		fmt.Println("")
	}
	fmt.Println("\n")
}

func post(grid [8][8]int) {
	gridString := ""
	for i := 0; i < 8; i++ {
		for j := 0; j < 8; j++ {
			gridString += strconv.Itoa(grid[i][j]) + ","
		}
		gridString += ";"
	}
	data := url.Values{
		"grid":     {gridString},
		"language": {"Go"},
	}
	resp, err := http.PostForm("https://odd-mole-91.serverless.social/CheckGrid", data)
	if err != nil {
		log.Fatalln(err)
	}

	defer resp.Body.Close()
	bodyBytes, _ := ioutil.ReadAll(resp.Body)
	bodyString := string(bodyBytes)
	fmt.Println(bodyString)
}

func main() {
	generations := 15
	var grid [8][8]int
	grid[1][0] = 1
	grid[2][1] = 1
	grid[2][2] = 1
	grid[1][2] = 1
	grid[0][2] = 1
	display(grid)
	for k := 0; k < generations; k++ {
		var newGrid [8][8]int
		for i := 0; i < 8; i++ {
			for j := 0; j < 8; j++ {
				newGrid[i][j] = rules(grid, i, j)
			}
		}
		display(newGrid)
		grid = newGrid
	}
	post(grid)
}
