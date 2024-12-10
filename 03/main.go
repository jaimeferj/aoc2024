package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
)

func findClosest(input int, findList []int) int {
	closest := -1
	for _, x := range findList {
		if x > input {
			break
		}
		closest = x
	}
	return closest
}

func main() {

	f, err := os.Open("input")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	realMulPattern := regexp.MustCompile(`mul\((\d+),(\d+)\)`)
	doPattern := regexp.MustCompile(`do\(\)`)
	dontPattern := regexp.MustCompile(`don't\(\)`)
	defaultIsDo := true
	sumOfMulOriginal := 0
	sumOfMul := 0
	for scanner.Scan() {
		lineString := scanner.Text()
		mulMatches := realMulPattern.FindAllStringSubmatchIndex(lineString, -1)
		doMatches := doPattern.FindAllStringIndex(lineString, -1)
		doIndexes := make([]int, len(doMatches))
		for i, doMatch := range doMatches {
			doIndexes[i] = doMatch[0]
		}
		dontMatches := dontPattern.FindAllStringIndex(lineString, -1)
		dontIndexes := make([]int, len(dontMatches))
		for i, dontMatch := range dontMatches {
			dontIndexes[i] = dontMatch[0]
		}
		// fmt.Println(lineString)

		closestDo := -1
		closestDont := -1

		for _, match := range mulMatches {
			closestDo = findClosest(match[0], doIndexes)
			closestDont = findClosest(match[0], dontIndexes)
			// fmt.Println(lineString[match[0]:match[1]])
			// fmt.Printf("GroupIdx: %s, Do: %s, Dont: %s\n", match[0], closestDo, closestDont)
			leftString := lineString[match[2]:match[3]]
			rightString := lineString[match[4]:match[5]]
			leftNumber, _ := strconv.Atoi(leftString)
			rightNumber, _ := strconv.Atoi(rightString)
			// fmt.Printf("`%s`, `%s`\n", leftNumber, rightNumber)
			sumOfMulOriginal += leftNumber * rightNumber

			if (closestDont > closestDo) || (!defaultIsDo && closestDo == -1 && closestDont == -1) {
				continue
			}
			sumOfMul += leftNumber * rightNumber
			// fmt.Printf("%s, %s\n", sumOfMulOriginal, sumOfMul)
		}
		defaultIsDo = closestDo >= closestDont
	}

	fmt.Println(sumOfMulOriginal)
	fmt.Println(sumOfMul)
}
