package main

import (
	"errors"
	"strconv"
)

// Calculates whether or not the grid is completed
func gridIsComplete(grid *grid) bool {
	for i := 0; i < len(*grid); i++ {
		rowIsComplete, _ := unitIsComplete(grid.getRow(i))
		columnIsComplete, _ := unitIsComplete(grid.getColumn(i))

		if !rowIsComplete || !columnIsComplete {
			return false
		}
	}

	for i := 0; i < len(*grid); i += 3 {
		for j := 0; j < len(*grid); j += 3 {
			blockIsComplete, _ := unitIsComplete(grid.getBlock(i, j))

			if !blockIsComplete {
				return false
			}
		}
	}

	return true
}

// Calculates if a slice of squares is completed (e.g. a row, column or block)
func unitIsComplete(squares []*square) (bool, error) {
	if len(squares) != 9 {
		return false, errors.New("an incorrect number of squares was provided. Expected 9, got: " + strconv.Itoa(len(squares)))
	}

	if !unitIsValid(squares) {
		return false, errors.New("the squares provided are an invalid set")
	}

	if unitContains(squares, "") {
		return false, nil
	}

	return true, nil
}

// Calculates if a slice of squares is valid (i.e. no duplicate or invalid values)
func unitIsValid(squares []*square) bool {
	var uniqueSquares []*square
	for _, square := range squares {
		value := square.getValue()
		if value != "" {
			if unitContains(uniqueSquares, value) {
				// the squares has a duplicate non-empty element
				return false
			}

			if !isValidValue(value) {
				// the squares has an invalid element
				return false
			}

			uniqueSquares = append(uniqueSquares, square)
		}
	}

	return true
}

// Calculates if a slice of squares contains a square with the given value
func unitContains(squares []*square, value string) bool {
	for _, square := range squares {
		if square.getValue() == value {
			return true
		}
	}
	return false
}
