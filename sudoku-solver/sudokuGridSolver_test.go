package main

import (
	"strings"
	"testing"
)

func TestUpdateUnitsContainingGroupsOfBlocksWithMatchingPossibleValues(t *testing.T) {
	t.Run("CorrectlyUpdatesUnit", func(t *testing.T) {
		blocksToTest := []*sudokuBlock{
			{possibleValues: []string{"1", "2"}},
			{possibleValues: []string{"1", "2"}},
			{possibleValues: []string{"1", "2", "3"}},
			{possibleValues: []string{"4"}},
			{possibleValues: []string{"5"}},
			{possibleValues: []string{"6"}},
			{possibleValues: []string{"7"}},
			{possibleValues: []string{"8"}},
			{possibleValues: []string{"9"}},
		}

		updateUnitsContainingGroupsOfBlocksWithMatchingPossibleValues(blocksToTest, blocksToTest[0])

		if !possibleValuesAreEqual(blocksToTest[0].possibleValues, []string{"1", "2"}) {
			t.Errorf("Block [0] should have possible values 1,2 but had possible values " + strings.Join(blocksToTest[0].possibleValues, ","))
		}

		if !possibleValuesAreEqual(blocksToTest[1].possibleValues, []string{"1", "2"}) {
			t.Errorf("Block [1] should have possible values 1,2 but had possible values " + strings.Join(blocksToTest[1].possibleValues, ","))
		}

		if !possibleValuesAreEqual(blocksToTest[2].possibleValues, []string{"3"}) {
			t.Errorf("Block [2] should have possible values 3 but had possible values " + strings.Join(blocksToTest[2].possibleValues, ","))
		}
	})
}

