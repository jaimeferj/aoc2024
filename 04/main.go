package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
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

func isLegalChar(input string) bool {
	legalChars := []string{"X", "M", "A", "S"}
	for _, x := range legalChars {
		if x == input {
			return true
		}
	}
	return false
}

func isXMAS(input []string) bool {
	jointString := strings.Join(input, "")
	return jointString == "XMAS" || jointString == "SAMX"
}

func isX_MAS(input []string) bool {
	jointString := strings.Join(input, "")
	return jointString == "MAS" || jointString == "SAM"
}

func main() {

	f, err := os.Open("input")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	board := make([][]string, 0)

	for scanner.Scan() {
		lineString := scanner.Text()
		board = append(board, strings.Split(lineString, ""))
	}
	width := len(board[0])
	height := len(board)
	totalXMAS := 0
	totalX_MAS := 0
	for i, line := range board {
		for j, x := range line {
			if !isLegalChar(x) {
				continue
			}
			if j+3 < width && isXMAS(line[j:j+4]) {
				totalXMAS++
			}
			if i+3 < height && isXMAS([]string{board[i][j], board[i+1][j], board[i+2][j], board[i+3][j]}) {
				totalXMAS++
			}
			if j+3 < width && i+3 < height && isXMAS([]string{board[i+1][j+1], board[i+2][j+2], board[i+3][j+3]}) {
				totalXMAS++
			}
			if j-3 >= 0 && i+3 < height && isXMAS([]string{board[i][j], board[i+1][j-1], board[i+2][j-2], board[i+3][j-3]}) {
				totalXMAS++
			}
			if j+1 < width && i+1 < height && j-1 >= 0 && i-1 >= 0 && isX_MAS([]string{board[i+1][j+1], board[i][j], board[i-1][j-1]}) &&
				isX_MAS([]string{board[i+1][j-1], board[i][j], board[i-1][j+1]}) {
				totalX_MAS++
			}

		}
	}
	fmt.Println(totalXMAS)
	fmt.Println(totalX_MAS)
}
