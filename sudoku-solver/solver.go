package main

import (
	"errors"
)

// Solves a given sudoku grid
func solveGrid(grid *grid) bool {
	for !gridIsComplete(grid) {
		numSquaresSolved := grid.countSolvedSquares()
		squareUpdated := false
		for _, val := range validValues {
			squareUpdated = updateGrid(grid, val)
		}

		// No further squares have been solved, so the grid cannot be solved
		if numSquaresSolved == grid.countSolvedSquares() && !squareUpdated {
			return false
		}
	}

	return true
}

func updateGrid(grid *grid, value string) bool {
	rowUpdated := false
	rows := grid.getRows()
	for _, row := range rows {
		rowUpdated, _ = updateUnit(row, value)
	}

	columnUpdated := false
	columns := grid.getColumns()
	for _, column := range columns {
		columnUpdated, _ = updateUnit(column, value)
	}

	blockUpdated := false
	blocks := grid.getBlocks()
	for _, block := range blocks {
		blockUpdated, _ = updateUnit(block, value)
	}

	return rowUpdated || columnUpdated || blockUpdated
}

func updateUnit(row []*square, value string) (bool, error) {
	solvedSquare, squaresWithValue, squaresWithoutValue, err := categorizeSquares(row, value)
	if err != nil {
		return false, err
	}

	// One square has no other possible values so all other squares can exclude this as a possible value
	valueExcluded := false
	if solvedSquare != nil {
		for _, square := range squaresWithValue {
			valueExcluded = square.exclude(value)
		}
		return valueExcluded, nil
	}

	// Only one square has the value as possible so it can have no other possible values
	if len(squaresWithValue) == 1 {
		squaresWithValue[0].setValue(value)
		return true, nil
	}

	// Multiple squares have the value as possible forming a mini-group of matching squares so it can be excluded from all other squares
	for _, square := range squaresWithValue {
		matchingSquares, nonMatchingSquares := bucketMatchingSquares(squaresWithValue, square)
		if len(matchingSquares) == len(square.possibleValues) {
			for _, square := range nonMatchingSquares {
				valueExcluded = square.exclude(value)
			}

			return valueExcluded, nil
		}
	}

	// Multiple squares have the value as possible forming a mini-group if non-matching squares that contain 1 or more other identical possible values not present in the nonMatchingSquares meaning all other values can be excluded from all other squares
	var nonMatchingSquaresPossibleValues []string // contains unique list of possible values from the squares that did not have the value provided
	for _, square := range squaresWithoutValue {
		for _, val := range square.possibleValues {
			for _, possVal := range nonMatchingSquaresPossibleValues {
				if val == possVal {
					break
				}
			}

			nonMatchingSquaresPossibleValues = append(nonMatchingSquaresPossibleValues, val)
		}
	}

	var valContainingSquaresWithOtherPossibleValue []*square // contains list of all squares containing the value provided and another possible value that was not possible for any squares without the provided value
	var interestingVals []string
	var unInterestingVals []string
	for _, square := range squaresWithValue {
		for _, val := range square.possibleValues {
			for _, possVal := range nonMatchingSquaresPossibleValues {
				if val == possVal {
					unInterestingVals = append(unInterestingVals, val)
					break
				}
			}

			// Both of these need to only add if unique
			valContainingSquaresWithOtherPossibleValue = append(valContainingSquaresWithOtherPossibleValue, square)
			interestingVals = append(interestingVals, val)
		}
	}

	if len(valContainingSquaresWithOtherPossibleValue) == len(interestingVals) {
		for _, square := range valContainingSquaresWithOtherPossibleValue {
			for _, val := range unInterestingVals {
				valueExcluded = square.exclude(val)
			}
		}
	}

	return valueExcluded, nil
}

func categorizeSquares(row []*square, value string) (*square, []*square, []*square, error) {
	var solvedSquare *square
	var squaresWithValue []*square
	var squaresWithoutValue []*square
	for _, square := range row {
		if square.getValue() == value {
			if solvedSquare == nil {
				solvedSquare = square
			} else {
				return nil, nil, nil, errors.New("multiple squares are solved with the same value")
			}
		} else if square.isPossibleValue(value) {
			squaresWithValue = append(squaresWithValue, square)
		} else {
			squaresWithoutValue = append(squaresWithoutValue, square)
		}
	}
	return solvedSquare, squaresWithValue, squaresWithoutValue, nil
}

// Determine which squares in a slice have a matching set of possible values vs a non-matching set and return buckets
func bucketMatchingSquares(squares []*square, targetSquare *square) ([]*square, []*square) {
	var matchingSquares []*square
	var nonMatchingSquares []*square
	for _, square := range squares {
		if square == targetSquare {
			matchingSquares = append(matchingSquares, targetSquare)
		} else {
			if valuesAreMatching(square.possibleValues, targetSquare.possibleValues) {
				matchingSquares = append(matchingSquares, square)
			} else {
				nonMatchingSquares = append(nonMatchingSquares, square)
			}
		}
	}
	return matchingSquares, nonMatchingSquares
}
