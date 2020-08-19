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

func TestCreatePuzzleSquare(t *testing.T) {
	tables := []struct {
		val       int
		checkPoss bool
	}{
		{0, true},
		{1, false},
		{7, false},
	}

	for _, table := range tables {
		result := createPuzzleSquare(table.val)
		if result.value != table.val {
			t.Errorf("PuzzleSquare value not expected, got: %d, want: %d.", result.value, table.val)
		}

		if table.checkPoss {
			possReduction := true
			for _, p := range result.possibilities {
				possReduction = possReduction && p
			}

			if possReduction != true {
				t.Errorf("PuzzleSquare possibilites not all true, value: %d", table.val)
			}
		}
	}
}
