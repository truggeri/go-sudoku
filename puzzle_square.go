package main

//PuzzleSquare comment for docs
type PuzzleSquare struct {
	value         int
	possibilities [lineWidth]bool
}

func createPuzzleSquare(v int) PuzzleSquare {
	var sqr PuzzleSquare
	if v > 0 {
		sqr.value = v
	} else {
		sqr.possibilities = [lineWidth]bool{true, true, true, true, true, true, true, true, true}
	}
	return sqr
}

func (sqr PuzzleSquare) solved() bool {
	if sqr.value > 0 {
		return true
	}
	return false
}
