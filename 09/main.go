package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	file, err := os.Open("input")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	numberList := []int{}
	for scanner.Scan() {
		for _, x := range scanner.Text() {
			number := int(x - '0')
			numberList = append(numberList, number)
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	emptySpacePositions := make([]int, (len(numberList)-1)/2)
	initialFilePositions := make([]int, (len(numberList)+1)/2)
	representation := make([]string, 0)
	idx := 0
	j := 0
	for i := range numberList {
		if i%2 != 0 {
			emptySpacePositions[j] = idx
			for range numberList[i] {
				representation = append(representation, ".")
			}
			j++
		} else {
			initialFilePositions[i/2] = idx
			for range numberList[i] {
				representation = append(representation, strconv.Itoa(i/2))
			}
		}
		idx += numberList[i]
	}

	// fmt.Println(numberList)
	// fmt.Println(initialFilePositions, emptySpacePositions)
	// fmt.Println(representation)
	filesCount := (len(numberList) + 1) / 2
	checksum := 0
	fmt.Println(representation)
	for rawFileIdx := range filesCount {
		fileId := filesCount - 1 - rawFileIdx
		fileIdx := fileId * 2
		fileSize := numberList[fileIdx]
		// fmt.Println(fileId, fileSize)
		foundSpace := false
		for rawEmptyIdx := range filesCount - 1 {
			emptyIdx := rawEmptyIdx*2 + 1
			emptySpace := numberList[emptyIdx]
			if fileSize > emptySpace || emptyIdx > fileIdx {
				continue
			} else {
				emptySpacePosition := emptySpacePositions[rawEmptyIdx]
				fileChecksum := fileId * (fileSize*emptySpacePosition + (fileSize*(fileSize-1))/2)
				if fileChecksum < 0 {
					panic("BOF")
				}
				checksum += fileChecksum
				// fmt.Println(fileId, emptySpacePosition, checksum)
				emptySpacePositions[rawEmptyIdx] += fileSize
				numberList[emptyIdx] -= fileSize
				foundSpace = true

				filePosition := initialFilePositions[fileId]
				for k := range fileSize {
					representation[filePosition+k] = "."
					representation[emptySpacePosition+k] = strconv.Itoa(fileId)
				}

				break
			}
		}
		if !foundSpace {
			filePosition := initialFilePositions[fileId]
			fileChecksum := fileId * (fileSize*filePosition + (fileSize*(fileSize-1))/2)
			checksum += fileChecksum
			if fileChecksum < 0 {
				panic("BOF")
			}
			// fmt.Println("No space found for", fileId, "at", filePosition, checksum)
		} else {
			// fmt.Println(representation)
		}

	}

	fmt.Println(checksum)
	checksum2 := 0
	for i, x := range representation {
		if x == "." {
			continue
		} else {
			fileId, _ := strconv.Atoi(x)
			checksum2 += i * fileId
		}
	}
	fmt.Println(checksum2)
}
