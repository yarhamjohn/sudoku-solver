package main

func SolveGrid(grid *sudokuGrid) {
	for row := 0; row < len(*grid); row++ {
		for col := 0; col < len((*grid)[row]); col++ {
			blockValue := (*grid)[row][col].GetBlockValue()

			if blockValue != " " {
				// update blocks in row
				updateRow(grid, row, blockValue)

				// update blocks in col
				updateCol(grid, col, blockValue)

				// update blocks in square
			}
		}
	}
}

func updateRow(grid *sudokuGrid, row int, value string) {
	for idx := range (*grid)[row] {
		if (*grid)[row][idx].GetBlockValue() == " " {
			for _, val := range (*grid)[row][idx].possibleValues {
				if val == value {
					(*grid)[row][idx].excludePossibleValue(val)
					break
				}
			}
		}
	}
}

func updateCol(grid *sudokuGrid, col int, value string) {
	var column []*sudokuBlock
	for row := 0; row < len(*grid); row++ {
		column = append(column, &(*grid)[row][col])
	}

	for _, block := range column {
		if block.GetBlockValue() == " " {
			for _, val := range block.possibleValues {
				if val != value {
					block.excludePossibleValue(val)
					break
				}
			}
		}
	}
}
