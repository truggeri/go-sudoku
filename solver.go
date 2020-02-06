package main

func solve(puz Puzzle) Puzzle {
	updated := findAllPossibilites(puz)

	for i := 0; i < lineWidth; i++ {
		for j := 0; j < lineWidth; j++ {
			if !(updated[i][j].solved()) && hasOnlyOne(updated[i][j].possibilities) {
				updated[i][j].value = findOnlyValue(updated[i][j].possibilities)
			}
		}
	}

	return updated
}

func findAllPossibilites(puz Puzzle) Puzzle {
	for i := 0; i < lineWidth; i++ {
		for j := 0; j < lineWidth; j++ {
			if !puz[i][j].solved() {
				puz[i][j].possibilities = findPossibilities(puz, i, j)
			}
		}
	}
	return puz
}

func findPossibilities(puz Puzzle, x, y int) [lineWidth]bool {
	var row [lineWidth]PuzzleSqure
	for i := 0; i < lineWidth; i++ {
		row[i] = puz[i][y]
	}
	rowPoss := findRowPossbilities(row, x)

	colPoss := findColPossbilities(puz[x], y)

	xOffset := x / cubeWidth
	yOffset := y / cubeWidth
	var cube [cubeWidth][cubeWidth]PuzzleSqure
	for i, v := range cube {
		for j := range v {
			cube[i][j] = puz[i+(xOffset*cubeWidth)][j+(yOffset*cubeWidth)]
		}
	}
	cubePoss := findSqrPossbilities(cube, x%cubeWidth, y%cubeWidth)

	return mergePoss(rowPoss, colPoss, cubePoss)
}

func findRowPossbilities(row [lineWidth]PuzzleSqure, x int) [lineWidth]bool {
	poss := [lineWidth]bool{true, true, true, true, true, true, true, true, true}

	for i := range row {
		if i == x {
			continue
		}

		if row[i].solved() {
			poss[row[i].value-1] = false
		}
	}

	return poss
}

func findColPossbilities(col [lineWidth]PuzzleSqure, y int) [lineWidth]bool {
	poss := [lineWidth]bool{true, true, true, true, true, true, true, true, true}

	for i := range col {
		if i == y {
			continue
		}

		if col[i].solved() {
			poss[col[i].value-1] = false
		}
	}

	return poss
}

func findSqrPossbilities(cube [cubeWidth][cubeWidth]PuzzleSqure, x, y int) [lineWidth]bool {
	poss := [lineWidth]bool{true, true, true, true, true, true, true, true, true}

	for i, v := range cube {
		for j := range v {
			if i == x && j == y {
				continue
			}

			if cube[i][j].solved() {
				poss[cube[i][j].value-1] = false
			}
		}
	}

	return poss
}

func mergePoss(row, col, cube [lineWidth]bool) [lineWidth]bool {
	poss := [lineWidth]bool{false, false, false, false, false, false, false, false, false}

	for i := 0; i < lineWidth; i++ {
		if row[i] == true && col[i] == true && cube[i] == true {
			poss[i] = true
		}
	}
	return poss
}

func hasOnlyOne(poss [lineWidth]bool) bool {
	counter := 0
	for _, v := range poss {
		if v == true {
			counter++
		}

		if counter > 1 {
			return false
		}
	}
	return true
}

func findOnlyValue(poss [lineWidth]bool) int {
	for i, v := range poss {
		if v == true {
			return i + 1
		}
	}
	return 0
}
