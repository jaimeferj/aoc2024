package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func displayGrid(grid [][]bool) {
	height := len(grid)
	width := len(grid[0])
	for j := range width {
		for i := range height {
			if grid[i][j] {
				fmt.Printf("x")
			} else {
				fmt.Printf(".")
			}
		}
		fmt.Println()
	}
}

func main() {
	file, err := os.Open("input")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	positions := make([]Vector, 0)
	velocities := make([]Vector, 0)
	for scanner.Scan() {

		lineString := scanner.Text()

		separatedString := strings.Split(lineString, " ")
		positionString, velocityString := separatedString[0], separatedString[1]

		positionX, _ := strconv.Atoi(strings.Split(strings.Split(positionString, "=")[1], ",")[0])
		positionY, _ := strconv.Atoi(strings.Split(strings.Split(positionString, "=")[1], ",")[1])
		position := Vector{positionX, positionY}
		positions = append(positions, position)

		velocityX, _ := strconv.Atoi(strings.Split(strings.Split(velocityString, "=")[1], ",")[0])
		velocityY, _ := strconv.Atoi(strings.Split(strings.Split(velocityString, "=")[1], ",")[1])
		velocity := Vector{velocityX, velocityY}
		velocities = append(velocities, velocity)

	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	// bounds := Vector{11, 7}
	bounds := Vector{101, 103}
	time := 100
	cuadrantsCount := [4]int{}
	// visitedTiles := make(map[Vector]bool)

	for i := range positions {
		newPosition := positions[i].Add(velocities[i].times(time)).Mod(bounds)
		// if visitedTiles[newPosition] || newPosition.x == (bounds.x-1)/2 || newPosition.y == (bounds.y-1)/2 {
		if newPosition.x == (bounds.x-1)/2 || newPosition.y == (bounds.y-1)/2 {
			continue
		}
		// visitedTiles[newPosition] = true

		cuadrant := 0
		if newPosition.x*2/bounds.x == 1 {
			cuadrant++
		}
		if newPosition.y*2/bounds.y == 1 {
			cuadrant += 2
		}
		cuadrantsCount[cuadrant]++
	}
	xCuadrants := 1
	for _, cuadrantSum := range cuadrantsCount {
		xCuadrants *= cuadrantSum
	}
	fmt.Println(cuadrantsCount)
	fmt.Println(xCuadrants)

	//Part 2

	time = 1
out:
	for {
		visitedTiles := make([][]bool, bounds.x)
		for i := range bounds.x {
			visitedTiles[i] = make([]bool, bounds.y)
		}
		for i := range positions {
			newPosition := positions[i].Add(velocities[i].times(time)).Mod(bounds)
			// if visitedTiles[newPosition] || newPosition.x == (bounds.x-1)/2 || newPosition.y == (bounds.y-1)/2 {
			visitedTiles[newPosition.x][newPosition.y] = true
		}
		for i := range bounds.x - 2 {
			for j := range bounds.y - 2 {
				if visitedTiles[i][j] && visitedTiles[i+1][j] && visitedTiles[i+2][j] && visitedTiles[i][j+1] && visitedTiles[i+1][j+1] && visitedTiles[i+2][j+1] && visitedTiles[i][j+2] && visitedTiles[i+1][j+2] && visitedTiles[i+2][j+2] {
					displayGrid(visitedTiles)
					break out
				}
			}
		}
		time++
	}
	fmt.Println(time)
}
