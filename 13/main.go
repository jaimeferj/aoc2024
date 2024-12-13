package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func mcd(x int, y int) int {
	if q := x % y; q != 0 {
		return mcd(y, q)
	} else {
		return y
	}
}

func main() {
	file, err := os.Open("input")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	listA := make([]Vector, 0)
	listB := make([]Vector, 0)
	listGoals := make([]Vector, 0)
	for scanner.Scan() {

		lineString := scanner.Text()
		if lineString == "" {
			continue
		}
		if strings.Contains(lineString, "Button A") {
			aXString := strings.Split(strings.Split(lineString, "X+")[1], ",")[0]
			aYString := strings.Split(lineString, "Y+")[1]
			aX, _ := strconv.Atoi(aXString)
			aY, _ := strconv.Atoi(aYString)
			listA = append(listA, Vector{aX, aY})
		} else if strings.Contains(lineString, "Button B") {
			bXString := strings.Split(strings.Split(lineString, "X+")[1], ",")[0]
			bYString := strings.Split(lineString, "Y+")[1]
			bX, _ := strconv.Atoi(bXString)
			bY, _ := strconv.Atoi(bYString)
			listB = append(listB, Vector{bX, bY})
		} else if strings.Contains(lineString, "Prize") {
			cXString := strings.Split(strings.Split(lineString, "X=")[1], ",")[0]
			cYString := strings.Split(lineString, "Y=")[1]
			cX, _ := strconv.Atoi(cXString)
			cY, _ := strconv.Atoi(cYString)
			listGoals = append(listGoals, Vector{cX, cY})
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	fmt.Println(listA, listB, listGoals)

	tokens := 0
	for i := range listA {
		vecA := listA[i]
		vecB := listB[i]
		vecC := listGoals[i]
		mcdX := mcd(vecA.x, vecB.x)
		mcdY := mcd(vecA.y, vecB.y)
		if vecC.x%mcdX != 0 || vecC.y%mcdY != 0 {
			continue
		}
		modA := vecA.Mod2()
		modB := vecB.Mod2()
		if vecA.Dot(vecB)*vecA.Dot(vecB) == modA*modB {
			if modA > modB && (modA/modB > 9 || (modA/modB == 9 && modA%modB == 0)) {
				tokens += 3 * vecC.x / vecA.x
			} else {
				tokens += vecC.x / vecB.x
			}

		} else {
			beta := (vecA.y*vecC.x - vecA.x*vecC.y) / (vecB.x*vecA.y - vecA.x*vecB.y)
			alpha := 0
			if vecA.x != 0 {
				alpha = (vecC.x - beta*vecB.x) / vecA.x
			} else {
				alpha = (vecC.y - beta*vecB.y) / vecA.y
			}
			assertCorrect := vecA.times(alpha).Add(vecB.times(beta)).Equals(vecC)
			if !assertCorrect {
				continue
			}
			tokens += 3*alpha + beta

		}
	}
	fmt.Println(tokens)

	tokens2 := 0
	for i := range listA {
		vecA := listA[i]
		vecB := listB[i]
		vecC := listGoals[i]
		vecC = Vector{vecC.x + 10000000000000, vecC.y + 10000000000000}
		mcdX := mcd(vecA.x, vecB.x)
		mcdY := mcd(vecA.y, vecB.y)
		if vecC.x%mcdX != 0 || vecC.y%mcdY != 0 {
			continue
		}
		modA := vecA.Mod2()
		modB := vecB.Mod2()
		if vecA.Dot(vecB)*vecA.Dot(vecB) == modA*modB {
			if modA > modB && (modA/modB > 9 || (modA/modB == 9 && modA%modB == 0)) {
				tokens2 += 3 * vecC.x / vecA.x
			} else {
				tokens2 += vecC.x / vecB.x
			}

		} else {
			beta := (vecA.y*vecC.x - vecA.x*vecC.y) / (vecB.x*vecA.y - vecA.x*vecB.y)
			alpha := 0
			if vecA.x != 0 {
				alpha = (vecC.x - beta*vecB.x) / vecA.x
			} else {
				alpha = (vecC.y - beta*vecB.y) / vecA.y
			}
			assertCorrect := vecA.times(alpha).Add(vecB.times(beta)).Equals(vecC)
			if !assertCorrect {
				continue
			}
			tokens2 += 3*alpha + beta

		}
	}
	fmt.Println(tokens2)

}
