package main

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func main() {
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
