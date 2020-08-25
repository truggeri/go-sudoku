package main

import (
	"errors"
)

type solveTechnique struct {
	set   func(int, int) PuzzleSet
	index func(int, int) int
}

// Solve Solves a given puzzle
func Solve(puz Puzzle) Puzzle {
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
	set := func(x, y int) PuzzleSet {
		return puz.GetRow(x)
	}
	index := func(x, y int) int {
		return y
	}
	return solveByElement(puz, solveTechnique{set, index})
}

func solveByElement(puz Puzzle, st solveTechnique) (Puzzle, error) {
	for x := 0; x < lineWidth; x++ {
		for y := 0; y < lineWidth; y++ {
			if puz[x][y].solved() {
				continue
			}

			updated, result := solveSet(st.set(x, y), st.index(x, y))
			if updated {
				puz[x][y] = result
				return puz, nil
			}
		}
	}
	return puz, errors.New("No elements solved by row")
}

func solveSet(set PuzzleSet, i int) (bool, PuzzleSquare) {
	if set[i].solved() {
		return false, set[i]
	}

	setOptions := set.Possibilities(i)
	for j, possInSet := range setOptions {
		if set[i].possibilities[j] && !possInSet {
			set[i].value = j + 1
			return true, set[i]
		}
	}

	return false, set[i]
}

func solveColumns(puz Puzzle) (Puzzle, error) {
	set := func(x, y int) PuzzleSet {
		return puz.GetColumn(y)
	}
	index := func(x, y int) int {
		return x
	}
	return solveByElement(puz, solveTechnique{set, index})
}

func solveCubes(puz Puzzle) (Puzzle, error) {
	set := func(x, y int) PuzzleSet {
		return puz.GetCube(x, y)
	}
	index := func(x, y int) int {
		return puz.getCubeIndex(x, y)
	}
	return solveByElement(puz, solveTechnique{set, index})
}
