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
	for _, line := range strings.Split(input, "\n") {
		var a, b, c int
		_, err := fmt.Sscanf(line, "%dx%dx%d", &a, &b, &c)
		if err != nil {
			panic(err)
		}
		firstSide := a * b * 2
		secondtSide := a * c * 2
		thirdtSide := c * b * 2

		additional := MinInt(a*b, a*c, b*c)
		result += firstSide + secondtSide + thirdtSide + additional
	}
	fmt.Printf("Part 1: %d \n", result)
}

func part2(input string) {
	result := 0
	for _, line := range strings.Split(input, "\n") {
		var a, b, c int
		_, err := fmt.Sscanf(line, "%dx%dx%d", &a, &b, &c)
		if err != nil {
			panic(err)
		}
		bow := a * b * c
		feet := MinInt((2 * (a + b)), (2 * (a + c)), (2 * (c + b)))
		result += bow + feet
	}

	fmt.Printf("Part 2: %d \n", result)
}

func MinInt(numbers ...int) int {
	if numbers[0] <= numbers[1] && numbers[0] <= numbers[2] {
		return numbers[0]
	} else if numbers[1] <= numbers[0] && numbers[1] <= numbers[2] {
		return numbers[1]
	} else {
		return numbers[2]
	}
}
