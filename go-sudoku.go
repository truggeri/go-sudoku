package main

import (
	"errors"
	"io/ioutil"
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
	file, err := parseInput()
	if err != nil {
		return err
	}

	log.Println(file)
	return nil
}

func parseInput() (string, error) {
	argsWithoutProg := os.Args[1:]
	if len(argsWithoutProg) == 0 {
		return "", errors.New("No input args provided")
	}

	file, err := ioutil.ReadFile(argsWithoutProg[0])
	if err != nil {
		return "", err
	}

	return string(file), nil
}
