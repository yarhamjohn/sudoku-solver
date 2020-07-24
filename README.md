A Sudoku solver implemented in Golang

This is a command line application and is published as a docker image here: https://hub.docker.com/r/yarhamjohn/sudoku-solver
To run the program, pull the docker image and then execute:

`docker run sudoku-solver --grid ",,,,,,,,,3,,4,,,,5,,2,,1,,8,,2,,3,,,,,,,,,,,,4,1,6,,7,8,9,,8,,3,,,,7,,6,,,5,,7,,9,,,,,9,3,,1,2,,,,6,2,9,,8,1,4,,"`

There is a single argument required:

`--grid`
A comma-separated list representing the sudoku grid to be solved. The digits are ordered by row (top -> bottom) then column (left -> right). In the sample above, this corresponds to a sudoku-grid that looks like this:

```
       |       |       
 3   4 |       | 5   2 
   1   | 8   2 |   3   
-----------------------
       |       |       
   4 1 | 6   7 | 8 9        
 8   3 |       | 7   6 
-----------------------
     5 |   7   | 9     
     9 | 3   1 | 2     
   6 2 | 9   8 | 1 4    
```
