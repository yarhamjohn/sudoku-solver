﻿namespace SudokuSolver;

public class SudokuGrid
{
    private readonly Cell[][] _grid;

    public SudokuGrid(string input)
    {
        _grid = input
            .Chunk(9)
            .Select(x => x
                .Select(y => new Cell(Convert.ToInt32(y.ToString())))
                .ToArray())
            .ToArray();
    }

    public Cell GetCell(int row, int col) => _grid[row][col];

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
        foreach (var row in _grid.Select(r => r))
        {
            Console.WriteLine(string.Join(" ", row.Select(x => x.GetValue())));
        }
    }
    
    public IEnumerable<Cell[]> GetRows() =>
        _grid.Select(row => row.Select(n => n).ToArray());

    private IEnumerable<SudokuBlock> GetRowBlocks() => GetRows().Select(r => new SudokuBlock(r));

    public IEnumerable<Cell[]> GetColumns() =>
        Enumerable.Range(0, 9).Select(x => _grid.Select(row => row[x]).ToArray());

    private IEnumerable<SudokuBlock> GetColumnBlocks() => GetColumns().Select(c => new SudokuBlock(c));

    public IEnumerable<Cell[]> GetSquares()
    {
        for (var row = 0; row < 9; row += 3)
        {
            for (var col = 0; col < 9; col += 3)
            {
                yield return new[]
                {
                    _grid[row][col], _grid[row][col + 1], _grid[row][col + 2],
                    _grid[row + 1][col], _grid[row + 1][col + 1], _grid[row + 1][col + 2],
                    _grid[row + 2][col], _grid[row + 2][col + 1], _grid[row + 2][col + 2]
                };
            }
        }
    }

    private IEnumerable<SudokuBlock> GetSquareBlocks() => GetSquares().Select(s => new SudokuBlock(s));
}