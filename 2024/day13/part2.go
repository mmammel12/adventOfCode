// Package day13 -
package day13

import "github.com/mmammel12/adventOfCode/2024/utils"

func processMachine(a, b, target utils.Point) int {
	axMult := a.X * b.Y
	ayMult := -(a.Y * b.X)
	targetxMult := target.X * b.Y
	targetyMult := -(target.Y * b.X)

	aMult := axMult + ayMult
	targetMult := targetxMult + targetyMult

	aPresses := targetMult / aMult

	if targetMult%aPresses != 0 {
		return -1
	}

	if (target.X-(a.X*aPresses))%b.X != 0 {
		return -1
	}

	bPresses := (target.X - (a.X * aPresses)) / b.X

	return aPresses*3 + bPresses
}

// Part2 -
func Part2(lines []string) (int, error) {
	ans := 0
	for i := 0; i < len(lines); i += 4 {
		a, b, target := getInputs(lines[i:i+3], 10_000_000_000_000)

		presses := processMachine(a, b, target)
		if presses != -1 {
			ans += presses
		}
	}

	return ans, nil
}
