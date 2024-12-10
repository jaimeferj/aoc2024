package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {

	f, err := os.Open("input")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)

	leftList := []int{}
	rightList := []int{}

	for scanner.Scan() {
		splitLine := strings.Split(scanner.Text(), "   ")
		left, _ := strconv.Atoi(splitLine[0])
		right, _ := strconv.Atoi(splitLine[1])
		leftList = append(leftList, left)
		rightList = append(rightList, right)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	sort.Ints(leftList)
	sort.Ints(rightList)

	totalDistance := 0
	similarity := 0
	for i := 0; i < len(leftList); i++ {
		leftValue := leftList[i]
		distance := leftValue - rightList[i]
		if distance < 0 {
			distance = -distance
		}
		totalDistance += distance

		rightSeenTimes := 0
		for j := 0; j < len(rightList); j++ {
			rightValue := rightList[j]
			if rightValue > leftValue {
				break
			}
			if rightValue < leftValue {
				continue
			}
			rightSeenTimes++

		}
		similarity += rightSeenTimes * leftValue
	}

	fmt.Println("Total distance: ", totalDistance)
	fmt.Println("Total similarity: ", similarity)

}
