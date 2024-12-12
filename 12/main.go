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

type Vector struct {
	x int
	y int
}

func checkInBounds(X Vector, bounds Vector) bool {
	return X.x >= 0 && X.x < bounds.x && X.y >= 0 && X.y < bounds.y
}

func (X Vector) Equals(Y Vector) bool {
	return X.x == Y.x && X.y == Y.y
}

func (X Vector) Add(Y Vector) Vector {
	return Vector{X.x + Y.x, X.y + Y.y}
}

func (X Vector) Revert() Vector {
	return Vector{-X.x, -X.y}
}

func (X Vector) Dot(Y Vector) int {
	return X.x*Y.y + X.y*Y.y
}

func (X Vector) Orientation(Y Vector) int {
	product := X.x*Y.y + X.y*Y.x
	if product > 0 {
		return 1
	} else if product == 0 {
		return 0
	} else {
		return -1
	}
}

func (X Vector) Rotate() Vector {
	return Vector{-X.y, X.x}
}

func (X Vector) RotateInverse() Vector {
	return Vector{X.y, -X.x}
}

func calcSides(region [][]bool, bounds [2]int) int {
	sides := 0
	initialPosition := Vector{}
out:
	for i, row := range region {
		for j, val := range row {
			if val {
				initialPosition = Vector{i, j}
				break out
			}
		}
	}
	// Calculate initialDirection
	occList := make([]Vector, 0)
	i := initialPosition.x
	j := initialPosition.y
	if i+1 < bounds[0] {
		if region[i+1][j] {
			occList = append(occList, Vector{i + 1, j})
		}
	}
	if j+1 < bounds[1] {
		if region[i][j+1] {
			occList = append(occList, Vector{i, j + 1})
		}

	}
	if i-1 >= 0 {
		if region[i-1][j] {
			occList = append(occList, Vector{i - 1, j})
		}
	}
	if j-1 >= 0 {
		if region[i][j-1] {
			occList = append(occList, Vector{i, j - 1})
		}
	}
	if len(occList) == 0 {
		return 4
	}

	firstDirection := occList[0].Add(initialPosition.Revert())
	downDirection := Vector{}
	if len(occList) == 1 {
		downDirection = firstDirection.Rotate()
	} else {
		downDirection = occList[1].Add(initialPosition.Revert())
	}

	currentDirection := firstDirection
	currentPosition := initialPosition
	boundsVec := Vector{bounds[0], bounds[1]}
	// TODO: It is easier if we convert from the center of the squares to a vertex approximation!
	for {
		nextPosition := currentPosition.Add(currentDirection)
		if nextPosition.Equals(initialPosition) {
			break
		}
		if !checkInBounds(nextPosition, boundsVec) || !region[nextPosition.x][nextPosition.y] {
			newDownDirection := Vector{}
			if currentDirection.Orientation(downDirection) == 1 {
				newDownDirection = downDirection.Rotate()
			} else {
				newDownDirection = downDirection.RotateInverse()
			}
			currentDirection = downDirection
			downDirection = newDownDirection
			sides++
			continue
		}
		currentPosition = nextPosition
	}
	return sides
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
					region = append(region, make([]bool, gameWidth))
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
