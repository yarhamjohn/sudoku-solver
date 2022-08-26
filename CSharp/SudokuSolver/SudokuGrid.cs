namespace SudokuSolver;

public class SudokuGrid
{
    private readonly List<int> _expectedValues = new() { 1, 2, 3, 4, 5, 6, 7, 8, 9 };
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
        GetRows()
            .Concat(GetColumns())
            .Concat(GetBoxes())
            .All(BlockIsComplete);

    public bool CanBeCompleted() =>
        GetRows()
            .Concat(GetColumns())
            .Concat(GetBoxes())
            .All(BlockIsCompletable);

    public void Print()
    {
        foreach (var row in _grid.Select(r => r))
        {
            Console.WriteLine(string.Join(" ", row.Select(x => x.GetValue())));
        }
    }
    
    public IEnumerable<Cell[]> GetRows() =>
        _grid.Select(row => row.Select(n => n).ToArray());

    public IEnumerable<Cell[]> GetColumns() =>
        Enumerable.Range(0, 9).Select(x => _grid.Select(row => row[x]).ToArray());

    public IEnumerable<Cell[]> GetBoxes()
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
    
    private bool BlockIsComplete(IEnumerable<Cell> block) =>
        !_expectedValues.Except(block.Select(c => c.GetValue()).Distinct()).Any();

    private bool BlockIsCompletable(IEnumerable<Cell> block)
    {
        var knownValues = block.Select(c => c.GetValue()).Where(v => v != 0).ToList();

        var anyInvalidValues = !knownValues.Except(_expectedValues).Any();

        var anyDuplicateValues = !knownValues.GroupBy(v => v).Any(g => g.Count() > 1);
        
        return anyInvalidValues && anyDuplicateValues;
    }
}