package main

import (
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	log.Println("puzzle 1 example input:", solveFirstPuzzle("./example.txt"))
	log.Println("puzzle 1 real input:", solveFirstPuzzle("./input.txt"))
	log.Println("puzzle 2 example input:", solveSecondPuzzle("./example2.txt"))
	log.Println("puzzle 2 real input:", solveSecondPuzzle("./input.txt"))
}

func solveFirstPuzzle(filePath string) int {
	input, err := os.ReadFile(filePath)
	checkErr(err)

	return calculateSection(string(input))
}

func solveSecondPuzzle(filePath string) int {
	input, err := os.ReadFile(filePath)
	checkErr(err)

	var total int
	for i, section := range strings.Split(string(input), "don't()") {
		if i == 0 {
			total += calculateSection(section)
		}

		_, enabledSection, _ := strings.Cut(section, "do()")

		if len(enabledSection) < 1 {
			continue
		}

		total += calculateSection(enabledSection)

	}

	return total
}

func calculateSection(input string) int {
	re := regexp.MustCompile(`mul\((\d+),(\d+)\)`)

	matches := re.FindAllStringSubmatch(input, -1)

	var total int
	for _, match := range matches {
		firstMult, err := strconv.Atoi(match[1])
		checkErr(err)
		secondMult, err := strconv.Atoi(match[2])
		checkErr(err)

		total += firstMult * secondMult
	}

	return total
}

func checkErr(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
