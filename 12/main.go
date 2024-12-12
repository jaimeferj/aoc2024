package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func walk(gameMap [][]rune, bounds [2]int, i int, j int, visitedPositions [][]bool, onVisitCallback func(int, int)) [2]int {
	onVisitCallback(i, j)
	visitedPositions[i][j] = true
	currentValue := gameMap[i][j]
	area := 1
	perimeter := 4
	if i+1 < bounds[0] && gameMap[i+1][j] == currentValue {
		if !visitedPositions[i+1][j] {
			result := walk(gameMap, bounds, i+1, j, visitedPositions, onVisitCallback)
			area += result[0]
			perimeter += result[1]
		}
		perimeter--
	}
	if j+1 < bounds[1] && gameMap[i][j+1] == currentValue {
		if !visitedPositions[i][j+1] {
			result := walk(gameMap, bounds, i, j+1, visitedPositions, onVisitCallback)
			area += result[0]
			perimeter += result[1]
		}
		perimeter--
	}
	if i-1 >= 0 && gameMap[i-1][j] == currentValue {
		if !visitedPositions[i-1][j] {
			result := walk(gameMap, bounds, i-1, j, visitedPositions, onVisitCallback)
			area += result[0]
			perimeter += result[1]
		}
		perimeter--
	}
	if j-1 >= 0 && gameMap[i][j-1] == currentValue {
		if !visitedPositions[i][j-1] {
			result := walk(gameMap, bounds, i, j-1, visitedPositions, onVisitCallback)
			area += result[0]
			perimeter += result[1]
		}
		perimeter--
	}
	return [2]int{area, perimeter}
}

func main() {
	file, err := os.Open("input")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	gameMap := make([][]rune, 0)
	gameWidth := 0
	gameHeight := 0
	for scanner.Scan() {

		lineString := scanner.Text()
		if gameWidth == 0 {
			gameWidth = len(lineString)
		}
		runeList := make([]rune, gameWidth)
		for i, x := range lineString {
			runeList[i] = x
		}
		gameMap = append(gameMap, runeList)
		gameHeight++
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	bounds := [2]int{gameHeight, gameWidth}
	visitedPositions := make([][]bool, gameHeight)
	for i := range gameHeight {
		visitedPositions[i] = make([]bool, gameWidth)
	}

	totalCost := 0
	uuidRegion := 0
	regions := make(map[int][][2]int)

	for i := range gameHeight {
		for j := range gameWidth {
			if !visitedPositions[i][j] {
				onVisit := func(i int, j int) {
					regions[uuidRegion] = append(regions[uuidRegion], [2]int{i, j})
				}
				regions[uuidRegion] = make([][2]int, 0)
				result := walk(gameMap, bounds, i, j, visitedPositions, onVisit)
				totalCost += result[0] * result[1]
				uuidRegion++
			}
		}
	}
	fmt.Println(totalCost)

}
