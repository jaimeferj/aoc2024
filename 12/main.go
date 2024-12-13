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

func (X Vector) inList(vecList []Vector) bool {
	for _, pos := range vecList {
		if X.x == pos.x && X.y == pos.y {
			return true
		}
	}
	return false

}

func countNeighbours(vertex Vector, centers [][]bool, bounds Vector) int {
	count := 0
	if vec := vertex.Add(Vector{0, 0}); checkInBounds(vec, bounds) && centers[vec.x][vec.y] {
		count++
	}
	if vec := vertex.Add(Vector{-1, 0}); checkInBounds(vec, bounds) && centers[vec.x][vec.y] {
		count++
	}
	if vec := vertex.Add(Vector{0, -1}); checkInBounds(vec, bounds) && centers[vec.x][vec.y] {
		count++
	}
	if vec := vertex.Add(Vector{-1, -1}); checkInBounds(vec, bounds) && centers[vec.x][vec.y] {
		count++
	}
	return count
}

func isDiagonalNeighbour(vertex Vector, centers [][]bool, bounds Vector) bool {
	if vec := vertex.Add(Vector{0, 0}); checkInBounds(vec, bounds) && centers[vec.x][vec.y] {
		if vec := vertex.Add(Vector{-1, -1}); checkInBounds(vec, bounds) && centers[vec.x][vec.y] {
			return true
		}
	}
	if vec := vertex.Add(Vector{-1, 0}); checkInBounds(vec, bounds) && centers[vec.x][vec.y] {
		if vec := vertex.Add(Vector{0, -1}); checkInBounds(vec, bounds) && centers[vec.x][vec.y] {
			return true
		}
	}
	return false
}

func calcSides(regionVertex [][]bool, regionCenters [][]bool, centerBounds [2]int) int {
	sides := 0
	boundsVec := Vector{centerBounds[0], centerBounds[1]}
	for i, row := range regionVertex {
		for j, val := range row {
			if val {
				valVec := Vector{i, j}
				neighbours := countNeighbours(valVec, regionCenters, boundsVec)
				if neighbours%2 != 0 {
					sides++
				} else if neighbours == 2 && isDiagonalNeighbour(valVec, regionCenters, boundsVec) {
					sides += 2
				}
			}
		}
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
	totalCost2 := 0
	uuidRegion := 0

	for i := range gameHeight {
		for j := range gameWidth {
			if !visitedPositions[i][j] {
				regionVertex := make([][]bool, gameHeight+1)
				for i := range gameHeight + 1 {
					regionVertex[i] = make([]bool, gameWidth+1)
				}
				regionCenters := make([][]bool, gameHeight)
				for i := range gameHeight {
					regionCenters[i] = make([]bool, gameWidth)
				}
				onVisit := func(i int, j int) {
					regionCenters[i][j] = true
					regionVertex[i][j] = true
					regionVertex[i][j+1] = true
					regionVertex[i+1][j] = true
					regionVertex[i+1][j+1] = true
				}
				result := walk(gameMap, bounds, i, j, visitedPositions, onVisit)
				area, perimeter := result[0], result[1]
				totalCost += area * perimeter
				sides := calcSides(regionVertex, regionCenters, bounds)
				totalCost2 += sides * area
				uuidRegion++
			}
		}
	}
	fmt.Println(totalCost)
	fmt.Println(totalCost2)

}
