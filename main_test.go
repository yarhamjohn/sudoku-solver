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
			" ", " ", " ", " ", " ", " ", " ", " ", " ",
			"3", " ", "4", " ", " ", " ", "5", " ", "2",
			" ", "1", " ", "8", " ", "2", " ", "3", " ",
			" ", " ", " ", " ", " ", " ", " ", " ", " ",
			" ", "4", "1", "6", " ", "7", "8", "9", " ",
			"8", " ", "3", " ", " ", " ", "7", " ", "6",
			" ", " ", "5", " ", "7", " ", "9", " ", " ",
			" ", " ", "9", "3", " ", "1", "2", " ", " ",
			" ", "6", "2", "9", " ", "8", "1", "4", " ",
		}

		fmt.Println(strings.Join(input, ","))

		cmd := exec.Command("./sudoku-solver.exe", "--grid", strings.Join(input, ","))

		output, err := cmd.CombinedOutput()
		if err != nil {
			t.Fatal(err)
		}

		fmt.Println(output)
	})
}
