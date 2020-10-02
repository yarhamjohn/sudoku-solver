package main

import (
	"errors"
	"strconv"
)

func gridIsComplete(sudokuGrid *sudokuGrid) bool {
	for i := 0; i < len(*sudokuGrid); i++ {
		rowIsComplete, _ := unitIsComplete(sudokuGrid.getRow(i))
		columnIsComplete, _ := unitIsComplete(sudokuGrid.getColumn(i))

		if !rowIsComplete || !columnIsComplete {
			return false
		}
	}

	for i := 0; i < len(*sudokuGrid); i += 3 {
		for j := 0; j < len(*sudokuGrid); j += 3 {
			squareIsComplete, _ := unitIsComplete(sudokuGrid.getSquare(i, j))

			if !squareIsComplete {
				return false
			}
		}
	}

	return true
}

func unitIsComplete(blocks []*sudokuBlock) (bool, error) {
	if len(blocks) != 9 {
		return false, errors.New("an incorrect number of blocks was provided. Expected 9, got: " + strconv.Itoa(len(blocks)))
	}

	if !unitIsValid(blocks) {
		return false, errors.New("the blocks provided are an invalid set")
	}

	if unitContains(blocks, "") {
		return false, nil
	}

	return true, nil
}

func unitIsValid(blocks []*sudokuBlock) bool {
	var uniqueBlocks []*sudokuBlock
	for _, block := range blocks {
		blockValue := block.GetBlockValue()
		if blockValue != "" {
			if unitContains(uniqueBlocks, blockValue) {
				// the blocks has a duplicate non-empty element
				return false
			}

			if !isPossibleValue(blockValue) {
				// the blocks has an invalid element
				return false
			}

			uniqueBlocks = append(uniqueBlocks, block)
		}
	}

	return true
}

func unitContains(blocks []*sudokuBlock, value string) bool {
	for _, block := range blocks {
		if block.GetBlockValue() == value {
			return true
		}
	}
	return false
}
