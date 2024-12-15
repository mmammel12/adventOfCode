package day12

import (
	"github.com/mmammel12/adventOfCode/2024/utils"
)

var diagonalDirections = map[string]utils.Direction{}

func countCorners(locations map[utils.Location]map[string]bool) int {
	diagonalStrings := []string{"NW", "NE", "SW", "SE"}
	visited := make(map[utils.Location]bool)
	count := 0
	for k, v := range locations {
		if _, exists := visited[k]; exists {
			continue
		}

		visited[k] = true
		for _, checks := range diagonalStrings {
			side1IsEdge := v[string(checks[0])]
			side2IsEdge := v[string(checks[1])]
			if side1IsEdge && side2IsEdge {
				// inner corner
				count++
			} else if !side1IsEdge && !side2IsEdge {
				if _, exists := locations[utils.Location{Row: k.Row + diagonalDirections[checks].Row, Col: k.Col + diagonalDirections[checks].Col}]; !exists {
					// outer corner
					count++
				}
			}
		}
	}

	return count
}

// Part2 -
func Part2(lines []string) (int, error) {
	splitLines = utils.SplitLines(lines, "")
	_, directions, diagonalDirections = utils.BuildDirections(splitLines)

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
		sides := countCorners(reg.locations)
		ans += reg.area * sides
	}

	return ans, nil
}
