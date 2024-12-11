package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"time"
)

type Position struct {
	X      int
	Y      int
	Height int
}

type Trailhead struct {
	Score int
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
	trailheads := findTrailheads(&world, true)

	// Part 2
	distinctTrailheads := findTrailheads(&world, false)

	executionTime := time.Since(executionStartTime)

	fmt.Println("Part 1:", len(trailheads))
	fmt.Println("Part 2:", len(distinctTrailheads))
	fmt.Println("Execution time:", executionTime)
}

func readMap(file *os.File) []Position {
	scanner := bufio.NewScanner(file)

	world := []Position{}

	y := 0

	for scanner.Scan() {
		line := scanner.Text()

		for x, char := range line {
			value, err := strconv.Atoi(string(char))

			if err != nil {
				panic(err)
			}

			position := Position{
				X:      x,
				Y:      y,
				Height: value,
			}

			world = append(world, position)
		}

		y++
	}

	return world
}

func findTrailheads(world *[]Position, partOne bool) []Position {
	startPositions := []Position{}

	for _, position := range *world {
		if position.Height == 0 {
			startPositions = append(startPositions, position)
		}
	}

	trailheads := []Position{}

	for _, position := range startPositions {
		visited := map[Position]bool{}
		positionTrailheads := []Position{}

		countPaths(position, world, &positionTrailheads, &visited, partOne)
		trailheads = append(trailheads, positionTrailheads...)
	}

	return trailheads
}

func countPaths(position Position, world *[]Position, trailheads *[]Position, visited *map[Position]bool, partOne bool) {
	if partOne {
		if _, ok := (*visited)[position]; ok {
			return
		}

		(*visited)[position] = true
	}

	if position.Height == 9 {
		*trailheads = append(*trailheads, position)
		return
	}

	// Find neighbors
	neighbors := findNeighbors(position, world)

	for _, neighbor := range neighbors {
		countPaths(neighbor, world, trailheads, visited, partOne)
	}
}

func findTrailheadAtPosition(world *[]Position, x int, y int) *Position {
	for _, position := range *world {
		if position.X == x && position.Y == y {
			return &position
		}
	}

	return nil
}

func findNeighbors(position Position, world *[]Position) []Position {
	neighbors := []Position{}
	neighborHeight := position.Height + 1

	for _, neighbor := range *world {
		if neighbor.Height == neighborHeight {
			if neighbor.X == position.X && neighbor.Y == position.Y-1 {
				// Top
				neighbors = append(neighbors, neighbor)
			} else if neighbor.X == position.X && neighbor.Y == position.Y+1 {
				// Bottom
				neighbors = append(neighbors, neighbor)
			} else if neighbor.X == position.X-1 && neighbor.Y == position.Y {
				// Left
				neighbors = append(neighbors, neighbor)
			} else if neighbor.X == position.X+1 && neighbor.Y == position.Y {
				// Right
				neighbors = append(neighbors, neighbor)
			}
		}

		if len(neighbors) == 4 {
			return neighbors
		}
	}

	return neighbors
}
