package main

import "testing"

func TestSudokuUnitIsComplete(t *testing.T) {
	t.Run("ReturnsTrueGivenCompleteUnit", func(t *testing.T) {
		unitToTest := []string{"1", "2", "3", "4", "5", "6", "7", "8", "9"}

		rowIsComplete, err := UnitIsComplete(unitToTest)

		if err != nil {
			t.Errorf("UnitIsComplete() unexpectedly failed with an error: " + err.Error())
		}

		if !rowIsComplete {
			t.Errorf("The unitToTest should be valid. UnitIsComplete() returned invalid")
		}
	})

	t.Run("ReturnsFalseGivenIncompleteUnit", func(t *testing.T) {
		unitToTest := []string{"1", "2", " ", " ", " ", " ", " ", " ", " "}

		rowIsComplete, err := UnitIsComplete(unitToTest)

		if err != nil {
			t.Errorf("UnitIsComplete() unexpectedly failed with an error: " + err.Error())
		}

		if rowIsComplete {
			t.Errorf("The unitToTest should not be valid. UnitIsComplete() returned valid")
		}
	})

	t.Run("ThrowsErrorGivenIncompleteUnitWithDuplicateElements", func(t *testing.T) {
		unitToTest := []string{" ", " ", "3", "3", " ", " ", " ", " ", " "}

		_, err := UnitIsComplete(unitToTest)

		if err == nil {
			t.Errorf("UnitIsComplete() should have failed with an error but did not.")
		}
	})

	t.Run("ThrowsErrorGivenCompleteUnitWithDuplicateElements", func(t *testing.T) {
		unitToTest := []string{"1", "2", "3", "3", "5", "6", "7", "8", "9"}

		_, err := UnitIsComplete(unitToTest)

		if err == nil {
			t.Errorf("UnitIsComplete() should have failed with an error but did not.")
		}
	})

	t.Run("ThrowsErrorGivenUnitWithTooFewElements", func(t *testing.T) {
		unitToTest := []string{"1"}

		_, err := UnitIsComplete(unitToTest)

		if err == nil {
			t.Errorf("UnitIsComplete() should have failed with an error but did not.")
		}
	})

	t.Run("ThrowsErrorGivenUnitWithTooManyElements", func(t *testing.T) {
		unitToTest := []string{"1", "2", "3", "4", "5", "6", "7", "8", "9", "10", "12", "13"}

		_, err := UnitIsComplete(unitToTest)

		if err == nil {
			t.Errorf("UnitIsComplete() should have failed with an error but did not.")
		}
	})

	t.Run("ThrowsErrorGivenUnitContainingInvalidNumberElement", func(t *testing.T) {
		unitToTest := []string{"0", "2", "3", "4", "5", "6", "7", "8", "9"}

		_, err := UnitIsComplete(unitToTest)

		if err == nil {
			t.Errorf("UnitIsComplete() should have failed with an error but did not.")
		}
	})

	t.Run("ThrowsErrorGivenUnitContainingNonNumberElement", func(t *testing.T) {
		unitToTest := []string{"abc", "2", "3", "4", "5", "6", "7", "8", "9"}

		_, err := UnitIsComplete(unitToTest)

		if err == nil {
			t.Errorf("UnitIsComplete() should have failed with an error but did not.")
		}
	})
}

func TestGridIsComplete(t *testing.T) {
	t.Run("ReturnsTrueGivenCompleteGrid", func(t *testing.T) {
		gridToTest := sudokuArray{
			[]string{"5", "3", "4", "6", "7", "8", "9", "1", "2"},
			[]string{"6", "7", "2", "1", "9", "5", "3", "4", "8"},
			[]string{"1", "9", "8", "3", "4", "2", "5", "6", "7"},
			[]string{"8", "5", "9", "7", "6", "1", "4", "2", "3"},
			[]string{"4", "2", "6", "8", "5", "3", "7", "9", "1"},
			[]string{"7", "1", "3", "9", "2", "4", "8", "5", "6"},
			[]string{"9", "6", "1", "5", "3", "7", "2", "8", "4"},
			[]string{"2", "8", "7", "4", "1", "9", "6", "3", "5"},
			[]string{"3", "4", "5", "2", "8", "6", "1", "7", "9"},
		}

		if !GridIsComplete(gridToTest) {
			t.Errorf("GridIsComplete() should have returned true but actually returned false.")
		}
	})

	t.Run("ReturnsFalseGivenIncompleteCompleteGrid", func(t *testing.T) {
		gridToTest := sudokuArray{
			[]string{"5", "3", "4", "6", "7", "8", "9", "1", "2"},
			[]string{"6", "7", "2", "1", "9", "5", "3", "4", "8"},
			[]string{"1", "9", "8", " ", " ", "2", "5", " ", "7"},
			[]string{"8", "5", "9", "7", "6", "1", "4", " ", "3"},
			[]string{"4", "2", "6", "8", "5", "3", "7", "9", "1"},
			[]string{"7", "1", "3", "9", "2", "4", "8", "5", "6"},
			[]string{"9", "6", "1", "5", "3", "7", "2", "8", "4"},
			[]string{"2", "8", "7", "4", "1", " ", "6", "3", "5"},
			[]string{"3", "4", "5", "2", "8", "6", "1", "7", "9"},
		}

		if GridIsComplete(gridToTest) {
			t.Errorf("GridIsComplete() should have returned false but actually returned true.")
		}
	})
}
