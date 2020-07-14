package main

import (
	"errors"
	"flag"
	"fmt"
	"strings"
)

type sudokuArray [][]string

func (i *sudokuArray) String() string {
	return "join it somehow"
}

func (i *sudokuArray) Set(value string) error {
	fullArray := strings.Split(value, ",")

	if len(fullArray) != 81 {
		return errors.New("Array not correct length")
	}

	for row := 0; row < len(fullArray); row += 9 {
		end := row + 9

		*i = append(*i, fullArray[row:end])
	}
	return nil
}

var sudokuInput sudokuArray

func main() {
	flag.Var(&sudokuInput, "grid", "Sudoku grid ")
	flag.Parse()

	for i := 0; i < len(sudokuInput); i++ {
		fmt.Println("test: ", sudokuInput[i])
	}

	fmt.Println(sudokuInput.String())
}
