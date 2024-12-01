// Package day1 -
package day1

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)

// Part1 -
func Part1(data string) error {
	lines := strings.Split(data, "\n")
	left := make([]int, len(lines)-1)
	right := make([]int, len(lines)-1)
	for i, line := range lines {
		nums := strings.Split(line, "   ")
		leftNum, err := strconv.Atoi(nums[0])
		if err != nil {
			break
		}
		rightNum, err := strconv.Atoi(nums[1])
		if err != nil {
			return err
		}
		left[i] = leftNum
		right[i] = rightNum
	}

	sort.Ints(left)
	sort.Ints(right)

	sum := 0
	for i, leftNum := range left {
		rightNum := right[i]
		sum += max(leftNum, rightNum) - min(leftNum, rightNum)
	}

	fmt.Printf("sum: %v\n", sum)

	return nil
}
