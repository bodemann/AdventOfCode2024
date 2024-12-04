package main

import (
	"fmt"
	"os"
	"slices"
	"sort"
	"strconv"
	"strings"
)

func main() {
	PartOne()
	PartTwo()
}

func PartOne() {
	rawData, err := os.ReadFile("data.txt")
	if err != nil {
		panic(err)
	}
	lines := strings.Split(string(rawData), "\n")
	var safe int
	for _, line := range lines {
		data := convertStringOfNumbersToIntegers(line, " ")
		// Make all "reports" (lines in the file) increasing
		if data[0] > data[1] {
			slices.Reverse(data)
		}
		var differences []int
		for idx := 0; idx < len(data)-1; idx++ {
			differences = append(differences, data[idx+1]-data[idx])
		}
		// Sorting the differences makes it easy to check for too big or too small differences, since we only have to find one error.
		sort.Ints(differences)
		if differences[0] >= 1 && differences[len(differences)-1] < 4 {
			safe++
		}
	}
	fmt.Println("Safe reports:", safe)
}

func PartTwo() {
	rawData, err := os.ReadFile("data.txt")
	if err != nil {
		panic(err)
	}
	lines := strings.Split(string(rawData), "\n")
	var safe int
	for _, line := range lines {
		data := convertStringOfNumbersToIntegers(line, " ")
		var valueChange []int
		var signChange []string
		var countErrors int
		for idx := 0; idx < len(data)-1; idx++ {
			valueChange = append(valueChange, absInt(data[idx+1]-data[idx]))
			if idx > 1 && data[idx-1] < data[idx] {
				signChange = append(signChange, "")
			}

		}
		// Sorting the differences makes it easy to check for too big or too small differences, since we only have to find one error.
		sort.Ints(differences)
		for idx, difference := range differences {
			if difference < 1 || difference > 3 {
				countErrors++
			} else {
				if (data[idx] < 0 && data[idx+1] > 0) || (data[idx] > 0 && data[idx+1] < 0) {
					countErrors++
				}
			}
		}
		if countErrors <= 1 {
			safe++
			//fmt.Println(data)
		}
	}
	fmt.Println("Safe reports part two:", safe)
}

func convertStringOfNumbersToIntegers(line string, sep string) []int {
	data := strings.Split(line, sep)
	var result []int
	for _, str := range data {
		number, err := strconv.Atoi(str)
		if err != nil {
			panic(err)
		}
		result = append(result, number)
	}
	return result
}

func absInt(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
