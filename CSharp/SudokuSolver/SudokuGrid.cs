namespace SudokuSolver;

public class SudokuGrid
{
    public readonly Cell[][] Grid;

    public SudokuGrid(Cell[][] grid)
    {
        Grid = grid;
    }

    public bool IsComplete() =>
        GetRowBlocks()
            .Concat(GetColumnBlocks())
            .Concat(GetSquareBlocks())
            .All(x => x.IsComplete());

    public bool CanBeCompleted() =>
        GetRowBlocks()
            .Concat(GetColumnBlocks())
            .Concat(GetSquareBlocks())
            .All(x => x.IsCompletable());

    public void Print()
    {
        foreach (var row in Grid.Select(r => r))
        {
            Console.WriteLine(string.Join(" ", row.Select(x => x.GetValue())));
        }
    }
    
    public IEnumerable<Cell[]> GetRows() =>
        Grid.Select(row => row.Select(n => n).ToArray());

    private IEnumerable<SudokuBlock> GetRowBlocks() => GetRows().Select(r => new SudokuBlock(r));

    public IEnumerable<Cell[]> GetColumns() =>
        Enumerable.Range(0, 9).Select(x => Grid.Select(row => row[x]).ToArray());

    private IEnumerable<SudokuBlock> GetColumnBlocks() => GetColumns().Select(c => new SudokuBlock(c));

    public IEnumerable<Cell[]> GetSquares()
    {
        for (var row = 0; row < 9; row += 3)
        {
            for (var col = 0; col < 9; col += 3)
            {
                yield return new[]
                {
                    Grid[row][col], Grid[row][col + 1], Grid[row][col + 2],
                    Grid[row + 1][col], Grid[row + 1][col + 1], Grid[row + 1][col + 2],
                    Grid[row + 2][col], Grid[row + 2][col + 1], Grid[row + 2][col + 2]
                };
            }
        }
    }

    private IEnumerable<SudokuBlock> GetSquareBlocks() => GetSquares().Select(s => new SudokuBlock(s));
}