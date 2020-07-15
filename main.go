package main

import (
	"errors"
	"flag"
	"fmt"
	"strings"
)

// Set up a type to handle the 2-d sudoku grid
type sudokuArray [][]string

// Method for generating a string representation of the sudokuArray type
func (i *sudokuArray) String() string {
	var grid []string

	for row := 0; row < len(*i); row++ {
		var rowToPrint []string

		for col := 0; col < len(*i); col++ {
			if col > 0 && col%3 == 0 {
				rowToPrint = append(rowToPrint, "|")
			}
			rowToPrint = append(rowToPrint, (*i)[row][col])
		}

		if row > 0 && row%3 == 0 {
			grid = append(grid, "---------------------\n")
		}

		grid = append(grid, strings.Join(rowToPrint, " "), "\n")
	}

	return strings.Join(grid, "")
}

// Method for parsing a string into a sudokuArray type
func (i *sudokuArray) Set(value string) error {
	fullArray := strings.Split(value, ",")

	if len(fullArray) != 81 {
		return errors.New("Array not correct length")
	}

	for row := 0; row < len(fullArray); row += 9 {
		end := row + 9

		*i = append(*i, fullArray[row:end])
	}
	return nil
}

var sudokuInput sudokuArray

func main() {
	flag.Var(&sudokuInput, "grid", "Sudoku grid ")
	flag.Parse() // Calls the Set() method on the sudokuInput

	fmt.Println("Input grid:")
	fmt.Println(sudokuInput.String())

	for !GridIsComplete(sudokuInput) {
		SolveGrid(sudokuInput)
	}

	fmt.Println("The grid has been solved!")
	fmt.Println(sudokuInput.String())
}

func SolveGrid(sudokuArray sudokuArray) {
	for row := 0; row < len(sudokuArray); row++ {
		var nonEmptyElements []string
		var elemToFill int
		for col := 0; col < len(sudokuArray[row]); col++ {
			if sudokuArray[row][col] != " " {
				nonEmptyElements = append(nonEmptyElements, sudokuArray[row][col])
			} else {
				elemToFill = col
			}
		}

		if len(nonEmptyElements) == 8 {
			for _, elem := range []string{"1", "2", "3", "4", "5", "6", "7", "8", "9"} {
				if !contains(nonEmptyElements, elem) {
					sudokuArray[row][elemToFill] = elem
				}
			}
		}
	}
}

func GridIsComplete(sudokuArray sudokuArray) bool {
	for i := 0; i < len(sudokuArray); i++ {
		rowIsComplete, _ := UnitIsComplete(sudokuArray[i])

		// If every row is complete then the grid must be complete
		if !rowIsComplete {
			return false
		}
	}

	return true
}

func UnitIsComplete(unit []string) (bool, error) {
	if !SudokuUnitIsValid(unit) {
		return false, errors.New("The unit provided is invalid: " + strings.Join(unit, ","))
	}

	if contains(unit, " ") {
		return false, nil
	}

	return true, nil
}

func SudokuUnitIsValid(unit []string) bool {
	if len(unit) != 9 {
		return false
	}

	var unitSet []string
	for _, u := range unit {
		if u != " " {
			if contains(unitSet, u) {
				// the unit has a duplicate non-empty element
				return false
			}

			if !contains([]string{"1", "2", "3", "4", "5", "6", "7", "8", "9"}, u) {
				// the unit has an invalid element
				return false
			}

			unitSet = append(unitSet, u)
		}
	}

	return true
}

func contains(slice []string, element string) bool {
	for _, a := range slice {
		if a == element {
			return true
		}
	}
	return false
}
