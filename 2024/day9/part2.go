// Package day9 -
package day9

import (
	"strconv"
	"strings"
)

type mem2 struct {
	idx    int
	amount int
	used   bool
}

// Part2 -
func Part2(lines []string) (int, error) {
	line := lines[0]
	nums := strings.Split(line, "")

	mems := make([]mem2, 0)

	idx := 0
	totalLen := 0
	for i, numStr := range nums {
		num, err := strconv.Atoi(numStr)
		if err != nil {
			return 0, err
		}

		totalLen += num
		if i%2 == 0 {
			// block
			mems = append(mems, mem2{idx, num, false})
			idx++
		} else {
			// empty
			mems = append(mems, mem2{-1, num, false})
		}
	}

	memory := make([]mem2, 0)
	for i, loc := range mems {
		if loc.idx >= 0 {
			memory = append(memory, loc)
		} else {
			for j := len(mems) - 1; j > i; j-- {
				if mems[j].idx == -1 || mems[j].used == true {
					continue
				}

				if mems[j].amount <= loc.amount {
					memory = append(memory, mems[j])
					mems[j].used = true
					loc.amount -= mems[j].amount
				}
			}

			if loc.amount > 0 {
				memory = append(memory, mem2{-1, loc.amount, true})
			}
		}
	}

	memoryInts := make([]int, 0)
	for _, loc := range memory {
		if loc.idx == -1 || loc.used {
			for i := 0; i < loc.amount; i++ {
				memoryInts = append(memoryInts, -1)
			}
		} else {
			for i := 0; i < loc.amount; i++ {
				memoryInts = append(memoryInts, loc.idx)
			}
		}
	}

	ans := 0
	for i, val := range memoryInts {
		if val != -1 {
			ans += i * val
		}
	}

	return ans, nil
}
