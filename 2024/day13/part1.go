// Package day13 -
package day13

import (
	"regexp"
	"strconv"
	"strings"
)

type point struct {
	x int
	y int
}

func getInputs(lines []string, targetOffset int) (a point, b point, target point) {
	xButtonRegex := `X\+\d+`
	yButtonRegex := `Y\+\d+`
	xTargetRegex := `X=\d+`
	yTargetRegex := `Y=\d+`
	xb := regexp.MustCompile(xButtonRegex)
	yb := regexp.MustCompile(yButtonRegex)
	xt := regexp.MustCompile(xTargetRegex)
	yt := regexp.MustCompile(yTargetRegex)

	axStr, _ := strings.CutPrefix(xb.FindString(lines[0]), "X+")
	ax, _ := strconv.Atoi(axStr)
	ayStr, _ := strings.CutPrefix(yb.FindString(lines[0]), "Y+")
	ay, _ := strconv.Atoi(ayStr)
	a = point{x: ax, y: ay}

	bxStr, _ := strings.CutPrefix(xb.FindString(lines[1]), "X+")
	bx, _ := strconv.Atoi(bxStr)
	byStr, _ := strings.CutPrefix(yb.FindString(lines[1]), "Y+")
	by, _ := strconv.Atoi(byStr)
	b = point{x: bx, y: by}

	txStr, _ := strings.CutPrefix(xt.FindString(lines[2]), "X=")
	tx, _ := strconv.Atoi(txStr)
	tyStr, _ := strings.CutPrefix(yt.FindString(lines[2]), "Y=")
	ty, _ := strconv.Atoi(tyStr)
	target = point{x: tx + targetOffset, y: ty + targetOffset}

	return a, b, target
}

func findSingleAxisSolutions(aMove, bmove, target int) [][]int {
	solutions := make([][]int, 0)
	maxMove1 := target/aMove + 1
	for aPresses := 0; aPresses < maxMove1; aPresses++ {
		remaining := target - aPresses*aMove
		if remaining < 0 {
			break
		}

		if remaining%bmove == 0 {
			bPresses := remaining / bmove
			solutions = append(solutions, []int{aPresses, bPresses})
		}
	}

	return solutions
}

func findMinPrice(a, b, target point) int {
	xSolutions := findSingleAxisSolutions(a.x, b.x, target.x)
	ySolutions := findSingleAxisSolutions(a.y, b.y, target.y)

	solutions := make([]int, 0)
	for _, xSolution := range xSolutions {
		for _, ySolution := range ySolutions {
			if xSolution[0] == ySolution[0] && xSolution[1] == ySolution[1] {
				solutions = append(solutions, (xSolution[0]*3)+xSolution[1])
			}
		}
	}

	if len(solutions) == 0 {
		return 0
	}

	minPrice := solutions[0]
	for _, solution := range solutions {
		if solution < minPrice {
			minPrice = solution
		}
	}

	return minPrice
}

// Part1 -
func Part1(lines []string) (int, error) {
	ans := 0
	for i := 0; i < len(lines); i += 4 {
		a, b, target := getInputs(lines[i:i+3], 0)
		ans += findMinPrice(a, b, target)
	}

	return ans, nil
}
