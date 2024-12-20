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
	log.Println("Input distance:", getDistanceFromFile("./input.txt"))
	log.Println("Example similarity score:", calculateSimilarityScore("./example.txt"))
	log.Println("Input similarity score:", calculateSimilarityScore("./input.txt"))
}

func getDistanceFromFile(filename string) int {
	left, right := getLeftAndRightLists(filename)
	sort.Slice(left, func(i, j int) bool {
		return left[i] < left[j]
	})
	sort.Slice(right, func(i, j int) bool {
		return right[i] < right[j]
	})

	var tally int

	for i := 0; i < len(left); i++ {
		if left[i] > right[i] {
			tally += left[i] - right[i]
		} else {
			tally += right[i] - left[i]
		}
	}

	return tally
}

func calculateSimilarityScore(filename string) int {
	left, right := getLeftAndRightLists(filename)

	var total int
	leftOccurrences := getOccurrenceMap(left)
	rightOccurrences := getOccurrenceMap(right)

	for key, leftValue := range leftOccurrences {
		if rightValue, ok := rightOccurrences[key]; ok {
			total += key * leftValue * rightValue
		}
	}

	return total
}

func checkErr(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func getLeftAndRightLists(filename string) ([]int, []int) {
	file, err := os.Open(filename)
	checkErr(err)
	defer file.Close()

	scanner := bufio.NewScanner(file)

	left := make([]int, 0)
	right := make([]int, 0)

	for scanner.Scan() {
		splitLine := strings.Split(scanner.Text(), "   ")

		leftInt, err := strconv.Atoi(splitLine[0])
		checkErr(err)
		rightInt, err := strconv.Atoi(splitLine[1])
		checkErr(err)

		left = append(left, leftInt)
		right = append(right, rightInt)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return left, right
}

func getOccurrenceMap(slice []int) map[int]int {
	occurrenceMap := make(map[int]int)

	for _, value := range slice {
		if _, ok := occurrenceMap[value]; ok {
			occurrenceMap[value]++
		} else {
			occurrenceMap[value] = 1
		}
	}

	return occurrenceMap
}
