package main

import (
	"fmt"
	"strings"
)

const lineWidth = 9
const cubeWidth = lineWidth / 3

// Puzzle Representation of a Sudoku puzzle
type Puzzle [lineWidth][lineWidth]PuzzleSquare

func (p Puzzle) getCube(x, y int) [lineWidth]PuzzleSquare {
	xOffset := x / cubeWidth
	yOffset := y / cubeWidth

	var result [lineWidth]PuzzleSquare
	xIndex := xOffset * cubeWidth
	yIndex := yOffset * cubeWidth
	copy(result[0:3], p[xIndex][yIndex:yIndex+3])
	copy(result[3:6], p[1+xIndex][yIndex:yIndex+3])
	copy(result[6:9], p[2+xIndex][yIndex:yIndex+3])
	return result
}

func (p Puzzle) getCubeIndex(x, y int) int {
	return (x%cubeWidth)*cubeWidth + (y % cubeWidth)
}

func (p Puzzle) String() string {
	var output []string

	for _, r := range p {
		var line []string
		for _, v := range r {
			if v.solved() {
				line = append(line, fmt.Sprintf("%v", v.value))
			} else {
				line = append(line, "-")
			}
		}
		output = append(output, strings.Join(line[:], " "))
	}

	return strings.Join(output[:], "\n")
}
