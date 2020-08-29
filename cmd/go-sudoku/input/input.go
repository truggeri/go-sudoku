package input

import (
	"errors"
	"io/ioutil"
	"os"
	"strconv"
	"strings"

	"github.com/truggeri/go-sudoku/cmd/go-sudoku/puzzle"
)

// LoadInput Creates Puzzle from a file path
func LoadInput() (puzzle.Puzzle, error) {
	var nums puzzle.Puzzle

	argsWithoutProg := os.Args[1:]
	if len(argsWithoutProg) == 0 {
		return nums, errors.New("No input args provided")
	}

	bytes, err := ioutil.ReadFile(argsWithoutProg[0])
	if err != nil {
		return nums, err
	}

	lines := strings.Split(string(bytes), "\n")
	for lineNum, line := range lines {
		if len(line) == 0 {
			continue
		}

		row, err := splitLine(line)
		if err != nil {
			return nums, err
		}
		nums[lineNum] = row
	}

	return nums, nil
}

func splitLine(line string) ([puzzle.LineWidth]puzzle.Square, error) {
	var row [puzzle.LineWidth]puzzle.Square
	values := strings.Split(line, " ")

	for i, v := range values {
		n, err := strconv.Atoi(v)
		if err != nil {
			return row, err
		}

		row[i] = puzzle.CreatePuzzleSquare(n)
	}

	return row, nil
}
