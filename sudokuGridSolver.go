package main

// TODO update blocks in square
func SolveGrid(grid *sudokuGrid) {
	for row := 0; row < len(*grid); row++ {
		for col := 0; col < len((*grid)[row]); col++ {
			blockValue := (*grid)[row][col].GetBlockValue()

			// block is solved so update containing units
			if blockValue != " " {
				updateOtherBlocksInUnit(grid.getRow(row), blockValue)
				updateOtherBlocksInUnit(grid.getColumn(col), blockValue)
			} else {
				updateSelf(grid.getRow(row), &(*grid)[row][col])
				updateSelf(grid.getColumn(col), &(*grid)[row][col])

				// TODO need to evaluate squares now
				updateSelfIfOnlyBlockInUnitWithAPossibleValue(grid.getRow(row), &(*grid)[row][col])
				updateSelfIfOnlyBlockInUnitWithAPossibleValue(grid.getColumn(col), &(*grid)[row][col])
			}
		}
	}
}

func updateSelf(unit []*sudokuBlock, block *sudokuBlock) {
	for _, b := range unit {
		if b.GetBlockValue() != " " {
			block.excludePossibleValue(b.GetBlockValue())
		}
	}
}

func updateOtherBlocksInUnit(unit []*sudokuBlock, value string) {
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

func updateSelfIfOnlyBlockInUnitWithAPossibleValue(unt []*sudokuBlock, block *sudokuBlock) {
	for _, val := range block.possibleValues {
		numOccurences := 0

		for _, block := range unt {
			if block.containsPossibleValue(val) && block.GetBlockValue() != val {
				numOccurences += 1
			}

			if numOccurences > 1 {
				break
			}
		}

		if numOccurences == 1 {
			block.possibleValues = []string{val}
		}
	}
}
