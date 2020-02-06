package main

import "fmt"

const lineWidth = 9
const cubeWidth = lineWidth / 3

type PuzzleSqure struct {
	value         int
	possibilities [lineWidth]bool
}
type Puzzle [lineWidth][lineWidth]PuzzleSqure

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

func printPuzzle(result Puzzle) {
	for _, r := range result {
		for _, v := range r {
			if v.solved() {
				fmt.Printf(" %v", v.value)
			} else {
				fmt.Print(" -")
			}
		}
		fmt.Println("")
	}
	fmt.Println("-----")
}
