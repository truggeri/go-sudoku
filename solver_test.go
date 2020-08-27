package main

import "testing"

func TestSolverEasy(t *testing.T) {
	puzzle := CreateTestPuzzleEasy()
	result := Solve(puzzle)
	expectedString := `6 3 8 2 9 1 7 4 5
4 1 9 6 7 5 2 8 3
5 2 7 3 8 4 9 6 1
3 8 1 4 6 2 5 7 9
9 6 2 5 1 7 8 3 4
7 4 5 8 3 9 6 1 2
1 9 4 7 2 6 3 5 8
8 5 6 9 4 3 1 2 7
2 7 3 1 5 8 4 9 6`

	if result.String() != expectedString {
		t.Errorf("Solver didn't solve (1) correctly, got: \n%s", result)
	}
}

func TestSolverMedium(t *testing.T) {
	puzzle := CreateTestPuzzleMedium()
	result := Solve(puzzle)
	expectedString := `6 5 3 7 9 2 8 1 4
1 9 8 5 3 4 7 2 6
7 2 4 8 6 1 5 9 3
9 8 6 4 1 7 3 5 2
4 3 7 2 5 9 1 6 8
5 1 2 3 8 6 4 7 9
8 7 1 9 2 3 6 4 5
3 4 9 6 7 5 2 8 1
2 6 5 1 4 8 9 3 7`

	if result.String() != expectedString {
		t.Errorf("Solver didn't solve (3) correctly, got: \n%s", result)
	}
}

func TestOnlyOnePossibility(t *testing.T) {
	tables := []struct {
		possibility possibilies
		onlyOne     bool
		value       int
	}{
		{possibilies{true, true, false, false, false, false, false, false, false}, false, 0},
		{possibilies{true, false, false, false, false, false, false, false, true}, false, 0},
		{possibilies{true, true, true, true, true, true, true, true, true}, false, 0},
		{possibilies{true, false, false, false, false, false, false, false, false}, true, 1},
		{possibilies{false, false, false, false, false, false, false, false, true}, true, 9},
		{possibilies{false, false, false, false, true, false, false, false, false}, true, 5},
	}

	for _, table := range tables {
		onlyOne, value := onlyOnePossibility(table.possibility)
		if onlyOne != table.onlyOne {
			t.Errorf("onlyOnePossibility error for onlyOne, expected: %t, got: %t", table.onlyOne, onlyOne)
		}
		if value != table.value {
			t.Errorf("onlyOnePossibility error for value, expected: %d, got: %d", table.value, value)
		}
	}
}
