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
		puz = puz.calculatePossibilities()

		newSolution := false
		var answer solution
		techniques := [4]func(Puzzle) (bool, solution){solveUniques, solveRows, solveColumns, solveCubes}
		for _, technique := range techniques {
			newSolution, answer = technique(puz)
			if newSolution {
				puz[answer.x][answer.y] = answer.square
				break
			}
		}

		if !newSolution {
			break
		}
	}
	return puz
}

func solveUniques(puz Puzzle) (bool, solution) {
	for x := 0; x < lineWidth; x++ {
		for y := 0; y < lineWidth; y++ {
			if puz[x][y].solved() {
				continue
			}

			onlyOne, value := onlyOnePossibility(puz[x][y].possibilities)
			if onlyOne {
				return true, solution{x: x, y: y, square: createPuzzleSquare(value)}
			}
		}
	}
	return false, solution{}
}

func onlyOnePossibility(poss possibilies) (bool, int) {
	counter, onlyOption := 0, 0

	for i, v := range poss {
		if v {
			onlyOption = i + 1
			counter++
		}

		if counter > 1 {
			return false, 0
		}
	}
	return true, onlyOption
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
	index := positionInCube
	return solveByElement(puz, solveTechnique{set, index})
}
