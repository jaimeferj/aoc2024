package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

type BlockType uint8

const (
	Border BlockType = iota
	Void
	Robot
	Box
)

type MoveId uint8

var (
	UpVector    Vector = Vector{1, 0}
	DownVector  Vector = Vector{-1, 0}
	RightVector Vector = Vector{0, 1}
	LeftVector  Vector = Vector{0, -1}
)

const (
	Up MoveId = iota
	Down
	Right
	Left
)

func tryMove(gameMap *[][]BlockType, position Vector, moveVector Vector) bool {
	currentBlock := (*gameMap)[position.x][position.y]
	newPosition := position.Add(moveVector)
	block := (*gameMap)[newPosition.x][newPosition.y]
	if block == Void || (block == Box && tryMove(gameMap, newPosition, moveVector)) {
		(*gameMap)[newPosition.x][newPosition.y] = currentBlock
		(*gameMap)[position.x][position.y] = Void
		return true

	}
	return false
}

func displayMap(gameMap [][]BlockType) {
	blockMapping := map[BlockType]rune{
		Border: '#',
		Void:   '.',
		Robot:  '@',
		Box:    'O',
	}
	for _, row := range gameMap {
		for _, x := range row {
			fmt.Printf(string(blockMapping[x]))
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
	gameMap := make([][]BlockType, 0)
	robotMovements := make([]MoveId, 0)

	blockMapping := map[rune]BlockType{
		'#': Border,
		'.': Void,
		'@': Robot,
		'O': Box,
	}
	moveMapping := map[rune]MoveId{
		'^': Down,
		'>': Right,
		'<': Left,
		'v': Up,
	}
	moveToVector := make([]Vector, 4)
	moveToVector[Down] = DownVector
	moveToVector[Up] = UpVector
	moveToVector[Left] = LeftVector
	moveToVector[Right] = RightVector

	mapWidth := 0
	mapHeight := 0
	initialPosition := Vector{}

	for scanner.Scan() {

		lineString := scanner.Text()
		if lineString == "" {
			continue
		}
		if firstChar := []rune(lineString)[0]; firstChar == 'v' || firstChar == '>' || firstChar == '<' || firstChar == '^' {
			for _, x := range lineString {
				robotMovements = append(robotMovements, moveMapping[x])
			}
		} else {
			if mapWidth == 0 {
				mapWidth = len(lineString)
			}

			row := make([]BlockType, mapWidth)
			for i, x := range lineString {
				blockType := blockMapping[x]
				row[i] = blockType
				if blockType == Robot {
					initialPosition = Vector{mapHeight, i}
				}
			}
			gameMap = append(gameMap, row)
			mapHeight++

		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	// fmt.Println(gameMap)
	// fmt.Println(robotMovements)
	position := initialPosition

	for _, movement := range robotMovements {

		moveVector := moveToVector[movement]
		if tryMove(&gameMap, position, moveVector) {
			position = position.Add(Vector(moveVector))
		}
		// displayMap(gameMap)
	}
	acc := 0
	for i, row := range gameMap {
		for j, x := range row {
			if x != Box {
				continue
			}
			acc += 100*i + j
		}
	}
	fmt.Println(acc)
}
