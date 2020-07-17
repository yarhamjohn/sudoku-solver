package main

import (
	"errors"
	"flag"
	"fmt"
	"strconv"
	"strings"
)

var possibleBlockValues = []string{"1", "2", "3", "4", "5", "6", "7", "8", "9"}

// Set up a types to handle the 2-d sudoku grid
type sudokuBlock struct {
	possibleValues []string
}

func (i *sudokuBlock) GetBlockValue() string {
	if len(i.possibleValues) == 1 {
		return i.possibleValues[0]
	}

	return " "
}

func getBlockFromString(value string) sudokuBlock {
	if value == " " {
		return sudokuBlock{possibleValues: possibleBlockValues}
	}

	return sudokuBlock{possibleValues: []string{value}}
}

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
			sudokuBlockRow = append(sudokuBlockRow, getBlockFromString(e))
		}

		*i = append(*i, sudokuBlockRow)
	}
	return nil
}

var input sudokuGrid

func main() {
	flag.Var(&input, "grid", "Sudoku grid ")
	flag.Parse() // Calls the Set() method on the input

	fmt.Println("Input grid:")
	fmt.Println(input.String())

	for !GridIsComplete(input) {
		SolveGrid(input)
	}

	fmt.Println("The grid has been solved!")
	fmt.Println(input.String())
}

func SolveGrid(sudokuGrid sudokuGrid) {
	for i := 0; i < len(sudokuGrid); i++ {
		rowHasOneMissingElement, colIndex := UnitHasOneMissingElement(sudokuGrid[i])
		if rowHasOneMissingElement {
			for _, value := range possibleBlockValues {
				if !sudokuBlocksContain(sudokuGrid[i], value) {
					sudokuGrid[i][colIndex] = sudokuBlock{possibleValues: []string{value}}
				}
			}
		}

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

func isPossibleBlockValue(value string) bool {
	for _, blockValue := range possibleBlockValues {
		if blockValue == value {
			return true
		}
	}
	return false
}
