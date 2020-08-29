package puzzle

// PuzzleSet One linewidth of puzzle squares. Can represent a row, column or cube
type PuzzleSet [LineWidth]PuzzleSquare

// Possibilities Gives which numbers are possible in the set
// @param exclude [int] Index of element to exclude from calculation
func (set PuzzleSet) Possibilities(exclude int) [LineWidth]bool {
	var possibilities [LineWidth]bool

	for i, element := range set {
		if i == exclude {
			continue
		}

		for n := range possibilities {
			if possibilities[n] {
				continue
			}

			if element.Solved() {
				if element.Value == n+1 {
					possibilities[n] = true
				}
			} else if element.Possibilities[n] {
				possibilities[n] = true
			}
		}
	}

	return possibilities
}
