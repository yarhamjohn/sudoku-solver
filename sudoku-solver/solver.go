package main

import "fmt"

// Solves a given sudoku grid
func solveGrid(grid *grid) bool {
	for !gridIsComplete(grid) {
		numSquaresSolved := grid.countSolvedSquares()
		fmt.Print(grid)
		for row := 0; row < len(*grid); row++ {
			for col := 0; col < len((*grid)[row]); col++ {
				value := (*grid)[row][col].getValue()
				if value != "" {
					// current block is solved, so update related squares
					updateRelatedSquaresGivenKnownValue(grid, row, col, value)
				}

				if value == "" {
					solveSquareContainingUniquePossibleValue(grid.getRow(row), &(*grid)[row][col])
					solveSquareContainingUniquePossibleValue(grid.getColumn(col), &(*grid)[row][col])
					solveSquareContainingUniquePossibleValue(grid.getBlock(row, col), &(*grid)[row][col])

					updateRelatedSquaresThatDoNotFormMiniGroupsOfMatchingPossibleValues(grid.getRow(row), &(*grid)[row][col])
					updateRelatedSquaresThatDoNotFormMiniGroupsOfMatchingPossibleValues(grid.getColumn(col), &(*grid)[row][col])
					updateRelatedSquaresThatDoNotFormMiniGroupsOfMatchingPossibleValues(grid.getBlock(row, col), &(*grid)[row][col])

					//TODO:
					// if two possible value both occur only in the same two blocks in a unit, those blocks can have no other possible values
					//https://www.thonky.com/sudoku/y-wing
					//http://www.sudokusnake.com/xwings.php
					// could also try guessing...
				}
			}
		}

		// No further squares have been solved, so the grid cannot be solved
		if numSquaresSolved == grid.countSolvedSquares() {
			return false
		}
	}

	return true
}

// Updates all related squares by excluding the given value from their possible values
func updateRelatedSquaresGivenKnownValue(grid *grid, row int, col int, value string) {
	relatedSquares := grid.getAllRelatedSquares(row, col)

	for _, squares := range relatedSquares {
		if squares.getValue() == "" {
			for _, possibleValue := range squares.possibleValues {
				if possibleValue == value {
					squares.exclude(possibleValue)
					break
				}
			}
		}
	}
}

// Update all squares in a slice that do not form a 'mini-group' of squares with matching possible values to exclude those values
func updateRelatedSquaresThatDoNotFormMiniGroupsOfMatchingPossibleValues(squares []*square, targetSquare *square) {
	matchingSquares, nonMatchingSquares := bucketMatchingSquares(squares, targetSquare)

	// If the number of squares with matching possible values is the same as the number of their possible values,
	// we know they form a 'mini-group' and that the non-matching squares cannot therefore have those possible values
	if len(matchingSquares) == len(targetSquare.possibleValues) {
		for _, square := range nonMatchingSquares {
			for _, value := range targetSquare.possibleValues {
				square.exclude(value)
			}
		}
	}
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

// Solves the specified square if that square is the only square in the provided slice of squares that has a particular possible value
func solveSquareContainingUniquePossibleValue(squares []*square, square *square) {
	for _, value := range square.possibleValues {
		numOccurrences := 0

		for _, b := range squares {
			if b.isPossibleValue(value) {
				numOccurrences += 1
			}

			if numOccurrences > 1 {
				break
			}
		}

		if numOccurrences == 1 {
			square.possibleValues = []string{value}
			break
		}
	}
}
