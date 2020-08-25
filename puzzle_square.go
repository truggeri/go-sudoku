package main

//PuzzleSquare Representation of a single square in the puzzle
type PuzzleSquare struct {
	value         int
	possibilities possibilies
}

type possibilies [lineWidth]bool

func createPuzzleSquare(v int) PuzzleSquare {
	var sqr PuzzleSquare
	if v > 0 {
		sqr.value = v
	} else {
		sqr.possibilities = possibilies{true, true, true, true, true, true, true, true, true}
	}
	return sqr
}

func (sqr PuzzleSquare) solved() bool {
	return sqr.value > 0
}
