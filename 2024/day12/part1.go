// Package day12 -
package day12

import "github.com/mmammel12/adventOfCode/2024/utils"

type region struct {
	area      int
	perimiter int
	locations map[utils.Location]map[string]bool
}

var splitLines [][]string
var directions = map[string]utils.Direction{}

func buildRegion(visited map[utils.Location]bool, current utils.Location, reg *region) {
	char := splitLines[current.Row][current.Col]
	visited[current] = true
	reg.area++

	for k, v := range directions {
		if v.IsBoundary(current) || splitLines[current.Row+v.Row][current.Col+v.Col] != char {
			reg.perimiter++
			if _, exists := reg.locations[current]; !exists {
				reg.locations[current] = make(map[string]bool)
			}
			reg.locations[current][k] = true

		} else if !v.IsBoundary(current) {
			if _, exists := reg.locations[current]; !exists {
				reg.locations[current] = make(map[string]bool)
			}
			reg.locations[current][k] = false
			if _, exists := visited[utils.Location{Row: current.Row + v.Row, Col: current.Col + v.Col}]; !exists && splitLines[current.Row+v.Row][current.Col+v.Col] == char {
				buildRegion(visited, utils.Location{Row: current.Row + v.Row, Col: current.Col + v.Col}, reg)
			}
		}
	}
}

// Part1 -
func Part1(lines []string) (int, error) {
	splitLines = utils.SplitLines(lines, "")
	_, directions, _ = utils.BuildDirections(splitLines)

	visited := make(map[utils.Location]bool)
	regions := make([]region, 0)
	for row, line := range splitLines {
		for col := range line {
			current := utils.Location{Row: row, Col: col}
			if _, exists := visited[current]; !exists {
				reg := region{0, 0, make(map[utils.Location]map[string]bool)}
				buildRegion(visited, current, &reg)
				regions = append(regions, reg)
			}
		}
	}

	ans := 0
	for _, reg := range regions {
		ans += reg.area * reg.perimiter
	}

	return ans, nil
}
