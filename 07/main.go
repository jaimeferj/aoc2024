package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func consumeList(acc int, data []int, start int, stopConditionCallback func(int, bool) bool, operations []string) {
	if stopConditionCallback(acc, start >= len(data)) {
		return
	}
	value := data[start]
	consumeList(acc+value, data, start+1, stopConditionCallback, append(operations, "+"))
	consumeList(acc*value, data, start+1, stopConditionCallback, append(operations, "*"))
}

func consumeListPart2(acc int, data []int, start int, stopConditionCallback func(int, bool) bool, operations []string) {
	if stopConditionCallback(acc, start >= len(data)) {
		return
	}
	value := data[start]
	concatProduct := 10
	for {
		if concatProduct <= value {
			concatProduct *= 10
		} else {
			break
		}
	}
	concatNumber := acc*concatProduct + value
	consumeListPart2(acc+value, data, start+1, stopConditionCallback, append(operations, "+"))
	consumeListPart2(acc*value, data, start+1, stopConditionCallback, append(operations, "*"))
	consumeListPart2(concatNumber, data, start+1, stopConditionCallback, append(operations, "|"))
}

func printEquation(lhs int, rhs []int) {
	fmt.Printf("%d:", lhs)
	for _, value := range rhs {
		fmt.Printf(" %d", value)
	}
	fmt.Printf("\n")
}

func main() {

	f, err := os.Open("input")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	stopConditionFactory := func(lhs int, lhsAccumulator *int, equationClaimed *bool) func(int, bool) bool {
		return func(acc int, isCompleteEquation bool) bool {
			if acc <= 0 {
				fmt.Println("YELL")
			}
			if acc == lhs && isCompleteEquation {
				if !*equationClaimed {
					*lhsAccumulator += lhs
					*equationClaimed = true
				}
				return true
			} else if isCompleteEquation {
				return true
			}
			return acc > lhs
		}
	}

	accumulatedLhs := 0
	accumulatedLhsPart2 := 0

	for scanner.Scan() {
		lineString := scanner.Text()
		splittedString := strings.Split(lineString, ":")
		lhs, _ := strconv.Atoi(splittedString[0])

		splittedRightHand := strings.Split(strings.Trim(splittedString[1], " "), " ")
		rhs := make([]int, len(splittedRightHand))
		for i, value := range splittedRightHand {
			number, _ := strconv.Atoi(value)
			rhs[i] = number
		}
		equationClaimed := false
		equationClaimedPart2 := false
		stopCondition1 := stopConditionFactory(lhs, &accumulatedLhs, &equationClaimed)
		stopCondition2 := stopConditionFactory(lhs, &accumulatedLhsPart2, &equationClaimedPart2)
		consumeList(rhs[0], rhs, 1, stopCondition1, []string{})
		consumeListPart2(rhs[0], rhs, 1, stopCondition2, []string{})
	}

	fmt.Println(accumulatedLhs)
	fmt.Println(accumulatedLhsPart2)
}
