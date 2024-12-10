// Package day8 -
package day8

type coord struct {
	row int
	col int
}

func createCoordIsValid(rowLen, colLen int) func(coord) bool {
	return func(checkCoord coord) bool {
		return checkCoord.row >= 0 && checkCoord.row < rowLen && checkCoord.col >= 0 && checkCoord.col < colLen
	}
}

// Part1 -
func Part1(lines []string) (int, error) {
	antennas := make(map[rune][]coord)
	for row, line := range lines {
		for col, char := range line {
			if char != '.' {
				if _, exists := antennas[char]; !exists {
					antennas[char] = make([]coord, 0)
				}

				antennas[char] = append(antennas[char], coord{row, col})
			}
		}
	}

	antinodes := make(map[coord]bool)
	rowLen := len(lines)
	colLen := len(lines[0])
	coordIsValid := createCoordIsValid(rowLen, colLen)
	count := 0

	for _, antennaCoords := range antennas {
		if len(antennaCoords) <= 1 {
			continue
		}

		for i := 0; i < len(antennaCoords)-1; i++ {
			for j := i + 1; j < len(antennaCoords); j++ {
				absRowDiff := max(antennaCoords[i].row, antennaCoords[j].row) - min(antennaCoords[i].row, antennaCoords[j].row)
				colDiff := antennaCoords[i].col - antennaCoords[j].col

				firstAntinode := coord{antennaCoords[i].row - absRowDiff, antennaCoords[i].col + colDiff}
				secondAntinode := coord{antennaCoords[j].row + absRowDiff, antennaCoords[j].col - colDiff}

				if _, exists := antinodes[firstAntinode]; !exists && coordIsValid(firstAntinode) {
					antinodes[firstAntinode] = true
					count++
				}
				if _, exists := antinodes[secondAntinode]; !exists && coordIsValid(secondAntinode) {
					antinodes[secondAntinode] = true
					count++
				}
			}
		}
	}

	return count, nil
}
