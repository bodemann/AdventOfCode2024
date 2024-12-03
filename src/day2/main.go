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
	rawData, err := os.ReadFile("/Users/jb/tmp/GoTmp/aoc2024day2/src/data.txt")
	if err != nil {
		panic(err)
	}
	lines := strings.Split(string(rawData), "\n")
	var safePartOne int
	var safePartTwo int
	var safeWithDampener int
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
		// Sorting the differences makes it easy to check for to big or small differences, since we only have to find one error.
		sort.Ints(differences)
		// This is the code for part 1
		if differences[0] >= 1 && differences[len(differences)-1] < 4 {
			safePartOne++
		}

		// This is the code for part 2
		var countErrors int
		for _, difference := range differences {
			if difference < 1 || difference > 3 {
				countErrors++
			}
		}
		switch countErrors {
		case 0:
			safePartTwo++
		case 1:
			safeWithDampener++
		case 2:
			fmt.Println(differences)
		}
		//fmt.Println(differences, countErrors, safePartTwo, safeWithDampener)
	}
	fmt.Println("Safe Part One:", safePartOne, "   Safe Part Two: ", safePartTwo, "   SafeWithDampener: ", safePartTwo+safeWithDampener)
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
