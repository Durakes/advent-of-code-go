package main

import (
	_ "embed"
	"fmt"
	"strings"
)

//go:embed input.txt
var input string

func main() {
	input = strings.TrimRight(input, "\n")
	part1(input)
	part2(input)

}

func part1(input string) {
	result := 0
	for _, letter := range input {
		if letter == '(' {
			result++
		} else {
			result--
		}
	}
	fmt.Printf("Output 1: %d \n", result)
}

func part2(input string) {
	result := 0
	firstFloor := 0
	for i, letter := range input {
		if letter == '(' {
			result++
		} else {
			result--
		}

		if result == -1 {
			firstFloor = i + 1
			break
		}
	}
	fmt.Printf("Output 2: %d", firstFloor)
}
