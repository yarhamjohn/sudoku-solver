package main

func SolveGrid(sudokuGrid sudokuGrid) {
	for i := 0; i < len(sudokuGrid); i++ {
		rowHasOneMissingElement, colIndex := UnitHasOneMissingElement(sudokuGrid[i])
		if rowHasOneMissingElement {
			for _, value := range possibleBlockValues {
				if !sudokuBlocksContain(sudokuGrid[i], value) {
					sudokuGrid[i][colIndex] = sudokuBlock{possibleValues: []string{value}}
				}
			}
		}

		var column []sudokuBlock
		for j := 0; j < len(sudokuGrid[i]); j++ {
			column = append(column, sudokuGrid[j][i])
		}

		colHasOneMissingElement, rowIndex := UnitHasOneMissingElement(column)
		if colHasOneMissingElement {
			for _, value := range possibleBlockValues {
				if !sudokuBlocksContain(sudokuGrid[i], value) {
					sudokuGrid[rowIndex][i] = sudokuBlock{possibleValues: []string{value}}
				}
			}
		}
	}
}

func UnitHasOneMissingElement(blocks []sudokuBlock) (bool, int) {
	var numEmptyElements int
	var index int
	for i := 0; i < len(blocks); i++ {
		if blocks[i].GetBlockValue() == " " {
			numEmptyElements += 1
			index = i
		}

		if numEmptyElements > 1 {
			return false, index
		}
	}

	return numEmptyElements == 1, index
}
