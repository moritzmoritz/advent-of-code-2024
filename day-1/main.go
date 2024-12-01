package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"slices"
	"strconv"
	"strings"
	"time"
)

func main() {
	inputFile := "input-1.txt"

	file, err := os.Open(inputFile)

	if err != nil {
		panic(err)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	leftList := []int{}
	rightList := []int{}

	for scanner.Scan() {
		line := scanner.Text()
		lineNumbers := strings.Split(line, "   ")

		if len(lineNumbers) == 2 {
			leftNumber, err := strconv.Atoi(lineNumbers[0])

			if err != nil {
				panic(err)
			}

			rightNumber, err := strconv.Atoi(lineNumbers[1])

			if err != nil {
				panic(err)
			}

			leftList = append(leftList, leftNumber)
			rightList = append(rightList, rightNumber)
		}
	}

	if err := scanner.Err(); err != nil {
		panic(err)
	}

	slices.Sort(leftList)
	slices.Sort(rightList)

	if len(leftList) != len(rightList) {
		panic("List lengths are not equal")
	}

	// Part 1 & 2
	executionStartTime := time.Now()

	totalDistance := 0
	totalSimalarityScore := 0
	for i := 0; i < len(leftList); i++ {
		distance := math.Abs(float64(leftList[i] - rightList[i]))
		totalDistance += int(distance)

		similarityScore := 0

		for j := 0; j < len(rightList); j++ {
			if rightList[j] == leftList[i] {
				similarityScore += leftList[i]
			} else if rightList[j] > leftList[i] {
				break
			}
		}

		totalSimalarityScore += similarityScore
	}

	executionTime := time.Since(executionStartTime)

	fmt.Println("Part 1:", totalDistance)
	fmt.Println("Part 2:", totalSimalarityScore)

	fmt.Println("Execution time:", executionTime)
}
