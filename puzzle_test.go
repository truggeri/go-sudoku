package main

import (
	"testing"
)

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

func TestPositionInCube(t *testing.T) {
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
		result := positionInCube(v.x, v.y)
		if result != v.index {
			t.Errorf("positionInCube error for (%d, %d), expected: %d, got: %d", v.x, v.y, v.index, result)
		}
	}
}

func TestCubePosition(t *testing.T) {
	tables := []struct {
		index int
		x     int
		y     int
	}{
		{0, 0, 0},
		{1, 3, 0},
		{2, 6, 0},
		{3, 0, 3},
		{4, 3, 3},
		{5, 6, 3},
		{6, 0, 6},
		{7, 3, 6},
		{8, 6, 6},
	}

	for _, table := range tables {
		x, y := cubePosition(table.index)
		if x != table.x {
			t.Errorf("cubePosition error for %d x, expected: %d, got: %d", table.index, table.x, x)
		}
		if y != table.y {
			t.Errorf("cubePosition error for %d y, expected: %d, got: %d", table.index, table.y, y)
		}
	}
}

func TestCubeNumber(t *testing.T) {
	tables := []struct {
		x     int
		y     int
		index int
	}{
		{0, 0, 0}, {1, 1, 0}, {2, 2, 0},
		{3, 0, 1}, {4, 1, 1}, {5, 2, 1},
		{6, 0, 2}, {7, 1, 2}, {8, 2, 2},
		{0, 3, 3}, {1, 4, 3}, {2, 5, 3},
		{3, 3, 4}, {4, 4, 4}, {5, 5, 4},
		{6, 3, 5}, {7, 4, 5}, {8, 5, 5},
		{0, 6, 6}, {1, 7, 6}, {2, 8, 6},
		{3, 6, 7}, {4, 7, 7}, {5, 8, 7},
		{6, 6, 8}, {7, 7, 8}, {8, 8, 8},
	}

	for _, table := range tables {
		index := cubeNumber(table.x, table.y)
		if index != table.index {
			t.Errorf("cubeNumber error for (%d, %d), expected: %d, got: %d", table.x, table.y, table.index, index)
		}
	}
}

func TestFindElementPossbilities(t *testing.T) {
	tables := []struct {
		elements PuzzleSet
		expected possibilies
	}{
		{PuzzleSet{
			createPuzzleSquare(1), createPuzzleSquare(2), createPuzzleSquare(0),
			createPuzzleSquare(4), createPuzzleSquare(0), createPuzzleSquare(6),
			createPuzzleSquare(0), createPuzzleSquare(8), createPuzzleSquare(9),
		}, possibilies{false, false, true, false, true, false, true, false, false}},
	}

	for _, table := range tables {
		poss := findElementPossbilities(table.elements)
		for i := range poss {
			if poss[i] != table.expected[i] {
				t.Errorf("findElementPossbilities error at index %d, expected: %t, got: %t", i, table.expected[i], poss[i])
			}
		}
	}
}

func TestString(t *testing.T) {
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
