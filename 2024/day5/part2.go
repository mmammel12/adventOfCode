// Package day5 -
package day5

import (
	"slices"
	"strconv"
	"strings"
)

func checkLine(line []string, rules map[string][]string) (int, int, bool) {
	seen := make(map[string]bool)
	seen[line[0]] = true
	for i := 1; i < len(line); i++ {
		if beforeSlice, exists := rules[line[i]]; exists {
			for _, before := range beforeSlice {
				if _, invalid := seen[before]; invalid {
					return i, slices.Index(line, before), false
				}
			}
		}
		seen[line[i]] = true
	}

	return -1, -1, true
}

func fixLine(line []string, rules map[string][]string, right, left int) []string {
	temp := line[right]
	line[right] = line[left]
	line[left] = temp

	right, left, isValid := checkLine(line, rules)
	if !isValid {
		line = fixLine(line, rules, right, left)
	}

	return line
}

// Part2 -
func Part2(lines []string) (int, error) {
	rules := make(map[string][]string)
	updates := make([][]string, 0)

	isRule := true
	for _, line := range lines {
		if len(line) == 0 {
			if !isRule {
				break
			}
			isRule = false
			continue
		}

		if isRule {
			left, right, _ := strings.Cut(line, "|")
			if _, exists := rules[left]; !exists {
				rules[left] = make([]string, 0)
			}
			rules[left] = append(rules[left], right)
		} else {
			updates = append(updates, strings.Split(line, ","))
		}
	}

	ans := 0
	for _, line := range updates {
		right, left, isValid := checkLine(line, rules)

		if !isValid {
			fixedLine := make([]string, len(line))
			copy(fixedLine, line)
			fixedLine = fixLine(fixedLine, rules, right, left)

			middleIdx := len(fixedLine) / 2
			num, err := strconv.Atoi(fixedLine[middleIdx])
			if err != nil {
				return 0, err
			}
			ans += num
		}
	}

	return ans, nil
}
