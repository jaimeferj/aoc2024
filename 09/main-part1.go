package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {

	f, err := os.Open("test")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)

	// fileSizes := []int{}
	// emptySizes := []int{}
	numberList := []int{}

	for scanner.Scan() {
		lineString := scanner.Text()
		for _, x := range lineString {
			number := int(x - '0')
			numberList = append(numberList, number)
			// if i%2 == 0 {
			//
			// 	fileSizes = append(fileSizes, number)
			// } else {
			// 	emptySizes = append(emptySizes, number)
			// }
		}
	}
	numberListCopy := make([]int, len(numberList))
	copy(numberListCopy, numberList)

	i := 0
	j := 0
	k := (len(numberList)+1)/2 - 1

	checksum := 0
	for {
		largestFileIdx := k * 2
		if i+1 >= len(numberList) {
			break
		}
		if numberList[i] == 0 {
			i++
			continue
		} else if i%2 == 0 {
			checksum += j * i / 2
			// fmt.Println(1, j, j*i/2, checksum)
		} else if numberList[largestFileIdx] > 0 {
			checksum += j * k
			// fmt.Println(2, j, j*k, checksum, k, numberList[largestFileIdx])
			numberList[largestFileIdx]--
			if numberList[largestFileIdx] <= 0 {
				k--
			}
		} else {
			break
		}
		numberList[i]--
		j++
	}
	fmt.Println(checksum)
	// Part 2

	k = (len(numberList)+1)/2 - 1
	emptySpacesMax := (len(numberList) - 1) / 2

	checksum = 0
	for {
		largestFileIdx := k * 2
		if numberListCopy[largestFileIdx] > 0 {
			fileSpace := numberListCopy[largestFileIdx]
			idx := 0
			foundEmptySpace := false
			for i := range emptySpacesMax - 1 {
				if i*2+1 > largestFileIdx {
					checksum += fileSpace * (largestFileIdx + (largestFileIdx-1)/2)
					break
				}
				//Sumamos los espacios rellenos
				idx += numberListCopy[i*2]
				emptyIdx := i*2 + 1
				emptySpace := numberListCopy[emptyIdx]
				if emptySpace <= 0 {
					continue
				} else if emptySpace >= fileSpace {
					numberListCopy[largestFileIdx] = 0
					numberListCopy[emptyIdx] -= fileSpace
					checksum += idx * fileSpace
					foundEmptySpace = true
					break
				} else {
					//Sumamos los espacios vacios no rellenados
					idx += emptySpace
				}
			}
			if !foundEmptySpace {
				checksum += idx * fileSpace
			}
		}
		k--
		if k < 0 {
			break
		}
	}
	fmt.Println(checksum)
}
