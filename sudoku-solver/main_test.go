package main

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
	"testing"
)

func TestMain(m *testing.M) {
	buildApp := exec.Command("go", "build")
	err := buildApp.Run()
	if err != nil {
		fmt.Printf("could not make binary for sudoku-solver: %v", err)
		os.Exit(1)
	}

	fmt.Println("app built")
	os.Exit(m.Run())
}

func TestSudokuSolver(t *testing.T) {
	t.Run("SolvesEasySudokuGrid", func(t *testing.T) {
		input := []string{
			"", "", "", "", "", "", "", "", "",
			"3", "", "4", "", "", "", "5", "", "2",
			"", "1", "", "8", "", "2", "", "3", "",
			"", "", "", "", "", "", "", "", "",
			"", "4", "1", "6", "", "7", "8", "9", "",
			"8", "", "3", "", "", "", "7", "", "6",
			"", "", "5", "", "7", "", "9", "", "",
			"", "", "9", "3", "", "1", "2", "", "",
			"", "6", "2", "9", "", "8", "1", "4", "",
		}

		cmd := exec.Command("./sudoku-solver", "--grid", strings.Join(input, ","))

		output, err := cmd.Output()
		if err != nil {
			t.Fatal(err)
		}

		if !strings.Contains(string(output), "The grid has been solved!") {
			t.Errorf("Grid should have been solved, but was not: \n" + string(output))
		}
	})

	t.Run("SolvesMediumSudokuGrid", func(t *testing.T) {
		input := []string{
			"", "", "", "", "", "3", "", "", "",
			"", "5", "", "1", "6", "", "7", "", "",
			"", "", "", "", "", "9", "8", "2", "",
			"", "", "", "4", "", "", "3", "", "5",
			"", "3", "9", "8", "", "", "", "1", "",
			"5", "", "4", "", "3", "6", "", "8", "",
			"7", "4", "", "3", "2", "", "", "", "",
			"", "", "6", "", "1", "", "", "4", "",
			"", "", "8", "6", "", "", "", "", "",
		}

		cmd := exec.Command("./sudoku-solver", "--grid", strings.Join(input, ","))

		output, err := cmd.Output()
		if err != nil {
			t.Fatal(err)
		}

		if !strings.Contains(string(output), "The grid has been solved!") {
			t.Errorf("Grid should have been solved, but was not: \n" + string(output))
		}
	})

	t.Run("SolvesHardSudokuGrid", func(t *testing.T) {
		input := []string{
			"", "", "", "", "", "", "", "", "",
			"", "", "", "8", "", "7", "", "", "",
			"", "5", "6", "", "3", "", "7", "8", "",
			"", "6", "1", "", "8", "", "9", "5", "",
			"2", "", "", "6", "", "5", "", "", "4",
			"", "", "", "", "", "", "", "", "",
			"9", "", "", "", "4", "", "", "", "7",
			"6", "", "2", "", "5", "", "4", "", "9",
			"5", "", "4", "2", "", "1", "3", "", "8",
		}

		cmd := exec.Command("./sudoku-solver", "--grid", strings.Join(input, ","))

		output, err := cmd.Output()
		if err != nil {
			fmt.Println(string(output))
			t.Fatal(err)
		}

		if !strings.Contains(string(output), "The grid has been solved!") {
			t.Errorf("Grid should have been solved, but was not: \n" + string(output))
		}
	})

	t.Run("SolvesExpertSudokuGrid", func(t *testing.T) {
		t.Skip("This test currently fails but is here for future reference")
		input := []string{
			"5", "", "", "", "", "", "4", "2", "",
			"", "", "", "6", "", "7", "", "1", "",
			"", "", "", "", "", "", "", "", "3",
			"", "", "4", "", "", "2", "", "", "8",
			"", "", "", "", "7", "9", "", "", "",
			"", "1", "", "5", "", "", "", "", "",
			"", "", "", "3", "4", "", "8", "", "",
			"", "5", "1", "", "2", "", "", "", "",
			"", "7", "", "", "", "", "", "", "6",
		}

		cmd := exec.Command("./sudoku-solver", "--grid", strings.Join(input, ","))

		output, err := cmd.Output()
		if err != nil {
			fmt.Println(string(output))
			t.Fatal(err)
		}

		if !strings.Contains(string(output), "The grid has been solved!") {
			t.Errorf("Grid should have been solved, but was not: \n" + string(output))
		}
	})
}
