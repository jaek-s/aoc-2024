package main

import (
	"log"
	"os"
	"slices"
	"strconv"
	"strings"
)

func main() {
	log.Println("Puzzle 1 example output:", solveFirstPuzzle("./example.txt"))
	log.Println("Puzzle 1 real output:", solveFirstPuzzle("./input.txt"))
	log.Println("Puzzle 2 example output:", solveSecondPuzzle("./example.txt"))
	log.Println("Puzzle 2 real output:", solveSecondPuzzle("./input.txt"))
}

func checkErr(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

type update []string

func (u update) getMiddlePage() int {
	middleIndex := (len(u) - 1) / 2
	strMiddle := u[middleIndex]
	middle, err := strconv.Atoi(strMiddle)
	checkErr(err)
	return middle
}

type puzzleInput struct {
	orderingRules [][]string
	updates       []update
}

func newPuzzleInput(filePath string) puzzleInput {
	file, err := os.ReadFile(filePath)
	checkErr(err)

	inputParts := strings.Split(strings.TrimSpace(string(file)), "\n\n")

	orderingRules := make([][]string, 0)
	for _, line := range strings.Split(inputParts[0], "\n") {
		orderingRules = append(orderingRules, strings.Split(line, "|"))
	}

	updates := make([]update, 0)
	for _, line := range strings.Split(inputParts[1], "\n") {
		updates = append(updates, strings.Split(line, ","))
	}

	return puzzleInput{
		orderingRules: orderingRules,
		updates:       updates,
	}
}

func (p puzzleInput) isValidUpdate(updateIndex int) bool {
	update := p.updates[updateIndex]

	for _, rule := range p.orderingRules {
		formerPage, latterPage := rule[0], rule[1]

		if !slices.Contains(update, formerPage) || !slices.Contains(update, latterPage) {
			continue
		}

		if slices.Index(update, formerPage) > slices.Index(update, latterPage) {
			return false
		}
	}

	return true
}

func (p puzzleInput) fixUpdate(updateIndex int) update {
	update := p.updates[updateIndex]

	for _, rule := range p.orderingRules {
		formerPage, latterPage := rule[0], rule[1]

		if !slices.Contains(update, formerPage) || !slices.Contains(update, latterPage) {
			continue
		}

		if slices.Index(update, formerPage) < slices.Index(update, latterPage) {
			continue
		}

		formerIndex := slices.Index(update, formerPage)
		latterIndex := slices.Index(update, latterPage)

		update[formerIndex], update[latterIndex] = update[latterIndex], update[formerIndex]
	}

	if p.isValidUpdate(updateIndex) {
		return update
	}

	return p.fixUpdate(updateIndex)
}

func solveFirstPuzzle(filePath string) int {
	input := newPuzzleInput(filePath)

	var output int
	for i, update := range input.updates {
		if !input.isValidUpdate(i) {
			continue
		}

		output += update.getMiddlePage()
	}

	return output
}

func solveSecondPuzzle(filePath string) int {
	input := newPuzzleInput(filePath)

	var output int
	for i := range input.updates {
		if input.isValidUpdate(i) {
			continue
		}

		fixedUpdate := input.fixUpdate(i)
		output += fixedUpdate.getMiddlePage()
	}

	return output
}
