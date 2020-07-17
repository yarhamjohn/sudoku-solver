package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {
	var input sudokuGrid

	flag.Var(&input, "grid", "Sudoku grid ")
	flag.Parse() // Calls the Set() method on the input

	fmt.Println("Input grid: \n" + input.String())

	for !gridIsComplete(&input) {
		numBlocksSolved := input.countBlocksSolved()

		SolveGrid(&input)

		if numBlocksSolved == input.countBlocksSolved() {
			fmt.Println("This grid cannot currently be solved. Current status: \n" + input.String())
			os.Exit(1)
		}
	}

	fmt.Println("The grid has been solved! \n" + input.String())
}
