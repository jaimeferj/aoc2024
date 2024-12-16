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
	End
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

// func tryMove(gameMap *[][]BlockType, position Vector, moveVector Vector) bool {
// 	currentBlock := (*gameMap)[position.x][position.y]
// 	newPosition := position.Add(moveVector)
// 	block := (*gameMap)[newPosition.x][newPosition.y]
// 	if block == Void || (block == Box && tryMove(gameMap, newPosition, moveVector)) {
// 		(*gameMap)[newPosition.x][newPosition.y] = currentBlock
// 		(*gameMap)[position.x][position.y] = Void
// 		return true
//
// 	}
// 	return false
// }

func displayMap(gameMap [][]BlockType) {
	blockMapping := map[BlockType]rune{
		Border: '#',
		Void:   '.',
		Robot:  'S',
		End:    'E',
	}
	for _, row := range gameMap {
		for _, x := range row {
			fmt.Printf(string(blockMapping[x]))
		}
		fmt.Println()
	}
}

func isNode(gameMap [][]BlockType, position Vector) bool {
	i := position.x
	j := position.y
	voidSpaces := 0
	if gameMap[i][j] == Border {
		return false
	}
	di := 0
	dj := 0

	if gameMap[i+1][j] == Void {
		di++
		voidSpaces++
	}
	if gameMap[i-1][j] == Void {
		di--
		voidSpaces++
	}
	if gameMap[i][j+1] == Void {
		dj++
		voidSpaces++
	}
	if gameMap[i][j-1] == Void {
		dj--
		voidSpaces++
	}
	return voidSpaces > 2 || (voidSpaces == 2 && (di != 0 || dj != 0))
}

func main() {
	file, err := os.Open("test")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	gameMap := make([][]BlockType, 0)

	blockMapping := map[rune]BlockType{
		'#': Border,
		'.': Void,
		'S': Robot,
		'E': End,
	}

	moveToVector := make([]Vector, 4)
	moveToVector[Down] = DownVector
	moveToVector[Up] = UpVector
	moveToVector[Left] = LeftVector
	moveToVector[Right] = RightVector

	mapWidth := 0
	mapHeight := 0
	initialPosition := Vector{}
	// endPosition := Vector{}
	// initialDirection := Right

	for scanner.Scan() {

		lineString := scanner.Text()
		if lineString == "" {
			continue
		}
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
			if blockType == End {
				endPosition = Vector{mapHeight, i}
			}
		}
		gameMap = append(gameMap, row)
		mapHeight++
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	nodeList := make([]Vector, 0)
	nodeDistances := make(map[Vector]map[Vector]int)
	endDistance := make(map[Vector]int)
	maxInt := 9223372036854775807

	for i, row := range gameMap {
		if i == 0 || i+1 == mapWidth {
			continue
		}
		for j := range row {
			if j == 0 || j+1 == mapHeight {
				continue
			}
			node := Vector{i, j}
			if isNode(gameMap, node) {
				fmt.Println(node)
				nodeList = append(nodeList, node)
			}
		}
	}
	for i, node := range nodeList {
		if _, exists := nodeDistances[node]; !exists {
			nodeDistances[node] = make(map[Vector]int)
		}
		nodeDistances[node] = make(map[Vector]int)
		for _, otherNode := range nodeList[i+1:] {
			if node == otherNode || (node.x != otherNode.x && node.y != otherNode.y) {
				continue
			}
			distance := 0
			if node.x == otherNode.x {
				for i := range otherNode.y - node.y {
					if gameMap[otherNode.x][node.x+i] != Void {
						distance = -1
						break
					}
					distance++
				}
			} else {
				for i := range otherNode.x - node.x {
					if gameMap[otherNode.x+i][node.y] != Void {
						distance = -1
						break
					}
					distance++
				}
			}
			if distance == -1 {
				continue
			}
			nodeDistances[node][otherNode] = distance
			if _, exists := nodeDistances[otherNode]; !exists {
				nodeDistances[otherNode] = make(map[Vector]int)
			}
			nodeDistances[node][otherNode] = distance
			fmt.Println(node, otherNode, distance)
		}
		endDistance[node] = maxInt
	}
	endDistance[initialPosition] = 0

	// 	currentNode := initialPosition
	// 	unvisitedNodes := make([]Vector, len(nodeList))
	// 	copy(unvisitedNodes, nodeList)
	//
	// mainloop:
	// 	for {
	// 		currentEndDistance := maxInt
	// 		for _, node := range unvisitedNodes {
	// 			if endDistance[node] < currentEndDistance {
	// 				currentNode = node
	// 				currentEndDistance = endDistance[node]
	// 			}
	// 		}
	// 		if currentEndDistance == maxInt || currentNode.Equals(endPosition) {
	// 			break mainloop
	// 		}
	// 		for otherNode, distances := range nodeDistances {
	//
	// 			distance := currentEndDistance + distances[currentNode]
	// 			if distance < endDistance[otherNode] {
	// 				endDistance[otherNode] = distance
	// 			}
	// 		}
	// 	}

	// fmt.Println(gameMap)
	// fmt.Println(robotMovements)
	// position := initialPosition

	displayMap(gameMap)
	fmt.Println(nodeDistances)
	// for _, movement := range robotMovements {
	//
	// 	moveVector := moveToVector[movement]
	// 	if tryMove(&gameMap, position, moveVector) {
	// 		position = position.Add(Vector(moveVector))
	// 	}
	// }
}
