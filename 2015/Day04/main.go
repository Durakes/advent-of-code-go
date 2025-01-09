package main

import (
	"crypto/md5"
	_ "embed"
	"encoding/hex"
	"fmt"
	"strconv"
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
	number := 0
	for {
		result := input + strconv.Itoa(number)
		hash := md5.Sum([]byte(result))
		finalresult := hex.EncodeToString(hash[:])
		firstFive := finalresult[:5]
		if firstFive == "00000" {
			break
		}
		number++
	}

	fmt.Printf("Output 1: %d \n", number)
}

func part2(input string) {
	number := 0
	for {
		result := input + strconv.Itoa(number)
		hash := md5.Sum([]byte(result))
		finalresult := hex.EncodeToString(hash[:])
		firstFive := finalresult[:6]
		if firstFive == "000000" {
			break
		}
		number++
	}

	fmt.Printf("Output 1: %d \n", number)
}
