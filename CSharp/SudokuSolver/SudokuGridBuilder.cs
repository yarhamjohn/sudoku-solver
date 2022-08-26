namespace SudokuSolver;

public static class SudokuGridBuilder
{
    public static SudokuGrid Build(string input)
    {
        var grid = input.Chunk(9).Select(ConvertToRow).ToArray();

        return new SudokuGrid(grid);
    }
    
    private static Cell[] ConvertToRow(char[] input) =>
        input.Select(x => new Cell(ToInt(x))).ToArray();

    private static int ToInt(char x) => Convert.ToInt32(x.ToString());
}