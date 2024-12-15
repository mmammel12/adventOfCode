// Package utils -
package utils

import "strings"

// SplitLines -
func SplitLines(lines []string, sep string) [][]string {
	splitLines := make([][]string, len(lines))
	for i, line := range lines {
		splitLines[i] = strings.Split(line, sep)
	}

	return splitLines
}

// Location -
type Location struct {
	Row int
	Col int
}

// Direction -
type Direction struct {
	Row        int
	Col        int
	IsBoundary func(Location) bool
}

// BuildDirections -
func BuildDirections(lines [][]string) (map[string]Direction, map[string]Direction, map[string]Direction) {
	allDirections := make(map[string]Direction)
	cardinalDirecations := make(map[string]Direction)
	diagonalDirections := make(map[string]Direction)

	cardinalDirecations["N"] = Direction{Row: -1, Col: 0, IsBoundary: func(l Location) bool {
		return l.Row-1 < 0
	}}
	cardinalDirecations["S"] = Direction{Row: 1, Col: 0, IsBoundary: func(l Location) bool {
		return l.Row+1 >= len(lines)
	}}
	cardinalDirecations["W"] = Direction{Row: 0, Col: -1, IsBoundary: func(l Location) bool {
		return l.Col-1 < 0
	}}
	cardinalDirecations["E"] = Direction{Row: 0, Col: 1, IsBoundary: func(l Location) bool {
		return l.Col+1 >= len(lines)
	}}

	diagonalDirections["NW"] = Direction{Row: -1, Col: -1, IsBoundary: func(l Location) bool {
		return l.Row-1 < 0 || l.Col-1 < 0
	}}
	diagonalDirections["NE"] = Direction{Row: -1, Col: 1, IsBoundary: func(l Location) bool {
		return l.Row-1 < 0 || l.Col+1 >= len(lines)
	}}
	diagonalDirections["SW"] = Direction{Row: 1, Col: -1, IsBoundary: func(l Location) bool {
		return l.Row+1 >= len(lines) || l.Col-1 < 0
	}}
	diagonalDirections["SE"] = Direction{Row: 1, Col: 1, IsBoundary: func(l Location) bool {
		return l.Row+1 >= len(lines) || l.Col+1 >= len(lines)
	}}

	for k, v := range cardinalDirecations {
		allDirections[k] = v
	}
	for k, v := range diagonalDirections {
		allDirections[k] = v
	}

	return allDirections, cardinalDirecations, diagonalDirections
}
