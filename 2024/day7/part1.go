// Package day7 -
package day7

import (
	"fmt"
	"strconv"
	"strings"
)

type node struct {
	lhs   *node
	rhs   *node
	value int
}

func buildTree(root *node, nums []string, i, result int) {
	num, err := strconv.Atoi(nums[i])
	if err != nil {
		fmt.Printf("Error: %v", err)
	}
	root.lhs = &node{value: root.value * num, lhs: nil, rhs: nil}
	root.rhs = &node{value: root.value + num, lhs: nil, rhs: nil}

	i++
	if i < len(nums) {
		if root.lhs.value <= result {
			buildTree(root.lhs, nums, i, result)
		}
		if root.rhs.value <= result {
			buildTree(root.rhs, nums, i, result)
		}
	}
}

func dfs(root *node, result, depth, expectedDepth int) bool {
	if depth == expectedDepth && root.value == result {
		return true
	}

	depth++
	left := false
	right := false
	if root.lhs != nil {
		left = dfs(root.lhs, result, depth, expectedDepth)
	}
	if root.rhs != nil {
		right = dfs(root.rhs, result, depth, expectedDepth)
	}

	return left || right
}

// Part1 -
func Part1(lines []string) (int, error) {
	count := 0

	for _, line := range lines {
		resultStr, numsStr, _ := strings.Cut(line, ": ")
		result, err := strconv.Atoi(resultStr)
		if err != nil {
			return 0, err
		}

		nums := strings.Split(numsStr, " ")
		firstNum, err := strconv.Atoi(nums[0])
		root := node{value: firstNum, lhs: nil, rhs: nil}
		buildTree(&root, nums, 1, result)

		if dfs(&root, result, 0, len(nums)-1) {
			count += result
		}
	}

	return count, nil
}
