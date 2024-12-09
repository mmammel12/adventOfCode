// Package day6 -
package day6

import (
	"fmt"
	"slices"
)

const obstruction = '#'

func createGuardData() (map[rune][]int, map[rune]rune, []rune) {
	guardDirections := map[rune][]int{
		'^': {-1, 0},
		'v': {1, 0},
		'<': {0, -1},
		'>': {0, 1},
	}

	guardTurns := map[rune]rune{
		'^': '>',
		'>': 'v',
		'v': '<',
		'<': '^',
	}

	guardChars := []rune{'^', 'v', '>', '<'}

	return guardDirections, guardTurns, guardChars
}

// Part1 -
func Part1(lines []string) (int, error) {
	guardDirections, guardTurns, guardChars := createGuardData()

	guardRow := -1
	guardCol := -1
	guard := ' '
	visited := make(map[string]bool)
	count := 1

	for row, line := range lines {
		for col, char := range line {
			if slices.Contains(guardChars, char) {
				guardRow = row
				guardCol = col
				guard = char
				break
			}
		}
	}

	for true {
		direction := guardDirections[guard]
		if guardRow+direction[0] < 0 || guardCol+direction[1] < 0 || guardRow+direction[0] >= len(lines) || guardCol+direction[1] >= len(lines[0]) {
			break
		}
		location := fmt.Sprintf("%v,%v", guardRow, guardCol)
		if _, exists := visited[location]; !exists {
			count++
			visited[location] = true
		}

		if lines[guardRow+direction[0]][guardCol+direction[1]] == obstruction {
			guard = guardTurns[guard]
			direction = guardDirections[guard]
		}
		guardRow += direction[0]
		guardCol += direction[1]
	}

	return count, nil
}
