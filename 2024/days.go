package main

import (
	"github.com/mmammel12/adventOfCode/2024/day1"
	"github.com/mmammel12/adventOfCode/2024/day2"
	"github.com/mmammel12/adventOfCode/2024/day3"
	"github.com/mmammel12/adventOfCode/2024/day4"
)

func getCommands() map[string]func([]string) (int, error) {
	return map[string]func([]string) (int, error){
		"1-1": day1.Part1,
		"1-2": day1.Part2,
		"2-1": day2.Part1,
		"2-2": day2.Part2,
		"3-1": day3.Part1,
		"3-2": day3.Part2,
		"4-1": day4.Part1,
		"4-2": day4.Part2,
	}
}
