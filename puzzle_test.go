package main

import (
	"testing"
)

func CreateTestPuzzle() Puzzle {
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

func TestGetRow(t *testing.T) {
	puzzle := CreateTestPuzzle()
	row := puzzle.GetRow(0)
	expected := [lineWidth]int{0, 3, 8, 2, 0, 0, 0, 4, 5}

	for i := 0; i < lineWidth; i++ {
		if row[i].value != expected[i] {
			t.Errorf("GetRow value error at index: %d, expected: %d, got: %d", i, expected[i], row[i].value)
		}
	}
}

func TestGetColumn(t *testing.T) {
	puzzle := CreateTestPuzzle()
	column := puzzle.GetColumn(3)
	expected := [lineWidth]int{2, 6, 0, 0, 5, 8, 0, 9, 0}

	for i := 0; i < lineWidth; i++ {
		if column[i].value != expected[i] {
			t.Errorf("GetColumn value error at index: %d, expected: %d, got: %d", i, expected[i], column[i].value)
		}
	}
}

func TestGetCube(t *testing.T) {
	puzzle := CreateTestPuzzle()
	result := puzzle.GetCube(1, 1)
	expected := [lineWidth]int{0, 3, 8, 0, 1, 0, 5, 0, 7}

	for i, v := range expected {
		if result[i].value != v {
			t.Errorf("GetCube value error at index: %d, expected: %d, got: %d", i, v, result[i].value)
		}
	}
}

func TestGetCubeIndex(t *testing.T) {
	puzzle := CreateTestPuzzle()
	tables := []struct {
		x     int
		y     int
		index int
	}{
		{0, 0, 0},
		{0, 3, 0},
		{0, 4, 1},
		{0, 5, 2},
		{3, 6, 0},
		{4, 7, 4},
		{5, 8, 8},
		{0, 8, 2},
	}

	for _, v := range tables {
		result := puzzle.getCubeIndex(v.x, v.y)
		if result != v.index {
			t.Errorf("getCubeIndex error for (%d, %d), expected: %d, got: %d", v.x, v.y, v.index, result)
		}
	}
}

func TestPrint(t *testing.T) {
	puzzle := CreateTestPuzzle()

	expectedString := `- 3 8 2 - - - 4 5
- 1 - 6 - 5 - - -
5 - 7 - 8 - - - -
3 - 1 - 6 2 - 7 -
9 - - 5 - 7 - - 4
- 4 - 8 3 - 6 - 2
- - - - 2 - 3 - 8
- - - 9 - 3 - 2 -
2 7 - - - 8 4 9 -`

	if puzzle.String() != expectedString {
		t.Errorf("Puzzle string not as expected, got: \n%s", puzzle)
	}
}
