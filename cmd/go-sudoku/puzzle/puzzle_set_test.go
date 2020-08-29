package puzzle

import (
	"testing"
)

func TestSetPossibilities(t *testing.T) {
	puzzle := CreateTestPuzzleEasy()
	puzzle = puzzle.CalculatePossibilities()
	var set Set = puzzle[0]

	tables := []struct {
		index    int
		possible bool
	}{
		{0, true},
		{1, true},
		{2, true},
		{3, true},
		{4, true},
		{5, true},
		{6, true},
		{7, true},
		{8, true},
	}

	poss := set.Possibilities(9)
	for _, table := range tables {
		if poss[table.index] != table.possible {
			t.Errorf("Possibilities at index %d wrong, expected: %t, got %t", table.index, table.possible, poss[table.index])
		}
	}

	tables[5].possible = false
	poss = set.Possibilities(0)
	for _, table := range tables {
		if poss[table.index] != table.possible {
			t.Errorf("Possibilities at index %d wrong, expected: %t, got %t", table.index, table.possible, poss[table.index])
		}
	}

}
