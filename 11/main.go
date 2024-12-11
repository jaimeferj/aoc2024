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
	if currentStone == 0 {
		stoneList[idx] = 1
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
		return muteStones(stoneList, idx+1)
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

	for i := range 75 {
		stones = muteStones(stones, 0)
		fmt.Println(i, len(stones))
	}
	fmt.Println(len(stones))
}
