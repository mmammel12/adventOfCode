package main

import (
	"github.com/mmammel12/adventOfCode/2024/day1"
	"github.com/mmammel12/adventOfCode/2024/day10"
	"github.com/mmammel12/adventOfCode/2024/day11"
	"github.com/mmammel12/adventOfCode/2024/day12"
	"github.com/mmammel12/adventOfCode/2024/day2"
	"github.com/mmammel12/adventOfCode/2024/day3"
	"github.com/mmammel12/adventOfCode/2024/day4"
	"github.com/mmammel12/adventOfCode/2024/day5"
	"github.com/mmammel12/adventOfCode/2024/day6"
	"github.com/mmammel12/adventOfCode/2024/day7"
	"github.com/mmammel12/adventOfCode/2024/day8"
	"github.com/mmammel12/adventOfCode/2024/day9"
)

func getCommands() map[string]func([]string) (int, error) {
	return map[string]func([]string) (int, error){
		"1-1":  day1.Part1,
		"1-2":  day1.Part2,
		"2-1":  day2.Part1,
		"2-2":  day2.Part2,
		"3-1":  day3.Part1,
		"3-2":  day3.Part2,
		"4-1":  day4.Part1,
		"4-2":  day4.Part2,
		"5-1":  day5.Part1,
		"5-2":  day5.Part2,
		"6-1":  day6.Part1,
		"6-2":  day6.Part2,
		"7-1":  day7.Part1,
		"7-2":  day7.Part2,
		"8-1":  day8.Part1,
		"8-2":  day8.Part2,
		"9-1":  day9.Part1,
		"9-2":  day9.Part2,
		"10-1": day10.Part1,
		"10-2": day10.Part2,
		"11-1": day11.Part1,
		"11-2": day11.Part2,
		"12-1": day12.Part1,
		"12-2": day12.Part2,
	}
}
