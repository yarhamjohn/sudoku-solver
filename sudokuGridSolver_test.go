package main

import (
	"testing"
)

func TestSolveGrid(t *testing.T) {
	t.Run("CompletesGridWith_OneEmptyBlock", func(t *testing.T) {
		gridToTest := sudokuGrid{
			createBlocks([]string{" ", "3", "4", "6", "7", "8", "9", "1", "2"}),
			createBlocks([]string{"6", "7", "2", "1", "9", "5", "3", "4", "8"}),
			createBlocks([]string{"1", "9", "8", "3", "4", "2", "5", "6", "7"}),
			createBlocks([]string{"8", "5", "9", "7", "6", "1", "4", "2", "3"}),
			createBlocks([]string{"4", "2", "6", "8", "5", "3", "7", "9", "1"}),
			createBlocks([]string{"7", "1", "3", "9", "2", "4", "8", "5", "6"}),
			createBlocks([]string{"9", "6", "1", "5", "3", "7", "2", "8", "4"}),
			createBlocks([]string{"2", "8", "7", "4", "1", "9", "6", "3", "5"}),
			createBlocks([]string{"3", "4", "5", "2", "8", "6", "1", "7", "9"}),
		}

		SolveGrid(&gridToTest)

		if gridToTest[0][0].GetBlockValue() != "5" {
			t.Errorf("Block [0][0] should have value 5 but had value " + gridToTest[0][0].GetBlockValue())
		}

		if !gridIsComplete(&gridToTest) {
			t.Errorf("The gridToTest should be complete, but is not: \n" + gridToTest.String())
		}
	})

	t.Run("CompletesGridWith_TwoEmptyBlocks_InDifferentRows_InSameColumn", func(t *testing.T) {
		gridToTest := sudokuGrid{
			createBlocks([]string{" ", "3", "4", "6", "7", "8", "9", "1", "2"}),
			createBlocks([]string{" ", "7", "2", "1", "9", "5", "3", "4", "8"}),
			createBlocks([]string{"1", "9", "8", "3", "4", "2", "5", "6", "7"}),
			createBlocks([]string{"8", "5", "9", "7", "6", "1", "4", "2", "3"}),
			createBlocks([]string{"4", "2", "6", "8", "5", "3", "7", "9", "1"}),
			createBlocks([]string{"7", "1", "3", "9", "2", "4", "8", "5", "6"}),
			createBlocks([]string{"9", "6", "1", "5", "3", "7", "2", "8", "4"}),
			createBlocks([]string{"2", "8", "7", "4", "1", "9", "6", "3", "5"}),
			createBlocks([]string{"3", "4", "5", "2", "8", "6", "1", "7", "9"}),
		}

		SolveGrid(&gridToTest)

		if gridToTest[0][0].GetBlockValue() != "5" {
			t.Errorf("Block [0][0] should have value 5 but had value " + gridToTest[0][0].GetBlockValue())
		}

		if gridToTest[1][0].GetBlockValue() != "6" {
			t.Errorf("Block [1][0] should have value 6 but had value " + gridToTest[1][0].GetBlockValue())
		}

		if !gridIsComplete(&gridToTest) {
			t.Errorf("The gridToTest should be complete, but is not: \n" + gridToTest.String())
		}
	})

	t.Run("CompletesGridWith_TwoEmptyBlocks_InDifferentColumns_InSameRow", func(t *testing.T) {
		gridToTest := sudokuGrid{
			createBlocks([]string{" ", " ", "4", "6", "7", "8", "9", "1", "2"}),
			createBlocks([]string{"6", "7", "2", "1", "9", "5", "3", "4", "8"}),
			createBlocks([]string{"1", "9", "8", "3", "4", "2", "5", "6", "7"}),
			createBlocks([]string{"8", "5", "9", "7", "6", "1", "4", "2", "3"}),
			createBlocks([]string{"4", "2", "6", "8", "5", "3", "7", "9", "1"}),
			createBlocks([]string{"7", "1", "3", "9", "2", "4", "8", "5", "6"}),
			createBlocks([]string{"9", "6", "1", "5", "3", "7", "2", "8", "4"}),
			createBlocks([]string{"2", "8", "7", "4", "1", "9", "6", "3", "5"}),
			createBlocks([]string{"3", "4", "5", "2", "8", "6", "1", "7", "9"}),
		}

		SolveGrid(&gridToTest)

		if gridToTest[0][0].GetBlockValue() != "5" {
			t.Errorf("Block [0][0] should have value 5 but had value " + gridToTest[0][0].GetBlockValue())
		}

		if gridToTest[0][1].GetBlockValue() != "3" {
			t.Errorf("Block [0][1] should have value 3 but had value " + gridToTest[0][1].GetBlockValue())
		}

		if !gridIsComplete(&gridToTest) {
			t.Errorf("The gridToTest should be complete, but is not: \n" + gridToTest.String())
		}
	})

	t.Run("CompletesGridWith_TwoEmptyBlocks_InDiffentRows_InDifferentColumns_InSameSquare", func(t *testing.T) {
		gridToTest := sudokuGrid{
			createBlocks([]string{" ", "3", "4", "6", "7", "8", "9", "1", "2"}),
			createBlocks([]string{"6", " ", "2", "1", "9", "5", "3", "4", "8"}),
			createBlocks([]string{"1", "9", "8", "3", "4", "2", "5", "6", "7"}),
			createBlocks([]string{"8", "5", "9", "7", "6", "1", "4", "2", "3"}),
			createBlocks([]string{"4", "2", "6", "8", "5", "3", "7", "9", "1"}),
			createBlocks([]string{"7", "1", "3", "9", "2", "4", "8", "5", "6"}),
			createBlocks([]string{"9", "6", "1", "5", "3", "7", "2", "8", "4"}),
			createBlocks([]string{"2", "8", "7", "4", "1", "9", "6", "3", "5"}),
			createBlocks([]string{"3", "4", "5", "2", "8", "6", "1", "7", "9"}),
		}

		SolveGrid(&gridToTest)

		if gridToTest[0][0].GetBlockValue() != "5" {
			t.Errorf("Block [0][0] should have value 5 but had value " + gridToTest[0][0].GetBlockValue())
		}

		if gridToTest[1][1].GetBlockValue() != "7" {
			t.Errorf("Block [1][1] should have value 7 but had value " + gridToTest[1][1].GetBlockValue())
		}

		if !gridIsComplete(&gridToTest) {
			t.Errorf("The gridToTest should be complete, but is not: \n" + gridToTest.String())
		}
	})

	t.Run("CompletesGridWith_TwoEmptyBlocks_InDiffentRows_InDifferentColumns_InDifferentSquares", func(t *testing.T) {
		gridToTest := sudokuGrid{
			createBlocks([]string{" ", "3", "4", "6", "7", "8", "9", "1", "2"}),
			createBlocks([]string{"6", "7", "2", " ", "9", "5", "3", "4", "8"}),
			createBlocks([]string{"1", "9", "8", "3", "4", "2", "5", "6", "7"}),
			createBlocks([]string{"8", "5", "9", "7", "6", "1", "4", "2", "3"}),
			createBlocks([]string{"4", "2", "6", "8", "5", "3", "7", "9", "1"}),
			createBlocks([]string{"7", "1", "3", "9", "2", "4", "8", "5", "6"}),
			createBlocks([]string{"9", "6", "1", "5", "3", "7", "2", "8", "4"}),
			createBlocks([]string{"2", "8", "7", "4", "1", "9", "6", "3", "5"}),
			createBlocks([]string{"3", "4", "5", "2", "8", "6", "1", "7", "9"}),
		}

		SolveGrid(&gridToTest)

		if gridToTest[0][0].GetBlockValue() != "5" {
			t.Errorf("Block [0][0] should have value 5 but had value " + gridToTest[0][0].GetBlockValue())
		}

		if gridToTest[1][3].GetBlockValue() != "1" {
			t.Errorf("Block [1][3] should have value 1 but had value " + gridToTest[1][1].GetBlockValue())
		}

		if !gridIsComplete(&gridToTest) {
			t.Errorf("The gridToTest should be complete, but is not: \n" + gridToTest.String())
		}
	})
}
