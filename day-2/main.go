package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
	"time"
)

const LEVEL_MIN_DISTANCE = 1
const LEVEL_MAX_DISTANCE = 3

const (
	INCREASE = 1
	DECREASE = -1
)

func main() {
	inputFile := "input.txt"

	file, err := os.Open(inputFile)

	if err != nil {
		panic(err)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	reports := [][]int{}

	for scanner.Scan() {
		line := scanner.Text()
		lineLevels := strings.Split(line, " ")

		report := []int{}

		for _, lineLevel := range lineLevels {
			level, err := strconv.Atoi(lineLevel)

			if err != nil {
				panic(err)
			}

			report = append(report, level)
		}

		reports = append(reports, report)
	}

	if err := scanner.Err(); err != nil {
		panic(err)
	}

	executionStartTime := time.Now()

	// Part 1
	validReportsCount := 0
	unsafeReports := [][]int{}

	for _, report := range reports {
		safeReport := reportIsSafe(report)

		if safeReport {
			validReportsCount++
		} else {
			unsafeReports = append(unsafeReports, report)
		}
	}

	// Part 2
	validReportsWithProblemDampenerCount := validReportsCount

	for _, report := range unsafeReports {
		for i := 0; i < len(report); i++ {
			newReport := newReportRemovingIndex(report, i)

			safeReport := reportIsSafe(newReport)

			if safeReport {
				validReportsWithProblemDampenerCount++
				break
			}
		}
	}

	executionTime := time.Since(executionStartTime)
	fmt.Println("Part 1:", validReportsCount)
	fmt.Println("Part 2:", validReportsWithProblemDampenerCount)

	fmt.Println("Execution time:", executionTime)
}

func validDistance(level1, level2 int) bool {
	if level1 == level2 {
		return false
	}

	return math.Abs(float64(level1-level2)) <= LEVEL_MAX_DISTANCE
}

func reportIsSafe(report []int) bool {
	for i := 0; i < len(report)-1; i++ {
		validDistance := validDistance(report[i], report[i+1])

		if !validDistance {
			return false
		}
	}

	currentOrder := INCREASE

	if report[1] < report[0] {
		currentOrder = DECREASE
	}

	for i := 1; i < len(report)-1; i++ {
		if report[i+1] < report[i] && currentOrder == INCREASE {
			return false
		} else if report[i+1] > report[i] && currentOrder == DECREASE {
			return false
		}
	}

	return true
}

func newReportRemovingIndex(report []int, index int) []int {
	newReport := []int{}

	for i, level := range report {
		if i == index {
			continue
		}

		newReport = append(newReport, level)
	}

	return newReport
}
