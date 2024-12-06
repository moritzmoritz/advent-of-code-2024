package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

type OrderRule struct {
	First  int
	Second int
}

// func main() {
// 	inputFile := "input.txt"

// 	file, err := os.Open(inputFile)

// 	if err != nil {
// 		panic(err)
// 	}

// 	defer file.Close()

// 	orderRules, updates := extractOrderRulesAndUpdates(file)

// 	executionStartTime := time.Now()

// 	// lines := extractLines(file)

// 	// Part 1
// 	validUpdates := filterValidUpdates(orderRules, updates)
// 	total := 0

// 	for _, update := range validUpdates {
// 		middleValue := findMiddleValue(update)
// 		total += middleValue
// 	}

// 	// total := 0

// 	// total += horizontalMatches(lines)
// 	// total += verticalMatches(lines)
// 	// total += diagonalMatches(lines)

// 	// Part 2
// 	// xmasCount := xmasMatches(lines)

// 	executionTime := time.Since(executionStartTime)
// 	fmt.Println("Part 1:", total)
// 	// fmt.Println("Part 2:", xmasCount)

// 	fmt.Println("Execution time:", executionTime)
// }

// func extractOrderRulesAndUpdates(input *os.File) ([]OrderRule, [][]int) {
// 	updates := [][]int{}
// 	orderRules := []OrderRule{}

// 	re := regexp.MustCompile(`(\d+)\|(\d+)`)

// 	scanner := bufio.NewScanner(input)

// 	for scanner.Scan() {
// 		line := scanner.Text()

// 		ruleMatches := re.FindStringSubmatch(line)

// 		if len(ruleMatches) == 3 {
// 			first, err := strconv.Atoi(ruleMatches[1])

// 			if err != nil {
// 				panic(err)
// 			}

// 			second, err := strconv.Atoi(ruleMatches[2])

// 			if err != nil {
// 				panic(err)
// 			}

// 			orderRule := OrderRule{
// 				First:  first,
// 				Second: second,
// 			}

// 			orderRules = append(orderRules, orderRule)
// 		}

// 		update := strings.Split(line, ",")

// 		if len(update) > 1 {
// 			lineUpdates := []int{}

// 			for _, updateValue := range update {
// 				updateInt, err := strconv.Atoi(updateValue)

// 				if err != nil {
// 					panic(err)
// 				}

// 				lineUpdates = append(lineUpdates, updateInt)
// 			}

// 			updates = append(updates, lineUpdates)
// 		}
// 	}

// 	if err := scanner.Err(); err != nil {
// 		panic(err)
// 	}

// 	return orderRules, updates
// }

// func filterValidUpdates(orderRules []OrderRule, updates [][]int) [][]int {
// 	validUpdates := [][]int{}

// 	for _, update := range updates {
// 		valid := true

// 		for idx, value := range update {
// 			updatesAfter := []int{}
// 			updatesBefore := []int{}

// 			for _, orderRule := range orderRules {
// 				if orderRule.First == value {
// 					updatesAfter = append(updatesAfter, orderRule.Second)
// 				} else if orderRule.Second == value {
// 					updatesBefore = append(updatesBefore, orderRule.First)
// 				}
// 			}

// 			// Check values after
// 			for i := idx + 1; i < len(update); i++ {
// 				contains := false

// 				for _, updateAfter := range updatesAfter {
// 					if update[i] == updateAfter {
// 						contains = true
// 						break
// 					}
// 				}

// 				valid = contains
// 			}

// 			if !valid {
// 				break
// 			}

// 			// Check values before
// 			for i := idx - 1; i >= 0; i-- {
// 				contains := false

// 				for _, updateBefore := range updatesBefore {
// 					if update[i] == updateBefore {
// 						contains = true
// 						break
// 					}
// 				}

// 				valid = contains
// 			}

// 			if !valid {
// 				break
// 			}
// 		}

// 		if valid {
// 			validUpdates = append(validUpdates, update)
// 		}
// 	}

// 	return validUpdates
// }

// func findMiddleValue(update []int) int {
// 	middle := len(update) / 2

// 	return update[middle]
// }

func main() {
	rulesBefore := make(map[string]map[string]struct{})
	rulesAfter := make(map[string]map[string]struct{})
	sortFunc := func(a string, b string) int {
		_, ok := rulesAfter[a][b]
		if ok {
			return -1
		}
		_, ok = rulesBefore[a][b]
		if ok {
			return 1
		}
		return 0
	}

	sumP1 := 0
	sumP2 := 0

	inputFile := "input.txt"

	file, err := os.Open(inputFile)

	if err != nil {
		panic(err)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()

		if strings.Contains(line, "|") {
			rule := strings.Split(line, "|")
			if rulesBefore[rule[1]] == nil {
				rulesBefore[rule[1]] = make(map[string]struct{})
			}
			rulesBefore[rule[1]][rule[0]] = struct{}{}
			if rulesAfter[rule[0]] == nil {
				rulesAfter[rule[0]] = make(map[string]struct{})
			}
			rulesAfter[rule[0]][rule[1]] = struct{}{}
		}
		if strings.Contains(line, ",") {
			strs := strings.Split(line, ",")
			if slices.IsSortedFunc(strs, sortFunc) {
				value, err := strconv.Atoi(strs[len(strs)/2])

				if err != nil {
					panic(err)
				}

				sumP1 += value
			} else {
				slices.SortFunc(strs, sortFunc)

				value, err := strconv.Atoi(strs[len(strs)/2])

				if err != nil {
					panic(err)
				}

				sumP2 += value
			}
		}
	}

	fmt.Println(sumP1)
	fmt.Println(sumP2)
}
