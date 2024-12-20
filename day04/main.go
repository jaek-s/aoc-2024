package main

import (
	"log"
	"os"
	"strings"
)

func main() {
	log.Print("puzzle 1 example input:", solveFirstPuzzle("./example.txt"))
	log.Print("puzzle 1 real input:", solveFirstPuzzle("./input.txt"))
	log.Print("puzzle 2 example input:", solveSecondPuzzle("./example.txt"))
	log.Print("puzzle 2 real input:", solveSecondPuzzle("./input.txt"))
}

func checkErr(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func getPuzzleInput(filePath string) []string {
	rawFile, err := os.ReadFile(filePath)
	checkErr(err)

	return strings.Split(strings.TrimSpace(string(rawFile)), "\n")
}

func solveFirstPuzzle(filePath string) int {
	input := getPuzzleInput(filePath)
	var count int

	for y, line := range input {
		for x := range line {
			if checkNorth(x, y, input) {
				count++
			}
			if checkNorthEast(x, y, input) {
				count++
			}
			if checkEast(x, y, input) {
				count++
			}
			if checkSouthEast(x, y, input) {
				count++
			}
			if checkSouth(x, y, input) {
				count++
			}
			if checkSouthWest(x, y, input) {
				count++
			}
			if checkWest(x, y, input) {
				count++
			}
			if checkNorthWest(x, y, input) {
				count++
			}
		}
	}

	return count
}

func checkNorth(x int, y int, input []string) bool {
	return y-3 >= 0 &&
		input[y][x] == 'X' &&
		input[y-1][x] == 'M' &&
		input[y-2][x] == 'A' &&
		input[y-3][x] == 'S'
}

func checkNorthEast(x int, y int, input []string) bool {
	return x+3 < len(input[0]) &&
		y-3 >= 0 &&
		input[y][x] == 'X' &&
		input[y-1][x+1] == 'M' &&
		input[y-2][x+2] == 'A' &&
		input[y-3][x+3] == 'S'
}

func checkEast(x int, y int, input []string) bool {
	return x+3 < len(input[0]) &&
		input[y][x] == 'X' &&
		input[y][x+1] == 'M' &&
		input[y][x+2] == 'A' &&
		input[y][x+3] == 'S'
}

func checkSouthEast(x int, y int, input []string) bool {
	return x+3 < len(input[0]) &&
		y+3 < len(input) &&
		input[y][x] == 'X' &&
		input[y+1][x+1] == 'M' &&
		input[y+2][x+2] == 'A' &&
		input[y+3][x+3] == 'S'
}

func checkSouth(x int, y int, input []string) bool {
	return y+3 < len(input) &&
		input[y][x] == 'X' &&
		input[y+1][x] == 'M' &&
		input[y+2][x] == 'A' &&
		input[y+3][x] == 'S'
}

func checkSouthWest(x int, y int, input []string) bool {
	return x-3 >= 0 &&
		y+3 < len(input) &&
		input[y][x] == 'X' &&
		input[y+1][x-1] == 'M' &&
		input[y+2][x-2] == 'A' &&
		input[y+3][x-3] == 'S'
}

func checkWest(x int, y int, input []string) bool {
	return x-3 >= 0 &&
		input[y][x] == 'X' &&
		input[y][x-1] == 'M' &&
		input[y][x-2] == 'A' &&
		input[y][x-3] == 'S'
}

func checkNorthWest(x int, y int, input []string) bool {
	return x-3 >= 0 &&
		y-3 >= 0 &&
		input[y][x] == 'X' &&
		input[y-1][x-1] == 'M' &&
		input[y-2][x-2] == 'A' &&
		input[y-3][x-3] == 'S'
}

func solveSecondPuzzle(filePath string) int {
	input := getPuzzleInput(filePath)
	var count int

	for y, line := range input {
		if y == 0 || y == len(input)-1 {
			continue
		}

		for x, char := range line {
			if x == 0 || x == len(line)-1 || char != 'A' {
				continue
			}
			nwMatch := input[y-1][x-1] == 'M' && input[y+1][x+1] == 'S'
			neMatch := input[y-1][x+1] == 'M' && input[y+1][x-1] == 'S'
			swMatch := input[y+1][x-1] == 'M' && input[y-1][x+1] == 'S'
			seMatch := input[y+1][x+1] == 'M' && input[y-1][x-1] == 'S'

			matchExists := nwMatch && neMatch ||
				nwMatch && swMatch ||
				neMatch && seMatch ||
				swMatch && seMatch

			if matchExists {
				count++
			}
		}
	}

	return count
}
