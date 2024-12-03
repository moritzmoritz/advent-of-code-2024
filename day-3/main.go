package main

import (
	"fmt"
	"io/ioutil"
	"regexp"
	"strconv"
	"time"
)

type MultiplicationType int

const (
	MULTIPLICATION MultiplicationType = iota
	DO
	DONT
)

type Multiplication struct {
	FirstInteger  *int
	SecondInteger *int
	Sum           *int
	Type          MultiplicationType
}

func main() {
	inputFile := "input.txt"

	inputData, err := ioutil.ReadFile(inputFile)

	if err != nil {
		panic(err)
	}

	input := string(inputData)

	multiplications := extractMultiplactionsAndInstructions(input)

	executionStartTime := time.Now()

	// Part 1
	total := 0

	for _, multiplication := range multiplications {
		if multiplication.Type != MULTIPLICATION {
			continue
		}

		total += *multiplication.Sum
	}

	// Part 2
	doEnabled := true
	enabledTotal := 0

	for _, multiplication := range multiplications {
		switch multiplication.Type {
		case DO:
			doEnabled = true
		case DONT:
			doEnabled = false
		case MULTIPLICATION:
			if doEnabled {
				enabledTotal += *multiplication.Sum
			}
		}
	}

	executionTime := time.Since(executionStartTime)
	fmt.Println("Part 1:", total)
	fmt.Println("Part 2:", enabledTotal)

	fmt.Println("Execution time:", executionTime)
}

func extractMultiplactionsAndInstructions(input string) []Multiplication {
	multiplications := []Multiplication{}

	re := regexp.MustCompile(`mul\((\d+),(\d+)\)|do\(\)|don\'t\(\)`)
	matches := re.FindAllStringSubmatch(input, -1)

	for _, match := range matches {
		if len(match) == 3 {
			if match[0] == "do()" {
				multiplications = append(multiplications, Multiplication{
					FirstInteger:  nil,
					SecondInteger: nil,
					Type:          DO,
				})
				continue
			} else if match[0] == "don't()" {
				multiplications = append(multiplications, Multiplication{
					FirstInteger:  nil,
					SecondInteger: nil,
					Type:          DONT,
				})
				continue
			}

			firstInteger, err := strconv.Atoi(match[1])

			if err != nil {
				panic(err)
			}

			secondInteger, err := strconv.Atoi(match[2])

			if err != nil {
				panic(err)
			}

			sum := firstInteger * secondInteger
			multiplications = append(multiplications, Multiplication{
				FirstInteger:  &firstInteger,
				SecondInteger: &secondInteger,
				Sum:           &sum,
				Type:          MULTIPLICATION,
			})
		}
	}

	return multiplications
}
