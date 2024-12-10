// Package day8 -
package day8

// Part2 -
func Part2(lines []string) (int, error) {
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
				// each antenna is also an antinode
				if _, exists := antinodes[antennaCoords[i]]; !exists {
					antinodes[antennaCoords[i]] = true
					count++
				}
				if _, exists := antinodes[antennaCoords[j]]; !exists {
					antinodes[antennaCoords[j]] = true
					count++
				}

				absRowDiff := max(antennaCoords[i].row, antennaCoords[j].row) - min(antennaCoords[i].row, antennaCoords[j].row)
				colDiff := antennaCoords[i].col - antennaCoords[j].col

				firstAntinode := coord{antennaCoords[i].row - absRowDiff, antennaCoords[i].col + colDiff}
				secondAntinode := coord{antennaCoords[j].row + absRowDiff, antennaCoords[j].col - colDiff}

				for coordIsValid(firstAntinode) {
					if _, exists := antinodes[firstAntinode]; !exists {
						antinodes[firstAntinode] = true
						count++
					}
					firstAntinode = coord{firstAntinode.row - absRowDiff, firstAntinode.col + colDiff}
				}

				for coordIsValid(secondAntinode) {
					if _, exists := antinodes[secondAntinode]; !exists {
						antinodes[secondAntinode] = true
						count++
					}
					secondAntinode = coord{secondAntinode.row + absRowDiff, secondAntinode.col - colDiff}
				}
			}
		}
	}

	return count, nil
}
