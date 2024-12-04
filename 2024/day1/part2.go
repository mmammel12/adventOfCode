package day1

import (
	"strconv"
	"strings"
)

// Part2 -
func Part2(lines []string) (int, error) {
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
			return 0, err
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

	return sum, nil
}
