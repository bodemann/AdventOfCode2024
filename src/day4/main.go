package main

import (
	"fmt"
	"os"
	"strings"
)

type Pos struct {
	x, y int
}

func main() {
	partOne()
	//partTwo()
}

func partOne() {
	rawData, err := os.ReadFile("test.txt")
	if err != nil {
		panic(err)
	}
	rows := strings.Split(string(rawData), "\n")

	var XMASCounter int
	for y, row := range rows {
		for x, char := range row {
			if string(char) == "X" {
				ms := findMs(x, y, rows)
				XMASCounter += findASs(x, y, ms, rows)
			}
		}
	}
	fmt.Println(XMASCounter)
}

func findMs(x, y int, rows []string) []Pos {
	var mPos []Pos
	for yM := -1; yM <= 1; yM++ {
		for xM := -1; xM <= 1; xM++ {
			if getChar(x+xM, y+yM, rows) == "M" {
				mPos = append(mPos, Pos{x + xM, y + yM})
			}
		}
	}
	return mPos
}

func findASs(x, y int, ms []Pos, rows []string) int {
	var ASCounter int
	for _, pos := range ms {
		switch {
		case x == pos.x && y == pos.y-1:
			ASCounter += isAS(x, y-2, x, y-3, rows)
		case x == pos.x+1 && y == pos.y-1:
			ASCounter += isAS(x+2, y-2, x+3, y-3, rows)
		case x == pos.x+1 && y == pos.y:
			ASCounter += isAS(x+2, y, x+3, y, rows)
		case x == pos.x+1 && y == pos.y+1:
			ASCounter += isAS(x+2, y+2, x+3, y+3, rows)
		case x == pos.x && y == pos.y+1:
			ASCounter += isAS(x, y+2, x, y+3, rows)
		case x == pos.x-1 && y == pos.y+1:
			ASCounter += isAS(x-2, y-2, x-3, y-3, rows)
		case x == pos.x-1 && y == pos.y:
			ASCounter += isAS(x-2, y, x-3, y, rows)
		case x == pos.x-1 && y == pos.y-1:
			ASCounter += isAS(x-2, y-2, x-3, y-3, rows)
		}
	}
	if ASCounter > 0 {
		fmt.Println("ASCounter", x, y, ms)
	}
	return ASCounter
}

func isAS(x1, y1, x2, y2 int, rows []string) int {
	if getChar(x1, y1, rows) == "A" && getChar(x2, y2, rows) == "S" {
		return 1
	}
	return 0
}

func getChar(x, y int, rows []string) string {
	switch {
	case x < 0:
		return ""
	case y < 0:
		return ""
	case y >= len(rows):
		return ""
	case x >= len(rows[y]):
		return ""
	}
	return string(rows[y][x])
}
