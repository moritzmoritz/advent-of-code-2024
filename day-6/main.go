package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"time"
)

type Direction = int

const (
	Up Direction = iota
	Down
	Left
	Right
)

type Position struct {
	X int
	Y int
}

type Map struct {
	Width  int
	Height int
}

type Guard struct {
	Position  Position
	Direction Direction
}

type Obstacle struct {
	Position Position
}

func main() {
	inputFile := "input.txt"

	file, err := os.Open(inputFile)

	if err != nil {
		panic(err)
	}

	defer file.Close()

	guard, obstacles, worldMap := readMap(file)

	executionStartTime := time.Now()

	// Part 1
	steps, visitedPositions, forLoop := walkMap(guard, obstacles, worldMap, false)
	fmt.Println(steps, len(visitedPositions), forLoop)

	// Part 2
	positions := walkMapWithAddingObstacles(guard, obstacles, worldMap, visitedPositions)
	fmt.Println(positions)

	executionTime := time.Since(executionStartTime)

	fmt.Println("Execution time:", executionTime)
}

func readMap(file *os.File) (Guard, []Obstacle, Map) {
	guard := Guard{}
	obstacles := []Obstacle{}
	worldMap := Map{}

	scanner := bufio.NewScanner(file)
	y := 0

	for scanner.Scan() {
		line := scanner.Text()

		width := len(line)

		for x, char := range line {
			if char == '#' {
				obstacles = append(obstacles, Obstacle{Position: Position{X: x, Y: y}})
			} else if char == '^' {
				guard.Direction = Up
				guard.Position = Position{X: x, Y: y}
			} else if char == '>' {
				guard.Direction = Right
				guard.Position = Position{X: x, Y: y}
			} else if char == 'v' {
				guard.Direction = Down
				guard.Position = Position{X: x, Y: y}
			} else if char == '<' {
				guard.Direction = Left
				guard.Position = Position{X: x, Y: y}
			}
		}

		worldMap.Width = int(math.Max(float64(worldMap.Width), float64(width)))
		y++
	}

	if err := scanner.Err(); err != nil {
		panic(err)
	}

	worldMap.Height = y
	return guard, obstacles, worldMap
}

func walkMap(guard Guard, obstacles []Obstacle, worldMap Map, checkLoop bool) (int, []Position, bool) {
	steps := 1
	walkedPositions := map[Position]bool{
		guard.Position: true,
	}

	walking := true
	forLoop := false

	for walking {
		var nextPosition Position

		switch guard.Direction {
		case Up:
			nextPosition = Position{X: guard.Position.X, Y: guard.Position.Y - 1}
		case Down:
			nextPosition = Position{X: guard.Position.X, Y: guard.Position.Y + 1}
		case Left:
			nextPosition = Position{X: guard.Position.X - 1, Y: guard.Position.Y}
		case Right:
			nextPosition = Position{X: guard.Position.X + 1, Y: guard.Position.Y}
		}

		obstacleFound := false

		for _, obstacle := range obstacles {
			if obstacle.Position == nextPosition {
				var newDirection Direction

				switch guard.Direction {
				case Up:
					newDirection = Right
				case Down:
					newDirection = Left
				case Left:
					newDirection = Up
				case Right:
					newDirection = Down
				}

				guard.Direction = newDirection
				obstacleFound = true
				break
			}
		}

		if obstacleFound {
			continue
		}

		if nextPosition.X < 0 || nextPosition.X >= worldMap.Width || nextPosition.Y < 0 || nextPosition.Y >= worldMap.Height {
			walking = false
			break
		}

		if !checkLoop {
			if _, ok := walkedPositions[nextPosition]; !ok {
				steps++
				walkedPositions[nextPosition] = true
			}
		} else {
			steps++
			if steps > 99999 {
				walking = false
				forLoop = true
				break
			}
		}

		guard.Position = nextPosition
	}

	visitedPositions := []Position{}

	for position := range walkedPositions {
		visitedPositions = append(visitedPositions, position)
	}

	return steps, visitedPositions, forLoop
}

func walkMapWithAddingObstacles(guard Guard, obstacles []Obstacle, worldMap Map, walkingPath []Position) int {
	positions := 0

	for idx, position := range walkingPath {
		if idx == 0 {
			continue
		}
		newObstacles := append(obstacles, Obstacle{Position: position})

		_, _, forLoopDetected := walkMap(guard, newObstacles, worldMap, true)

		if forLoopDetected {
			positions++
		}
	}

	return positions
}
