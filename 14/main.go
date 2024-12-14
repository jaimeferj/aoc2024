package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

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
		positions[i] = positions[i].Add(velocities[i].times(time)).Mod(bounds)
		// if visitedTiles[positions[i]] || positions[i].x == (bounds.x-1)/2 || positions[i].y == (bounds.y-1)/2 {
		if positions[i].x == (bounds.x-1)/2 || positions[i].y == (bounds.y-1)/2 {
			continue
		}
		// visitedTiles[positions[i]] = true

		cuadrant := 0
		if positions[i].x*2/bounds.x == 1 {
			cuadrant++
		}
		if positions[i].y*2/bounds.y == 1 {
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

}
