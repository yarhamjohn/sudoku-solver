using System.Diagnostics;

namespace SudokuSolver;

public class SudokuGridSolver
{
    private readonly SudokuGrid _grid;

    private int _rowCount;
    private int _colCount;

    private readonly Stopwatch _stopwatch = new();

    public SudokuGridSolver(SudokuGrid grid)
    {
        _grid = grid;
    }
    
    public void ApplyRules()
    {
        var rows = _grid.GetRows();
        ExcludeValuesKnownInBlock(rows);
        
        var columns = _grid.GetColumns();
        ExcludeValuesKnownInBlock(columns);
        
        var squares = _grid.GetBoxes();
        ExcludeValuesKnownInBlock(squares);
    }

    public (int rowCount, int colCount, TimeSpan elapsedTime) GetMetrics() =>
        (_rowCount, _colCount, _stopwatch.Elapsed);

    public void Solve()
    {
        _stopwatch.Start();
        
        while (!_grid.IsComplete())
        {
            for (var row = 0; row < 9; row++)
            {
                for (var col = 0; col < 9; col++)
                {
                    var currentCell = _grid.GetCell(row, col);
                    
                    if (!currentCell.IsKnown())
                    {
                        currentCell.Increment();
                
                        if (!_grid.CanBeCompleted())
                        {
                            if (col == 0 && row == 0)
                            {
                                row -= 1;
                                break;
                            }
        
                            while (currentCell.GetValue() <= 9)
                            {
                                currentCell.Increment();
                                if (_grid.CanBeCompleted())
                                {
                                    break;
                                }
                            }
        
                            if (currentCell.GetValue() == 10)
                            {
                                currentCell.Reset();
                                var (nextRow, nextCol) = GetLastUnknownCell(row, col);
        
                                row = nextRow;
                                col = nextCol;
                            }
                        }
                    }
                    _colCount += 1;
                }
        
                _rowCount += 1;
            }
        }
        
        _stopwatch.Stop();
    }
    
    private (int nextRow, int nextCol) GetLastUnknownCell(int row, int col)
    {
        int startingRow;
        int startingCol;

        if (col == 0)
        {
            startingRow = row - 1;
            startingCol = 8;
        }
        else
        {
            startingRow = row;
            startingCol = col - 1;
        }
    
        for (var r = startingRow; r >= 0; r--)
        {
            for (var c = startingCol; c >= 0; c--)
            {
                if (!_grid.GetCell(r, c).IsKnown())
                {
                    return (r, c - 1);
                }
            }

            startingCol = 8;
        }

        return (0, 0);
    }

    private static void ExcludeValuesKnownInBlock(IEnumerable<Cell[]> rows)
    {
        foreach (var row in rows)
        {
            var knownValues = row.Select(r => r.GetValue()).Where(v => v != 0).ToList();
            foreach (var cell in row)
            {
                if (!cell.IsKnown())
                {
                    cell.RemovePossibleValues(knownValues);
                }
            }
        }
    }
}