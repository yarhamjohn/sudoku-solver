namespace SudokuSolver;

public class SudokuGrid
{
    public readonly Node[][] Grid;

    public SudokuGrid(Node[][] grid)
    {
        Grid = grid;
    }

    public bool IsComplete() =>
        GetRows()
            .Concat(GetColumns())
            .Concat(GetSquares())
            .All(x => x.IsComplete());

    public bool CanBeCompleted() =>
        GetRows()
            .Concat(GetColumns())
            .Concat(GetSquares())
            .All(x => x.IsCompletable());

    public void Print()
    {
        foreach (var row in Grid.Select(r => r))
        {
            Console.WriteLine(string.Join(" ", row.Select(x => x.Cell.GetValue())));
        }
    }

    private IEnumerable<SudokuBlock> GetRows() =>
        Grid.Select(row => new SudokuBlock(row.Select(n => n.Cell)));

    private IEnumerable<SudokuBlock> GetColumns() =>
        Enumerable.Range(0, 9).Select(x => new SudokuBlock(Grid.Select(row => row[x].Cell)));

    private IEnumerable<SudokuBlock> GetSquares()
    {
        for (var row = 0; row < 9; row += 3)
        {
            for (var col = 0; col < 9; col += 3)
            {
                yield return new SudokuBlock(new[]
                {
                    Grid[row][col].Cell, Grid[row][col + 1].Cell, Grid[row][col + 2].Cell,
                    Grid[row + 1][col].Cell, Grid[row + 1][col + 1].Cell, Grid[row + 1][col + 2].Cell,
                    Grid[row + 2][col].Cell, Grid[row + 2][col + 1].Cell, Grid[row + 2][col + 2].Cell
                });
            }
        }
    }
}