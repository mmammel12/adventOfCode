// Package day4 -
package day4

// Part2 -
func Part2(lines []string) (int, error) {
	count := 0
	for row := 1; row < len(lines)-1; row++ {
		for col := 1; col < len(lines[row])-1; col++ {
			if lines[row][col] == 'A' {
				topLeft := lines[row-1][col-1]
				topRight := lines[row-1][col+1]
				bottomLeft := lines[row+1][col-1]
				bottomRight := lines[row+1][col+1]

				if topLeft != 'M' && topLeft != 'S' {
					continue
				}
				if topRight != 'M' && topRight != 'S' {
					continue
				}

				if topLeft == 'M' && bottomRight != 'S' {
					continue
				} else if topLeft == 'S' && bottomRight != 'M' {
					continue
				}

				if topRight == 'M' && bottomLeft != 'S' {
					continue
				} else if topRight == 'S' && bottomLeft != 'M' {
					continue
				}
				count++
			}
		}
	}

	return count, nil
}
