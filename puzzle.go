package main

import "fmt"

const lineWidth = 9
const cubeWidth = lineWidth / 3

type Puzzle [lineWidth][lineWidth]PuzzleSqure

func (result Puzzle) print() {
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
