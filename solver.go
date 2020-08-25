package main

type solveTechnique struct {
	set   func(int, int) PuzzleSet
	index func(int, int) int
}

type solution struct {
	x, y   int
	square PuzzleSquare
}

// Solve Returns the given puzzle with all elements solved
func Solve(puz Puzzle) Puzzle {
	for true {
		puz = puz.UpdatePossibilites()

		solved, solution := solveUniques(puz)
		if solved {
			puz[solution.x][solution.y] = solution.square
			continue
		}

		solved, solution = solveRows(puz)
		if solved {
			puz[solution.x][solution.y] = solution.square
			continue
		}

		solved, solution = solveColumns(puz)
		if solved {
			puz[solution.x][solution.y] = solution.square
			continue
		}

		solved, solution = solveCubes(puz)
		if solved {
			puz[solution.x][solution.y] = solution.square
			continue
		}
		break
	}

	return puz
}

func solveUniques(puz Puzzle) (bool, solution) {
	for x := 0; x < lineWidth; x++ {
		for y := 0; y < lineWidth; y++ {
			if puz[x][y].solved() {
				continue
			}

			if hasOnlyOne(puz[x][y].possibilities) {
				return true, solution{x: x, y: y, square: createPuzzleSquare(findOnlyValue(puz[x][y].possibilities))}
			}
		}
	}
	return false, solution{}
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

func solveRows(puz Puzzle) (bool, solution) {
	set := func(x, y int) PuzzleSet {
		return puz.GetRow(x)
	}
	index := func(x, y int) int {
		return y
	}
	return solveByElement(puz, solveTechnique{set, index})
}

func solveByElement(puz Puzzle, st solveTechnique) (bool, solution) {
	for x := 0; x < lineWidth; x++ {
		for y := 0; y < lineWidth; y++ {
			if puz[x][y].solved() {
				continue
			}

			updated, result := solveSet(st.set(x, y), st.index(x, y))
			if updated {
				return true, solution{x: x, y: y, square: result}
			}
		}
	}
	return false, solution{}
}

func solveSet(set PuzzleSet, i int) (bool, PuzzleSquare) {
	if set[i].solved() {
		return false, set[i]
	}

	setOptions := set.Possibilities(i)
	for j := range setOptions {
		if set[i].possibilities[j] && !setOptions[j] {
			return true, createPuzzleSquare(j + 1)
		}
	}

	return false, set[i]
}

func solveColumns(puz Puzzle) (bool, solution) {
	set := func(x, y int) PuzzleSet {
		return puz.GetColumn(y)
	}
	index := func(x, y int) int {
		return x
	}
	return solveByElement(puz, solveTechnique{set, index})
}

func solveCubes(puz Puzzle) (bool, solution) {
	set := func(x, y int) PuzzleSet {
		return puz.GetCube(x, y)
	}
	index := func(x, y int) int {
		return puz.getCubeIndex(x, y)
	}
	return solveByElement(puz, solveTechnique{set, index})
}
