package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func isRecordSafe(splitLine []string) (bool, int) {
	lastNumber := 0
	ascending := false
	valid := true
	invalidIdx := 0
	for i, splitString := range splitLine {
		number, _ := strconv.Atoi(splitString)
		if i == 1 {
			ascending = number > lastNumber
		}
		if i != 0 {
			if (number == lastNumber) || (ascending != (number > lastNumber)) || (ascending && ((number - lastNumber) > 3)) || (!ascending && ((lastNumber - number) > 3)) {
				valid = false
				invalidIdx = i
				break
			}
		}
		lastNumber = number
	}
	return valid, invalidIdx
}

func main() {

	f, err := os.Open("input")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)

	validRecordsV1 := 0
	validRecordsV2 := 0
	for scanner.Scan() {
		splitLine := strings.Split(scanner.Text(), " ")
		if isValidRecord, _ := isRecordSafe(splitLine); isValidRecord {
			validRecordsV1++
			validRecordsV2++
		} else {
			newSplitLine := make([]string, len(splitLine))
			for i := range len(splitLine) {
				copy(newSplitLine, splitLine)
				isValidRecord, _ := isRecordSafe(append(newSplitLine[:i], newSplitLine[i+1:]...))
				if isValidRecord {
					validRecordsV2++
					break
				}
			}
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	fmt.Println("Valid Records V1:", validRecordsV1)
	fmt.Println("Valid Records V2:", validRecordsV2)
}
