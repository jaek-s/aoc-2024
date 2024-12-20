package main

import (
	"bufio"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	log.Println("Example distance:", getDistanceFromFile("./example.txt"))
	log.Println("Input1 distance:", getDistanceFromFile("./input1.txt"))
}

func getDistanceFromFile(filename string) int {
	file, err := os.Open(filename)
	checkErr(err)
	defer file.Close()

	scanner := bufio.NewScanner(file)

	left := make([]string, 0, 6)
	right := make([]string, 0, 6)

	for scanner.Scan() {
		splitLine := strings.Split(scanner.Text(), "   ")
		left = append(left, splitLine[0])
		right = append(right, splitLine[1])
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	file.Close()

	sort.Slice(left, func(i, j int) bool {
		return left[i] < left[j]
	})
	sort.Slice(right, func(i, j int) bool {
		return right[i] < right[j]
	})

	var tally int

	for i := 0; i < len(left); i++ {
		leftInt, err := strconv.Atoi(left[i])
		checkErr(err)

		rightInt, err := strconv.Atoi(right[i])
		checkErr(err)

		if leftInt > rightInt {
			tally += leftInt - rightInt
		} else {
			tally += rightInt - leftInt
		}
	}

	return tally
}

func checkErr(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
