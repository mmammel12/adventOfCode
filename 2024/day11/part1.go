// Package day11 -
package day11

import (
	"fmt"
	"strconv"
	"strings"
	"time"
)

type memoKey struct {
	num    string
	blinks int
}

func expandMemoR(stone string, memo map[memoKey]int, blinks int) int {
	key := memoKey{stone, blinks}
	if val, exists := memo[key]; exists {
		return val
	}

	if blinks == 0 {
		return 1
	}

	if stone == "0" {
		return expandMemoR("1", memo, blinks-1)
	}

	if len(stone)%2 == 0 {
		mid := len(stone) / 2
		left := stone[:mid]
		right := strings.TrimLeft(stone[mid:], "0")
		if right == "" {
			right = "0"
		}
		total := expandMemoR(left, memo, blinks-1) + expandMemoR(right, memo, blinks-1)
		memo[key] = total
		return total
	}

	num, err := strconv.Atoi(stone)
	if err != nil {
		fmt.Printf("error: %v\n", err)
		panic(err)
	}
	total := expandMemoR(fmt.Sprintf("%v", num*2024), memo, blinks-1)
	memo[key] = total
	return total
}

func expandMemo(nums []string, memo map[string][]string) []string {
	newNums := make([]string, 0)
	for _, numStr := range nums {
		if val, exists := memo[numStr]; exists {
			newNums = append(newNums, val...)
			continue
		}

		var iterNums []string
		if numStr == "0" {
			iterNums = []string{"1"}
		} else if len(numStr)%2 == 0 {
			mid := len(numStr) / 2
			left := numStr[:mid]
			right := strings.TrimLeft(numStr[mid:], "0")
			if right == "" {
				right = "0"
			}
			iterNums = []string{left, right}
		} else {
			num, err := strconv.Atoi(numStr)
			if err != nil {
				fmt.Printf("error: %v\n", err)
				panic(err)
			}
			iterNums = []string{strconv.Itoa(num * 2024)}
		}
		memo[numStr] = iterNums
		newNums = append(newNums, iterNums...)
	}

	return newNums
}

func expand(nums []string) []string {
	var newNums []string
	for _, numStr := range nums {
		if numStr == "0" {
			newNums = append(newNums, "1")
		} else if len(numStr)%2 == 0 {
			mid := len(numStr) / 2
			left := numStr[:mid]
			right := strings.TrimLeft(numStr[mid:], "0")
			if right == "" {
				right = "0"
			}
			newNums = append(newNums, left, right)
		} else {
			num, err := strconv.Atoi(numStr)
			if err != nil {
				fmt.Printf("error: %v\n", err)
				panic(err)
			}
			newNums = append(newNums, strconv.Itoa(num*2024))
		}
	}

	return newNums
}

// Part1 -
func Part1(lines []string) (int, error) {
	line := strings.Split(lines[0], " ")

	ans := 0
	noMemoStart := time.Now()
	for _, numStr := range line {
		nums := []string{numStr}
		for i := 0; i < 25; i++ {
			nums = expand(nums)
		}
		ans += len(nums)
	}
	noMemoElapsed := time.Since(noMemoStart)

	ans2 := 0
	memoStart := time.Now()
	memo := make(map[string][]string)
	for _, numStr := range line {
		nums := []string{numStr}
		for i := 0; i < 25; i++ {
			nums = expandMemo(nums, memo)
		}
		ans2 += len(nums)
	}
	memoElapsed := time.Since(memoStart)

	ans3 := 0
	memoRecursiveStart := time.Now()
	newMemo := make(map[memoKey]int)
	for _, numStr := range line {
		ans3 += expandMemoR(numStr, newMemo, 25)
	}
	memoRecursiveElapsed := time.Since(memoRecursiveStart)

	fmt.Printf("No memo: %v\n", noMemoElapsed)
	fmt.Printf("memo: %v\n", memoElapsed)
	fmt.Printf("memo R: %v\n", memoRecursiveElapsed)

	fmt.Printf("memo answer: %v\n", ans2)
	fmt.Printf("memo R answer: %v\n", ans3)

	return ans, nil
}
