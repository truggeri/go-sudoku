package solver

import "github.com/truggeri/go-sudoku/cmd/go-sudoku/puzzle"

type solveTechnique struct {
	set   func(int, int) puzzle.Set
	index func(int, int) int
}

type solution struct {
	x, y   int
	square puzzle.Square
}

// Solve Returns the given puzzle with all elements solved
func Solve(puz puzzle.Puzzle) puzzle.Puzzle {
	for true {
		puz = puz.CalculatePossibilities()

		newSolution := false
		var answer solution
		techniques := [4]func(puzzle.Puzzle) (bool, solution){solveUniques, solveRows, solveColumns, solveCubes}
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

func solveUniques(puz puzzle.Puzzle) (bool, solution) {
	for x := 0; x < puzzle.LineWidth; x++ {
		for y := 0; y < puzzle.LineWidth; y++ {
			if puz[x][y].Solved() {
				continue
			}

			onlyOne, value := onlyOnePossibility(puz[x][y].Possibilities)
			if onlyOne {
				return true, solution{x: x, y: y, square: puzzle.CreatePuzzleSquare(value)}
			}
		}
	}
	return false, solution{}
}

func onlyOnePossibility(poss puzzle.Possibilies) (bool, int) {
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

func solveRows(puz puzzle.Puzzle) (bool, solution) {
	set := func(x, y int) puzzle.Set {
		return puz.GetRow(x)
	}
	index := func(x, y int) int {
		return y
	}
	return solveByElement(puz, solveTechnique{set, index})
}

func solveByElement(puz puzzle.Puzzle, st solveTechnique) (bool, solution) {
	for x := 0; x < puzzle.LineWidth; x++ {
		for y := 0; y < puzzle.LineWidth; y++ {
			if puz[x][y].Solved() {
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

func solveSet(set puzzle.Set, i int) (bool, puzzle.Square) {
	if set[i].Solved() {
		return false, set[i]
	}

	setOptions := set.Possibilities(i)
	for j := range setOptions {
		if set[i].Possibilities[j] && !setOptions[j] {
			return true, puzzle.CreatePuzzleSquare(j + 1)
		}
	}

	return false, set[i]
}

func solveColumns(puz puzzle.Puzzle) (bool, solution) {
	set := func(x, y int) puzzle.Set {
		return puz.GetColumn(y)
	}
	index := func(x, y int) int {
		return x
	}
	return solveByElement(puz, solveTechnique{set, index})
}

func solveCubes(puz puzzle.Puzzle) (bool, solution) {
	set := func(x, y int) puzzle.Set {
		return puz.GetCube(x, y)
	}
	index := puzzle.PositionInCube
	return solveByElement(puz, solveTechnique{set, index})
}
