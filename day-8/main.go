package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"time"
)

type Position struct {
	X int
	Y int
}

type Antenna struct {
	Frequency string
	Position  Position
}

type World struct {
	Width    int
	Height   int
	Antennas []Antenna
}

func main() {
	inputFile := "input.txt"

	file, err := os.Open(inputFile)

	if err != nil {
		panic(err)
	}

	defer file.Close()

	world := readMap(file)

	executionStartTime := time.Now()

	// Part 1
	antinodes := findAntinodes(world, false)
	uniqueAntinodes := filterDuplicateAntinodes(antinodes)

	// Part 2
	inLineAntinodes := findAntinodes(world, true)
	uniqueInlineAntinodes := filterDuplicateAntinodes(inLineAntinodes)

	// partTwoTotal := findTotalCalibrationResult(equations, 3)

	executionTime := time.Since(executionStartTime)

	fmt.Println("Part 1:", len(uniqueAntinodes))
	fmt.Println("Part 2:", len(uniqueInlineAntinodes))
	fmt.Println("Execution time:", executionTime)
}

func readMap(file *os.File) World {
	world := World{}

	scanner := bufio.NewScanner(file)

	y := 0

	for scanner.Scan() {
		line := scanner.Text()

		world.Width = int(math.Max(float64(world.Width), float64(len(line))))

		for x, char := range line {
			if char != '.' {
				antenna := Antenna{
					Frequency: string(char),
					Position: Position{
						X: x,
						Y: y,
					},
				}

				world.Antennas = append(world.Antennas, antenna)
			}
		}

		y++
	}

	err := scanner.Err()

	if err != nil {
		panic(err)
	}

	world.Height = y
	return world
}

func findAntinodes(worldMap World, inLine bool) []Position {
	antinodes := []Position{}

	for _, antenna := range worldMap.Antennas {
		otherAntennas := []Antenna{}

		for _, otherAntenna := range worldMap.Antennas {
			if otherAntenna.Frequency == antenna.Frequency && antenna.Position != otherAntenna.Position {
				otherAntennas = append(otherAntennas, otherAntenna)
			}
		}

		for _, otherAntenna := range otherAntennas {
			positionIndex := 1

			if inLine {
				positionIndex = 0
			}

			validPosition := true

			for validPosition {
				xDistance := antenna.Position.X - otherAntenna.Position.X
				yDistance := antenna.Position.Y - otherAntenna.Position.Y

				antinode := Position{
					X: antenna.Position.X + xDistance*positionIndex,
					Y: antenna.Position.Y + yDistance*positionIndex,
				}

				if !inLine {
					validPosition = false

					if antinode.X < 0 || antinode.X >= worldMap.Width || antinode.Y < 0 || antinode.Y >= worldMap.Height {
						continue
					}

					antinodes = append(antinodes, antinode)
				} else {
					if antinode.X < 0 || antinode.X >= worldMap.Width || antinode.Y < 0 || antinode.Y >= worldMap.Height {
						validPosition = false
						continue
					}

					antinodes = append(antinodes, antinode)
					positionIndex++
				}
			}
		}
	}

	return antinodes
}

func filterDuplicateAntinodes(antinodes []Position) []Position {
	filteredAntinodes := []Position{}

	for _, antinode := range antinodes {
		isDuplicate := false

		for _, filteredAntinode := range filteredAntinodes {
			if filteredAntinode == antinode {
				isDuplicate = true
				break
			}
		}

		if !isDuplicate {
			filteredAntinodes = append(filteredAntinodes, antinode)
		}
	}

	return filteredAntinodes
}
