package main

// Solves a given sudoku grid
func solveGrid(grid *grid) {
	excludeKnownValuesFromRelatedSquares(grid)

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

// Checks every square with a known value in the grid and excludes that value from the possible values of all related squares (e.g. row, column, block)
func excludeKnownValuesFromRelatedSquares(grid *grid) {
	squareUpdated := false
	for row := 0; row < len(*grid); row++ {
		for col := 0; col < len((*grid)[row]); col++ {
			value := (*grid)[row][col].getValue()

			if value != "" {
				// current block is solved, so update related squares
				squareUpdated = updateRelatedSquares(grid, row, col, value)
			}
		}
	}

	// at least one related square got updated, so re-run
	if squareUpdated {
		excludeKnownValuesFromRelatedSquares(grid)
	}
}

// Updates all related squares by excluding the given value from their possible values
func updateRelatedSquares(grid *grid, row int, col int, value string) bool {
	relatedSquares := grid.getAllRelatedSquares(row, col)
	squareUpdated := false

	for _, squares := range relatedSquares {
		if squares.getValue() == "" {
			for _, possibleValue := range squares.possibleValues {
				if possibleValue == value {
					squares.exclude(possibleValue)
					squareUpdated = true
					break
				}
			}
		}
	}

	return squareUpdated
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
