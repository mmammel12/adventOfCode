package day2

import (
	"strconv"
	"strings"
)

func part2CheckLine(line []int) int {
	part1Safe := CheckLine(line)
	if part1Safe == 1 {
		return 1
	}

	for i := range line {
		copyLine := make([]int, len(line))
		copy(copyLine, line)
		copyLevel := make([]int, len(line)-1)
		copy(copyLevel, append(copyLine[:i], copyLine[i+1:]...))
		if CheckLine(copyLevel) == 1 {
			return 1
		}
	}
	return 0
}

// Part2 -
func Part2(lines []string) (int, error) {
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
		safe += part2CheckLine(level)
	}

	return safe, nil
}
