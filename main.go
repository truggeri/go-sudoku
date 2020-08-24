package main

import (
	"fmt"
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

	fmt.Println(puz)
	result := solve(puz)
	fmt.Println("-----------------")
	fmt.Println(result)

	return nil
}
