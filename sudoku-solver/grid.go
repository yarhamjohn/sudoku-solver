package main

import (
	"errors"
	"strings"
)

// Type definition for a sudoku grid
type grid [][]square

// Method for generating a string representation of the grid type
func (i *grid) String() string {
	var gridToPrint []string

	for row := 0; row < len(*i); row++ {
		var rowToPrint []string

		for col := 0; col < len(*i); col++ {
			if col > 0 && col%3 == 0 {
				rowToPrint = append(rowToPrint, "|")
			}

			value := (*i)[row][col].getValue()
			if value == "" {
				value = " "
			}

			rowToPrint = append(rowToPrint, value)
		}

		if row > 0 && row%3 == 0 {
			gridToPrint = append(gridToPrint, "---------------------\n")
		}

		gridToPrint = append(gridToPrint, strings.Join(rowToPrint, " "), "\n")
	}

	return strings.Join(gridToPrint, "")
}

// Extension method for a grid used by the command line parser to parse a string into an instance of a grid
func (i *grid) Set(value string) error {
	fullArray := strings.Split(value, ",")

	if len(fullArray) != 81 {
		return errors.New("array not correct length")
	}

	for startElem := 0; startElem < len(fullArray); startElem += 9 {
		endElem := startElem + 9

		var row []square
		for _, elem := range fullArray[startElem:endElem] {
			row = append(row, createSquare(elem))
		}

		*i = append(*i, row)
	}
	return nil
}

// Extension method on a grid to count the number of squares that have been solved
func (i *grid) countSolvedSquares() int {
	num := 0
	for _, row := range *i {
		for _, square := range row {
			if square.getValue() != "" {
				num += 1
			}
		}
	}

	return num
}

// Extension method on a grid that returns the row containing the specified square
func (i *grid) getRow(row int) []*square {
	var squares []*square
	for col := 0; col < len(*i); col++ {
		squares = append(squares, &(*i)[row][col])
	}

	return squares
}

// Extension method on a grid that returns the column containing the specified square
func (i *grid) getColumn(col int) []*square {
	var squares []*square
	for row := 0; row < len(*i); row++ {
		squares = append(squares, &(*i)[row][col])
	}

	return squares
}

// Extension method on a grid that returns the 3x3 sudoku block containing the specified square
func (i *grid) getBlock(row int, col int) []*square {
	var squares []*square

	// Gets the quotient only then turns it into a start row/col for the target square
	startRow := (row / 3) * 3
	startCol := (col / 3) * 3

	for r := startRow; r < startRow+3; r++ {
		for c := startCol; c < startCol+3; c++ {
			squares = append(squares, &(*i)[r][c])
		}
	}

	return squares
}

// Extension method on a grid to get all the squares related to the specified square (e.g. same row, column or block)
func (i *grid) getAllRelatedSquares(row int, col int) []*square {
	var squares []*square
	for _, square := range i.getRow(row) {
		squares = append(squares, square)
	}

	for _, square := range i.getColumn(col) {
		squares = append(squares, square)
	}

	for _, square := range i.getBlock(row, col) {
		squares = append(squares, square)
	}

	return unique(squares)
}

// Removes duplicate pointers from a slice of square pointers
func unique(squares []*square) []*square {
	keys := make(map[*square]bool)
	var uniqueSquares []*square
	for _, square := range squares {
		if _, alreadySeen := keys[square]; !alreadySeen {
			keys[square] = true
			uniqueSquares = append(uniqueSquares, square)
		}
	}

	return uniqueSquares
}
