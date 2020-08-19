package main

import "testing"

func TestSolved(t *testing.T) {
	var square PuzzleSquare

	tables := []struct {
		val    int
		solved bool
	}{
		{0, false},
		{1, true},
		{7, true},
	}

	for _, table := range tables {
		square.value = table.val
		if square.solved() != table.solved {
			t.Errorf("PuzzleSquare with value %d, not solved correctly, got: %t, want: %t.", square.value, square.solved(), table.solved)
		}
	}
}
