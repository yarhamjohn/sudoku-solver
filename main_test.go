package main

import "testing"

func TestSudokuUnitIsComplete(t *testing.T) {
	t.Run("ReturnsTrueGivenCompleteUnit", func(t *testing.T) {
		unitToTest := []string{"1", "2", "3", "4", "5", "6", "7", "8", "9"}

		rowIsComplete, err := SudokuUnitIsComplete(unitToTest)

		if err != nil {
			t.Errorf("SudokuUnitIsComplete() unexpectedly failed with an error: " + err.Error())
		}

		if !rowIsComplete {
			t.Errorf("The unitToTest should be valid. SudokuUnitIsComplete() returned invalid")
		}
	})

	t.Run("ReturnsFalseGivenIncompleteUnit", func(t *testing.T) {
		unitToTest := []string{"1", "2", " ", " ", " ", " ", " ", " ", " "}

		rowIsComplete, err := SudokuUnitIsComplete(unitToTest)

		if err != nil {
			t.Errorf("SudokuUnitIsComplete() unexpectedly failed with an error: " + err.Error())
		}

		if rowIsComplete {
			t.Errorf("The unitToTest should not be valid. SudokuUnitIsComplete() returned valid")
		}
	})

	t.Run("ThrowsErrorGivenIncompleteUnitWithDuplicateElements", func(t *testing.T) {
		unitToTest := []string{" ", " ", "3", "3", " ", " ", " ", " ", " "}

		_, err := SudokuUnitIsComplete(unitToTest)

		if err == nil {
			t.Errorf("SudokuUnitIsComplete() should have failed with an error but did not.")
		}
	})

	t.Run("ThrowsErrorGivenCompleteUnitWithDuplicateElements", func(t *testing.T) {
		unitToTest := []string{"1", "2", "3", "3", "5", "6", "7", "8", "9"}

		_, err := SudokuUnitIsComplete(unitToTest)

		if err == nil {
			t.Errorf("SudokuUnitIsComplete() should have failed with an error but did not.")
		}
	})

	t.Run("ThrowsErrorGivenUnitWithTooFewElements", func(t *testing.T) {
		unitToTest := []string{"1"}

		_, err := SudokuUnitIsComplete(unitToTest)

		if err == nil {
			t.Errorf("SudokuUnitIsComplete() should have failed with an error but did not.")
		}
	})

	t.Run("ThrowsErrorGivenUnitWithTooManyElements", func(t *testing.T) {
		unitToTest := []string{"1", "2", "3", "4", "5", "6", "7", "8", "9", "10", "12", "13"}

		_, err := SudokuUnitIsComplete(unitToTest)

		if err == nil {
			t.Errorf("SudokuUnitIsComplete() should have failed with an error but did not.")
		}
	})

	t.Run("ThrowsErrorGivenUnitContainingInvalidNumberElement", func(t *testing.T) {
		unitToTest := []string{"0", "2", "3", "4", "5", "6", "7", "8", "9"}

		_, err := SudokuUnitIsComplete(unitToTest)

		if err == nil {
			t.Errorf("SudokuUnitIsComplete() should have failed with an error but did not.")
		}
	})

	t.Run("ThrowsErrorGivenUnitContainingNonNumberElement", func(t *testing.T) {
		unitToTest := []string{"abc", "2", "3", "4", "5", "6", "7", "8", "9"}

		_, err := SudokuUnitIsComplete(unitToTest)

		if err == nil {
			t.Errorf("SudokuUnitIsComplete() should have failed with an error but did not.")
		}
	})
}
