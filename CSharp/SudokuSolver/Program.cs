using SudokuSolver;

var input = args[0];

var grid = new SudokuGrid(input);

Console.WriteLine();
Console.WriteLine("-------------- Starting grid --------------");
Console.WriteLine();

grid.Print();

var solver = new SudokuGridSolver(grid);

solver.ApplyRules();
solver.Solve();

var (rowCount, colCount, elapsedTime) = solver.GetMetrics();

Console.WriteLine();
Console.WriteLine($"-------------- Completed in {elapsedTime} with {rowCount} row and {colCount} column iterations --------------");
Console.WriteLine();

grid.Print();
