package main

import (
	"errors"
)

func solve(puz Puzzle) Puzzle {
	for true {
		var err error

		puz = puz.UpdatePossibilites()
		puz, err = solveUniques(puz)
		if err == nil {
			continue
		}

		puz, err = solveRows(puz)
		if err == nil {
			continue
		}

		puz, err = solveColumns(puz)
		if err == nil {
			continue
		}

		puz, err = solveCubes(puz)
		if err == nil {
			continue
		}

		break
	}

	return puz
}

func solveUniques(puz Puzzle) (Puzzle, error) {
	updatedFlag := false

	for x := 0; x < lineWidth; x++ {
		for y := 0; y < lineWidth; y++ {
			if puz[x][y].solved() {
				continue
			}

			if hasOnlyOne(puz[x][y].possibilities) {
				puz[x][y].value = findOnlyValue(puz[x][y].possibilities)
				updatedFlag = true
			}
		}
	}

	if !updatedFlag {
		return puz, errors.New("No elements solved")
	}

	return puz, nil
}

func hasOnlyOne(poss [lineWidth]bool) bool {
	counter := 0
	for _, v := range poss {
		if v {
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
		if v {
			return i + 1
		}
	}
	return 0
}

func solveRows(puz Puzzle) (Puzzle, error) {
	updatedFlag := false
	var result PuzzleSet

	for x := 0; x < lineWidth; x++ {
		for y := 0; y < lineWidth; y++ {
			if puz[x][y].solved() {
				continue
			}

			updatedFlag, result = solveSet(puz.GetRow(x), y)
			if updatedFlag {
				puz[x][y] = result[y]
				return puz, nil
			}
		}
	}
	return puz, errors.New("No elements solved by row")
}

func solveSet(set PuzzleSet, i int) (bool, PuzzleSet) {
	if set[i].solved() {
		return false, set
	}

	updatedFlag := false
	setOptions := set.Possibilities(i)
	for j, possInSet := range setOptions {
		if set[i].possibilities[j] && !possInSet {
			updatedFlag = true
			set[i].value = j + 1
		}
	}

	return updatedFlag, set
}

func solveColumns(puz Puzzle) (Puzzle, error) {
	updatedFlag := false
	var result PuzzleSet

	for x := 0; x < lineWidth; x++ {
		for y := 0; y < lineWidth; y++ {
			if puz[x][y].solved() {
				continue
			}

			updatedFlag, result = solveSet(puz.GetColumn(y), x)
			if updatedFlag {
				puz[x][y] = result[x]
				return puz, nil
			}
		}
	}
	return puz, errors.New("No elements solved by column")
}

func solveCubes(puz Puzzle) (Puzzle, error) {
	return puz, errors.New("No elements solved")
}
