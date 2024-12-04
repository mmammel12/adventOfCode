// Package day3 -
package day3

import (
	"strings"
)

// Part2 -
func Part2(lines []string) (int, error) {
	sum := 0
	joined := strings.Join(lines, "\n")
	splitStr := strings.Split(joined, "do()")
	for _, split := range splitStr {
		dos := strings.Split(split, "don't()")[0]
		val, err := Part1([]string{dos})
		if err != nil {
			return 0, err
		}
		sum += val
	}

	return sum, nil
}
