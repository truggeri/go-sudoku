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

func (p Puzzle) getCubeIndex(x, y int) int {
	return (x%cubeWidth)*cubeWidth + (y % cubeWidth)
}

// UpdatePossibilites Updates internal possibilities
func (p Puzzle) UpdatePossibilites() Puzzle {
	for i := 0; i < lineWidth; i++ {
		for j := 0; j < lineWidth; j++ {
			if !p[i][j].solved() {
				p[i][j].possibilities = findPossibilities(p, i, j)
			}
		}
	}
	return p
}

func findPossibilities(puz Puzzle, x, y int) [lineWidth]bool {
	rowPoss := findElementPossbilities(puz.GetRow(x), x)
	colPoss := findElementPossbilities(puz.GetColumn(y), y)
	cubePoss := findElementPossbilities(puz.GetCube(x, y), puz.getCubeIndex(x, y))

	return mergePoss(rowPoss, colPoss, cubePoss)
}

func findElementPossbilities(elements PuzzleSet, exclude int) [lineWidth]bool {
	poss := [lineWidth]bool{true, true, true, true, true, true, true, true, true}

	for i := range elements {
		if i == exclude {
			continue
		}

		if elements[i].solved() {
			poss[elements[i].value-1] = false
		}
	}

	return poss
}

func mergePoss(row, col, cube [lineWidth]bool) [lineWidth]bool {
	var poss [lineWidth]bool

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
