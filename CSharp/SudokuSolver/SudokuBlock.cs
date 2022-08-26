namespace SudokuSolver;

public class SudokuBlock
{
    private readonly IEnumerable<Cell> _cells;
    private readonly List<int> _expectedValues = new() { 1, 2, 3, 4, 5, 6, 7, 8, 9 };

    public SudokuBlock(IEnumerable<Cell> cells)
    {
        _cells = cells;
    }

    public bool IsComplete() =>
        !_expectedValues.Except(_cells.Select(c => c.GetValue()).Distinct()).Any();

    public bool IsCompletable()
    {
        var knownValues = _cells.Select(c => c.GetValue()).Where(v => v != 0).ToList();

        var anyInvalidValues = !knownValues.Except(_expectedValues).Any();

        var anyDuplicateValues = !knownValues.GroupBy(v => v).Any(g => g.Count() > 1);
        
        return anyInvalidValues && anyDuplicateValues;
    }
}