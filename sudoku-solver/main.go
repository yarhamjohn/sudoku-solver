package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {
	var input grid

	flag.Var(&input, "grid", "Sudoku grid ")
	flag.Parse() // Calls the Set() method on the input

	if input == nil {
		fmt.Println("No input grid was provided.")
		os.Exit(1)
	}

	fmt.Println("Input grid: \n" + input.String())

	for !gridIsComplete(&input) {
		numBlocksSolved := input.countSolvedSquares()

		SolveGrid(&input)

		if numBlocksSolved == input.countSolvedSquares() {
			fmt.Println("This grid cannot currently be solved. Current status: \n" + input.String())
			os.Exit(1)
		}
	}

	fmt.Println("The grid has been solved! \n" + input.String())
}
