package puzzle

import "testing"

func TestSolved(t *testing.T) {
	var square Square

	tables := []struct {
		val    int
		solved bool
	}{
		{0, false},
		{1, true},
		{7, true},
	}

	for _, table := range tables {
		square.Value = table.val
		if square.Solved() != table.solved {
			t.Errorf("Square with value %d, not solved correctly, got: %t, want: %t.", square.Value, square.Solved(), table.solved)
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
		result := CreatePuzzleSquare(table.val)
		if result.Value != table.val {
			t.Errorf("Square value not expected, got: %d, want: %d.", result.Value, table.val)
		}

		if table.checkPoss {
			possReduction := true
			for _, p := range result.Possibilities {
				possReduction = possReduction && p
			}

			if possReduction != true {
				t.Errorf("Square possibilites not all true, value: %d", table.val)
			}
		}
	}
}
