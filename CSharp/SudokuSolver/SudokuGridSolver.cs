namespace SudokuSolver;

public static class SudokuGridSolver
{
    public static void ApplyRules(SudokuGrid nodeGrid)
    {
        var rows = nodeGrid.GetRows();
        ExcludeValuesKnownInBlock(rows);
        
        var columns = nodeGrid.GetColumns();
        ExcludeValuesKnownInBlock(columns);
        
        var squares = nodeGrid.GetSquares();
        ExcludeValuesKnownInBlock(squares);
    }

    public static (int rowCount, int colCount) Solve(SudokuGrid grid)
    {
        var rowCount = 0;
        var colCount = 0;

        var startingGrid = grid.Grid;

        while (!grid.IsComplete())
        {
            for (var row = 0; row < startingGrid.Length; row++)
            {
                for (var col = 0; col < startingGrid.Length; col++)
                {
                    var currentNode = startingGrid[row][col];
                    
                    if (!currentNode.Cell.IsKnown())
                    {
                        currentNode.Cell.Increment();
                
                        if (!grid.CanBeCompleted())
                        {
                            if (col == 0 && row == 0)
                            {
                                row -= 1;
                                break;
                            }
        
                            while (currentNode.Cell.GetValue() <= 9)
                            {
                                currentNode.Cell.Increment();
                                if (grid.CanBeCompleted())
                                {
                                    break;
                                }
                            }
        
                            if (currentNode.Cell.GetValue() == 10)
                            {
                                currentNode.Cell.Reset();
                                var (nextRow, nextCol) = GetLastUnknownCell(startingGrid, row, col);
        
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
    
    private static (int nextRow, int nextCol) GetLastUnknownCell(Node[][] grid, int row, int col)
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
                if (!grid[r][c].Cell.IsKnown())
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