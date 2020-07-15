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

	fmt.Println(sudokuInput.String())
}
