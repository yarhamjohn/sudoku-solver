package main

func SolveGrid(grid *sudokuGrid) {
	for row := 0; row < len(*grid); row++ {
		for col := 0; col < len((*grid)[row]); col++ {
			blockValue := (*grid)[row][col].GetBlockValue()

			if blockValue != " " {
				// current block is solved, so update blocks in the same row, col and square
				updateBlocksInContainingUnits(grid, row, col, blockValue)
			}
		}
	}

	for row := 0; row < len(*grid); row++ {
		for col := 0; col < len((*grid)[row]); col++ {
			blockValue := (*grid)[row][col].GetBlockValue()

			if blockValue == " " {
				updateSelfIfOnlyBlockInUnitWithAPossibleValue(grid.getRow(row), &(*grid)[row][col])
				updateSelfIfOnlyBlockInUnitWithAPossibleValue(grid.getColumn(col), &(*grid)[row][col])
				updateSelfIfOnlyBlockInUnitWithAPossibleValue(grid.getSquare(row, col), &(*grid)[row][col])

				updateUnitsContainingGroupsOfBlocksWithMatchingPossibleValues(grid.getRow(row), &(*grid)[row][col])
				updateUnitsContainingGroupsOfBlocksWithMatchingPossibleValues(grid.getColumn(col), &(*grid)[row][col])
				updateUnitsContainingGroupsOfBlocksWithMatchingPossibleValues(grid.getSquare(row, col), &(*grid)[row][col])

				//TODO:
				// if two possible value both occur only in the same two blocks in a unit, those blocks can have no other possible values
				//https://www.thonky.com/sudoku/y-wing
				//http://www.sudokusnake.com/xwings.php
				// could also try guessing...
			}
		}
	}
}

func updateUnitsContainingGroupsOfBlocksWithMatchingPossibleValues(unit []*sudokuBlock, block *sudokuBlock) {
	var matchingBlocks []*sudokuBlock
	var nonMatchingBlocks []*sudokuBlock
	for _, b := range unit {
		if b == block {
			matchingBlocks = append(matchingBlocks, block)
		} else {
			if possibleValuesAreEqual(b.possibleValues, block.possibleValues) {
				matchingBlocks = append(matchingBlocks, b)
			} else {
				nonMatchingBlocks = append(nonMatchingBlocks, b)
			}
		}
	}

	if len(matchingBlocks) == len(block.possibleValues) {
		for _, b := range nonMatchingBlocks {
			for _, v := range block.possibleValues {
				b.excludePossibleValue(v)
			}
		}
	}
}

func updateBlocksInContainingUnits(grid *sudokuGrid, row int, col int, blockValue string) {
	blocksToUpdate := grid.getAllRelatedBlocks(row, col)

	for _, block := range blocksToUpdate {
		if block.GetBlockValue() == " " {
			for _, val := range block.possibleValues {
				if val == blockValue {
					block.excludePossibleValue(val)
					break
				}
			}
		}
	}
}

func updateSelfIfOnlyBlockInUnitWithAPossibleValue(blocks []*sudokuBlock, block *sudokuBlock) {
	for _, val := range block.possibleValues {
		numOccurences := 0

		for _, b := range blocks {
			if b.containsPossibleValue(val) {
				numOccurences += 1
			}

			if numOccurences > 1 {
				break
			}
		}

		if numOccurences == 1 {
			block.possibleValues = []string{val}
			break
		}
	}
}
