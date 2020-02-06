package main

import (
	"errors"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

func loadInput() ([lineWidth][lineWidth]int, error) {
	var nums [lineWidth][lineWidth]int

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

func splitLine(line string) ([lineWidth]int, error) {
	var row [lineWidth]int
	values := strings.Split(line, " ")

	for i, v := range values {
		n, err := strconv.Atoi(v)
		if err != nil {
			return row, err
		}

		row[i] = n
	}

	return row, nil
}