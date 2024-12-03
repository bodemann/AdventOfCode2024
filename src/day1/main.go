package main

import (
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

func main() {
	data, err := os.ReadFile("/home/jb/IT_Projects/GoProjects/AdventOfCode2024/src/day1/data.txt")
	if err != nil {
		panic(err)
	}
	lines := strings.Split(string(data), "\n")
	var leftColumn []int
	var rightColumn []int
	for _, line := range lines {
		row := strings.Split(line, " ")
		left, err := strconv.Atoi(row[0])
		if err != nil {
			panic(err)
		}
		leftColumn = append(leftColumn, left)

		right, err := strconv.Atoi(row[3])
		if err != nil {
			panic(err)
		}
		rightColumn = append(rightColumn, right)
	}
	var result int
	slices.Sort(leftColumn)
	slices.Sort(rightColumn)
	for idx, _ := range leftColumn {
		result += absDiffInt(leftColumn[idx], rightColumn[idx])
	}
	fmt.Println(result)

	//Part 2
	var result2 int
	for idx, _ := range leftColumn {
		var findCounter int
		for {
			pos, found := slices.BinarySearch(rightColumn, leftColumn[idx])
			if !found {
				break
			}
			findCounter++
			rightColumn = append(rightColumn[:pos], rightColumn[pos+1:]...)
		}
		result2 += findCounter * leftColumn[idx]
	}
	fmt.Println(result2)
}

func absDiffInt(x, y int) int {
	if x < y {
		return y - x
	}
	return x - y
}
