package main

// Slice containing all the possible valid values of a sudoku square
var validValues = []string{"1", "2", "3", "4", "5", "6", "7", "8", "9"} //TODO: why are these not numbers?

// Checks if two string arrays have matching sets of elements
func valuesAreMatching(a, b []string) bool {
	if len(a) != len(b) {
		return false
	}

	for i, v := range a {
		if v != b[i] {
			return false
		}
	}
	return true
}

// Checks if a given value is a valid sudoku value
func isValidValue(value string) bool {
	for _, blockValue := range validValues {
		if blockValue == value {
			return true
		}
	}
	return false
}

// Type definition for a square in a sudoku grid
type square struct {
	possibleValues []string
}

// Extension method for a square which gets the value of the square
// An empty string is returned if there are multiple possible values
func (i *square) getValue() string {
	if len(i.possibleValues) == 1 {
		return i.possibleValues[0]
	}

	return ""
}

// Extension method for a square which sets the possible values to a specified value
func (i *square) setValue(value string) {
	if isValidValue(value) {
		i.possibleValues = []string{value}
	}
}

// Extension method for a square which checks if a given value is still possible
func (i *square) isPossibleValue(value string) bool {
	for _, val := range i.possibleValues {
		if val == value {
			return true
		}
	}

	return false
}

// Extension method for a square which excludes a given value from the remaining possible values
func (i *square) exclude(value string) bool {
	if i.getValue() != "" {
		return false
	}

	var valuesToKeep []string
	valueExcluded := false
	for _, val := range i.possibleValues {
		if val != value {
			valuesToKeep = append(valuesToKeep, val)
		}
	}

	if len(valuesToKeep) != len(i.possibleValues) {
		i.possibleValues = valuesToKeep
		valueExcluded = true
	}

	return valueExcluded
}

// Creates an instance of a square with a given value
// If an empty string is provided, the square is created with all possible values
func createSquare(value string) square {
	if value == "" {
		return square{possibleValues: validValues}
	}

	return square{possibleValues: []string{value}}
}

// Creates a slice of squares from a given slice of values
// A square given an empty string value is created with all possible values
func createSquares(values []string) []square {
	var squares []square
	for _, value := range values {
		squares = append(squares, createSquare(value))
	}
	return squares
}
