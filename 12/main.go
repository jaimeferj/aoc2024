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

func calcSides(region [][]bool) int {
	sides := 0
	initialPosition := [2]int{}
	for i, row := range region {
		for j, val := range row {
			if val {
				initialPosition = [2]int{i, j}
				break
			}
		}
	}
	direction := [2]int{0, 1}
	upDirection := [2]int{1, 0}
	for {
		if initialPosition[0] + direction[0]
	}
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

	for i := range gameHeight {
		for j := range gameWidth {
			if !visitedPositions[i][j] {
				region := make([][]bool, gameHeight)
				for range gameHeight {
					region = append(region, make([]int, gameWidth))
				}
				onVisit := func(i int, j int) {
					region[i][j] = true
				}
				result := walk(gameMap, bounds, i, j, visitedPositions, onVisit)
				totalCost += result[0] * result[1]
				uuidRegion++
			}
		}
	}
	fmt.Println(totalCost)

}
