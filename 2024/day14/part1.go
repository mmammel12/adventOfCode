// Package day14 -
package day14

import (
	"strconv"
	"strings"

	"github.com/mmammel12/adventOfCode/2024/utils"
)

const width = 101
const height = 103

func strToPoint(str string) utils.Point {
	split := strings.Split(str, ",")
	x, _ := strconv.Atoi(split[0])
	y, _ := strconv.Atoi(split[1])

	return utils.Point{X: x, Y: y}
}

func getLineInfo(line string) (utils.Point, utils.Point) {
	parts := strings.Split(line, " ")
	p := strToPoint(strings.Replace(parts[0], "p=", "", 1))
	v := strToPoint(strings.Replace(parts[1], "v=", "", 1))

	return p, v
}

func getQuadrant(p utils.Point) int {
	if p.X == width/2 || p.Y == height/2 {
		return -1
	}

	if p.X < width/2 {
		if p.Y < height/2 {
			return 0
		}
		return 2
	}

	if p.Y < height/2 {
		return 1
	}

	return 3
}

func abs(num int, x bool) int {
	if num < 0 {
		if x {
			return width + num
		}
		return height + num
	}

	return num
}

// Part1 -
func Part1(lines []string) (int, error) {

	quadrants := make([]int, 4)
	for _, line := range lines {
		p, v := getLineInfo(line)
		x := (p.X + 100*v.X) % width
		y := (p.Y + 100*v.Y) % height
		final := utils.Point{X: abs(x, true), Y: abs(y, false)}
		q := getQuadrant(final)
		if q != -1 {
			quadrants[q]++
		}
	}

	ans := 1
	for _, q := range quadrants {
		ans *= q
	}

	return ans, nil
}
