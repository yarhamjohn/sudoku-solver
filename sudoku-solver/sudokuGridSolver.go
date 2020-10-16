package main

func SolveGrid(grid *grid) {
	for row := 0; row < len(*grid); row++ {
		for col := 0; col < len((*grid)[row]); col++ {
			blockValue := (*grid)[row][col].getValue()

			if blockValue != "" {
				// current block is solved, so update blocks in the same row, col and square
				updateBlocksInContainingUnits(grid, row, col, blockValue)
			}
		}
	}

	for row := 0; row < len(*grid); row++ {
		for col := 0; col < len((*grid)[row]); col++ {
			blockValue := (*grid)[row][col].getValue()

			if blockValue == "" {
				updateSelfIfOnlyBlockInUnitWithAPossibleValue(grid.getRow(row), &(*grid)[row][col])
				updateSelfIfOnlyBlockInUnitWithAPossibleValue(grid.getColumn(col), &(*grid)[row][col])
				updateSelfIfOnlyBlockInUnitWithAPossibleValue(grid.getBlock(row, col), &(*grid)[row][col])

				updateUnitsContainingGroupsOfBlocksWithMatchingPossibleValues(grid.getRow(row), &(*grid)[row][col])
				updateUnitsContainingGroupsOfBlocksWithMatchingPossibleValues(grid.getColumn(col), &(*grid)[row][col])
				updateUnitsContainingGroupsOfBlocksWithMatchingPossibleValues(grid.getBlock(row, col), &(*grid)[row][col])

				//TODO:
				// if two possible value both occur only in the same two blocks in a unit, those blocks can have no other possible values
				//https://www.thonky.com/sudoku/y-wing
				//http://www.sudokusnake.com/xwings.php
				// could also try guessing...
			}
		}
	}
}

func updateUnitsContainingGroupsOfBlocksWithMatchingPossibleValues(unit []*square, block *square) {
	var matchingBlocks []*square
	var nonMatchingBlocks []*square
	for _, b := range unit {
		if b == block {
			matchingBlocks = append(matchingBlocks, block)
		} else {
			if valuesAreMatching(b.possibleValues, block.possibleValues) {
				matchingBlocks = append(matchingBlocks, b)
			} else {
				nonMatchingBlocks = append(nonMatchingBlocks, b)
			}
		}
	}

	if len(matchingBlocks) == len(block.possibleValues) {
		for _, b := range nonMatchingBlocks {
			for _, v := range block.possibleValues {
				b.exclude(v)
			}
		}
	}
}

func updateBlocksInContainingUnits(grid *grid, row int, col int, blockValue string) {
	blocksToUpdate := grid.getAllRelatedSquares(row, col)

	for _, block := range blocksToUpdate {
		if block.getValue() == "" {
			for _, val := range block.possibleValues {
				if val == blockValue {
					block.exclude(val)
					break
				}
			}
		}
	}
}

func updateSelfIfOnlyBlockInUnitWithAPossibleValue(blocks []*square, block *square) {
	for _, val := range block.possibleValues {
		numOccurences := 0

		for _, b := range blocks {
			if b.isPossibleValue(val) {
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
