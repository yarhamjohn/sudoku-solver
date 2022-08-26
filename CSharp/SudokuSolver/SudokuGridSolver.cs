namespace SudokuSolver;

public static class SudokuGridSolver
{
    public static void ApplyRules(SudokuGrid grid)
    {
        var rows = grid.GetRows();
        ExcludeValuesKnownInBlock(rows);
        
        var columns = grid.GetColumns();
        ExcludeValuesKnownInBlock(columns);
        
        var squares = grid.GetSquares();
        ExcludeValuesKnownInBlock(squares);
    }

    public static (int rowCount, int colCount) Solve(SudokuGrid grid)
    {
        var rowCount = 0;
        var colCount = 0;

        while (!grid.IsComplete())
        {
            for (var row = 0; row < 9; row++)
            {
                for (var col = 0; col < 9; col++)
                {
                    var currentCell = grid.GetCell(row, col);
                    
                    if (!currentCell.IsKnown())
                    {
                        currentCell.Increment();
                
                        if (!grid.CanBeCompleted())
                        {
                            if (col == 0 && row == 0)
                            {
                                row -= 1;
                                break;
                            }
        
                            while (currentCell.GetValue() <= 9)
                            {
                                currentCell.Increment();
                                if (grid.CanBeCompleted())
                                {
                                    break;
                                }
                            }
        
                            if (currentCell.GetValue() == 10)
                            {
                                currentCell.Reset();
                                var (nextRow, nextCol) = GetLastUnknownCell(grid, row, col);
        
                                row = nextRow;
                                col = nextCol;
                            }
                        }
                    }
                    colCount += 1;
                }
        
                rowCount += 1;
            }
        }

        return (rowCount, colCount);
    }
    
    private static (int nextRow, int nextCol) GetLastUnknownCell(SudokuGrid grid, int row, int col)
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
                if (!grid.GetCell(r, c).IsKnown())
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