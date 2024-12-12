// Package day10 -
package day10

import (
	"strconv"
)

type loc struct {
	row int
	col int
}

func search(lines [][]int, start loc, next int, coords map[loc]bool) map[loc]bool {
	// check above
	if start.row-1 >= 0 {
		nextLoc := loc{start.row - 1, start.col}
		if lines[nextLoc.row][nextLoc.col] == next {
			if _, exists := coords[nextLoc]; !exists && next == 9 {
				coords[nextLoc] = true
			} else {
				search(lines, nextLoc, next+1, coords)
			}
		}
	}

	// check below
	if start.row+1 < len(lines) {
		nextLoc := loc{start.row + 1, start.col}
		if lines[nextLoc.row][nextLoc.col] == next {
			if _, exists := coords[nextLoc]; !exists && next == 9 {
				coords[nextLoc] = true
			} else {
				search(lines, nextLoc, next+1, coords)
			}
		}
	}

	// check left
	if start.col-1 >= 0 {
		nextLoc := loc{start.row, start.col - 1}
		if lines[nextLoc.row][nextLoc.col] == next {
			if _, exists := coords[nextLoc]; !exists && next == 9 {
				coords[nextLoc] = true
			} else {
				search(lines, nextLoc, next+1, coords)
			}
		}
	}

	// check right
	if start.col+1 < len(lines[0]) {
		nextLoc := loc{start.row, start.col + 1}
		if lines[nextLoc.row][nextLoc.col] == next {
			if _, exists := coords[nextLoc]; !exists && next == 9 {
				coords[nextLoc] = true
			} else {
				search(lines, nextLoc, next+1, coords)
			}
		}
	}

	return coords
}

func createCheckTrail(lines [][]int) func(loc) map[loc]bool {
	return func(start loc) map[loc]bool {
		coords := make(map[loc]bool, 0)
		return search(lines, start, 1, coords)
	}
}

// Part1 -
func Part1(lines []string) (int, error) {
	splitLines := make([][]int, len(lines))
	for i, line := range lines {
		for _, char := range line {
			num, err := strconv.Atoi(string(char))
			if err != nil {
				return 0, err
			}
			splitLines[i] = append(splitLines[i], num)
		}
	}

	checkTrail := createCheckTrail(splitLines)
	ans := 0
	for row, line := range splitLines {
		for col, num := range line {
			if num == 0 {
				currentLoc := loc{row, col}
				trailEnds := checkTrail(currentLoc)
				ans += len(trailEnds)
			}
		}
	}

	return ans, nil
}
