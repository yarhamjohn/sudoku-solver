package main

import (
	"errors"
	"strings"
)

type sudokuGrid [][]sudokuBlock

// Method for generating a string representation of the sudokuGrid type
func (i *sudokuGrid) String() string {
	var gridToPrint []string

	for row := 0; row < len(*i); row++ {
		var rowToPrint []string

		for col := 0; col < len(*i); col++ {
			if col > 0 && col%3 == 0 {
				rowToPrint = append(rowToPrint, "|")
			}
			rowToPrint = append(rowToPrint, (*i)[row][col].GetBlockValue())
		}

		if row > 0 && row%3 == 0 {
			gridToPrint = append(gridToPrint, "---------------------\n")
		}

		gridToPrint = append(gridToPrint, strings.Join(rowToPrint, " "), "\n")
	}

	return strings.Join(gridToPrint, "")
}

// Method for parsing a string into a sudokuGrid type
func (i *sudokuGrid) Set(value string) error {
	fullArray := strings.Split(value, ",")

	if len(fullArray) != 81 {
		return errors.New("Array not correct length")
	}

	for row := 0; row < len(fullArray); row += 9 {
		end := row + 9

		var sudokuBlockRow []sudokuBlock
		for _, e := range fullArray[row:end] {
			sudokuBlockRow = append(sudokuBlockRow, createBlock(e))
		}

		*i = append(*i, sudokuBlockRow)
	}
	return nil
}

func (i *sudokuGrid) getRow(row int) []*sudokuBlock {
	var blocks []*sudokuBlock
	for col := 0; col < len(*i); col++ {
		blocks = append(blocks, &(*i)[row][col])
	}

	return blocks
}

func (i *sudokuGrid) getColumn(col int) []*sudokuBlock {
	var blocks []*sudokuBlock
	for row := 0; row < len(*i); row++ {
		blocks = append(blocks, &(*i)[row][col])
	}

	return blocks
}
