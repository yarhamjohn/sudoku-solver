package main

import (
	"testing"
)

func TestUnitIsComplete(t *testing.T) {
	t.Run("ReturnsTrueGivenCompleteUnit", func(t *testing.T) {
		unitToTest := createBlocks([]string{"1", "2", "3", "4", "5", "6", "7", "8", "9"})
		pointersToTest := GetPointers(unitToTest)

		rowIsComplete, err := unitIsComplete(pointersToTest)

		if err != nil {
			t.Errorf("unitIsComplete() unexpectedly failed with an error: " + err.Error())
		}

		if !rowIsComplete {
			t.Errorf("The unitToTest should be valid. unitIsComplete() returned invalid")
		}
	})

	t.Run("ReturnsFalseGivenIncompleteUnit", func(t *testing.T) {
		unitToTest := createBlocks([]string{"1", "2", "", "", "", "", "", "", ""})
		pointersToTest := GetPointers(unitToTest)

		rowIsComplete, err := unitIsComplete(pointersToTest)

		if err != nil {
			t.Errorf("unitIsComplete() unexpectedly failed with an error: " + err.Error())
		}

		if rowIsComplete {
			t.Errorf("The unitToTest should not be valid. unitIsComplete() returned valid")
		}
	})

	t.Run("ThrowsErrorGivenIncompleteUnitWithDuplicateElements", func(t *testing.T) {
		unitToTest := createBlocks([]string{"", "", "3", "3", "", "", "", "", ""})
		pointersToTest := GetPointers(unitToTest)

		_, err := unitIsComplete(pointersToTest)

		if err == nil {
			t.Errorf("unitIsComplete() should have failed with an error but did not.")
		}
	})

	t.Run("ThrowsErrorGivenCompleteUnitWithDuplicateElements", func(t *testing.T) {
		unitToTest := createBlocks([]string{"1", "2", "3", "3", "5", "6", "7", "8", "9"})
		pointersToTest := GetPointers(unitToTest)

		_, err := unitIsComplete(pointersToTest)

		if err == nil {
			t.Errorf("unitIsComplete() should have failed with an error but did not.")
		}
	})

	t.Run("ThrowsErrorGivenUnitWithTooFewElements", func(t *testing.T) {
		unitToTest := createBlocks([]string{"1"})
		pointersToTest := GetPointers(unitToTest)

		_, err := unitIsComplete(pointersToTest)

		if err == nil {
			t.Errorf("unitIsComplete() should have failed with an error but did not.")
		}
	})

	t.Run("ThrowsErrorGivenUnitWithTooManyElements", func(t *testing.T) {
		unitToTest := createBlocks([]string{"1", "2", "3", "4", "5", "6", "7", "8", "9", "10", "12", "13"})
		pointersToTest := GetPointers(unitToTest)

		_, err := unitIsComplete(pointersToTest)

		if err == nil {
			t.Errorf("unitIsComplete() should have failed with an error but did not.")
		}
	})

	t.Run("ThrowsErrorGivenUnitContainingInvalidNumberElement", func(t *testing.T) {
		unitToTest := createBlocks([]string{"0", "2", "3", "4", "5", "6", "7", "8", "9"})
		pointersToTest := GetPointers(unitToTest)

		_, err := unitIsComplete(pointersToTest)

		if err == nil {
			t.Errorf("unitIsComplete() should have failed with an error but did not.")
		}
	})

	t.Run("ThrowsErrorGivenUnitContainingNonNumberElement", func(t *testing.T) {
		unitToTest := createBlocks([]string{"abc", "2", "3", "4", "5", "6", "7", "8", "9"})
		pointersToTest := GetPointers(unitToTest)

		_, err := unitIsComplete(pointersToTest)

		if err == nil {
			t.Errorf("unitIsComplete() should have failed with an error but did not.")
		}
	})
}

func GetPointers(unitToTest []sudokuBlock) []*sudokuBlock {
	var pointersToTest []*sudokuBlock
	for i := 0; i < len(unitToTest); i++ {
		pointersToTest = append(pointersToTest, &unitToTest[i])
	}
	return pointersToTest
}

func TestGridIsComplete(t *testing.T) {
	t.Run("ReturnsTrueGivenCompleteGrid", func(t *testing.T) {
		gridToTest := sudokuGrid{
			createBlocks([]string{"5", "3", "4", "6", "7", "8", "9", "1", "2"}),
			createBlocks([]string{"6", "7", "2", "1", "9", "5", "3", "4", "8"}),
			createBlocks([]string{"1", "9", "8", "3", "4", "2", "5", "6", "7"}),
			createBlocks([]string{"8", "5", "9", "7", "6", "1", "4", "2", "3"}),
			createBlocks([]string{"4", "2", "6", "8", "5", "3", "7", "9", "1"}),
			createBlocks([]string{"7", "1", "3", "9", "2", "4", "8", "5", "6"}),
			createBlocks([]string{"9", "6", "1", "5", "3", "7", "2", "8", "4"}),
			createBlocks([]string{"2", "8", "7", "4", "1", "9", "6", "3", "5"}),
			createBlocks([]string{"3", "4", "5", "2", "8", "6", "1", "7", "9"}),
		}

		if !gridIsComplete(&gridToTest) {
			t.Errorf("gridIsComplete() should have returned true but actually returned false.")
		}
	})

	t.Run("ReturnsFalseGivenIncompleteCompleteGrid", func(t *testing.T) {
		gridToTest := sudokuGrid{
			createBlocks([]string{"5", "3", "4", "6", "7", "8", "9", "1", "2"}),
			createBlocks([]string{"6", "7", "2", "1", "9", "5", "3", "4", "8"}),
			createBlocks([]string{"1", "9", "8", "", "", "2", "5", "", "7"}),
			createBlocks([]string{"8", "5", "9", "7", "6", "1", "4", "", "3"}),
			createBlocks([]string{"4", "2", "6", "8", "5", "3", "7", "9", "1"}),
			createBlocks([]string{"7", "1", "3", "9", "2", "4", "8", "5", "6"}),
			createBlocks([]string{"9", "6", "1", "5", "3", "7", "2", "8", "4"}),
			createBlocks([]string{"2", "8", "7", "4", "1", "", "6", "3", "5"}),
			createBlocks([]string{"3", "4", "5", "2", "8", "6", "1", "7", "9"}),
		}

		if gridIsComplete(&gridToTest) {
			t.Errorf("gridIsComplete() should have returned false but actually returned true.")
		}
	})
}
