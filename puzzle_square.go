package main

type PuzzleSqure struct {
	value         int
	possibilities [lineWidth]bool
}

func createPuzzleSquare(v int) PuzzleSqure {
	var sqr PuzzleSqure
	if v > 0 {
		sqr.value = v
	} else {
		sqr.possibilities = [lineWidth]bool{true, true, true, true, true, true, true, true, true}
	}
	return sqr
}

func (sqr PuzzleSqure) solved() bool {
	if sqr.value > 0 {
		return true
	}
	return false
}
