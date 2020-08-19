package main

import "testing"

func TestPrint(t *testing.T) {
	var puzzle Puzzle
	data := [lineWidth][lineWidth]int{
		{1, 2, 3, 4, 5, 6, 7, 8, 9},
		{0, 2, 3, 4, 5, 6, 7, 8, 0},
		{0, 2, 3, 4, 5, 6, 0, 8, 9},
		{1, 2, 3, 4, 5, 6, 7, 8, 9},
		{1, 2, 3, 4, 5, 6, 7, 8, 9},
		{1, 2, 3, 4, 5, 6, 7, 8, 9},
		{1, 2, 3, 4, 5, 6, 7, 8, 9},
		{0, 0, 0, 0, 0, 0, 0, 0, 0},
		{1, 2, 3, 4, 5, 6, 7, 8, 9}}

	for i, v := range data {
		for j, c := range v {
			puzzle[i][j] = createPuzzleSquare(c)
		}
	}

	expectedString := `1 2 3 4 5 6 7 8 9
- 2 3 4 5 6 7 8 -
- 2 3 4 5 6 - 8 9
1 2 3 4 5 6 7 8 9
1 2 3 4 5 6 7 8 9
1 2 3 4 5 6 7 8 9
1 2 3 4 5 6 7 8 9
- - - - - - - - -
1 2 3 4 5 6 7 8 9`

	if puzzle.String() != expectedString {
		t.Errorf("Puzzle string not as expected, got: \n%s", puzzle)
	}
}
