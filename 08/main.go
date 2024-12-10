package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func isAsciiLetterOrNumber(r rune) bool {
	return (r >= 'A' && r <= 'Z') || (r >= 'a' && r <= 'z') || (r >= '0' && r <= '9')
}

func checkInBounds(X [2]int, bounds [2]int) bool {
	x, y := X[0], X[1]
	width, height := bounds[0], bounds[1]
	return x >= 0 && x < height && y >= 0 && y < width
}

func main() {

	f, err := os.Open("input")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)

	// numberOfAntinodes := 0

	gameMap := make([][]rune, 0)
	for scanner.Scan() {
		lineString := scanner.Text()
		lineRunes := []rune(lineString)
		gameMap = append(gameMap, lineRunes)
	}

	mapWidth := len(gameMap[0])
	mapHeight := len(gameMap)
	mapBounds := [2]int{mapWidth, mapHeight}
	antennaLocations := make(map[rune][][2]int, 0)
	antinodesLocations := make(map[[2]int]bool, 0)
	antinodesLocations2 := make(map[[2]int]bool, 0)

	for i, row := range gameMap {
		for j, location := range row {
			if isAsciiLetterOrNumber(location) {
				_, ok := antennaLocations[location]
				if !ok {
					antennaLocations[location] = [][2]int{{i, j}}
				} else {
					for _, antenna := range antennaLocations[location] {
						dX, dY := i-antenna[0], j-antenna[1]
						antinode1 := [2]int{antenna[0] + 2*dX, antenna[1] + 2*dY}
						antinode2 := [2]int{antenna[0] - dX, antenna[1] - dY}
						if checkInBounds(antinode1, mapBounds) {
							antinodesLocations[antinode1] = true
						}
						if checkInBounds(antinode2, mapBounds) {
							antinodesLocations[antinode2] = true
						}
						k := 0
						for {
							antinode := [2]int{antenna[0] + (k)*dX, antenna[1] + (k)*dY}
							if !checkInBounds(antinode, mapBounds) {
								break
							} else {
								antinodesLocations2[antinode] = true
							}
							k++
						}
						k = 0
						for {
							antinode := [2]int{antenna[0] - k*dX, antenna[1] - k*dY}
							if !checkInBounds(antinode, mapBounds) {
								break
							} else {
								antinodesLocations2[antinode] = true
							}
							k++
						}
					}
					antennaLocations[location] = append(antennaLocations[location], [2]int{i, j})
				}

			}
		}
	}

	fmt.Println(len(antinodesLocations))
	fmt.Println(len(antinodesLocations2))
}
