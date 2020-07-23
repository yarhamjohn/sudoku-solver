package main

//TODO use sets?
var possibleBlockValues = []string{"1", "2", "3", "4", "5", "6", "7", "8", "9"}

func possibleValuesAreEqual(a, b []string) bool {
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

func isPossibleValue(value string) bool {
	for _, blockValue := range possibleBlockValues {
		if blockValue == value {
			return true
		}
	}
	return false
}

type sudokuBlock struct {
	possibleValues []string
}

func (i *sudokuBlock) GetBlockValue() string {
	if len(i.possibleValues) == 1 {
		return i.possibleValues[0]
	}

	return " "
}

func (i *sudokuBlock) containsPossibleValue(value string) bool {
	for _, val := range i.possibleValues {
		if val == value {
			return true
		}
	}

	return false
}

func (i *sudokuBlock) excludePossibleValue(value string) {
	if i.GetBlockValue() != " " {
		return
	}

	var valuesToKeep []string
	for _, val := range i.possibleValues {
		if val != value {
			valuesToKeep = append(valuesToKeep, val)
		}
	}
	i.possibleValues = valuesToKeep
}

func createBlock(value string) sudokuBlock {
	if value == " " {
		return sudokuBlock{possibleValues: possibleBlockValues}
	}

	return sudokuBlock{possibleValues: []string{value}}
}

func createBlocks(values []string) []sudokuBlock {
	var blocks []sudokuBlock
	for _, value := range values {
		if value == " " {
			blocks = append(blocks, sudokuBlock{possibleValues: possibleBlockValues})
		} else {
			blocks = append(blocks, sudokuBlock{possibleValues: []string{value}})
		}
	}
	return blocks
}
