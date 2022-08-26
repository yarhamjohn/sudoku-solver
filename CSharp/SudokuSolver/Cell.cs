namespace SudokuSolver;

public class Cell
{
    private readonly List<int> _possibleValues;

    private int _currentValue;
    
    public Cell(int input)
    {
        if (input > 9)
        {
            throw new InvalidOperationException("Cells can't have a value greater than 9.");
        }

        _possibleValues = input == 0 ? new List<int> { 1, 2, 3, 4, 5, 6, 7, 8, 9 } : new List<int> { input };
        _currentValue = input;
    }

    public bool IsKnown() =>
        _possibleValues.Count == 1;

    public void RemovePossibleValues(IEnumerable<int> values)
    {
        foreach (var value in values)
        {
            _possibleValues.Remove(value);
        }

        if (_possibleValues.Count == 1)
        {
            _currentValue = _possibleValues.Single();
        }
    }

    public int GetValue() => _currentValue;
    public void Increment() => _currentValue += 1;

    public void Reset() => _currentValue = 0;
}