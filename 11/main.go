package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"slices"
	"strconv"
	"strings"
)

var memo = map[int]int{0: 1}

type Node struct {
	value int
	next  *Node
	prev  *Node
}

type LinkedList struct {
	head   *Node
	tail   *Node
	length int
}

func (l *LinkedList) InsertAfter(current *Node, value int) *Node {
	newNode := &Node{value: value}
	if current == nil {
		l.head = newNode
		l.tail = newNode
		return newNode
	}
	newNode.next = current.next
	newNode.prev = current
	if current.next != nil {
		current.next.prev = newNode
	}
	current.next = newNode
	if current == l.tail {
		l.tail = newNode
	}
	l.length++
	return newNode
}

func (l *LinkedList) Display() {
	for current := l.head; current != nil; current = current.next {
		fmt.Print(current.value, " ")
	}
	fmt.Println()
}

func numberDigits(x int) int {
	i := 0
	for {
		i++
		x = x / 10
		if x < 1 {
			return i
		}
	}
}

func muteStones(stoneList []int, idx int) []int {
	if idx >= len(stoneList) {
		return stoneList
	}
	currentStone := stoneList[idx]
	if result, ok := memo[currentStone]; ok {
		stoneList[idx] = result
		return muteStones(stoneList, idx+1)
	}
	if currentStone == 0 {
		stoneList[idx] = 1
		memo[0] = 1
		return muteStones(stoneList, idx+1)
	} else if stoneDigits := numberDigits(currentStone); stoneDigits%2 == 0 {
		tenPower := 1
		for range stoneDigits / 2 {
			tenPower *= 10
		}
		upperDigits := currentStone / tenPower
		lowerDigits := currentStone - upperDigits*tenPower
		stoneList[idx] = upperDigits
		stoneList = slices.Insert(stoneList, idx+1, lowerDigits)
		return muteStones(stoneList, idx+2)
	} else {
		stoneList[idx] *= 2024
		memo[currentStone] = stoneList[idx]
		return muteStones(stoneList, idx+1)
	}

}

func muteStonesNoRecursive(stoneList []int) []int {
	newStones := make([]int, len(stoneList))
	newIdx := 0
	for _, stone := range stoneList {
		if result, ok := memo[stone]; ok {
			newStones[newIdx] = result
		} else if stoneDigits := numberDigits(stone); stoneDigits%2 == 0 {
			tenPower := 1
			for range stoneDigits / 2 {
				tenPower *= 10
			}
			upperDigits := stone / tenPower
			lowerDigits := stone - upperDigits*tenPower
			newStones[newIdx] = upperDigits
			newStones = slices.Insert(newStones, newIdx+1, lowerDigits)
			newIdx++
		} else {
			newStones[newIdx] = stone * 2024
			memo[stone] = newStones[newIdx]
		}
		newIdx++
	}
	return newStones
}

func muteStonesNoRecursiveLL(stones *LinkedList) {
	justInserted := false
	for stone := stones.head; stone != nil; stone = stone.next {
		if justInserted {
			justInserted = false
			continue
		}
		stoneValue := stone.value
		if stoneValue == 0 {
			stone.value = 1
		} else if stoneDigits := numberDigits(stoneValue); stoneDigits%2 == 0 {
			tenPower := 1
			for range stoneDigits / 2 {
				tenPower *= 10
			}
			upperDigits := stoneValue / tenPower
			lowerDigits := stoneValue - upperDigits*tenPower
			stone.value = upperDigits
			stones.InsertAfter(stone, lowerDigits)
			justInserted = true
		} else {
			stone.value = stoneValue * 2024
		}
	}
}

func sumList(x []int) int {
	acc := 0
	for y := range x {
		acc += y
	}
	return acc

}

func main() {
	file, err := os.Open("input")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	stones := make([]int, 0)
	for scanner.Scan() {

		lineString := scanner.Text()
		for _, x := range strings.Split(lineString, " ") {
			number, _ := strconv.Atoi(x)
			stones = append(stones, number)
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	stonesLL := &LinkedList{}
	for _, stone := range stones {
		stonesLL.InsertAfter(stonesLL.tail, stone)
	}
	stonesLL.Display()
	for i := range 75 {
		// stones = muteStones(stones, 0)
		// stones = muteStonesNoRecursive(stones)
		muteStonesNoRecursiveLL(stonesLL)
		fmt.Println(i, stonesLL.length)
	}
	fmt.Println(stonesLL.length)
}
