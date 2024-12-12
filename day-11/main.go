package main

import (
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
	"time"
)

const STONE_MULTIPLIER = 2024

var blinkCache = make(map[string]int)

func main() {
	inputFile := "input.txt"

	file, err := os.Open(inputFile)

	if err != nil {
		panic(err)
	}

	defer file.Close()

	stones := readStones(file)

	executionStartTime := time.Now()

	// Part 1
	stonesCount := blink(stones, 25)

	// Part 2
	secondStonesCount := blink(stones, 75)

	executionTime := time.Since(executionStartTime)
	fmt.Println("Part 1:", stonesCount)
	fmt.Println("Part 2:", secondStonesCount)
	fmt.Println("Execution time:", executionTime)
}

func readStones(file *os.File) []int {
	content, err := io.ReadAll(file)

	if err != nil {
		panic(err)
	}

	line := string(content)

	splittedLine := strings.Split(line, " ")

	stones := []int{}

	for _, stone := range splittedLine {
		value, err := strconv.Atoi(stone)

		if err != nil {
			panic(err)
		}

		stones = append(stones, value)
	}

	return stones
}

func blink(stones []int, times int) int {
	var total int
	for _, stone := range stones {
		total += stonesAfterBlinks(stone, times)
	}
	return total
}

func stonesAfterBlinks(stone int, times int) int {
	if times == 0 {
		return 1
	}

	// Try to get cached value
	cacheKey := fmt.Sprintf("%d-%d", stone, times)

	if cachedValue, ok := blinkCache[cacheKey]; ok {
		return cachedValue
	}

	sum := 0
	if stone == 0 {
		sum = stonesAfterBlinks(1, times-1)
	} else if stringValue := strconv.Itoa(stone); len(stringValue)%2 == 0 {
		half := len(stringValue) / 2

		firstHalfNumber, err := strconv.Atoi(stringValue[:half])

		if err != nil {
			panic(err)
		}

		secondHalfNumber, err := strconv.Atoi(stringValue[half:])

		if err != nil {
			panic(err)
		}

		sum = stonesAfterBlinks(firstHalfNumber, times-1) + stonesAfterBlinks(secondHalfNumber, times-1)
	} else {
		sum = stonesAfterBlinks(stone*STONE_MULTIPLIER, times-1)
	}

	blinkCache[cacheKey] = sum
	return sum
}
