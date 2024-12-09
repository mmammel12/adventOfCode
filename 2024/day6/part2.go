// Package day6 -
package day6

import (
	"fmt"
	"slices"
)

func checkLoop(lines []string, guardRow, guardCol, obsRow, obsCol int, guard rune) int {
	guardDirections, guardTurns, _ := createGuardData()
	path := make(map[string]rune)

	for true {
		direction := guardDirections[guard]
		if guardRow+direction[0] < 0 || guardCol+direction[1] < 0 || guardRow+direction[0] >= len(lines) || guardCol+direction[1] >= len(lines[0]) {
			break
		}

		if lines[guardRow+direction[0]][guardCol+direction[1]] == obstruction || (guardRow+direction[0] == obsRow && guardCol+direction[1] == obsCol) {
			guard = guardTurns[guard]
			direction = guardDirections[guard]
			if prevGuard, exists := path[fmt.Sprintf("%v,%v", guardRow, guardCol)]; exists && prevGuard == guard {
				return 1
			}
			path[fmt.Sprintf("%v,%v", guardRow, guardCol)] = guard
		}

		guardRow += direction[0]
		guardCol += direction[1]
	}

	return 0
}

// Part2 -
func Part2(lines []string) (int, error) {
	writeLines := make([]string, len(lines))
	copy(writeLines, lines)

	guardDirections, guardTurns, guardChars := createGuardData()
	guardRow := -1
	guardCol := -1
	guard := ' '
	path := make([][2]int, 0)
	count := 0

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
	guardStartRow := guardRow
	guardStartCol := guardCol
	guardStart := guard

	for true {
		direction := guardDirections[guard]
		path = append(path, [2]int{guardRow, guardCol})

		if guardRow+direction[0] < 0 || guardCol+direction[1] < 0 || guardRow+direction[0] >= len(lines) || guardCol+direction[1] >= len(lines[0]) {
			break
		}
		if lines[guardRow+direction[0]][guardCol+direction[1]] == obstruction {
			guard = guardTurns[guard]
			direction = guardDirections[guard]
		}
		guardRow += direction[0]
		guardCol += direction[1]
	}

	checked := make(map[string]bool)
	for i := 1; i < len(path); i++ {
		location := path[i]
		obsRow := location[0]
		obsCol := location[1]

		checkedStr := fmt.Sprintf("%v,%v", obsRow, obsCol)
		if _, exists := checked[checkedStr]; exists {
			continue
		}
		checked[checkedStr] = true

		count += checkLoop(lines, guardStartRow, guardStartCol, obsRow, obsCol, guardStart)
	}

	return count, nil
}
