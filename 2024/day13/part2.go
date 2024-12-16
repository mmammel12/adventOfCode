// Package day13 -
package day13

func extendedGCD(a, b int) (gcd, x, y int) {
	if a == 0 {
		return b, 0, 1
	}

	gcd, x1, y1 := extendedGCD(b%a, a)
	x = y1 - (b/a)*x1
	y = x1
	return gcd, x, y
}

func findSingleAxisSolutionsP2(aMove, bMove, target int) []int {
	gcd, x0, y0 := extendedGCD(aMove, bMove)

	if target%gcd != 0 {
		return make([]int, 0)
	}

	scale := target / gcd
	x0 *= scale
	y0 *= scale

	// find minimum non-negative solution
	// number of complete cycles to add/subtract to make both numbers non-negative
	cycles := max((-x0*gcd)/bMove, (y0*gcd)/aMove)

	aPresses := x0 + (bMove/gcd)*cycles
	bPresses := y0 - (aMove/gcd)*cycles

	return []int{aPresses, bPresses}
}

// Part2 -
func Part2(lines []string) (int, error) {
	ans := 0
	for i := 0; i < len(lines); i += 4 {
		a, b, target := getInputs(lines[i:i+3], 10000000000000)
		xSolution := findSingleAxisSolutionsP2(a.x, b.x, target.x)
		ySolution := findSingleAxisSolutionsP2(a.y, b.y, target.y)

		if len(xSolution) == 0 || len(ySolution) == 0 {
			continue
		}

		aPresses := max(xSolution[0], ySolution[0])

		additionalX := aPresses - xSolution[0]
		additionalY := aPresses - ySolution[0]

		xb := xSolution[1] - (a.x*additionalX)/b.x
		yb := ySolution[1] - (a.y*additionalY)/b.y

		if xb == yb && xb >= 0 {
			ans += aPresses*3 + xb
		}
	}

	return ans, nil
}
