// Package main -
package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	args := os.Args
	if len(args) < 2 {
		fmt.Println("Error: not enough arguments\nFirst arg is day [1,2,3,...,25]\nSecond arg is part [1, 2]\nThird arg is for test data [t] - defaults to real data")
		os.Exit(1)
	}

	fileName := "input.txt"
	if len(args) == 4 && args[3] == "t" {
		fileName = "test.txt"
	}
	path := fmt.Sprintf("./day%v/%v", args[1], fileName)
	data, err := os.ReadFile(path)
	if err != nil {
		fmt.Printf("Error reading file: %v", err)
		os.Exit(1)
	}

	lines := strings.Split(string(data), "\n")

	commands := getCommands()
	commandName := fmt.Sprintf("%v-%v", args[1], args[2])

	if fn, exists := commands[commandName]; exists {
		ans, err := fn(lines)
		if err != nil {
			fmt.Printf("Error in part 1: %v", err)
			os.Exit(1)
		}

		fmt.Printf("Answer: %v\n", ans)
	}

	os.Exit(0)
}
