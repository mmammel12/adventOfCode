// Package day10 -
package day10

import (
	"strconv"
)

type node struct {
	up    *node
	down  *node
	left  *node
	right *node
	final bool
}

func countFinalLeaves(n *node) int {
	if n == nil {
		return 0
	}

	if n.final {
		return 1
	}

	return countFinalLeaves(n.up) +
		countFinalLeaves(n.down) +
		countFinalLeaves(n.left) +
		countFinalLeaves(n.right)
}

func searchP2(lines [][]int, start loc, next int, coordsNode *node) {
	// check above
	if start.row-1 >= 0 {
		nextLoc := loc{start.row - 1, start.col}
		if lines[nextLoc.row][nextLoc.col] == next {
			coordsNode.up = &node{final: true}
			if next != 9 {
				coordsNode.up.final = false
				searchP2(lines, nextLoc, next+1, coordsNode.up)
			}
		}
	}

	// check below
	if start.row+1 < len(lines) {
		nextLoc := loc{start.row + 1, start.col}
		if lines[nextLoc.row][nextLoc.col] == next {
			coordsNode.down = &node{final: true}
			if next != 9 {
				coordsNode.down.final = false
				searchP2(lines, nextLoc, next+1, coordsNode.down)
			}
		}
	}

	// check left
	if start.col-1 >= 0 {
		nextLoc := loc{start.row, start.col - 1}
		if lines[nextLoc.row][nextLoc.col] == next {
			coordsNode.left = &node{final: true}
			if next != 9 {
				coordsNode.left.final = false
				searchP2(lines, nextLoc, next+1, coordsNode.left)
			}
		}
	}

	// check right
	if start.col+1 < len(lines[0]) {
		nextLoc := loc{start.row, start.col + 1}
		if lines[nextLoc.row][nextLoc.col] == next {
			coordsNode.right = &node{final: true}
			if next != 9 {
				coordsNode.right.final = false
				searchP2(lines, nextLoc, next+1, coordsNode.right)
			}
		}
	}
}

func createCheckTrailP2(lines [][]int) func(loc) *node {
	return func(start loc) *node {
		root := &node{nil, nil, nil, nil, false}
		searchP2(lines, start, 1, root)
		return root
	}
}

// Part2 -
func Part2(lines []string) (int, error) {
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

	checkTrail := createCheckTrailP2(splitLines)
	ans := 0
	for row, line := range splitLines {
		for col, num := range line {
			if num == 0 {
				trailTree := checkTrail(loc{row, col})
				ans += countFinalLeaves(trailTree)
			}
		}
	}

	return ans, nil
}
