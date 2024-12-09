// Package day7 -
package day7

import (
	"fmt"
	"strconv"
	"strings"
)

type p2Node struct {
	lhs   *p2Node
	mhs   *p2Node
	rhs   *p2Node
	value int
}

func buildP2Tree(root *p2Node, nums []string, i, result int) {
	num, err := strconv.Atoi(nums[i])
	if err != nil {
		fmt.Printf("Error: %v\n", err)
	}
	root.lhs = &p2Node{value: root.value * num, lhs: nil, mhs: nil, rhs: nil}
	root.rhs = &p2Node{value: root.value + num, lhs: nil, mhs: nil, rhs: nil}

	mhsValue, err := strconv.Atoi(strconv.Itoa(root.value) + nums[i])
	if err != nil {
		fmt.Printf("Error: %v\n", err)
	}
	root.mhs = &p2Node{value: mhsValue, lhs: nil, mhs: nil, rhs: nil}

	i++
	if i < len(nums) {
		if root.lhs.value <= result {
			buildP2Tree(root.lhs, nums, i, result)
		}
		if root.mhs.value <= result {
			buildP2Tree(root.mhs, nums, i, result)
		}
		if root.rhs.value <= result {
			buildP2Tree(root.rhs, nums, i, result)
		}
	}
}

func dfsP2(root *p2Node, result, depth, expectedDepth int) bool {
	if depth == expectedDepth && root.value == result {
		return true
	}

	depth++
	left := false
	middle := false
	right := false
	if root.lhs != nil {
		left = dfsP2(root.lhs, result, depth, expectedDepth)
	}
	if root.mhs != nil {
		middle = dfsP2(root.mhs, result, depth, expectedDepth)
	}
	if root.rhs != nil {
		right = dfsP2(root.rhs, result, depth, expectedDepth)
	}

	return left || right || middle
}

// Part2 -
func Part2(lines []string) (int, error) {
	count := 0

	for _, line := range lines {
		resultStr, numsStr, _ := strings.Cut(line, ": ")
		result, err := strconv.Atoi(resultStr)
		if err != nil {
			return 0, err
		}

		nums := strings.Split(numsStr, " ")
		firstNum, err := strconv.Atoi(nums[0])
		root := p2Node{value: firstNum, lhs: nil, mhs: nil, rhs: nil}
		buildP2Tree(&root, nums, 1, result)

		if dfsP2(&root, result, 0, len(nums)-1) {
			count += result
		}
	}

	return count, nil
}
