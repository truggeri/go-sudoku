package main

import (
	"fmt"
	"strings"
)

const lineWidth = 9
const cubeWidth = lineWidth / 3

// Puzzle Representation of a Sudoku puzzle
type Puzzle [lineWidth][lineWidth]PuzzleSquare

// GetRow returns a given row
func (p Puzzle) GetRow(i int) PuzzleSet {
	return p[i]
}

// GetColumn returns a column by given index
func (p Puzzle) GetColumn(i int) PuzzleSet {
	var column PuzzleSet
	for x := 0; x < lineWidth; x++ {
		column[x] = p[x][i]
	}
	return column
}

// GetCube returns squares that form a cube around given point
func (p Puzzle) GetCube(x, y int) PuzzleSet {
	xOffset := x / cubeWidth
	yOffset := y / cubeWidth

	var result PuzzleSet
	xIndex := xOffset * cubeWidth
	yIndex := yOffset * cubeWidth
	copy(result[0:3], p[xIndex][yIndex:yIndex+3])
	copy(result[3:6], p[1+xIndex][yIndex:yIndex+3])
	copy(result[6:9], p[2+xIndex][yIndex:yIndex+3])
	return result
}

func (p Puzzle) calculatePossibilities() Puzzle {
	var possByRow [lineWidth]possibilies
	for row := 0; row < lineWidth; row++ {
		possByRow[row] = findElementPossbilities(p.GetRow(row))
	}

	var possByColumn [lineWidth]possibilies
	for col := 0; col < lineWidth; col++ {
		possByColumn[col] = findElementPossbilities(p.GetColumn(col))
	}

	var possByCube [lineWidth]possibilies
	for cube := 0; cube < lineWidth; cube++ {
		possByCube[cube] = findElementPossbilities(p.GetCube(cubePosition(cube)))
	}

	for x := 0; x < lineWidth; x++ {
		for y := 0; y < lineWidth; y++ {
			if !p[x][y].solved() {
				p[x][y].possibilities = mergePoss(possByRow[x], possByColumn[y], possByCube[cubeNumber(x, y)])
			}
		}
	}

	return p
}

func cubePosition(cubeNumber int) (int, int) {
	return (cubeNumber * cubeWidth) % lineWidth, (cubeNumber / cubeWidth) * cubeWidth
}

func cubeNumber(x, y int) int {
	return (x / cubeWidth) + (y/cubeWidth)*cubeWidth
}

func findElementPossbilities(elements PuzzleSet) possibilies {
	poss := possibilies{true, true, true, true, true, true, true, true, true}

	for _, element := range elements {
		if element.solved() {
			poss[element.value-1] = false
		}
	}

	return poss
}

func positionInCube(x, y int) int {
	return (x%cubeWidth)*cubeWidth + (y % cubeWidth)
}

func mergePoss(row, col, cube possibilies) possibilies {
	var poss possibilies
	for i := range poss {
		poss[i] = row[i] && col[i] && cube[i]
	}
	return poss
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
