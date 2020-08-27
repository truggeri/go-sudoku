package main

import "testing"

var result Puzzle

func BenchmarkSolverEasy(b *testing.B) {
	var r Puzzle
	puzzle := CreateTestPuzzleEasy()
	for n := 0; n < b.N; n++ {
		r = Solve(puzzle)
	}
	result = r
}

func BenchmarkSolverMedium(b *testing.B) {
	var r Puzzle
	puzzle := CreateTestPuzzleMedium()
	for n := 0; n < b.N; n++ {
		r = Solve(puzzle)
	}
	result = r
}
