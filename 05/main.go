package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func mapKeys(mapInput map[int]bool) []int {
	keys := make([]int, 0, len(mapInput))
	for k := range mapInput {
		keys = append(keys, k)
	}
	return keys
}

func checkInList(input int, checkList []int) bool {
	// fmt.Printf("Checking %s in %s...", input, checkList)
	for _, x := range checkList {
		if x == input {
			return true
		}
	}
	return false
}

func permuteInvalid(updates []string, mappingOrder map[int]map[int]bool) ([]string, bool) {
	badLeftIdx := 0
	badRightIdx := 0
	invalidUpdate := false
	for i, updateLeft := range updates {
		leftNumber, _ := strconv.Atoi(updateLeft)
		for j, updateRight := range updates[i+1:] {
			rightNumber, _ := strconv.Atoi(updateRight)
			if !(checkInList(rightNumber, mapKeys(mappingOrder[leftNumber])) && !checkInList(leftNumber, mapKeys(mappingOrder[rightNumber]))) {
				invalidUpdate = true
				badLeftIdx = i
				badRightIdx = j + i + 1
				break
			}
		}
		if invalidUpdate {
			break
		}
	}
	permutedList := make([]string, len(updates))
	copy(permutedList, updates)
	badLeftNumber := updates[badLeftIdx]
	badRightNumber := updates[badRightIdx]
	permutedList[badLeftIdx] = badRightNumber
	permutedList[badRightIdx] = badLeftNumber
	return permutedList, invalidUpdate
}

func main() {

	f, err := os.Open("input")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	orderMapDirect := make(map[int]map[int]bool)
	orderMapIndirect := make(map[int]map[int]bool)

	for scanner.Scan() {
		lineString := scanner.Text()
		if lineString == "" {
			break
		}
		splittedString := strings.Split(lineString, "|")
		leftNumber, _ := strconv.Atoi(splittedString[0])
		rightNumber, _ := strconv.Atoi(splittedString[1])
		_, ok := orderMapDirect[leftNumber] // check for existence
		if !ok {
			orderMapDirect[leftNumber] = map[int]bool{rightNumber: true}
		} else {
			orderMapDirect[leftNumber][rightNumber] = true
		}
		_, ok = orderMapIndirect[rightNumber] // check for existence
		if !ok {
			orderMapIndirect[rightNumber] = map[int]bool{leftNumber: true}
		} else {
			orderMapIndirect[rightNumber][leftNumber] = true
		}
	}
	sumMiddleUpdate := 0
	invalidUpdates := [][]string{}

	for scanner.Scan() {
		lineString := scanner.Text()
		updates := strings.Split(lineString, ",")
		invalidUpdate := false
		for i, updateLeft := range updates {
			leftNumber, _ := strconv.Atoi(updateLeft)
			for _, updateRight := range updates[i+1:] {
				rightNumber, _ := strconv.Atoi(updateRight)
				if !(checkInList(rightNumber, mapKeys(orderMapDirect[leftNumber])) && !checkInList(leftNumber, mapKeys(orderMapDirect[rightNumber]))) {
					invalidUpdate = true
					invalidUpdates = append(invalidUpdates, updates)
					break
				}
			}
			if invalidUpdate {
				break
			}
		}
		if !invalidUpdate {
			middleNumber, _ := strconv.Atoi(updates[len(updates)/2])
			sumMiddleUpdate += middleNumber
		}
	}
	fmt.Println(sumMiddleUpdate)

	permutedSumMiddle := 0
	for _, invalidUpdate := range invalidUpdates {
		isInvalid := true
		// fmt.Printf("Invalid permutation: %s\n", invalidUpdate)
		for {
			invalidUpdate, isInvalid = permuteInvalid(invalidUpdate, orderMapDirect)
			// fmt.Printf("Invalid permutation: %s\n", invalidUpdate, isInvalid)
			if !isInvalid {
				break
			}

		}
		fmt.Printf("Valid permutation: %s\n", invalidUpdate)
		middleNumber, _ := strconv.Atoi(invalidUpdate[len(invalidUpdate)/2])
		permutedSumMiddle += middleNumber
	}

	fmt.Println(permutedSumMiddle)
}
