package main

import (
	"flag"
	"fmt"
)

func main() {
	var input sudokuGrid

	flag.Var(&input, "grid", "Sudoku grid ")
	flag.Parse() // Calls the Set() method on the input

	fmt.Println("Input grid:")
	fmt.Println(input.String())

	for !gridIsComplete(&input) {
		SolveGrid(&input)
	}

	fmt.Println("The grid has been solved!")
	fmt.Println(input.String())
}
