package main

import (
	"log"
	"os"
)

func main() {
	if err := run(); err != nil {
		log.Println("error :", err)
		os.Exit(1)
	}
}

func run() error {
	puz, err := loadInput()
	if err != nil {
		return err
	}

	printPuzzle(puz)
	result := solve(puz)
	printPuzzle(result)

	return nil
}
