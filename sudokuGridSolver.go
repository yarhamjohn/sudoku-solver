package main

import (
	"fmt"
	"strconv"
	"strings"
)

func SolveGrid(sudokuGrid *sudokuGrid) {
	// update all possible values
	for row := 0; row < len(*sudokuGrid); row++ {
		for col := 0; col < len((*sudokuGrid)[row]); col++ {
			blockValue := (*sudokuGrid)[row][col].GetBlockValue()

			fmt.Println("blockValue: " + blockValue + "; Row: " + strconv.Itoa(row) + "; Col: " + strconv.Itoa(col))

			if blockValue != " " {
				// update blocks in row
				sudokuGrid = updateRow(sudokuGrid, row, blockValue)
				fmt.Println("blockValAfter: " + strings.Join((*sudokuGrid)[0][0].possibleValues, ","))

				// update blocks in col
				//updateCol(sudokuGrid, col, blockValue)

				// update blocks in square
			}

			fmt.Println("target value: " + (*sudokuGrid)[0][0].GetBlockValue())
			fmt.Println("possible values: " + strings.Join((*sudokuGrid)[0][0].possibleValues, ","))
		}
	}

	/*
		for i := 0; i < len(sudokuGrid); i++ {
			solveRow(sudokuGrid[i])

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
		}*/
}

func updateRow(grid *sudokuGrid, row int, value string) sudokuGrid {
	fmt.Println("update row")

	var newRow []sudokuBlock
	for _, block := range (*grid)[row] {
		// TODO: use set instead?
		var newValues []string
		for _, val := range block.possibleValues {
			if val != value {
				newValues = append(newValues, val)
			}
		}
		// add method to remove value from possible values
		fmt.Println("newValues: " + strings.Join(newValues, ","))
		newRow = append(newRow, sudokuBlock{possibleValues: newValues})
		fmt.Println("blockVal:" + strings.Join(block.possibleValues, ","))
	}

	var newGrid sudokuGrid
	for i := 0; i < len(grid); i++ {
		if i == row {
			newGrid = append(newGrid, newRow)
		} else {
			newGrid = append(newGrid, grid[i])
		}
	}

	return newGrid
}

func updateCol(sudokuGrid sudokuGrid, col int, value string) {
	fmt.Println("update col")

	var column []sudokuBlock
	for row := 0; row < len(sudokuGrid); row++ {
		column = append(column, sudokuGrid[row][col])
	}

	for _, block := range column {
		// TODO: use set instead?
		var newValues []string
		for _, val := range block.possibleValues {
			if val != value {
				newValues = append(newValues, val)
			}
		}
		block = sudokuBlock{possibleValues: newValues}
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

func solveRow(row []sudokuBlock) {
	rowHasOneMissingElement, colIndex := UnitHasOneMissingElement(row)
	if rowHasOneMissingElement {
		for _, value := range possibleBlockValues {
			if !sudokuBlocksContain(row, value) {
				row[colIndex] = sudokuBlock{possibleValues: []string{value}}
			}
		}
	}
}
