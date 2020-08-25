# Go Sudoku

![CI status badge](https://github.com/truggeri/go-sudoku/workflows/CI/badge.svg)

A very simple learning application for Go.

## Background

When learning a new language, I often go for a Sudoku solver application. This is my attempt to put some of my recent Golang knowledge to work.

## Run

```bash
go build
./go-sudoku <file>
```

or

```bash
go run . <file>
```

## Test

```bash
go test
```

## Algorithm

This section will describe the algorithm used to solve the puzzels. The basic idea is this.

1. Find the possible options for each empty space
2. Examine space-by-space for places with:
    1. Only one possible solution
    2. Only space in the row with a possible solution
    3. Only space in the column with a possible solution
    4. Only space in the cube with a possible solution

## Input

The application requires a file path parameter to a puzzle. The puzzle file should be nine lines long. Each line contains a space seperated list of entries. Each entry is either a number if given in the puzzle or a zero (0) to indicate the space is unsolved.

```txt
0 3 8 2 0 0 0 4 5
0 1 0 6 0 5 0 0 0
5 0 7 0 8 0 0 0 0
3 0 1 0 6 2 0 7 0
9 0 0 5 0 7 0 0 4
0 4 0 8 3 0 6 0 2
0 0 0 0 2 0 3 0 8
0 0 0 9 0 3 0 2 0
2 7 0 0 0 8 4 9 0
```
