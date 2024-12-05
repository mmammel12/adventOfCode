// Package day4 -
package day4

import (
	"strings"
)

const xmas = "XMAS"
const samx = "SAMX"

// Part1 -
func Part1(lines []string) (int, error) {
	count := 0
	for row, line := range lines {
		count += strings.Count(line, xmas) + strings.Count(line, samx)
		for col, ch := range line {
			if ch == 'X' {
				// check up
				xmasCharIdx := 1
				if row-3 >= 0 {
					found := 1
					for j := 1; j < 4; j++ {
						if lines[row-j][col] != xmas[xmasCharIdx] {
							found = 0
							break
						}
						xmasCharIdx++
					}
					count += found
				}

				// check down
				xmasCharIdx = 1
				if row+3 < len(lines) {
					found := 1
					for j := 1; j < 4; j++ {
						if lines[row+j][col] != xmas[xmasCharIdx] {
							found = 0
							break
						}
						xmasCharIdx++
					}
					count += found
				}

				// check diagonal down right
				xmasCharIdx = 1
				if row+3 < len(lines) && col+3 < len(lines) {
					found := 1
					for j := 1; j < 4; j++ {
						if lines[row+j][col+j] != xmas[xmasCharIdx] {
							found = 0
							break
						}
						xmasCharIdx++
					}
					count += found
				}

				// check diagonal up left
				xmasCharIdx = 1
				if row-3 >= 0 && col-3 >= 0 {
					found := 1
					for j := 1; j < 4; j++ {
						if lines[row-j][col-j] != xmas[xmasCharIdx] {
							found = 0
							break
						}
						xmasCharIdx++
					}
					count += found
				}

				// check diagonal down left
				xmasCharIdx = 1
				if row+3 < len(lines) && col-3 >= 0 {
					found := 1
					for j := 1; j < 4; j++ {
						if lines[row+j][col-j] != xmas[xmasCharIdx] {
							found = 0
							break
						}
						xmasCharIdx++
					}
					count += found
				}

				// check diagonal up right
				xmasCharIdx = 1
				if row-3 >= 0 && col+3 < len(line) {
					found := 1
					for j := 1; j < 4; j++ {
						if lines[row-j][col+j] != xmas[xmasCharIdx] {
							found = 0
							break
						}
						xmasCharIdx++
					}
					count += found
				}
			}
		}
	}

	return count, nil
}
