package day14

import (
	"fmt"
	"strings"

	"github.com/mmammel12/adventOfCode/2024/utils"
)

func iterate(p, v utils.Point) utils.Point {
	x := (p.X + v.X) % width
	y := (p.Y + v.Y) % height
	return utils.Point{X: abs(x, true), Y: abs(y, false)}
}

func draw(points []utils.Point) {
	lines := make([][]string, 0)
	for i := 0; i < height; i++ {
		lines = append(lines, make([]string, 0))
		for j := 0; j < width; j++ {
			lines[i] = append(lines[i], " ")
		}
	}

	for _, p := range points {
		lines[p.Y][p.X] = "X"
	}

	for _, line := range lines {
		fmt.Println(strings.Join(line, ""))
	}
	fmt.Println("~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~")
}

func calculateSafetyFactor(points []utils.Point) int {
	quadrants := make([]int, 4)
	for _, p := range points {
		q := getQuadrant(p)
		if q != -1 {
			quadrants[q]++
		}
	}

	safety := 1
	for _, q := range quadrants {
		safety *= q
	}
	return safety
}

// Part2 -
func Part2(lines []string) (int, error) {

	points := make([]utils.Point, 0)
	velocities := make([]utils.Point, 0)
	for _, line := range lines {
		p, v := getLineInfo(line)
		points = append(points, p)
		velocities = append(velocities, v)
	}

	minSafetyFactor := 231019008
	ans := 0
	for i := 0; i < width*height; i++ {
		for j, p := range points {
			points[j] = iterate(p, velocities[j])
		}
		safetyFactor := calculateSafetyFactor(points)
		if safetyFactor < minSafetyFactor {
			minSafetyFactor = safetyFactor
			ans = i + 1
			draw(points)
		}
	}

	return ans, nil
}
