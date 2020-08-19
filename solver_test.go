package main

import "testing"

func TestSolver(t *testing.T) {
	puzzle := CreateTestPuzzle()
	result := solve(puzzle)
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
		t.Errorf("Solver didn't solve correctly, got: \n%s", result)
	}
}
