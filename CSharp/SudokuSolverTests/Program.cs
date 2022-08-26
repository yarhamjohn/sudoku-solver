namespace SudokuSolverTests;

public class Tests
{
    [Test]
    public void Test1()
    {
        var entryPoint = typeof(Program).Assembly.EntryPoint!;
        
        var input = new [] { "530070000600195000098000060800060003400803001700020006060000280000419005000080079" };
        
        entryPoint.Invoke(null, new object[] { input });
    }
}