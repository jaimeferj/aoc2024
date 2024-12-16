package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

type BlockType uint8
type BlockType2 uint8

const (
	Border BlockType = iota
	Void
	Robot
	Box
)
const (
	Border2 BlockType2 = iota
	Void2
	Robot2
	BoxOpen
	BoxClose
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

func processQueue(gameMap *[][]BlockType2, changesQueue *map[Vector]BlockType2) {
	for vec, block := range *changesQueue {
		(*gameMap)[vec.x][vec.y] = block
	}
}

func tryMove2(gameMap *[][]BlockType2, position Vector, moveVector Vector, changesMap *map[Vector]BlockType2) bool {
	currentBlock := (*gameMap)[position.x][position.y]
	newPosition := position.Add(moveVector)
	block := (*gameMap)[newPosition.x][newPosition.y]
	upDir := moveVector.y == 0
	rightBlock := Border2
	isBoxUp := false
	if upDir {
		if currentBlock == BoxClose {
			isBoxUp = true
			position.y--
			newPosition.y--
			rightBlock = block
			block = (*gameMap)[newPosition.x][newPosition.y]
		} else if currentBlock == BoxOpen {
			isBoxUp = true
			rightBlock = (*gameMap)[newPosition.x][newPosition.y+1]
		}
	}
	if !isBoxUp && (block == Void2 || ((block == BoxOpen || block == BoxClose) && tryMove2(gameMap, newPosition, moveVector, changesMap))) {
		(*changesMap)[newPosition] = currentBlock
		if _, exists := (*changesMap)[position]; !exists {
			(*changesMap)[position] = Void2

		}
		if currentBlock == Robot2 {
			processQueue(gameMap, changesMap)
		}
		return true
	} else if isBoxUp && ((block == Void2 && rightBlock == Void2) || (block == Void2 && rightBlock == BoxOpen && tryMove2(gameMap, newPosition.Add(Vector{0, 1}), moveVector, changesMap)) || (block == BoxClose && rightBlock == Void2 && tryMove2(gameMap, newPosition, moveVector, changesMap)) || (block == BoxClose && rightBlock == BoxOpen && tryMove2(gameMap, newPosition, moveVector, changesMap) && tryMove2(gameMap, newPosition.Add(Vector{0, 1}), moveVector, changesMap) || (block == BoxOpen && rightBlock == BoxClose && tryMove2(gameMap, newPosition, moveVector, changesMap)))) {
		(*changesMap)[newPosition] = BoxOpen
		(*changesMap)[newPosition.Add(Vector{0, 1})] = BoxClose
		if _, exists := (*changesMap)[position]; !exists {
			(*changesMap)[position] = Void2
		}
		if _, exists := (*changesMap)[position.Add(Vector{0, 1})]; !exists {
			(*changesMap)[position.Add(Vector{0, 1})] = Void2
		}
		if currentBlock == Robot2 {
			processQueue(gameMap, changesMap)
		}
		return true
	}
	return false
}

func displayMap2(gameMap [][]BlockType2) {
	blockMapping := map[BlockType2]rune{
		Border2:  '#',
		Void2:    '.',
		Robot2:   '@',
		BoxOpen:  '[',
		BoxClose: ']',
	}
	for _, row := range gameMap {
		for _, x := range row {
			fmt.Printf(string(blockMapping[x]))
		}
		fmt.Println()
	}
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
	gameMap2 := make([][]BlockType2, 0)
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
	inverseMoveMapping := map[MoveId]rune{
		Down:  '^',
		Right: '>',
		Left:  '<',
		Up:    'v',
	}
	moveToVector := make([]Vector, 4)
	moveToVector[Down] = DownVector
	moveToVector[Up] = UpVector
	moveToVector[Left] = LeftVector
	moveToVector[Right] = RightVector

	mapWidth := 0
	mapHeight := 0
	initialPosition := Vector{}
	initialPosition2 := Vector{}

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
			row2 := make([]BlockType2, mapWidth*2)
			for i, x := range lineString {
				blockType := blockMapping[x]
				row[i] = blockType
				if blockType == Border || blockType == Void {
					row2[i*2] = BlockType2(blockType)
					row2[i*2+1] = BlockType2(blockType)
				} else if blockType == Robot {
					row2[i*2] = Robot2
					row2[i*2+1] = Void2
				} else if blockType == Box {
					row2[i*2] = BoxOpen
					row2[i*2+1] = BoxClose
				}
				if blockType == Robot {
					initialPosition = Vector{mapHeight, i}
					initialPosition2 = Vector{mapHeight, i * 2}
				}
			}
			gameMap = append(gameMap, row)
			gameMap2 = append(gameMap2, row2)
			mapHeight++

		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	// fmt.Println(gameMap)
	// fmt.Println(robotMovements)
	position := initialPosition
	position2 := initialPosition2

	displayMap2(gameMap2)
	for _, movement := range robotMovements {

		moveVector := moveToVector[movement]
		if tryMove(&gameMap, position, moveVector) {
			position = position.Add(Vector(moveVector))
		}
		queue := make(map[Vector]BlockType2)
		if tryMove2(&gameMap2, position2, moveVector, &queue) {
			position2 = position2.Add(Vector(moveVector))
		}
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
	acc2 := 0
	for i, row := range gameMap2 {
		for j, x := range row {
			if x != BoxOpen {
				continue
			}
			acc2 += 100*i + j
		}
	}
	fmt.Println(acc)
	fmt.Println(acc2)
}
