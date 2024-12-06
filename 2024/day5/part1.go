// Package day5 -
package day5

import (
	"strconv"
	"strings"
)

// Part1 -
func Part1(lines []string) (int, error) {
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
	for _, update := range updates {
		isValid := true
		seen := make(map[string]bool)
		seen[update[0]] = true
		for i := 1; i < len(update); i++ {
			if beforeSlice, exists := rules[update[i]]; exists {
				for _, before := range beforeSlice {
					if _, invalid := seen[before]; invalid {
						isValid = false
						break
					}
				}
			}
			seen[update[i]] = true

			if !isValid {
				break
			}
		}

		if isValid {
			middleIdx := len(update) / 2
			num, err := strconv.Atoi(update[middleIdx])
			if err != nil {
				return 0, err
			}
			ans += num
		}
	}

	return ans, nil
}
