package solver

import "github.com/truggeri/go-sudoku/cmd/go-sudoku/puzzle"

func CreateTestPuzzleEasy() puzzle.Puzzle {
	var puz puzzle.Puzzle
	data := [puzzle.LineWidth][puzzle.LineWidth]int{
		{0, 3, 8, 2, 0, 0, 0, 4, 5},
		{0, 1, 0, 6, 0, 5, 0, 0, 0},
		{5, 0, 7, 0, 8, 0, 0, 0, 0},
		{3, 0, 1, 0, 6, 2, 0, 7, 0},
		{9, 0, 0, 5, 0, 7, 0, 0, 4},
		{0, 4, 0, 8, 3, 0, 6, 0, 2},
		{0, 0, 0, 0, 2, 0, 3, 0, 8},
		{0, 0, 0, 9, 0, 3, 0, 2, 0},
		{2, 7, 0, 0, 0, 8, 4, 9, 0}}

	for i, v := range data {
		for j, c := range v {
			puz[i][j] = puzzle.CreatePuzzleSquare(c)
		}
	}

	return puz
}

func CreateTestPuzzleMedium() puzzle.Puzzle {
	var puz puzzle.Puzzle
	data := [puzzle.LineWidth][puzzle.LineWidth]int{
		{0, 0, 3, 0, 0, 2, 8, 0, 0},
		{1, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 2, 0, 8, 6, 0, 0, 9, 0},
		{9, 0, 0, 4, 1, 0, 0, 0, 0},
		{4, 3, 0, 0, 5, 0, 0, 6, 8},
		{0, 0, 0, 0, 8, 6, 0, 0, 9},
		{0, 7, 0, 0, 2, 3, 0, 4, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 1},
		{0, 0, 5, 1, 0, 0, 9, 0, 0}}

	for i, v := range data {
		for j, c := range v {
			puz[i][j] = puzzle.CreatePuzzleSquare(c)
		}
	}

	return puz
}

func CreateTestPuzzleHard() puzzle.Puzzle {
	var puz puzzle.Puzzle
	data := [puzzle.LineWidth][puzzle.LineWidth]int{
		{7, 0, 2, 0, 3, 0, 0, 0, 0},
		{6, 0, 0, 1, 0, 2, 0, 8, 0},
		{0, 1, 0, 0, 0, 0, 4, 0, 0},
		{0, 0, 0, 0, 0, 0, 9, 6, 0},
		{0, 0, 6, 8, 4, 1, 5, 0, 0},
		{0, 3, 5, 0, 0, 0, 0, 0, 0},
		{0, 0, 9, 0, 0, 0, 0, 3, 0},
		{0, 5, 0, 3, 0, 7, 0, 0, 9},
		{0, 0, 0, 0, 2, 0, 6, 0, 8}}

	for i, v := range data {
		for j, c := range v {
			puz[i][j] = puzzle.CreatePuzzleSquare(c)
		}
	}

	return puz
}
