package input

import (
	"io/ioutil"
	"strconv"
	"strings"

	"github.com/truggeri/go-sudoku/cmd/go-sudoku/puzzle"
)

// LoadInput Creates Puzzle from a file path
func LoadInput(filepath string) (puzzle.Puzzle, error) {
	var nums puzzle.Puzzle

	bytes, err := ioutil.ReadFile(filepath)
	if err != nil {
		return nums, err
	}

	lines := strings.Split(string(bytes), "\n")
	for lineNum, line := range lines {
		if len(line) == 0 {
			continue
		}

		set, err := splitLine(line)
		if err != nil {
			return nums, err
		}
		nums[lineNum] = set
	}

	return nums, nil
}

func splitLine(line string) (puzzle.Set, error) {
	var set puzzle.Set
	values := strings.Fields(line)

	for i, v := range values {
		n, err := strconv.Atoi(v)
		if err != nil {
			return set, err
		}

		set[i] = puzzle.CreatePuzzleSquare(n)
	}

	return set, nil
}
