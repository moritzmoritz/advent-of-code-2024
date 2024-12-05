package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"time"
)

type Position struct {
	X int
	Y int
}

type Letter struct {
	Letter   string
	Position Position
}

func main() {
	inputFile := "input.txt"

	file, err := os.Open(inputFile)

	if err != nil {
		panic(err)
	}

	defer file.Close()

	executionStartTime := time.Now()

	lines := extractLines(file)

	// Part 1
	total := 0

	total += horizontalMatches(lines)
	total += verticalMatches(lines)
	total += diagonalMatches(lines)

	// Part 2
	xmasCount := xmasMatches(lines)

	executionTime := time.Since(executionStartTime)
	fmt.Println("Part 1:", total)
	fmt.Println("Part 2:", xmasCount)

	fmt.Println("Execution time:", executionTime)
}

func extractLines(input *os.File) []string {
	lines := []string{}

	scanner := bufio.NewScanner(input)

	for scanner.Scan() {
		line := scanner.Text()
		lines = append(lines, line)
	}

	if err := scanner.Err(); err != nil {
		panic(err)
	}

	return lines
}

func horizontalMatches(lines []string) int {
	total := 0

	// Find all "XMAS"
	for _, line := range lines {
		re := regexp.MustCompile(`XMAS`)
		matches := re.FindAllStringSubmatch(line, -1)
		total += len(matches)
	}

	// Find all "SAMX"
	for _, line := range lines {
		re := regexp.MustCompile(`SAMX`)
		matches := re.FindAllStringSubmatch(line, -1)
		total += len(matches)
	}

	return total
}

func verticalMatches(lines []string) int {
	total := 0

	for y, line := range lines {
		for x, letter := range line {
			if letter == 'X' {
				if y < len(lines)-3 {
					if lines[y+1][x] == 'M' && lines[y+2][x] == 'A' && lines[y+3][x] == 'S' {
						total++
					}

					if x >= 3 {
						if lines[y+1][x-1] == 'M' && lines[y+2][x-2] == 'A' && lines[y+3][x-3] == 'S' {
							total++
						}
					}
				}

				if y >= 3 {
					if lines[y-1][x] == 'M' && lines[y-2][x] == 'A' && lines[y-3][x] == 'S' {
						total++
					}

					if x < len(line)-3 {
						if lines[y-1][x+1] == 'M' && lines[y-2][x+2] == 'A' && lines[y-3][x+3] == 'S' {
							total++
						}
					}
				}
			}
		}
	}

	return total
}

func diagonalMatches(lines []string) int {
	total := 0

	for y, line := range lines {
		for x, letter := range line {
			if letter == 'X' {
				if y < len(lines)-3 && x < len(line)-3 {
					if lines[y+1][x+1] == 'M' && lines[y+2][x+2] == 'A' && lines[y+3][x+3] == 'S' {
						total++
					}
				}

				if y >= 3 && x >= 3 {
					if lines[y-1][x-1] == 'M' && lines[y-2][x-2] == 'A' && lines[y-3][x-3] == 'S' {
						total++
					}
				}
			}
		}
	}

	return total
}

func xmasMatches(lines []string) int {
	centerLocations := []Position{}

	totalLines := len(lines)
	for y, line := range lines {
		for x, letter := range line {
			if letter == 'M' {
				// M
				// .A
				// ..S
				if x < len(line)-2 && y < totalLines-2 {
					if lines[y+1][x+1] == 'A' && lines[y+2][x+2] == 'S' {
						centerLocations = append(centerLocations, Position{x + 1, y + 1})
					}
				}

				// ..M
				// .A
				// S
				if x > 1 && y < totalLines-2 {
					if lines[y+1][x-1] == 'A' && lines[y+2][x-2] == 'S' {
						centerLocations = append(centerLocations, Position{x - 1, y + 1})
					}
				}

				// ..S
				// .A
				// M
				if x < len(line)-2 && y > 1 {
					if lines[y-1][x+1] == 'A' && lines[y-2][x+2] == 'S' {
						centerLocations = append(centerLocations, Position{x + 1, y - 1})
					}
				}

				// S
				// .A
				// ..M
				if x > 1 && y > 1 {
					if lines[y-1][x-1] == 'A' && lines[y-2][x-2] == 'S' {
						centerLocations = append(centerLocations, Position{x - 1, y - 1})
					}
				}

			}
		}
	}

	counts := map[Position]int{}

	for _, location := range centerLocations {
		counts[location]++
	}

	total := 0
	for _, count := range counts {
		if count == 2 {
			total++
		}
	}

	return total
}
