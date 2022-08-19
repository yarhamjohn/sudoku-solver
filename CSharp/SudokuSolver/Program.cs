var inputDefinition = args[0];

var startingGrid = inputDefinition.Chunk(9).Select(ConvertToRow).ToArray();

PrintGrid(startingGrid);

var rowCounter = 0;
var colCounter = 0;

while (!GridIsComplete(startingGrid))
{
    for (var row = 0; row < startingGrid.Length; row++)
    {
        for (var col = 0; col < startingGrid.Length; col++)
        {
            if (!startingGrid[row][col].IsKnown)
            {
                startingGrid[row][col].Increment();

                var gridIsValid = GridIsValid(startingGrid);

                if (!gridIsValid)
                {
                    if (col == 0 && row == 0)
                    {
                        row -= 1;
                        break;
                    }

                    while (startingGrid[row][col].Value <= 9)
                    {
                        startingGrid[row][col].Increment();
                        if (GridIsValid(startingGrid))
                        {
                            break;
                        }
                    }

                    if (startingGrid[row][col].Value == 10)
                    {
                        startingGrid[row][col].Reset();
                        var (nextRow, nextCol) = GetLastUnknownCell(startingGrid, row, col);

                        row = nextRow;
                        col = nextCol;
                    }
                }
            }
            
            colCounter += 1;
        }
        
        rowCounter += 1;
    }
}

Console.WriteLine();
Console.WriteLine($"-------------- Complete in {rowCounter} row and {colCounter} column iterations ----------------");
Console.WriteLine();

PrintGrid(startingGrid);


(int nextRow, int nextCol) GetLastUnknownCell(Cell[][] grid, int row, int col)
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
            if (!grid[r][c].IsKnown)
            {
                return (r, c - 1);
            }
        }

        startingCol = 8;
    }

    return (0, 0);
}

Cell[] ConvertToRow(char[] input) => input.Select(x => new Cell(Convert.ToInt32(x.ToString()))).ToArray();

bool GridIsComplete(Cell[][] grid) => GridIsValid(grid) && !grid.SelectMany(row => row.Where(c => c.Value == 0)).Any();

bool GridIsValid(Cell[][] grid) => !AnyRowIsInvalid(grid) && !AnyColIsInvalid(grid) && !AnySquareIsInvalid(grid);

bool AnyRowIsInvalid(Cell[][] grid) => grid.Any(row => !BlockIsValid(row));

bool AnyColIsInvalid(Cell[][] grid) => Enumerable.Range(0, 9).Any(x => !BlockIsValid(grid.Select(r => r[x]).ToArray()));

bool AnySquareIsInvalid(Cell[][] grid) {
    for (var i = 0; i < 9; i += 3)
    {
        for (var j = 0; j < 9; j += 3)
        {
            var block = new[]
            {
                grid[i][j], grid[i][j + 1], grid[i][j + 2],
                grid[i + 1][j], grid[i + 1][j + 1], grid[i + 1][j + 2],
                grid[i + 2][j], grid[i + 2][j + 1], grid[i + 2][j + 2]
            };

            if (!BlockIsValid(block))
            {
                return true;
            }
        }
    }

    return false;
}

bool BlockIsValid(Cell[] block) => 
    !block.Any(x => x.Value > 9) 
    && block
        .Where(c => c.Value != 0)
        .GroupBy(c => c.Value)
        .All(g => g.Count() == 1);


void PrintGrid(Cell[][] grid)
{
    foreach (var row in startingGrid)
    {
        Console.WriteLine(string.Join(" ", row.Select(x => x)));
    }
}

public class Cell
{
    public int Value;
    public readonly bool IsKnown;
    
    public Cell(int input)
    {
        Value = input;
        IsKnown = input != 0;
    }

    public void Increment() => Value += 1;

    public void Reset() => Value = 0;

    public override string ToString()
    {
        return Value.ToString();
    }
}
