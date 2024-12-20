package main

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	log.Print("Safe report count for example.txt: ", getSafeReportCount("./example.txt"))
	log.Print("Safe report count for input.txt: ", getSafeReportCount("./input.txt"))
	log.Print("Safe report count with dampener for example.txt: ", getSafeReportCountWithDampener("./example.txt"))
	log.Print("Safe report count with dampener for input.txt: ", getSafeReportCountWithDampener("./input.txt"))
}

func checkErr(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func getSafeReportCount(filepath string) int {
	reports := getReports(filepath)
	safeReportCount := 0

	for _, report := range reports {
		if isReportSafe(report) {
			safeReportCount++
		}
	}

	return safeReportCount
}

func getSafeReportCountWithDampener(filepath string) int {
	reports := getReports(filepath)
	safeReportCount := 0

	for _, report := range reports {
		if isReportSafe(report) {
			safeReportCount++
			continue
		}

		for i := 0; i < len(report); i++ {
			removedReport := removeIndex(report, i)

			if isReportSafe(removedReport) {
				safeReportCount++
				break
			}
		}

	}

	return safeReportCount
}

func isReportSafe(report []int) bool {

	isIncreasing := report[0] < report[1]

	const maxDifference = 3

	for i, level := range report {
		if i == 0 {
			continue
		}

		previousLevel := report[i-1]

		if level == previousLevel {
			return false
		}

		if isIncreasing {
			if level < previousLevel {
				return false
			}

			if level-previousLevel > maxDifference {
				return false
			}
		}

		if !isIncreasing {
			if level > previousLevel {
				return false
			}

			if previousLevel-level > maxDifference {
				return false
			}
		}
	}

	return true
}

func getReports(filepath string) [][]int {
	file, err := os.Open(filepath)
	checkErr(err)
	defer file.Close()

	reportSlice := make([][]int, 0)
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()
		levels := strings.Split(line, " ")
		report := make([]int, 0)

		for _, level := range levels {
			convertedLevel, err := strconv.Atoi(level)
			checkErr(err)
			report = append(report, convertedLevel)
		}

		reportSlice = append(reportSlice, report)
	}

	return reportSlice
}

func removeIndex(s []int, index int) []int {
	ret := make([]int, 0)
	ret = append(ret, s[:index]...)
	return append(ret, s[index+1:]...)
}
