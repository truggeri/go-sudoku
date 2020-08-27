package main

func CreateTestPuzzleEasy() Puzzle {
	var puzzle Puzzle
	data := [lineWidth][lineWidth]int{
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
			puzzle[i][j] = createPuzzleSquare(c)
		}
	}

	return puzzle
}

func CreateTestPuzzleMedium() Puzzle {
	var puzzle Puzzle
	data := [lineWidth][lineWidth]int{
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
			puzzle[i][j] = createPuzzleSquare(c)
		}
	}

	return puzzle
}
