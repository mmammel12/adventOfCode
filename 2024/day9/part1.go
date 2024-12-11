// Package day9 -
package day9

import (
	"strconv"
	"strings"
)

type mem struct {
	idx    int
	amount int
}

// Part1 -
func Part1(lines []string) (int, error) {
	line := lines[0]
	nums := strings.Split(line, "")

	mems := make([]mem, 0)

	idx := 0
	totalLen := 0
	for i, numStr := range nums {
		num, err := strconv.Atoi(numStr)
		if err != nil {
			return 0, err
		}

		if i%2 == 0 {
			// block
			mems = append(mems, mem{idx, num})
			idx++
			totalLen += num
		} else {
			// empty
			mems = append(mems, mem{-1, num})
		}
	}

	memory := make([]int, totalLen)
	j := 0
	k := len(mems) - 1
	if mems[k].idx == -1 {
		k--
	}
	for i := 0; i < len(mems); i++ {
		if j >= totalLen {
			break
		}
		if mems[i].idx != -1 {
			for mems[i].amount > 0 {
				memory[j] = mems[i].idx
				mems[i].amount--
				j++
				if j >= totalLen {
					break
				}
			}
		} else {
			for mems[i].amount > 0 {
				memory[j] = mems[k].idx
				mems[k].amount--
				mems[i].amount--
				j++

				if mems[k].amount == 0 {
					k -= 2
				}
				if j >= totalLen {
					break
				}
			}
		}
	}

	ans := 0
	for i, val := range memory {
		ans += i * val
	}

	return ans, nil
}
