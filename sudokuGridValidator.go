package main

import (
	"errors"
	"strconv"
)

func GridIsComplete(sudokuGrid sudokuGrid) bool {
	for i := 0; i < len(sudokuGrid); i++ {
		rowIsComplete, _ := UnitIsComplete(sudokuGrid[i])

		// If every row is complete then the grid must be complete
		if !rowIsComplete {
			return false
		}
	}

	return true
}

func UnitIsComplete(blocks []sudokuBlock) (bool, error) {
	if len(blocks) != 9 {
		return false, errors.New("an incorrect number of blocks was provided. Expected 9, got: " + strconv.Itoa(len(blocks)))
	}

	if !SudokuUnitIsValid(blocks) {
		return false, errors.New("the blocks provided are an invalid set")
	}

	if sudokuBlocksContain(blocks, " ") {
		return false, nil
	}

	return true, nil
}

func SudokuUnitIsValid(blocks []sudokuBlock) bool {
	var uniqueBlocks []sudokuBlock
	for _, block := range blocks {
		blockValue := block.GetBlockValue()
		if blockValue != " " {
			if sudokuBlocksContain(uniqueBlocks, blockValue) {
				// the blocks has a duplicate non-empty element
				return false
			}

			if !isPossibleBlockValue(blockValue) {
				// the blocks has an invalid element
				return false
			}

			uniqueBlocks = append(uniqueBlocks, block)
		}
	}

	return true
}

func sudokuBlocksContain(blocks []sudokuBlock, value string) bool {
	for _, block := range blocks {
		if block.GetBlockValue() == value {
			return true
		}
	}
	return false
}
