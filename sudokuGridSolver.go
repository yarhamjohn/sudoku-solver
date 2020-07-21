package main

import (
	"fmt"
	"strconv"
	"strings"
)

// TODO update blocks in square
func SolveGrid(grid *sudokuGrid) {
	for row := 0; row < len(*grid); row++ {
		for col := 0; col < len((*grid)[row]); col++ {
			blockValue := (*grid)[row][col].GetBlockValue()
			fmt.Println("Row: " + strconv.Itoa(row) + "; Col: " + strconv.Itoa(col) + "; Value: " + blockValue + "; Possible Values: " + strings.Join((*grid)[row][col].possibleValues, ","))

			if blockValue != " " {
				// current block is solved, so update blocks in the same row, col and square
				updateBlocksInContainingUnits(grid, row, col, blockValue)
			}

			fmt.Println("Row: " + strconv.Itoa(row) + "; Col: " + strconv.Itoa(col) + "; Value: " + blockValue + "; Possible Values: " + strings.Join((*grid)[row][col].possibleValues, ","))
		}
	}

	fmt.Println("Update related blocks: \n" + grid.String())

	for row := 0; row < len(*grid); row++ {
		for col := 0; col < len((*grid)[row]); col++ {
			blockValue := (*grid)[row][col].GetBlockValue()

			if blockValue == " " {
				updateSelfIfOnlyBlockInUnitWithAPossibleValue(grid.getRow(row), &(*grid)[row][col])
				updateSelfIfOnlyBlockInUnitWithAPossibleValue(grid.getColumn(col), &(*grid)[row][col])
				updateSelfIfOnlyBlockInUnitWithAPossibleValue(grid.getSquare(row, col), &(*grid)[row][col])
			}
		}
	}

	fmt.Println("Update self: \n" + grid.String())
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
			if b.containsPossibleValue(val) && b.GetBlockValue() != val {
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
