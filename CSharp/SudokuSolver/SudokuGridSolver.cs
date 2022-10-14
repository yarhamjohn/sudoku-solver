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
        var rows = _grid.GetRows().ToArray();
        ExcludeValuesKnownInBlock(rows);
        
        var columns = _grid.GetColumns().ToArray();
        ExcludeValuesKnownInBlock(columns);
        
        var boxes = _grid.GetBoxes().ToArray();
        ExcludeValuesKnownInBlock(boxes);

        var doWhile = true;
        while (doWhile)
        {
            var rowChanged = false;
            var colChanged = false;
            var boxChanged = false;
            
            foreach (var row in rows)
            {
                rowChanged = ExcludeValuesOnlyPossibleInOneCell(row);
            }

            foreach (var col in columns)
            {
                colChanged = ExcludeValuesOnlyPossibleInOneCell(col);
            }

            foreach (var box in boxes)
            {
                boxChanged = ExcludeValuesOnlyPossibleInOneCell(box);
            }

            doWhile = rowChanged || colChanged || boxChanged;
        }
    }

    public (int rowCount, int colCount, TimeSpan elapsedTime) GetMetrics() =>
        (_rowCount, _colCount, _stopwatch.Elapsed);

    public void Solve()
    {
        _stopwatch.Start();
        
        DepthFirstSearch();
        
        _stopwatch.Stop();
    }

    private void DepthFirstSearch()
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

    private static void ExcludeValuesKnownInBlock(IEnumerable<Cell[]> blocks)
    {
        foreach (var block in blocks)
        {
            var knownValues = GetKnownValues(block);
            foreach (var cell in block)
            {
                if (!cell.IsKnown())
                {
                    cell.RemovePossibleValues(knownValues);
                }
            }
        }
    }

    private static List<int> GetKnownValues(Cell[] block) => 
        block.Where(c => c.IsKnown()).Select(c => c.GetValue()).ToList();

    private static bool ExcludeValuesOnlyPossibleInOneCell(Cell[] block)
    {
        var valuesPossibleInOneCell = block
            .SelectMany(r => r.GetPossibleValues())
            .GroupBy(v => v)
            .Where(g => g.Count() == 1)
            .Select(g => g.Key);

        var valuesToUpdate = valuesPossibleInOneCell.Except(GetKnownValues(block)).ToArray();

        foreach (var val in valuesToUpdate)
        {
            block.Single(c => c.HasPossibleValue(val)).SetKnownValue(val);
        }

        return valuesToUpdate.Any();
    }
}