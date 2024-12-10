package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func walk(gameMap [][]int, bounds [2]int, i int, j int, finishCallback func([2]int)) {
	currentValue := gameMap[i][j]
	if currentValue == 9 {
		finishCallback([2]int{i, j})
		return
	}
	if i+1 < bounds[0] && gameMap[i+1][j]-1 == currentValue {
		walk(gameMap, bounds, i+1, j, finishCallback)
	}
	if j+1 < bounds[1] && gameMap[i][j+1]-1 == currentValue {
		walk(gameMap, bounds, i, j+1, finishCallback)
	}
	if i-1 >= 0 && gameMap[i-1][j]-1 == currentValue {
		walk(gameMap, bounds, i-1, j, finishCallback)
	}
	if j-1 >= 0 && gameMap[i][j-1]-1 == currentValue {
		walk(gameMap, bounds, i, j-1, finishCallback)
	}
}

func main() {
	file, err := os.Open("input")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	gameMap := make([][]int, 0)
	gameWidth := 0
	gameHeight := 0
	for scanner.Scan() {

		lineString := scanner.Text()
		if gameWidth == 0 {
			gameWidth = len(lineString)
		}
		lineNumbers := make([]int, gameWidth)
		for i, x := range lineString {
			if x == []rune(".")[0] {
				lineNumbers[i] = 99
				continue
			}
			number := int(x - '0')
			lineNumbers[i] = number
		}
		gameMap = append(gameMap, lineNumbers)
		gameHeight++
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	bounds := [2]int{gameHeight, gameWidth}
	totalTrails := 0
	totalRating := 0
	for i := range gameHeight {
		for j := range gameWidth {
			if gameMap[i][j] == 0 {
				visitedNines := map[[2]int]bool{}
				// visitedNinesAmounts := map[[2]int]int{}
				onFinish := func(position [2]int) {
					if !visitedNines[position] {
						visitedNines[position] = true
						totalTrails++
						// fmt.Println(i, j, totalTrails)

					}
					totalRating++

				}
				walk(gameMap, bounds, i, j, onFinish)
			}
		}
	}
	fmt.Println(totalTrails)
	fmt.Println(totalRating)

}
