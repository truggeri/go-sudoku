package puzzle

//Square Representation of a single square in the puzzle
type Square struct {
	Value         int
	Possibilities Possibilies
}

// Possibilies Possible values for this square
type Possibilies [LineWidth]bool

// CreatePuzzleSquare Creates a single square
func CreatePuzzleSquare(v int) Square {
	var sqr Square
	if v > 0 {
		sqr.Value = v
	} else {
		sqr.Possibilities = Possibilies{true, true, true, true, true, true, true, true, true}
	}
	return sqr
}

// Solved Does this square have a known value
func (sqr Square) Solved() bool {
	return sqr.Value > 0
}
