package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func updateDirection(currentDirection int) int {
	newDirection := currentDirection + 1
	if newDirection < 4 {
		return newDirection
	} else {
		return 0
	}
}

func findPosition(gameMap [][]string) [2]int {
	for i, line := range gameMap {
		for j, x := range line {
			if x == "^" {
				return [2]int{i, j}
			}
		}
	}
	return [2]int{-1, -1}
}

var gameDirection map[int][2]int
var mapHeight int
var mapWidth int

func runMap(gameMap [][]string, onWalkCallback func(x int, y int, direction int) bool) {
	currentPosition := findPosition(gameMap)
	nextPosition := currentPosition
	currentDirection := 0
	nextMapItem := ""

	for {
		nextPosition[0] = currentPosition[0] + gameDirection[currentDirection][0]
		nextPosition[1] = currentPosition[1] + gameDirection[currentDirection][1]
		if nextPosition[0] < 0 || nextPosition[0] >= mapHeight || nextPosition[1] < 0 || nextPosition[1] >= mapWidth {
			break
		}
		nextMapItem = gameMap[nextPosition[0]][nextPosition[1]]
		if nextMapItem == "#" {
			currentDirection = updateDirection(currentDirection)
		} else {
			repeatedPath := onWalkCallback(nextPosition[0], nextPosition[1], currentDirection)
			if repeatedPath {
				break
			}
			currentPosition = nextPosition
		}
	}
}

func main() {

	f, err := os.Open("input")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	gameMap := make([][]string, 0)
	gameDirection = make(map[int][2]int)
	gameDirection[0] = [2]int{-1, 0}
	gameDirection[1] = [2]int{0, 1}
	gameDirection[2] = [2]int{1, 0}
	gameDirection[3] = [2]int{0, -1}

	for scanner.Scan() {
		lineString := scanner.Text()
		splittedString := strings.Split(lineString, "")
		gameMap = append(gameMap, splittedString)
	}
	mapWidth = len(gameMap[0])
	mapHeight = len(gameMap)
	visitedPositions := make([][]bool, mapHeight)
	visitedPositionsDirections := make([][][]bool, mapHeight)
	for i := range visitedPositions {
		visitedPositions[i] = make([]bool, mapWidth)
		visitedPositionsDirections[i] = make([][]bool, mapWidth)
		for j := range visitedPositionsDirections {
			visitedPositionsDirections[i][j] = make([]bool, 4)
		}
	}
	onWalkCallback := func(x int, y int, direction int) bool {
		visitedPositions[x][y] = true
		if visitedPositionsDirections[x][y][direction] {
			return true
		} else {
			visitedPositionsDirections[x][y][direction] = true
			return false

		}
	}
	runMap(gameMap, onWalkCallback)
	totalVisitedPositions := 0
	for i := range visitedPositions {
		for j := range visitedPositions[i] {
			if visitedPositions[i][j] {
				totalVisitedPositions++
			}
		}
	}
	fmt.Println(totalVisitedPositions)

	initialPosition := findPosition(gameMap)
	// Part 2
	totalLoopPosibilities := 0
	newGameMap := make([][]string, mapHeight)
	for k := range newGameMap {
		newGameMap[k] = make([]string, mapWidth)
		copy(newGameMap[k], gameMap[k])
	}
	runMap(gameMap, onWalkCallback)
	for i, row := range visitedPositions {
		for j, isVisited := range row {
			if isVisited && !(i == initialPosition[0] && j == initialPosition[1]) {
				visitedPositionsDirectionsLoop := make([][][]bool, mapHeight)
				for k := range visitedPositionsDirectionsLoop {
					visitedPositionsDirectionsLoop[k] = make([][]bool, mapWidth)
					for l := range visitedPositionsDirectionsLoop {
						visitedPositionsDirectionsLoop[k][l] = make([]bool, 4)
					}
				}
				onWalkCallbackLoop := func(x int, y int, direction int) bool {
					if visitedPositionsDirectionsLoop[x][y][direction] {
						totalLoopPosibilities++
						return true
					} else {
						visitedPositionsDirectionsLoop[x][y][direction] = true
						return false

					}
				}
				for k := range newGameMap {
					copy(newGameMap[k], gameMap[k])
				}
				newGameMap[i][j] = "#"
				runMap(newGameMap, onWalkCallbackLoop)
			}
		}
	}
	fmt.Println(totalLoopPosibilities)

}
