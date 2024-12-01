// Package main -
package main

import (
	"fmt"
	"os"

	"github.com/mmammel12/adventOfCode/2024/day1"
)

func main() {
	data, err := os.ReadFile("./day1/input.txt")
	if err != nil {
		fmt.Printf("Error reading file: %v", err)
		os.Exit(1)
	}

	err = day1.Part2(string(data))
	if err != nil {
		fmt.Printf("Error in part 1: %v", err)
		os.Exit(1)
	}

	os.Exit(0)
}
