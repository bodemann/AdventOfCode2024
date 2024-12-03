package main

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	// partOne and partTwo are independent.
	partOne()
	partTwo()
}

func partOne() {
	rawData, err := os.ReadFile("data.txt")
	if err != nil {
		panic(err)
	}
	lines := strings.Split(string(rawData), "\n")

	rMul, _ := regexp.Compile(`mul\(\d+,\d+\)`)
	rNum, _ := regexp.Compile(`\d+`)
	var sum int
	for _, line := range lines {
		mulsStr := (rMul.FindAllString(line, -1))
		for _, mul := range mulsStr {
			numStrs := (rNum.FindAllString(mul, -1))
			numInts := convertStringsToIntegers(numStrs)
			sum += numInts[0] * numInts[1]
		}
	}
	fmt.Println(sum)
}

func partTwo() {
	rawData, err := os.ReadFile("data.txt")
	if err != nil {
		panic(err)
	}

	instructionsRegEx, _ := regexp.Compile(`mul\(\d+,\d+\)|do\(\)|don't\(\)`)
	NumRegEx, _ := regexp.Compile(`\d+`)
	instructions := (instructionsRegEx.FindAllString(string(rawData), -1))

	var sum int
	instructionEnabled := true
	for _, instruction := range instructions {
		switch instruction {
		case "do()":
			instructionEnabled = true
		case "don't()":
			instructionEnabled = false
		default:
			if instructionEnabled {
				numStrs := (NumRegEx.FindAllString(instruction, -1))
				numInts := convertStringsToIntegers(numStrs)
				sum += numInts[0] * numInts[1]
			}

		}
	}
	fmt.Println(sum)
}

func convertStringsToIntegers(strs []string) []int {
	var result []int
	for _, str := range strs {
		number, err := strconv.Atoi(str)
		if err != nil {
			panic(err)
		}
		result = append(result, number)
	}
	return result
}
