package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

type Operation = int

const (
	Add Operation = iota
	Multiply
	Combine
)

type Equation struct {
	Value int
	Colon []int
}

func main() {
	inputFile := "input.txt"

	file, err := os.Open(inputFile)

	if err != nil {
		panic(err)
	}

	defer file.Close()

	equations := extractEquations(file)

	executionStartTime := time.Now()

	// Part 1
	partOneTotal := findTotalCalibrationResult(equations, 2)

	// Part 2
	partTwoTotal := findTotalCalibrationResult(equations, 3)

	executionTime := time.Since(executionStartTime)

	fmt.Println("Part 1:", partOneTotal)
	fmt.Println("Part 2:", partTwoTotal)
	fmt.Println("Execution time:", executionTime)
}

func extractEquations(input *os.File) []Equation {
	equations := []Equation{}

	scanner := bufio.NewScanner(input)

	for scanner.Scan() {
		line := scanner.Text()

		splittedLine := strings.Split(line, ":")

		value, err := strconv.Atoi(splittedLine[0])

		if err != nil {
			panic(err)
		}

		colon := splittedLine[1]
		splittedColon := strings.Split(colon, " ")

		equation := Equation{
			Value: value,
			Colon: []int{},
		}

		for _, value := range splittedColon {
			if value == "" {
				continue
			}

			colonValue, err := strconv.Atoi(value)

			if err != nil {
				panic(err)
			}

			equation.Colon = append(equation.Colon, colonValue)
		}

		equations = append(equations, equation)
	}

	if err := scanner.Err(); err != nil {
		panic(err)
	}

	return equations
}

func findTotalCalibrationResult(equations []Equation, operations int) int {
	total := 0

	for i := 0; i < len(equations); i++ {
		equation := equations[i]

		numOps := len(equation.Colon) - 1

		for i := 0; i < (operations << (2 * numOps)); i++ {
			// Generate the operations for this combination
			ops := make([]int, numOps)
			for j := 0; j < numOps; j++ {
				// Determine the operation using modulo
				ops[j] = (i / (operations << (2 * j))) % operations
			}

			// Apply the operations
			result := applyOperations(equation.Colon, ops)

			if result == equation.Value {
				total += result
				break
			}
		}
	}

	return total
}

func applyOperations(arr []int, ops []Operation) int {
	// Start with the first element
	result := arr[0]

	for i := 1; i < len(arr); i++ {
		switch ops[i-1] {
		case Add:
			result += arr[i]
		case Multiply:
			result *= arr[i]
		case Combine:
			resultStr := strconv.Itoa(result) + strconv.Itoa(arr[i])
			var err error
			result, err = strconv.Atoi(resultStr)
			if err != nil {
				panic(err)
			}
		}
	}
	return result
}
