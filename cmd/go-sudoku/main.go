package main

import (
	"fmt"
	"log"
	"os"

	"github.com/truggeri/go-sudoku/cmd/go-sudoku/input"
	"github.com/truggeri/go-sudoku/cmd/go-sudoku/solver"
)

func main() {
	if err := run(); err != nil {
		log.Println("error :", err)
		os.Exit(1)
	}
}

func run() error {
	puz, err := input.LoadInput()
	if err != nil {
		return err
	}

	fmt.Println(puz)
	result := solver.Solve(puz)
	fmt.Println("-----------------")
	fmt.Println(result)

	return nil
}
