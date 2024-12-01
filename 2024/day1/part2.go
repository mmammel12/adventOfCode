// Package day1 -
package day1

import (
	"fmt"
	"strconv"
	"strings"
)

// Part2 -
func Part2(data string) error {
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

	rightMap := make(map[int]int)
	for i, num := range right {
		if _, exists := rightMap[right[i]]; !exists {
			rightMap[num] = 1
		} else {
			rightMap[num]++
		}
	}

	sum := 0
	for _, leftNum := range left {
		if rightCount, exists := rightMap[leftNum]; exists {
			sum += leftNum * rightCount
		}
	}

	fmt.Printf("sum: %v\n", sum)

	return nil
}
