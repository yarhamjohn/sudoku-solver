using System.Diagnostics;
using SudokuSolver;

// Read the input
var inputDefinition = args[0];

var stopwatch = new Stopwatch();
stopwatch.Start();

// construct a 2D array of double-linked nodes
var grid = SudokuGridBuilder.Build(inputDefinition);

grid.Print();

// Apply basic rules to reduce known values in each cell
SudokuGridSolver.ApplyRules(grid);

// Use depth-first search with backtracking to brute force unknown cells
var (rowCount, colCount) = SudokuGridSolver.Solve(grid);

Console.WriteLine();
Console.WriteLine($"-------------- Complete in {rowCount} row and {colCount} column iterations ----------------");
Console.WriteLine();

grid.Print();

Console.WriteLine();
Console.WriteLine($"Time elapsed: {stopwatch.Elapsed}");

stopwatch.Stop();

