package main

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

func isPossibleBlockValue(value string) bool {
	for _, blockValue := range possibleBlockValues {
		if blockValue == value {
			return true
		}
	}
	return false
}
