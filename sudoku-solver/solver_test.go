package main

import (
	"testing"
)

func TestSolveGrid(t *testing.T) {
	t.Run("CompletesGridWith_OneEmptyBlock", func(t *testing.T) {
		gridToTest := grid{
			createSquares([]string{"", "3", "4", "6", "7", "8", "9", "1", "2"}),
			createSquares([]string{"6", "7", "2", "1", "9", "5", "3", "4", "8"}),
			createSquares([]string{"1", "9", "8", "3", "4", "2", "5", "6", "7"}),
			createSquares([]string{"8", "5", "9", "7", "6", "1", "4", "2", "3"}),
			createSquares([]string{"4", "2", "6", "8", "5", "3", "7", "9", "1"}),
			createSquares([]string{"7", "1", "3", "9", "2", "4", "8", "5", "6"}),
			createSquares([]string{"9", "6", "1", "5", "3", "7", "2", "8", "4"}),
			createSquares([]string{"2", "8", "7", "4", "1", "9", "6", "3", "5"}),
			createSquares([]string{"3", "4", "5", "2", "8", "6", "1", "7", "9"}),
		}

		solveGrid(&gridToTest)

		if gridToTest[0][0].getValue() != "5" {
			t.Errorf("Block [0][0] should have value 5 but had value " + gridToTest[0][0].getValue())
		}

		if !gridIsComplete(&gridToTest) {
			t.Errorf("The gridToTest should be complete, but is not: \n" + gridToTest.String())
		}
	})

	t.Run("CompletesGridWith_TwoEmptyBlocks_InTwoRows_InOneColumn_InOneSquare", func(t *testing.T) {
		gridToTest := grid{
			createSquares([]string{"", "3", "4", "6", "7", "8", "9", "1", "2"}),
			createSquares([]string{"", "7", "2", "1", "9", "5", "3", "4", "8"}),
			createSquares([]string{"1", "9", "8", "3", "4", "2", "5", "6", "7"}),
			createSquares([]string{"8", "5", "9", "7", "6", "1", "4", "2", "3"}),
			createSquares([]string{"4", "2", "6", "8", "5", "3", "7", "9", "1"}),
			createSquares([]string{"7", "1", "3", "9", "2", "4", "8", "5", "6"}),
			createSquares([]string{"9", "6", "1", "5", "3", "7", "2", "8", "4"}),
			createSquares([]string{"2", "8", "7", "4", "1", "9", "6", "3", "5"}),
			createSquares([]string{"3", "4", "5", "2", "8", "6", "1", "7", "9"}),
		}

		solveGrid(&gridToTest)

		if gridToTest[0][0].getValue() != "5" {
			t.Errorf("Block [0][0] should have value 5 but had value " + gridToTest[0][0].getValue())
		}

		if gridToTest[1][0].getValue() != "6" {
			t.Errorf("Block [1][0] should have value 6 but had value " + gridToTest[1][0].getValue())
		}

		if !gridIsComplete(&gridToTest) {
			t.Errorf("The gridToTest should be complete, but is not: \n" + gridToTest.String())
		}
	})

	t.Run("CompletesGridWith_TwoEmptyBlocks_InTwoColumns_InOneRow_InOneSquare", func(t *testing.T) {
		gridToTest := grid{
			createSquares([]string{"", "", "4", "6", "7", "8", "9", "1", "2"}),
			createSquares([]string{"6", "7", "2", "1", "9", "5", "3", "4", "8"}),
			createSquares([]string{"1", "9", "8", "3", "4", "2", "5", "6", "7"}),
			createSquares([]string{"8", "5", "9", "7", "6", "1", "4", "2", "3"}),
			createSquares([]string{"4", "2", "6", "8", "5", "3", "7", "9", "1"}),
			createSquares([]string{"7", "1", "3", "9", "2", "4", "8", "5", "6"}),
			createSquares([]string{"9", "6", "1", "5", "3", "7", "2", "8", "4"}),
			createSquares([]string{"2", "8", "7", "4", "1", "9", "6", "3", "5"}),
			createSquares([]string{"3", "4", "5", "2", "8", "6", "1", "7", "9"}),
		}

		solveGrid(&gridToTest)

		if gridToTest[0][0].getValue() != "5" {
			t.Errorf("Block [0][0] should have value 5 but had value " + gridToTest[0][0].getValue())
		}

		if gridToTest[0][1].getValue() != "3" {
			t.Errorf("Block [0][1] should have value 3 but had value " + gridToTest[0][1].getValue())
		}

		if !gridIsComplete(&gridToTest) {
			t.Errorf("The gridToTest should be complete, but is not: \n" + gridToTest.String())
		}
	})

	t.Run("CompletesGridWith_TwoEmptyBlocks_InTwoRows_InTwoColumns_InOneSquare", func(t *testing.T) {
		gridToTest := grid{
			createSquares([]string{"", "3", "4", "6", "7", "8", "9", "1", "2"}),
			createSquares([]string{"6", "", "2", "1", "9", "5", "3", "4", "8"}),
			createSquares([]string{"1", "9", "8", "3", "4", "2", "5", "6", "7"}),
			createSquares([]string{"8", "5", "9", "7", "6", "1", "4", "2", "3"}),
			createSquares([]string{"4", "2", "6", "8", "5", "3", "7", "9", "1"}),
			createSquares([]string{"7", "1", "3", "9", "2", "4", "8", "5", "6"}),
			createSquares([]string{"9", "6", "1", "5", "3", "7", "2", "8", "4"}),
			createSquares([]string{"2", "8", "7", "4", "1", "9", "6", "3", "5"}),
			createSquares([]string{"3", "4", "5", "2", "8", "6", "1", "7", "9"}),
		}

		solveGrid(&gridToTest)

		if gridToTest[0][0].getValue() != "5" {
			t.Errorf("Block [0][0] should have value 5 but had value " + gridToTest[0][0].getValue())
		}

		if gridToTest[1][1].getValue() != "7" {
			t.Errorf("Block [1][1] should have value 7 but had value " + gridToTest[1][1].getValue())
		}

		if !gridIsComplete(&gridToTest) {
			t.Errorf("The gridToTest should be complete, but is not: \n" + gridToTest.String())
		}
	})

	t.Run("CompletesGridWith_TwoEmptyBlocks_InTwoRows_InTwoColumns_InTwoSquares", func(t *testing.T) {
		gridToTest := grid{
			createSquares([]string{"", "3", "4", "6", "7", "8", "9", "1", "2"}),
			createSquares([]string{"6", "7", "2", "", "9", "5", "3", "4", "8"}),
			createSquares([]string{"1", "9", "8", "3", "4", "2", "5", "6", "7"}),
			createSquares([]string{"8", "5", "9", "7", "6", "1", "4", "2", "3"}),
			createSquares([]string{"4", "2", "6", "8", "5", "3", "7", "9", "1"}),
			createSquares([]string{"7", "1", "3", "9", "2", "4", "8", "5", "6"}),
			createSquares([]string{"9", "6", "1", "5", "3", "7", "2", "8", "4"}),
			createSquares([]string{"2", "8", "7", "4", "1", "9", "6", "3", "5"}),
			createSquares([]string{"3", "4", "5", "2", "8", "6", "1", "7", "9"}),
		}

		solveGrid(&gridToTest)

		if gridToTest[0][0].getValue() != "5" {
			t.Errorf("Block [0][0] should have value 5 but had value " + gridToTest[0][0].getValue())
		}

		if gridToTest[1][3].getValue() != "1" {
			t.Errorf("Block [1][3] should have value 1 but had value " + gridToTest[1][3].getValue())
		}

		if !gridIsComplete(&gridToTest) {
			t.Errorf("The gridToTest should be complete, but is not: \n" + gridToTest.String())
		}
	})

	t.Run("CompletesGridWith_FourEmptyBlocks_InTwoRows_InTwoColumns_InOneSquare", func(t *testing.T) {
		gridToTest := grid{
			createSquares([]string{"", "3", "4", "", "7", "8", "9", "1", "2"}),
			createSquares([]string{"6", "7", "2", "1", "9", "5", "3", "4", "8"}),
			createSquares([]string{"1", "9", "8", "3", "4", "2", "5", "6", "7"}),
			createSquares([]string{"", "5", "9", "", "6", "1", "4", "2", "3"}),
			createSquares([]string{"4", "2", "6", "8", "5", "3", "7", "9", "1"}),
			createSquares([]string{"7", "1", "3", "9", "2", "4", "8", "5", "6"}),
			createSquares([]string{"9", "6", "1", "5", "3", "7", "2", "8", "4"}),
			createSquares([]string{"2", "8", "7", "4", "1", "9", "6", "3", "5"}),
			createSquares([]string{"3", "4", "5", "2", "8", "6", "1", "7", "9"}),
		}

		solveGrid(&gridToTest)

		if gridToTest[0][0].getValue() != "5" {
			t.Errorf("Block [0][0] should have value 5 but had value " + gridToTest[0][0].getValue())
		}

		if gridToTest[3][0].getValue() != "8" {
			t.Errorf("Block [3][0] should have value 8 but had value " + gridToTest[3][0].getValue())
		}

		if gridToTest[0][3].getValue() != "6" {
			t.Errorf("Block [0][3] should have value 6 but had value " + gridToTest[0][3].getValue())
		}

		if gridToTest[3][3].getValue() != "7" {
			t.Errorf("Block [3][3]] should have value 7 but had value " + gridToTest[3][3].getValue())
		}

		if !gridIsComplete(&gridToTest) {
			t.Errorf("The gridToTest should be complete, but is not: \n" + gridToTest.String())
		}
	})

	t.Run("CompletesGridWith_FourEmptyBlocks_InTwoRows_InTwoColumns_InOneSquare", func(t *testing.T) {
		gridToTest := grid{
			createSquares([]string{"", "3", "", "6", "7", "8", "9", "1", "2"}),
			createSquares([]string{"6", "7", "2", "1", "9", "5", "3", "4", "8"}),
			createSquares([]string{"", "9", "", "3", "4", "2", "5", "6", "7"}),
			createSquares([]string{"8", "5", "9", "7", "6", "1", "4", "2", "3"}),
			createSquares([]string{"4", "2", "6", "8", "5", "3", "7", "9", "1"}),
			createSquares([]string{"7", "1", "3", "9", "2", "4", "8", "5", "6"}),
			createSquares([]string{"9", "6", "1", "5", "3", "7", "2", "8", "4"}),
			createSquares([]string{"2", "8", "7", "4", "1", "9", "6", "3", "5"}),
			createSquares([]string{"3", "4", "5", "2", "8", "6", "1", "7", "9"}),
		}

		solveGrid(&gridToTest)

		if gridToTest[0][0].getValue() != "5" {
			t.Errorf("Block [0][0] should have value 5 but had value " + gridToTest[0][0].getValue())
		}

		if gridToTest[2][0].getValue() != "1" {
			t.Errorf("Block [2][0] should have value 1 but had value " + gridToTest[2][0].getValue())
		}

		if gridToTest[0][2].getValue() != "4" {
			t.Errorf("Block [0][2] should have value 4 but had value " + gridToTest[0][2].getValue())
		}

		if gridToTest[2][2].getValue() != "8" {
			t.Errorf("Block [2][2]] should have value 8 but had value " + gridToTest[2][2].getValue())
		}

		if !gridIsComplete(&gridToTest) {
			t.Errorf("The gridToTest should be complete, but is not: \n" + gridToTest.String())
		}
	})

	t.Run("CompletesGridWith_FourEmptyBlocks_InTwoRows_InTwoColumns_InTwoSquares", func(t *testing.T) {
		gridToTest := grid{
			createSquares([]string{"", "3", "4", "", "7", "8", "9", "1", "2"}),
			createSquares([]string{"6", "7", "2", "1", "9", "5", "3", "4", "8"}),
			createSquares([]string{"", "9", "8", "", "4", "2", "5", "6", "7"}),
			createSquares([]string{"8", "5", "9", "7", "6", "1", "4", "2", "3"}),
			createSquares([]string{"4", "2", "6", "8", "5", "3", "7", "9", "1"}),
			createSquares([]string{"7", "1", "3", "9", "2", "4", "8", "5", "6"}),
			createSquares([]string{"9", "6", "1", "5", "3", "7", "2", "8", "4"}),
			createSquares([]string{"2", "8", "7", "4", "1", "9", "6", "3", "5"}),
			createSquares([]string{"3", "4", "5", "2", "8", "6", "1", "7", "9"}),
		}

		solveGrid(&gridToTest)

		if gridToTest[0][0].getValue() != "5" {
			t.Errorf("Block [0][0] should have value 5 but had value " + gridToTest[0][0].getValue())
		}

		if gridToTest[2][0].getValue() != "1" {
			t.Errorf("Block [2][0] should have value 1 but had value " + gridToTest[2][0].getValue())
		}

		if gridToTest[0][3].getValue() != "6" {
			t.Errorf("Block [0][3] should have value 6 but had value " + gridToTest[0][3].getValue())
		}

		if gridToTest[2][3].getValue() != "3" {
			t.Errorf("Block [2][3]] should have value 3 but had value " + gridToTest[2][3].getValue())
		}

		if !gridIsComplete(&gridToTest) {
			t.Errorf("The gridToTest should be complete, but is not: \n" + gridToTest.String())
		}
	})

	t.Run("CompletesGridWith_SixEmptyBlocks_InThreeRows_InTwoColumns_InOneSquare", func(t *testing.T) {
		gridToTest := grid{
			createSquares([]string{"", "3", "", "6", "7", "8", "9", "1", "2"}),
			createSquares([]string{"", "7", "", "1", "9", "5", "3", "4", "8"}),
			createSquares([]string{"", "9", "", "3", "4", "2", "5", "6", "7"}),
			createSquares([]string{"8", "5", "9", "7", "6", "1", "4", "2", "3"}),
			createSquares([]string{"4", "2", "6", "8", "5", "3", "7", "9", "1"}),
			createSquares([]string{"7", "1", "3", "9", "2", "4", "8", "5", "6"}),
			createSquares([]string{"9", "6", "1", "5", "3", "7", "2", "8", "4"}),
			createSquares([]string{"2", "8", "7", "4", "1", "9", "6", "3", "5"}),
			createSquares([]string{"3", "4", "5", "2", "8", "6", "1", "7", "9"}),
		}

		solveGrid(&gridToTest)

		if gridToTest[0][0].getValue() != "5" {
			t.Errorf("Block [0][0] should have value 5 but had value " + gridToTest[0][0].getValue())
		}

		if gridToTest[1][0].getValue() != "6" {
			t.Errorf("Block [1][0] should have value 6 but had value " + gridToTest[1][0].getValue())
		}

		if gridToTest[2][0].getValue() != "1" {
			t.Errorf("Block [2][0] should have value 1 but had value " + gridToTest[2][0].getValue())
		}

		if gridToTest[0][2].getValue() != "4" {
			t.Errorf("Block [0][2]] should have value 4 but had value " + gridToTest[0][2].getValue())
		}

		if gridToTest[1][2].getValue() != "2" {
			t.Errorf("Block [1][2]] should have value 2 but had value " + gridToTest[1][2].getValue())
		}

		if gridToTest[2][2].getValue() != "8" {
			t.Errorf("Block [2][2]] should have value 8 but had value " + gridToTest[2][2].getValue())
		}

		if !gridIsComplete(&gridToTest) {
			t.Errorf("The gridToTest should be complete, but is not: \n" + gridToTest.String())
		}
	})

	t.Run("CompletesGridWith_NineEmptyBlocks_InThreeRows_InThreeColumns_InOneSquare", func(t *testing.T) {
		gridToTest := grid{
			createSquares([]string{"", "", "", "6", "7", "8", "9", "1", "2"}),
			createSquares([]string{"", "", "", "1", "9", "5", "3", "4", "8"}),
			createSquares([]string{"", "", "", "3", "4", "2", "5", "6", "7"}),
			createSquares([]string{"8", "5", "9", "7", "6", "1", "4", "2", "3"}),
			createSquares([]string{"4", "2", "6", "8", "5", "3", "7", "9", "1"}),
			createSquares([]string{"7", "1", "3", "9", "2", "4", "8", "5", "6"}),
			createSquares([]string{"9", "6", "1", "5", "3", "7", "2", "8", "4"}),
			createSquares([]string{"2", "8", "7", "4", "1", "9", "6", "3", "5"}),
			createSquares([]string{"3", "4", "5", "2", "8", "6", "1", "7", "9"}),
		}

		solveGrid(&gridToTest)

		if gridToTest[0][0].getValue() != "5" {
			t.Errorf("Block [0][0] should have value 5 but had value " + gridToTest[0][0].getValue())
		}

		if gridToTest[1][0].getValue() != "6" {
			t.Errorf("Block [1][0] should have value 6 but had value " + gridToTest[1][0].getValue())
		}

		if gridToTest[2][0].getValue() != "1" {
			t.Errorf("Block [2][0] should have value 1 but had value " + gridToTest[2][0].getValue())
		}

		if gridToTest[0][1].getValue() != "3" {
			t.Errorf("Block [0][1]] should have value 3 but had value " + gridToTest[0][1].getValue())
		}

		if gridToTest[1][1].getValue() != "7" {
			t.Errorf("Block [1][1]] should have value 7 but had value " + gridToTest[1][1].getValue())
		}

		if gridToTest[2][1].getValue() != "9" {
			t.Errorf("Block [2][1]] should have value 9 but had value " + gridToTest[2][1].getValue())
		}

		if gridToTest[0][2].getValue() != "4" {
			t.Errorf("Block [0][2]] should have value 4 but had value " + gridToTest[0][2].getValue())
		}

		if gridToTest[1][2].getValue() != "2" {
			t.Errorf("Block [1][2]] should have value 2 but had value " + gridToTest[1][2].getValue())
		}

		if gridToTest[2][2].getValue() != "8" {
			t.Errorf("Block [2][2]] should have value 8 but had value " + gridToTest[2][2].getValue())
		}

		if !gridIsComplete(&gridToTest) {
			t.Errorf("The gridToTest should be complete, but is not: \n" + gridToTest.String())
		}
	})
}
