using SudokuSolver;

namespace SudokuSolverTests;

[TestFixture]
public class SudokuBlockTests
{
    [TestCase(new [] {1, 2, 3, 4, 5, 6, 7, 8, 9}, true)]
    [TestCase(new [] {1, 2, 3, 4, 5, 6, 7, 8}, false)]
    [TestCase(new [] {0, 2, 3, 4, 5, 6, 7, 8, 9}, false)]
    public void IsComplete_ReturnsCorrectAnswer(int[] input, bool expected)
    {
        var block = new SudokuBlock(input.Select(x => new Cell(x)));

        var result = block.IsComplete();
        
        Assert.That(result, Is.EqualTo(expected));
    }
    
    [TestCase(new [] {1, 2, 3, 4, 5, 6, 7, 8, 9}, true)]
    [TestCase(new [] {0, 1, 2, 3, 4, 5, 6, 7, 8}, true)]
    [TestCase(new [] {0, 0, 0, 0, 5, 6, 7, 8, 9}, true)]
    [TestCase(new [] {0, 1, 2, 3, 4, 5, 6, 7, 7}, false)]
    [TestCase(new [] {0, 0, 2, 3, 4, 5, 6, 7, 7}, false)]
    public void IsCompletable_ReturnsCorrectAnswer(int[] input, bool expected)
    {
        var block = new SudokuBlock(input.Select(x => new Cell(x)));

        var result = block.IsCompletable();
        
        Assert.That(result, Is.EqualTo(expected));
    }
}