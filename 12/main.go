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

func isBorderPosition(position Vector, positionList []Vector, bounds Vector) bool {
	i, j := position.x, position.y
	if (i+1 < bounds.x && !(Vector{i + 1, j}.inList(positionList))) || (i-1 > 0 && !(Vector{i - 1, j}.inList(positionList))) || (j+1 < bounds.y && !(Vector{i, j + 1}.inList(positionList))) || (j-1 > 0 && !(Vector{i, j - 1}.inList(positionList)) || i-1 < 0 || j-1 < 0 || i+1 >= bounds.x || j+1 >= bounds.y) {
		return true
	}
	return false
}

func walkBorder(initialPosition Vector, initialDirection Vector, isValidPosition func(Vector) bool, onWalkCallback func(Vector), onSteerCallback func(Vector)) {
	onWalkCallback(initialPosition)
	currentDirection := initialDirection
	currentPosition := initialPosition
	for {
		if leftPos := currentPosition.Add(currentDirection.Rotate()); isValidPosition(leftPos) {
			currentDirection = currentDirection.Rotate()
			currentPosition = leftPos
			onSteerCallback(currentPosition)
		} else if frontPos := currentPosition.Add(currentDirection); isValidPosition(frontPos) {
			currentPosition = frontPos
		} else if rightPost := currentPosition.Add(currentDirection.RotateInverse()); isValidPosition(rightPost) {
			currentDirection = currentDirection.RotateInverse()
			currentPosition = rightPost
			onSteerCallback(currentPosition)
		} else {
			break
		}
		if currentPosition.Equals(initialPosition) {
			break
		}
		onWalkCallback(currentPosition)
	}
}

func isGoodOriented(position Vector, direction Vector, positionList []Vector, bounds Vector) bool {
	emptyDirection := direction.Rotate()
	emptyPosition := position.Add(emptyDirection)
	if emptyPosition.x < 0 || emptyDirection.y < 0 || emptyPosition.x >= bounds.x || emptyPosition.y >= bounds.y || !emptyDirection.inList(positionList) {
		return true
	} else {
		return false
	}
}

func findNewBorderDirection(positionList []Vector, position Vector, bounds Vector, borders []Vector, nonBorders []Vector) Vector {
	if tempPos := position.Add(Vector{1, 0}); tempPos.x < bounds.x && !tempPos.inList(borders) && !tempPos.inList(nonBorders) && isBorderPosition(tempPos, positionList, bounds) && isGoodOriented(position, Vector{1, 0}, positionList, bounds) {
		return Vector{1, 0}
	}
	if tempPos := position.Add(Vector{-1, 0}); tempPos.x >= 0 && !tempPos.inList(borders) && !tempPos.inList(nonBorders) && isBorderPosition(tempPos, positionList, bounds) && isGoodOriented(position, Vector{-1, 0}, positionList, bounds) {
		return Vector{-1, 0}
	}
	if tempPos := position.Add(Vector{0, 1}); tempPos.y < bounds.y && !tempPos.inList(borders) && !tempPos.inList(nonBorders) && isBorderPosition(tempPos, positionList, bounds) && isGoodOriented(position, Vector{0, 1}, positionList, bounds) {
		return Vector{0, 1}
	}
	if tempPos := position.Add(Vector{0, -1}); tempPos.y >= 0 && !tempPos.inList(borders) && !tempPos.inList(nonBorders) && isBorderPosition(tempPos, positionList, bounds) && isGoodOriented(position, Vector{0, -1}, positionList, bounds) {
		return Vector{0, -1}
	}
	panic("WHAT")
}

func calcSides(region [][]bool, bounds [2]int) int {
	sides := 0
	bordersVisited := make([]Vector, 0)
	nonBordersVisited := make([]Vector, 0)
	positionList := make([]Vector, 0)
	boundsVec := Vector{bounds[0], bounds[1]}
	for i, row := range region {
		for j, val := range row {
			if val {
				valVec := Vector{i, j}
				positionList = append(positionList, valVec)
			}
		}
	}
	onChangeSideCallback := func(vec Vector) {
		sides++
	}
	onWalkCallback := func(vec Vector) {
		bordersVisited = append(bordersVisited, vec)
	}
	isValidPosition := func(vec Vector) bool {
		return checkInBounds(vec, boundsVec) && region[vec.x][vec.y] && !vec.inList(bordersVisited) && !vec.inList(nonBordersVisited) && isBorderPosition(vec, positionList, boundsVec)
	}
	for _, position := range positionList {
		isBorderVisited := position.inList(bordersVisited)
		isNonBorderVisited := position.inList(nonBordersVisited)
		if !isBorderVisited && !isNonBorderVisited {
			if isBorderPosition(position, positionList, boundsVec) {
				initalDirection := findNewBorderDirection(positionList, position, boundsVec, bordersVisited, nonBordersVisited)
				walkBorder(position, initalDirection, isValidPosition, onWalkCallback, onChangeSideCallback)
			} else {
				nonBordersVisited = append(nonBordersVisited, position)
			}
		}

	}
	return sides
}

func main() {
	file, err := os.Open("test")
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
	vertexBounds := [2]int{gameHeight*2 - 1, gameWidth*2 - 1}
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
				region := make([][]bool, gameHeight*2-1)
				for i := range gameHeight*2 - 1 {
					region[i] = make([]bool, gameWidth*2-1)
				}
				onVisit := func(i int, j int) {
					if i*2 >= 0 {
						if j*2+1 < gameWidth*2-1 {
							region[i*2][j*2+1] = true
						}
						if j*2 >= 0 {
							region[i*2][j*2] = true
						}
					}
					if i*2+1 < gameHeight*2-1 {
						if j*2+1 < gameWidth*2-1 {
							region[i*2+1][j*2+1] = true
						}
						if j*2 >= 0 {
							region[i*2+1][j*2] = true
						}
					}
				}
				result := walk(gameMap, bounds, i, j, visitedPositions, onVisit)
				area, perimeter := result[0], result[1]
				totalCost += area * perimeter
				sides := calcSides(region, vertexBounds)
				totalCost2 += sides * area
				uuidRegion++
			}
		}
	}
	fmt.Println(totalCost)
	fmt.Println(totalCost2)

}
