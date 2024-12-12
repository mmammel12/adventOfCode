// Package day11 -
package day11

import (
	"fmt"
	"strings"
	"time"
)

// Part2 -
func Part2(lines []string) (int, error) {
	line := strings.Split(lines[0], " ")

	ans := 0
	start := time.Now()
	newMemo := make(map[memoKey]int)
	for _, numStr := range line {
		ans += expandMemoR(numStr, newMemo, 75)
	}
	elapsed := time.Since(start)
	fmt.Printf("elapsed: %v\n", elapsed)

	return ans, nil
}
