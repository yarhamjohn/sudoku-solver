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
//TODO doNT USE SPACCE (double space silently did not work)
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

func (i *sudokuGrid) countBlocksSolved() int {
	num := 0
	for _, row := range *i {
		for _, block := range row {
			if block.GetBlockValue() != " " {
				num += 1
			}
		}
	}

	return num
}

func (i *sudokuGrid) getSquare(row int, col int) []*sudokuBlock {
	var blocks []*sudokuBlock

	// Gets the quotient only then turns it into a start row/col for the target square
	startRow := (row / 3) * 3
	startCol := (col / 3) * 3

	for r := startRow; r < startRow+3; r++ {
		for c := startCol; c < startCol+3; c++ {
			blocks = append(blocks, &(*i)[r][c])
		}
	}

	return blocks
}

func (i *sudokuGrid) getAllRelatedBlocks(row int, col int) []*sudokuBlock {
	var blocks []*sudokuBlock
	for _, b := range i.getRow(row) {
		blocks = append(blocks, b)
	}

	for _, b := range i.getColumn(col) {
		blocks = append(blocks, b)
	}

	for _, b := range i.getSquare(row, col) {
		blocks = append(blocks, b)
	}

	return blocks
}
