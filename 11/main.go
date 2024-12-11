package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"time"
)

var start time.Time

func init() {
	start = time.Now()
}

var memoBlinks = map[[2]int]int{}

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

func muteStonesBlinks(stone int, blinks int) int {
	totalLength := 1
	originalStone := stone
	if blinks == 0 {
		return 1
	}
	if result, ok := memoBlinks[[2]int{stone, blinks}]; ok {
		return result
	}
	for i := range blinks {
		if result, ok := memoBlinks[[2]int{stone, blinks - i}]; ok {
			totalLength += result
			break
		}
		if stone == 0 {
			stone = 1
		} else if stoneDigits := numberDigits(stone); stoneDigits%2 == 0 {
			tenPower := 1
			for range stoneDigits / 2 {
				tenPower *= 10
			}
			upperDigits := stone / tenPower
			lowerDigits := stone - upperDigits*tenPower
			stone = upperDigits
			totalLength += muteStonesBlinks(lowerDigits, blinks-i-1)
		} else {
			stone *= 2024
		}
	}
	memoBlinks[[2]int{originalStone, blinks}] = totalLength
	return totalLength
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

	totalLength := 0
	fmt.Println("main execution started at time", time.Since(start))
	for _, stone := range stones {
		totalLength += muteStonesBlinks(stone, 75)
	}
	fmt.Println(totalLength)
	fmt.Println("\nmain execution stopped at time", time.Since(start))
}
