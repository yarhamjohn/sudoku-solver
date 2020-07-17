package main

func SolveGrid(grid *sudokuGrid) {
	for row := 0; row < len(*grid); row++ {
		for col := 0; col < len((*grid)[row]); col++ {
			blockValue := (*grid)[row][col].GetBlockValue()

			if blockValue != " " {
				updateUnit(grid.getRow(row), blockValue)
				updateUnit(grid.getColumn(col), blockValue)
				// update blocks in square
			}
		}
	}
}

func updateUnit(unit []*sudokuBlock, value string) {
	for _, block := range unit {
		if block.GetBlockValue() == " " {
			for _, val := range block.possibleValues {
				if val == value {
					block.excludePossibleValue(val)
					break
				}
			}
		}
	}
}
