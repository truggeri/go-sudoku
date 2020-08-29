package puzzle

import (
	"fmt"
	"strings"
)

// LineWidth the number of cells in each row, column and cube
const LineWidth = 9
const cubeWidth = LineWidth / 3

// Puzzle Representation of a Sudoku puzzle
type Puzzle [LineWidth][LineWidth]PuzzleSquare

// GetRow returns a given row
func (p Puzzle) GetRow(i int) PuzzleSet {
	return p[i]
}

// GetColumn returns a column by given index
func (p Puzzle) GetColumn(i int) PuzzleSet {
	var column PuzzleSet
	for x := 0; x < LineWidth; x++ {
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

// CalculatePossibilities Gives puzzle with updated possibility values
func (p Puzzle) CalculatePossibilities() Puzzle {
	var possByRow [LineWidth]Possibilies
	for row := 0; row < LineWidth; row++ {
		possByRow[row] = findElementPossbilities(p.GetRow(row))
	}

	var possByColumn [LineWidth]Possibilies
	for col := 0; col < LineWidth; col++ {
		possByColumn[col] = findElementPossbilities(p.GetColumn(col))
	}

	var possByCube [LineWidth]Possibilies
	for cube := 0; cube < LineWidth; cube++ {
		possByCube[cube] = findElementPossbilities(p.GetCube(cubePosition(cube)))
	}

	for x := 0; x < LineWidth; x++ {
		for y := 0; y < LineWidth; y++ {
			if !p[x][y].Solved() {
				p[x][y].Possibilities = mergePoss(possByRow[x], possByColumn[y], possByCube[cubeNumber(x, y)])
			}
		}
	}

	return p
}

func cubePosition(cubeNumber int) (int, int) {
	return (cubeNumber * cubeWidth) % LineWidth, (cubeNumber / cubeWidth) * cubeWidth
}

func cubeNumber(x, y int) int {
	return (x / cubeWidth) + (y/cubeWidth)*cubeWidth
}

func findElementPossbilities(elements PuzzleSet) Possibilies {
	poss := Possibilies{true, true, true, true, true, true, true, true, true}

	for _, element := range elements {
		if element.Solved() {
			poss[element.Value-1] = false
		}
	}

	return poss
}

// PositionInCube Given an x,y, gives the position in a set
func PositionInCube(x, y int) int {
	return (x%cubeWidth)*cubeWidth + (y % cubeWidth)
}

func mergePoss(row, col, cube Possibilies) Possibilies {
	var result Possibilies
	for i := range result {
		result[i] = row[i] && col[i] && cube[i]
	}
	return result
}

func (p Puzzle) String() string {
	var output []string

	for _, r := range p {
		var line []string
		for _, v := range r {
			if v.Solved() {
				line = append(line, fmt.Sprintf("%v", v.Value))
			} else {
				line = append(line, "-")
			}
		}
		output = append(output, strings.Join(line[:], " "))
	}

	return strings.Join(output[:], "\n")
}
