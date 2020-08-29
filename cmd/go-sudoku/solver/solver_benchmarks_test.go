package solver

import (
	"testing"

	"github.com/truggeri/go-sudoku/cmd/go-sudoku/puzzle"
)

var result puzzle.Puzzle

func BenchmarkSolverEasy(b *testing.B) {
	var r puzzle.Puzzle
	puzzle := CreateTestPuzzleEasy()
	for n := 0; n < b.N; n++ {
		r = Solve(puzzle)
	}
	result = r
}

func BenchmarkSolverMedium(b *testing.B) {
	var r puzzle.Puzzle
	puzzle := CreateTestPuzzleMedium()
	for n := 0; n < b.N; n++ {
		r = Solve(puzzle)
	}
	result = r
}

func BenchmarkSolverHard(b *testing.B) {
	var r puzzle.Puzzle
	puzzle := CreateTestPuzzleHard()
	for n := 0; n < b.N; n++ {
		r = Solve(puzzle)
	}
	result = r
}
