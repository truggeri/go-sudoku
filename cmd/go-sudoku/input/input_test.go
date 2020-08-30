package input

import (
	"testing"

	"github.com/truggeri/go-sudoku/cmd/go-sudoku/puzzle"
)

func TestLoadInput(t *testing.T) {
	filepath := "../../../puzzles/easy.puzzle"
	puz, err := LoadInput(filepath)

	if err != nil {
		t.Errorf("Error returned while testing LoadInput")
	}

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

	for i := range data {
		for j := range data[i] {
			if data[i][j] != puz[i][j].Value {
				t.Errorf("LoadInput error, expected %d at %d,%d, but got %d", data[i][j], i, j, puz[i][j].Value)
			}
		}
	}
}

func TestFailedLoadInput(t *testing.T) {
	_, err := LoadInput("")

	if err == nil {
		t.Errorf("Error not returned while testing LoadInput with bad input")
	}
}
