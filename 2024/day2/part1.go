// Package day2 -
package day2

import (
	"math"
	"strconv"
	"strings"
)

// CheckLine -
func CheckLine(line []int) int {
	minDiff := 1.0
	maxDiff := 3.0
	inc := true
	init := false

	for i := 1; i < len(line); i++ {
		if !init {
			init = true
			if line[i-1] > line[i] {
				inc = false
			}
		}

		if inc {
			if line[i-1] >= line[i] {
				return 0
			}
		} else {
			if line[i-1] <= line[i] {
				return 0
			}
		}

		diff := math.Abs(float64(line[i-1] - line[i]))
		if diff < minDiff || diff > maxDiff {
			return 0
		}
	}

	return 1
}

// Part1 -
func Part1(lines []string) (int, error) {
	levels := make([][]int, len(lines)-1)
	for i, line := range lines {
		if len(line) == 0 {
			break
		}
		strNums := strings.Split(line, " ")
		for _, strNum := range strNums {
			num, err := strconv.Atoi(strNum)
			if err != nil {
				return 0, err
			}
			levels[i] = append(levels[i], num)
		}
	}

	safe := 0

	for _, level := range levels {
		safe += CheckLine(level)
	}

	return safe, nil
}
