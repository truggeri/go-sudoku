package main

import (
	"errors"
	"fmt"
)

func solve(puz Puzzle) Puzzle {
	var updatedFlag bool = true
	// var err error

	for updatedFlag == true {
		puz2, err := solveUniques(puz)

		if puz2 != puz {
			fmt.Println("Puzzle has been updated")
		} else {
			fmt.Println("Puzzle has NOT been updated")
		}

		puz = puz2

		if err == nil {
			continue
		}

		//...

		if err != nil {
			break
		}
	}

	return puz
}

func solveUniques(puz Puzzle) (Puzzle, error) {
	puz = findAllPossibilites(puz)
	updatedFlag := false

	for i := 0; i < lineWidth; i++ {
		for j := 0; j < lineWidth; j++ {
			if !(puz[i][j].solved()) && hasOnlyOne(puz[i][j].possibilities) {
				puz[i][j].value = findOnlyValue(puz[i][j].possibilities)
				updatedFlag = true
			}
		}
	}

	if !updatedFlag {
		return puz, errors.New("No elements solved")
	}

	return puz, nil
}

func findAllPossibilites(puz Puzzle) Puzzle {
	for i := 0; i < lineWidth; i++ {
		for j := 0; j < lineWidth; j++ {
			if !puz[i][j].solved() {
				puz[i][j].possibilities = findPossibilities(puz, i, j)
			}
		}
	}
	return puz
}

func findPossibilities(puz Puzzle, x, y int) [lineWidth]bool {
	var row [lineWidth]PuzzleSquare
	for i := 0; i < lineWidth; i++ {
		row[i] = puz[i][y]
	}
	rowPoss := findElementPossbilities(row, x)
	colPoss := findElementPossbilities(puz[x], y)

	xOffset := x / cubeWidth
	yOffset := y / cubeWidth
	var cube [cubeWidth][cubeWidth]PuzzleSquare
	for i, v := range cube {
		for j := range v {
			cube[i][j] = puz[i+(xOffset*cubeWidth)][j+(yOffset*cubeWidth)]
		}
	}
	cubePoss := findSqrPossbilities(cube, x%cubeWidth, y%cubeWidth)

	return mergePoss(rowPoss, colPoss, cubePoss)
}

func findElementPossbilities(elements [lineWidth]PuzzleSquare, x int) [lineWidth]bool {
	poss := [lineWidth]bool{true, true, true, true, true, true, true, true, true}

	for i := range elements {
		if i == x {
			continue
		}

		if elements[i].solved() {
			poss[elements[i].value-1] = false
		}
	}

	return poss
}

func findSqrPossbilities(cube [cubeWidth][cubeWidth]PuzzleSquare, x, y int) [lineWidth]bool {
	poss := [lineWidth]bool{true, true, true, true, true, true, true, true, true}

	for i, v := range cube {
		for j := range v {
			if i == x && j == y {
				continue
			}

			if cube[i][j].solved() {
				poss[cube[i][j].value-1] = false
			}
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

func hasOnlyOne(poss [lineWidth]bool) bool {
	counter := 0
	for _, v := range poss {
		if v == true {
			counter++
		}

		if counter > 1 {
			return false
		}
	}
	return true
}

func findOnlyValue(poss [lineWidth]bool) int {
	for i, v := range poss {
		if v == true {
			return i + 1
		}
	}
	return 0
}
