package main

// PuzzleSet One linewidth of puzzle squares
// Can represent a row, column or cube
type PuzzleSet [lineWidth]PuzzleSquare

// Possibilities Gives which numbers are possible in the set
// @param exclude [int] Index of element to exclude from calculation
func (set PuzzleSet) Possibilities(exclude int) [lineWidth]bool {
	var possibilities [lineWidth]bool

	for i, element := range set {
		if i == exclude {
			continue
		}

		for n := range possibilities {
			if possibilities[n] {
				continue
			}

			if element.solved() {
				if element.value == n+1 {
					possibilities[n] = true
				}
			} else if element.possibilities[n] {
				possibilities[n] = true
			}
		}
	}

	return possibilities
}
