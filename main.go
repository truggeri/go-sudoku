package main

import (
	"fmt"
	"log"
	"os"
)

const lineWidth = 9

type PuzzleSqure int

func main() {
	if err := run(); err != nil {
		log.Println("error :", err)
		os.Exit(1)
	}
}

func run() error {
	nums, err := loadInput()
	if err != nil {
		return err
	}

	for _, v := range nums {
		fmt.Println(v)
	}
	return nil
}
