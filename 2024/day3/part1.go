// Package day3 -
package day3

import (
	"regexp"
	"strconv"
	"strings"
)

// CleanMatch -
func CleanMatch(match string) ([]int, error) {
	nums := match[4 : len(match)-1]
	splitNums := strings.Split(nums, ",")
	x, err := strconv.Atoi(splitNums[0])
	if err != nil {
		return nil, err
	}
	y, err := strconv.Atoi(splitNums[1])
	if err != nil {
		return nil, err
	}

	return []int{x, y}, nil
}

// Part1 -
func Part1(lines []string) (int, error) {
	regex, err := regexp.Compile("mul\\(\\d{1,3},\\d{1,3}\\)")
	if err != nil {
		return 0, err
	}

	sum := 0
	for _, line := range lines {
		matches := regex.FindAllString(line, -1)
		if len(matches) == 0 {
			continue
		}
		for _, match := range matches {
			nums, err := CleanMatch(match)
			if err != nil {
				return 0, err
			}
			sum += nums[0] * nums[1]
		}
	}

	return sum, nil
}
