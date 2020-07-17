A Sudoku solver implemented in Golang

This is a command line application that takes a single parameter:

`--grid`
A comma-separated list representing the sudoku grid to be solved.

For example:

`./sudoku-solver.exe --grid "1,2,3,4,5,6,7,8,9,1,2,3,4,5,6,7,8,9,1,2,3,4,5,6,7,8,9,1,2,3,4,5,6,7,8,9,1,2,3,4,5,6,7,8,9,1,2,3,4,5,6,7,8,9,1,2,3,4,5,6,7,8,9,1,2,3,4,5,6,7,8,9,1,2,3,4,5,6,7,8,9`

Will print the grid:

```
1 2 3 | 4 5 6 | 7 8 9
1 2 3 | 4 5 6 | 7 8 9
1 2 3 | 4 5 6 | 7 8 9
---------------------
1 2 3 | 4 5 6 | 7 8 9
1 2 3 | 4 5 6 | 7 8 9
1 2 3 | 4 5 6 | 7 8 9
---------------------
1 2 3 | 4 5 6 | 7 8 9
1 2 3 | 4 5 6 | 7 8 9
1 2 3 | 4 5 6 | 7 8 9
```

Next task is to actually solve a provided sudoku grid