func TestSolveGrid(t *testing.T) {
	t.Run("CompletesGridWith_OneEmptyBlock", func(t *testing.T) {
		gridToTest := sudokuGrid{
			createBlocks([]string{"", "3", "4", "6", "7", "8", "9", "1", "2"}),
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

	t.Run("CompletesGridWith_TwoEmptyBlocks_InTwoRows_InOneColumn_InOneSquare", func(t *testing.T) {
		gridToTest := sudokuGrid{
			createBlocks([]string{"", "3", "4", "6", "7", "8", "9", "1", "2"}),
			createBlocks([]string{"", "7", "2", "1", "9", "5", "3", "4", "8"}),
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

	t.Run("CompletesGridWith_TwoEmptyBlocks_InTwoColumns_InOneRow_InOneSquare", func(t *testing.T) {
		gridToTest := sudokuGrid{
			createBlocks([]string{"", "", "4", "6", "7", "8", "9", "1", "2"}),
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

	t.Run("CompletesGridWith_TwoEmptyBlocks_InTwoRows_InTwoColumns_InOneSquare", func(t *testing.T) {
		gridToTest := sudokuGrid{
			createBlocks([]string{"", "3", "4", "6", "7", "8", "9", "1", "2"}),
			createBlocks([]string{"6", "", "2", "1", "9", "5", "3", "4", "8"}),
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

	t.Run("CompletesGridWith_TwoEmptyBlocks_InTwoRows_InTwoColumns_InTwoSquares", func(t *testing.T) {
		gridToTest := sudokuGrid{
			createBlocks([]string{"", "3", "4", "6", "7", "8", "9", "1", "2"}),
			createBlocks([]string{"6", "7", "2", "", "9", "5", "3", "4", "8"}),
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
			t.Errorf("Block [1][3] should have value 1 but had value " + gridToTest[1][3].GetBlockValue())
		}

		if !gridIsComplete(&gridToTest) {
			t.Errorf("The gridToTest should be complete, but is not: \n" + gridToTest.String())
		}
	})

	t.Run("CompletesGridWith_FourEmptyBlocks_InTwoRows_InTwoColumns_InOneSquare", func(t *testing.T) {
		gridToTest := sudokuGrid{
			createBlocks([]string{"", "3", "4", "", "7", "8", "9", "1", "2"}),
			createBlocks([]string{"6", "7", "2", "1", "9", "5", "3", "4", "8"}),
			createBlocks([]string{"1", "9", "8", "3", "4", "2", "5", "6", "7"}),
			createBlocks([]string{"", "5", "9", "", "6", "1", "4", "2", "3"}),
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

		if gridToTest[3][0].GetBlockValue() != "8" {
			t.Errorf("Block [3][0] should have value 8 but had value " + gridToTest[3][0].GetBlockValue())
		}

		if gridToTest[0][3].GetBlockValue() != "6" {
			t.Errorf("Block [0][3] should have value 6 but had value " + gridToTest[0][3].GetBlockValue())
		}

		if gridToTest[3][3].GetBlockValue() != "7" {
			t.Errorf("Block [3][3]] should have value 7 but had value " + gridToTest[3][3].GetBlockValue())
		}

		if !gridIsComplete(&gridToTest) {
			t.Errorf("The gridToTest should be complete, but is not: \n" + gridToTest.String())
		}
	})

	t.Run("CompletesGridWith_FourEmptyBlocks_InTwoRows_InTwoColumns_InOneSquare", func(t *testing.T) {
		gridToTest := sudokuGrid{
			createBlocks([]string{"", "3", "", "6", "7", "8", "9", "1", "2"}),
			createBlocks([]string{"6", "7", "2", "1", "9", "5", "3", "4", "8"}),
			createBlocks([]string{"", "9", "", "3", "4", "2", "5", "6", "7"}),
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

		if gridToTest[2][0].GetBlockValue() != "1" {
			t.Errorf("Block [2][0] should have value 1 but had value " + gridToTest[2][0].GetBlockValue())
		}

		if gridToTest[0][2].GetBlockValue() != "4" {
			t.Errorf("Block [0][2] should have value 4 but had value " + gridToTest[0][2].GetBlockValue())
		}

		if gridToTest[2][2].GetBlockValue() != "8" {
			t.Errorf("Block [2][2]] should have value 8 but had value " + gridToTest[2][2].GetBlockValue())
		}

		if !gridIsComplete(&gridToTest) {
			t.Errorf("The gridToTest should be complete, but is not: \n" + gridToTest.String())
		}
	})

	t.Run("CompletesGridWith_FourEmptyBlocks_InTwoRows_InTwoColumns_InTwoSquares", func(t *testing.T) {
		gridToTest := sudokuGrid{
			createBlocks([]string{"", "3", "4", "", "7", "8", "9", "1", "2"}),
			createBlocks([]string{"6", "7", "2", "1", "9", "5", "3", "4", "8"}),
			createBlocks([]string{"", "9", "8", "", "4", "2", "5", "6", "7"}),
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

		if gridToTest[2][0].GetBlockValue() != "1" {
			t.Errorf("Block [2][0] should have value 1 but had value " + gridToTest[2][0].GetBlockValue())
		}

		if gridToTest[0][3].GetBlockValue() != "6" {
			t.Errorf("Block [0][3] should have value 6 but had value " + gridToTest[0][3].GetBlockValue())
		}

		if gridToTest[2][3].GetBlockValue() != "3" {
			t.Errorf("Block [2][3]] should have value 3 but had value " + gridToTest[2][3].GetBlockValue())
		}

		if !gridIsComplete(&gridToTest) {
			t.Errorf("The gridToTest should be complete, but is not: \n" + gridToTest.String())
		}
	})

	t.Run("CompletesGridWith_SixEmptyBlocks_InThreeRows_InTwoColumns_InOneSquare", func(t *testing.T) {
		gridToTest := sudokuGrid{
			createBlocks([]string{"", "3", "", "6", "7", "8", "9", "1", "2"}),
			createBlocks([]string{"", "7", "", "1", "9", "5", "3", "4", "8"}),
			createBlocks([]string{"", "9", "", "3", "4", "2", "5", "6", "7"}),
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

		if gridToTest[2][0].GetBlockValue() != "1" {
			t.Errorf("Block [2][0] should have value 1 but had value " + gridToTest[2][0].GetBlockValue())
		}

		if gridToTest[0][2].GetBlockValue() != "4" {
			t.Errorf("Block [0][2]] should have value 4 but had value " + gridToTest[0][2].GetBlockValue())
		}

		if gridToTest[1][2].GetBlockValue() != "2" {
			t.Errorf("Block [1][2]] should have value 2 but had value " + gridToTest[1][2].GetBlockValue())
		}

		if gridToTest[2][2].GetBlockValue() != "8" {
			t.Errorf("Block [2][2]] should have value 8 but had value " + gridToTest[2][2].GetBlockValue())
		}

		if !gridIsComplete(&gridToTest) {
			t.Errorf("The gridToTest should be complete, but is not: \n" + gridToTest.String())
		}
	})

	t.Run("CompletesGridWith_NineEmptyBlocks_InThreeRows_InThreeColumns_InOneSquare", func(t *testing.T) {
		gridToTest := sudokuGrid{
			createBlocks([]string{"", "", "", "6", "7", "8", "9", "1", "2"}),
			createBlocks([]string{"", "", "", "1", "9", "5", "3", "4", "8"}),
			createBlocks([]string{"", "", "", "3", "4", "2", "5", "6", "7"}),
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

		if gridToTest[2][0].GetBlockValue() != "1" {
			t.Errorf("Block [2][0] should have value 1 but had value " + gridToTest[2][0].GetBlockValue())
		}

		if gridToTest[0][1].GetBlockValue() != "3" {
			t.Errorf("Block [0][1]] should have value 3 but had value " + gridToTest[0][1].GetBlockValue())
		}

		if gridToTest[1][1].GetBlockValue() != "7" {
			t.Errorf("Block [1][1]] should have value 7 but had value " + gridToTest[1][1].GetBlockValue())
		}

		if gridToTest[2][1].GetBlockValue() != "9" {
			t.Errorf("Block [2][1]] should have value 9 but had value " + gridToTest[2][1].GetBlockValue())
		}

		if gridToTest[0][2].GetBlockValue() != "4" {
			t.Errorf("Block [0][2]] should have value 4 but had value " + gridToTest[0][2].GetBlockValue())
		}

		if gridToTest[1][2].GetBlockValue() != "2" {
			t.Errorf("Block [1][2]] should have value 2 but had value " + gridToTest[1][2].GetBlockValue())
		}

		if gridToTest[2][2].GetBlockValue() != "8" {
			t.Errorf("Block [2][2]] should have value 8 but had value " + gridToTest[2][2].GetBlockValue())
		}

		if !gridIsComplete(&gridToTest) {
			t.Errorf("The gridToTest should be complete, but is not: \n" + gridToTest.String())
		}
	})
}
