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

	puz.print()
	result := solve(puz)
	result.print()

	return nil
}
